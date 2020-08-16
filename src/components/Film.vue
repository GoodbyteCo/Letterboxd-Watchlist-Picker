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
      <div v-else id="hello">
        <img v-bind:src="img_url" alt="image or film" />
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
      let userlist = inputted.filter(function(el) { return el; });
      if(userlist.length < 1) return;
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
</style>
