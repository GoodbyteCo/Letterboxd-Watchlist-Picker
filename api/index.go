package film

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

// Film struct for http response
type film struct {
	Slug  string `json:"slug"`      //url of film
	Image string `json:"image_url"` //url of image
	Name  string `json:"film_name"`
}

// struct for channel to send film and whether it has finished a user
type filmSend struct {
	film string //film to be sent over channel
	done bool   //if user is done
}

type nothingReason int

const (
	INTERSECT = iota
	UNION
)

const FEATURELENGTH = 60

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

const urlscrape = "https://letterboxd.com/ajax/poster" //first part of url for getting full info on film
const urlEnd = "std/125x187/"                          // second part of url for getting full info on film
const site = "https://letterboxd.com"

// func main() {
// 	getFilmHandler := http.HandlerFunc(Handler)
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

// Main handler func for request
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
	ignore, _ := query["ignore"]
	filmFilter := ""
	if len(ignore) > 0 {
		filmFilter = generateFilmFilter(ignore[0])
	}

	var userFilm film
	var err error

	if inter && len(users) > 1 {
		userFilm, err = scrapeMain(users, true, filmFilter)
	} else {
		userFilm, err = scrapeMain(users, false, filmFilter)
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

// main scraping function
func scrapeMain(users []string, intersect bool, filmFilter string) (film, error) {
	var user int = 0          //conuter for number of users increses by one when a users page starts being scraped decreses when user has finished think kinda like a semaphore
	var totalFilms []string   //final list to hold all films
	ch := make(chan filmSend) //channel to send films over
	// start go routine to scrape each user
	for _, a := range users {
		user++
		url := getListURL(a)
		go scrape(url, filmFilter, ch)
	}
	for {
		userFilm := <-ch
		if userFilm.done { //if channel is done then the scrape for that user has finished so decrease the user count
			user--
			if user == 0 {
				break
			}
		} else {
			totalFilms = append(totalFilms, userFilm.film) //append film received over channel to list
		}

	}

	//chose random film from list
	if len(totalFilms) == 0 {
		return film{}, &nothingError{reason: UNION}
	}
	log.Print("results")
	var filmList []string
	if intersect {
		intersectList := getIntersect(totalFilms, len(users))
		length := len(intersectList)
		if length == 0 {
			return film{}, &nothingError{reason: INTERSECT}
		}
		filmList = intersectList
	} else {
		filmList = totalFilms
	}

	n := rand.Intn(len(filmList))
	log.Println("total: ", len(filmList))
	log.Println("random index: ", n)
	chosenFilm := filmList[n]
	log.Println("chosen film: ", chosenFilm)

	finalFilm := getFilmDetails(chosenFilm)
	if strings.Contains(finalFilm.Image, "https://s.ltrbxd.com/static/img/empty-poster") {
		finalFilm.Image = "https://watchlistpicker.com/noimagefound.jpg"
	}
	return finalFilm, nil
}

func getListURL(listString string) string {
	listString = strings.ToLower(listString)
	log.Println(listString)
	if strings.Contains(listString, "/") {
		if strings.Contains(listString, "actor/") || strings.Contains(listString, "director/") {
			url := site + "/" + listString
			return url
		} else {
			url := ""
			if strings.Contains(listString, "/list/") {
				url = site + "/" + listString
			} else {
				strslice := strings.Split(listString, "/") //strslice[0] is user name strslice[1] is listname
				url = site + "/" + strslice[0] + "/list/" + strslice[1]
			}
			return url
		}
	} else {
		url := site + "/" + listString + "/watchlist"
		return url
	}
}

func getFilmDetails(chosenFilm string) film {
	var filmDetails film
	ajc := colly.NewCollector(
		colly.Async(true),
	)
	ajc.OnHTML("div.film-poster", func(e *colly.HTMLElement) {
		name := e.Attr("data-film-name")
		img := e.ChildAttr("img", "src")
		filmDetails = film{
			Slug:  (site + chosenFilm),
			Image: makeBigger(img),
			Name:  name,
		}
	})

	ajc.Visit(urlscrape + chosenFilm + urlEnd) //start go routine to collect film data)
	ajc.Wait()
	return filmDetails
}

func scrape(url string, filmFilter string, ch chan filmSend) {
	siteToVisit := url

	c := colly.NewCollector(
		colly.Async(true),
	)
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 100})
	if len(filmFilter) > 0 {
		log.Println("adding cookie for filter: ", filmFilter)
		c.SetCookies(siteToVisit, []*http.Cookie{{Name: "filmFilter", Value: filmFilter, Domain: "letterboxd.com", Path: "/"}})
	}
	c.OnHTML(".poster-container", func(e *colly.HTMLElement) { //primary scarer to get url of each film that contian full information
		e.ForEach("div.film-poster", func(i int, ein *colly.HTMLElement) {
			filmLink := ein.Attr("data-target-link")
			ch <- ok(filmLink)
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

func ok(f string) filmSend {
	return filmSend{
		film: f,
		done: false,
	}
}

func done() filmSend {
	return filmSend{
		film: "",
		done: true,
	}
}

func getIntersect(filmSlice []string, numOfUsers int) []string {
	keys := make(map[string]int)
	list := []string{}
	for _, entry := range filmSlice {
		i, _ := keys[entry]
		if i < (numOfUsers - 1) {
			keys[entry]++
		} else {
			list = append(list, entry)
		}
	}
	return list
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func makeBigger(url string) string {
	return strings.ReplaceAll(url, "-0-125-0-187-", "-0-230-0-345-")
}

func generateFilmFilter(ignoreString string) string {
	var filmFilters []string
	if strings.Contains(ignoreString, "unreleased") {
		filmFilters = append(filmFilters, "hide-unreleased")
	}
	if strings.Contains(ignoreString, "shorts") {
		filmFilters = append(filmFilters, "hide-shorts")
	}
	if strings.Contains(ignoreString, "feature") {
		filmFilters = append(filmFilters, "show-shorts")
	}
	var filmFilter = strings.Join(filmFilters, "%20")
	return filmFilter
}
