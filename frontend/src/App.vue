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
@font-face {
  font-family: "ofont"; 
  src: url("../fonts/ofont.ru_AC_Line.ttf") format("truetype"); 
  font-style: normal; 
  font-weight: normal; 
}
@font-face {
  font-family: "eurofurencelight"; 
  src: url("../fonts/eurofurencelight.ttf") format("truetype"); 
  font-style: normal; 
  font-weight: normal; 
}
body {
  margin: 0;
}
#app {
  font-family: ofont, serif;
  font-weight: 400;
  font-style: normal;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
.app {
  min-height: calc(100vh - 187px);
}
</style>
