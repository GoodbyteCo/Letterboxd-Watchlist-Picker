<template>
  <div id="app">
    <div id="main">
      <Film />
    </div>
    <Footer />
  </div>
</template>

<script>
import Film from "./components/Film.vue";
import Footer from "./components/Footer.vue";
export default {
  name: "App",
  components: {
    Film,
    Footer,
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 30px;
}

#main {
  min-height: -webkit-calc(75vh - 15px);
}
</style>
