package film

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

//Film struct for http response
type film struct {
	Slug  string `json:"slug"`      //url of film
	Image string `json:"image_url"` //url of image
	Year  string `json:"release_year"`
	Name  string `json:"film_name"`
	Length string `json:"film_length"`
}

//struct for channel to send film and whether is has finshed a user
type filmSend struct {
	film film //film to be sent over channel
	done bool //if user is done
}

type nothingReason int

const (
	INTERSECT = iota
	UNION
)

const FEATURELENGTH = 80

type nothingError struct {
	reason nothingReason
}

func (e *nothingError) ToString() string {
	switch e.reason {
	case INTERSECT:
		return "empty intersect"
	case UNION:
		return "empty union"
	default:
		return "big error"
	}

}

func (e *nothingError) Error() string {
	return e.ToString()
}

const url = "https://letterboxd.com/ajax/poster" //first part of url for getting full info on film
const urlEnd = "menu/linked/125x187/"            // second part of url for getting full info on film
const site = "https://letterboxd.com"

// func main() {
// 	getFilmHandler := http.HandlerFunc(getFilm)
// 	http.Handle("/film", getFilmHandler)
// 	log.Println("serving at :8080")
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8080"
// 		log.Printf("Defaulting to port %s", port)
// 	}

// 	log.Printf("Listening on port %s", port)
// 	http.ListenAndServe(":"+port, nil)
// }

var year int

func init() {
	year = time.Now().Year()
}



//Main handler func for request
func Handler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	log.Println(year)
	query := r.URL.Query() //Get URL Params(type map)
	users, ok := query["users"]
	log.Println(len(users))
	if !ok || len(users) == 0 {
		http.Error(w, "no users", 400)
		return
	}
	_, inter := query["intersect"]
	_, ignore := query["ignore_unreleased"]
	_, short := query["only_shorts"]
	_, feature := query["only_features"]

	var userFilm film
	var err error
	var scrapefunc func([]string, bool, bool) (film, error)
	
	if short {
		scrapefunc = shortScrape()
	} else if feature {
		scrapefunc = featureScrape()
	} else {
		scrapefunc = allScrape()
	}
	
	if ignore {
		if inter {
			if len(users) == 1 {
					userFilm, err = scrapefunc(users, false, true)

				userFilm, err = scrapefunc(users, false, true)
			} else {
				userFilm, err = scrapefunc(users, true, true)
			}
		} else {
			userFilm, err = scrapefunc(users, false, true)
		}
	} else {
		if inter {
			if len(users) == 1 {
				userFilm, err = scrapefunc(users, false, false)
			} else {
				userFilm, err = scrapefunc(users, true, false)
			}
		} else {
			userFilm, err = scrapefunc(users, false, false)
		}
	}
	if err != nil {
		var e *nothingError
		if errors.As(err, &e) {
			switch e.reason {
			case INTERSECT:
				http.Error(w, "Intersect error", 406)
				return
			case UNION:
				http.Error(w, "Union error", 404)
				return
			}
		}
	}

	js, err := json.Marshal(userFilm)
	if err != nil {
		http.Error(w, "internal error", 500)
		return
	}
	w.Write(js)
}


func shortScrape() func([]string, bool, bool) (film, error){
	return func(users []string, inter bool, ignore bool) (film, error) {
		return scrapeUser(users, inter, ignore, true, false)
	}
}

func featureScrape() func([]string, bool, bool) (film, error){
	return func(users []string, inter bool, ignore bool) (film, error) {
		return scrapeUser(users, inter, ignore, false, true)
	}
}

func allScrape() func([]string, bool, bool) (film, error){
	return func(users []string, inter bool, ignore bool) (film, error) {
		return scrapeUser(users, inter, ignore, false, false)
	}
}

