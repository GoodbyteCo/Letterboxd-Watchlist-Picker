<template>
  <div id="app">
    <Film />
  </div>
</template>

<script>
import Film from "./components/Film.vue";

export default {
  name: "App",
  components: {
    Film,
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
</style>
