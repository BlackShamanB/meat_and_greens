<template>
  <nav-bar></nav-bar>
  <div class="app">
    <router-view></router-view>
  </div>
  <my-footer></my-footer>
</template>

<script>
export default {
  name: "App",
};
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Bona+Nova+SC:ital,wght@0,400;0,700;1,400&display=swap');
@import "../node_modules/bootstrap/dist/css/bootstrap-grid.min.css";

body {
  margin: 0;
}
#app {
  font-family: "Bona Nova SC", serif;
  font-weight: 400;
  font-style: normal;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
