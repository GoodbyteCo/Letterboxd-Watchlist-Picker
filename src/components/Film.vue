<template>
  <div>
    <h1>Random Watchlist Picker</h1>
    <div class="hello">
      <p>Enter your Letterboxd username to get a random film off of your watchlist. Enter multiple usernames by seperating with a space or comma.</p>
      <div class="input">
        <label>Username(s):</label>
        <input type="text" v-model="users" />
        <button v-on:click="submit()">Submit</button>
      </div>
      <div v-if="loading">
        <h2>Loading Film</h2>
        <p>Sorry this may take a bit as we scrape letterboxd do to api restrictions</p>
      </div>
      <div v-else-if="pressed">
        <div v-if="notfound">
          <h2>Nothing Found</h2>
          <p>Sorry nothing was found is your watchlist empty</p>
        </div>
        <div v-else id="container">
          <a
            v-bind:href="url"
            :style="{ backgroundImage: 'url(' + img_url + ')' }"
            class="film-cover"
            alt="film poster"
          ></a>
          <div>
            <p class="you-should">You should watch</p>
            <a v-bind:href="url" class="title">{{ name }}</a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Film",
  data() {
    return {
      users: "",
      info: "",
      notfound: false,
      loading: false,
      pressed: false,
    };
  },
  methods: {
    submit() {
      document.body.className = "entered";
      this.pressed = true;
      this.loading = true;
      let inputted = this.users.split(/(?:,| )+/);
      let userlist = inputted.filter(function (el) {
        return el;
      });
      if (userlist.length < 1) return;
      console.log(userlist);
      let url = "https://letterboxd-random.ue.r.appspot.com/film?";
      for (let i = 0; i < userlist.length; i++) {
        if (i > 0) {
          url += "&";
        }
        url += `users=${userlist[i].trim()}`;
      }
      try {
        let vue = this;
        console.log(url);
        fetch(url)
          .then(function (res) {
            vue.loading = false;
            if (res.status != 200) {
              vue.notfound = true;
              return "";
            }

            return res.json();
          })
          .then((json) => (this.info = json));
      } catch (e) {
        this.$alert(
          "Something went wrong. Please try again in a moment. Error:" + e,
          "An error occured"
        );
      }
    },
  },
  computed: {
    filmNotFound: function () {
      return this.notfound;
    },
    url: function () {
      return this.info.slug;
    },
    img_url: function () {
      return this.info.image_url;
    },
    name: function () {
      return this.info.film_name;
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
a {
  font-size: 1.5rem;
  color: #415569;
}

a:hover {
  color: #40bcf4;
}

.hello {
  transform: translateY(70px);
  transition: transform 2s ease;
}

h1 {
  font-size: 1.5rem;
  transform: scale(3) translateY(15px);
  transition: transform 2s ease;
}

.entered h1,
.entered .hello {
  transform: none;
}

h3 {
  margin: 40px 0 0;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

.film-cover {
  margin: auto;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.35), 0 0 2px 1px rgb(0 0 0 / 5%);
  display: block;
  width: 230px;
  height: 345px;
  border-radius: 4px;
}

.film-cover:hover {
  box-shadow: inset 0 0 0 3px #40bcf4, 0 1px 3px rgba(0, 0, 0, 0.35),
    0 0 2px 1px rgb(0 0 0 / 5%);
}

.film-cover::after {
  content: "";
  background-image: linear-gradient(
    90deg,
    hsla(0, 0%, 100%, 0) 0,
    hsla(0, 0%, 100%, 0.5) 50%,
    hsla(0, 0%, 100%, 0)
  );
  display: block;
  width: 100%;
  height: 1px;
}

#container {
  display: grid;
  grid-template-rows: 345px auto;
  gap: 20px;
  text-align: center;
  max-width: 400px;
  width: 90%;
  margin: 60px auto;
}

a.title {
  font-weight: bold;
  font-size: 2rem;
  margin: 0;
  color: inherit;
}

button {
  padding: 0 1rem;
  line-height: 2.8rem;
  font-size: 1.1rem;
  border: 0;
  background: black;
  color: white;
  text-transform: uppercase;
  letter-spacing: 0.02em;
  cursor: pointer;
  border-radius: 0 4px 4px 0;
  outline: none;
}

button:hover,
button:focus,
button:focus-within {
  background: #40bcf4;
}

input {
  border: 0;
  background: #ebebeb;
  padding: 0 1rem;
  line-height: 2.8rem;
  font-size: 1.1rem;
  outline: none;
  border-radius: 4px 0 0 4px;
}

input:active,
input:focus,
input:focus-within {
  box-shadow: inset 0 0 0 3px #40bcf4;
}

label {
  visibility: hidden;
  display: block;
}

p {
  max-width: 60ch;
  margin: 0rem auto 1rem;
}

p.you-should {
  margin: 1rem auto 0.5rem;
}
</style>
