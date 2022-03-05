<a href="https://github.com/GoodbyteCo/Letterboxd-Watchlist-Picker">
  <img width="230" alt="Watchlist Picker Development" src="public/dev-logo.png">
</a>

<br><br><br>

# Letterboxd Watchlist Picker (Web App)

[[Frequently asked questions](#frequently-asked-questions)]
[[Operational notes](#operational-notes)]
[[Credits](#credits)]

A simple website that gives you a random film off your watchlist (or any list).
The site is built with Vue 2, and deployed at
[watchlistpicker.com](https://watchlistpicker.com) via Vercel. See also, the
[Watchlist Picker backend](https://github.com/GoodbyteCo/Watchlist-Picker-Backend)
and [CLI version](https://github.com/HoloPollock/watchlist-picker).

## Frequently asked questions

If you have a question that isn't listed below, send us an email at
[support@goodbyte.ca](mailto:support@goodbyte.ca).

#### How do I get a film?
Enter your Letterboxd username into the search bar and click "search".

#### Why does it say nothing found?
If your account is set to private, your watchlist cannot be read. If your account is
public, make sure your username is spelt correctly and that there are films in your
watchlist. If you are using the Advanced Options, make sure there are already-released
movies in your list (released before the current year, and have a date).

#### Why does it give me unreleased movies?
By default, Watchlist Picker returns any film in your Letterboxd watchlist. To only see
released films, click on "Advanced Options" (located right below the search bar), and
uncheck "unreleased films" from the included results. Note, as we can not actually see
the exact date of release, all movies released within the current year are also excluded.

#### How do I get a film from a list?
Enter the username of the list creator, followed by a <kbd>/</kbd> and the list
title. For example, you would search through
[this list](https://letterboxd.com/jack/list/its-someones-favorite-movie/), 
by entering: `jack/its-someones-favorite-movie`
([see it in action on Watchlist Picker](https://watchlistpicker.com/?u=jack/its-someones-favorite-movie)).
If you want to just copy the text straight out of the URL, `jack/list/its-someones-favorite-movie`
will work too.

#### Can I get a film from an actor's filmography?
Yes, enter `actor` followed by a <kbd>/</kbd> then the actor's slug. So for example,
if you wanted a Brad Pitt film, you would enter `actor/brad-pitt`.

#### Can I get a film from an directors's filmography?
Yes, enter `director` followed by a <kbd>/</kbd> then the director's slug. So for
example, if you wanted a Miranda July film, you would enter `director/miranda-july`.
	
#### How do I search multiple lists at once?
Enter all of the lists and usernames into the search bar, separated by a space or comma.

#### How do I only show films that are on all of the lists?
By default, the lists are combined into a giant pool of movies. To get only movies
that appear on all of the lists you've entered (the intersection), click on "Advanced
Options" (located right below the search bar), and change "Union" to "Intersection".

#### Why does it say the intersection between the lists is empty?
There was not a single movie that appeared on every single one of the lists you entered.
If you are certain there should be, make sure all of the lists and usernames are spelt correctly.

#### Can I only search for short films?
Yes, uncheck "feature-length films" from the included results in "Advanced Options".

#### Can I exclude short films from results?
Yes, uncheck "short films" from the included results in "Advanced Options".

#### Why does it keep loading a new movie when I refresh the page?
The URL of the page contains your search criteria, so you can easily bookmark your
searches or share with friends. To reset it, simply click on the logo.

#### How do I turn on dark-mode?
Click on the toggle switch in the top right corner.

#### How do I turn off dark-mode?
Click on the toggle switch in the top right corner.

#### Why is the logo spinning?
It's like the Letterboxd logo, but a slot machine.

#### Why is the loading bar stuck?
The site may have lost connection to the API, or it's just taking longer to search
through the list than anticipated (often the really big lists can take a minute or
two). Try opening the site again in a new tab.

#### How can I support the site?
If you wish to help cover operational costs, we have a [ko-fi](https://ko-fi.com/goodbyte).

#### I'm having trouble using the site.
If at any point you have issues using the website, please do not hesitate to contact
us at [support@goodbyte.ca](mailto:support@goodbyte.ca).


## Operational notes

[[Building locally](#building-locally)] [[So where are things?](#so-where-are-things)]

The status of the site can be checked on [the public dashboard](https://watchlistpicker.checklyhq.com/).

### Building locally

Clone the project and run `pnpm install`.

The API uses Vercel functions, so use `vercel dev` to test. However, this requires
a [Vercel](https://vercel.com) account. To test without the films propagating (everything
will work, just the picture of the film will not fill in), run:

```
pnpm run dev
```

### So where are things?

- **Backend:** the backend the site communicates with is located in the `api/` directory,
  and deployed alongside the rest of the app.
- **Frontend:** the frontend is written in Vue and located in the `src/` folder.
- **Assets:** the favicons and occasional image are located in the `public/` directory,
  while the actual on-site logo is located in `Logo.vue` as an svg. Watchlist Picker
  does not import any custom typefaces.

## Credits

Watchlist Picker is a [Goodbyte](https://goodbyte.ca) project. Development by
[Quinn Pollock](https://github.com/HoloPollock) and [Jack Guinane](https://github.com/qjack001).
Many thanks to Letterboxd for letting us scrape their website. 