//main scraping function
func scrapeUser(users []string, intersect bool, ignore bool, short bool, feature bool) (film, error) {
	var user int = 0          //conuter for number of users increses by one when a users page starts being scraped decreses when user has finished think kinda like a semaphore
	var totalFilms []film     //final list to hold all film
	ch := make(chan filmSend) //channel to send films over
	// start go routine to scrape each user
	for _, a := range users {
		log.Println(a)
		user++
		if strings.Contains(a, "/") {
			if strings.Contains(a,"actor/") {
				go scrapeActor(a, ch)
			} else {
				go scrapeList(a, ch)
			}
		} else {
			go scrape(a, ch)
		}
	}
	for {
		userFilm := <-ch
		if userFilm.done { //if users channel is don't then the scapre for that user has finished so decrease the user count
			user--
			if user == 0 {
				break
			}
		} else {
			totalFilms = append(totalFilms, userFilm.film) //append feilm recieved over channel to list
		}

	}

	//chose random film from list
	if len(totalFilms) == 0 {
		return film{}, &nothingError{reason: UNION}
	}
	log.Print("results")
	var finalFilm film
	var filmList []film
	if intersect {
		intersectList := getintersect(totalFilms,len(users))
		length := len(intersectList)
		if length == 0 {
			return film{}, &nothingError{reason: INTERSECT}
		}
		filmList = intersectList
	} else {
		filmList = totalFilms
	}
	if ignore {
		fmt.Println("ignore")
		filmList = removeCurrentYear(filmList)
	}
	if short {
		fmt.Println("only shorts")
		filmList = onlyShorts(filmList)
	} else if feature {
		fmt.Println("only Feature")
		filmList = onlyFeature(filmList)
	}
	n := rand.Intn(len(filmList))
	log.Println(len(filmList))
	log.Println(n)
	log.Println(filmList[n])
	finalFilm = filmList[n]
	return finalFilm, nil
}

//function to scapre an single user
func scrape(userName string, ch chan filmSend) {
	siteToVisit := site + "/" + userName + "/watchlist"

	ajc := colly.NewCollector(
		colly.Async(true),
	)
	ajc.OnHTML("div#film-page-wrapper", func(e *colly.HTMLElement) {
		name := e.ChildAttr("div.filmposter","data-film-name")
		slug := e.ChildAttr("div.filmposter","data-target-link")
		img := e.ChildAttr("img", "src")
		year := e.ChildAttr("div.filmposter","data-film-release-year")
		lenght := e.ChildText("p.text-footer")
		tempfilm := film{
			Slug:  (site + slug),
			Image: img,
			Year: year,
			Name:  name,
			Length: strings.TrimSpace(before(lenght,"mins")),
		}
		ch <- ok(tempfilm)
	})

	c := colly.NewCollector(
		colly.Async(true),
	)
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 100})
	c.OnHTML(".poster-container", func(e *colly.HTMLElement) { //primary scarer to get url of each film that contian full information
		e.ForEach("div.film-poster", func(i int, ein *colly.HTMLElement) {
			slug := ein.Attr("data-film-slug")
			ajc.Visit(site + slug) //start go routine to collect all film data
		})

	})
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, "watchlist/page") {
			e.Request.Visit(e.Request.AbsoluteURL(link))
		}
	})

	c.Visit(siteToVisit)
	c.Wait()
	ajc.Wait()
	ch <- done()

}


func scrapeList(listnameIn string, ch chan filmSend) {
	siteToVisit := ""
	listname := strings.ToLower(listnameIn)

	if strings.Contains(listname, "/list/") {
		siteToVisit = site + "/" + listname
	} else {
		strslice := strings.Split(listname, "/") //strslice[0] is user name strslice[1] is listname
		siteToVisit = site + "/" + strslice[0] + "/list/" + strslice[1]

	}
	log.Println(siteToVisit)

	ajc := colly.NewCollector(
		colly.Async(true),
	)
	ajc.OnHTML("div#film-page-wrapper", func(e *colly.HTMLElement) {
		name := e.ChildAttr("div.filmposter","data-film-name")
		slug := e.ChildAttr("div.filmposter","data-target-link")
		img := e.ChildAttr("img", "src")
		year := e.ChildAttr("div.filmposter","data-film-release-year")
		lenght := e.ChildText("p.text-footer")
		tempfilm := film{
			Slug:  (site + slug),
			Image: img,
			Year: year,
			Name:  name,
			Length: strings.TrimSpace(before(lenght,"mins")),
		}
		ch <- ok(tempfilm)
	})
	c := colly.NewCollector(
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 100})
	c.OnHTML(".poster-container", func(e *colly.HTMLElement) {
		e.ForEach("div.film-poster", func(i int, ein *colly.HTMLElement) {
			slug := ein.Attr("data-film-slug")
			ajc.Visit(site + slug)
		})

	})
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, "/page") {
			e.Request.Visit(e.Request.AbsoluteURL(link))
		}
	})

	c.Visit(siteToVisit)
	c.Wait()
	ajc.Wait()
	ch <- done()

}

