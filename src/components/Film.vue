<template>
  <div class="hello">
    <div class="input">
      <label>Username(s)</label>
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
        <a v-bind:href="url">{{ name }}</a>
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
</style>