func scrapeActor(actor string, ch chan filmSend) {
	siteToVisit := site + "/" + actor
	fmt.Println(siteToVisit)

	c := colly.NewCollector(
		colly.Async(true),
	)
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 100})
	ajc := colly.NewCollector(
		colly.Async(true),
	)
	ajc.OnHTML("div#film-page-wrapper", func(e *colly.HTMLElement) {
		name := e.ChildAttr("div.filmposter","data-film-name")
		slug := e.ChildAttr("div.filmposter","data-target-link")
		img := e.ChildAttr("img", "src")
		year := e.ChildAttr("div.filmposter","data-film-release-year")
		lenght := e.ChildText("p.text-footer")
		tempfilm := film{
			Slug:  (site + slug),
			Image: img,
			Year: year,
			Name:  name,
			Length: strings.TrimSpace(before(lenght,"mins")),
		}
		ch <- ok(tempfilm)
	})

	c.OnHTML(".poster-container", func(e *colly.HTMLElement) {
		e.ForEach("div.film-poster", func(i int, ein *colly.HTMLElement) {
			slug := ein.Attr("data-film-slug")
			ajc.Visit(site + slug)
		})

	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, "/page") {
			e.Request.Visit(e.Request.AbsoluteURL(link))
		}
	})

	c.Visit(siteToVisit)
	c.Wait()
	ch <- done() // users has finished so send done through channel

}

func ok(f film) filmSend {
	return filmSend{
		film: f,
		done: false,
	}
}

func done() filmSend {
	return filmSend{
		film: film{},
		done: true,
	}
}

func getintersect(filmSlice []film, numOfUsers int) []film {
	keys := make(map[film]int)
	list := []film{}
	for _, entry := range filmSlice {
		i, _ := keys[entry]
		if i < (numOfUsers - 1) {
			keys[entry] ++
		} else {
			list = append(list, entry)
		}
	}
	return list
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func makeBiggerActor(url string) string {
	return strings.ReplaceAll(url, "-0-150-0-225-", "-0-230-0-345-")

}
func makeBigger(url string) string {
	return strings.ReplaceAll(url, "-0-125-0-187-", "-0-230-0-345-")
}

func removeCurrentYear(filmSlice []film) []film {
	list := []film{}
	for _, entry := range filmSlice {
		if entry.Year == "" {
			continue
		}
		filmYear, _ := strconv.Atoi(entry.Year)
		if filmYear < year {
			list = append(list, entry)
		}
	}
	return list
}


func before(value string, a string) string {
    // Get substring before a string.
    pos := strings.Index(value, a)
    if pos == -1 {
        return ""
    }
    return value[0:pos]
}


func onlyShorts(filmSlice []film) []film {
	list := []film{}
	for _, entry := range filmSlice {
		length, err := strconv.Atoi(entry.Length)
		if err != nil {
			continue
		}
		if length < FEATURELENGTH {
			list = append(list, entry)
		}
	}
	return list
}

func onlyFeature(filmSlice []film) []film {
	list := []film{}
	for _, entry := range filmSlice {
		length, err := strconv.Atoi(entry.Length)
		if err != nil {
			continue
		}
		if length >= FEATURELENGTH {
			list = append(list, entry)
		}
	}
	return list
}



