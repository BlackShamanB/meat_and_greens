<template>
  <div class="navbar">
    <div @click="$router.push('/')">Мясо и зелень</div>
    <div class="navbar__btns">
      <div @click="$router.push('/about')">О нас</div>
      <div @click="$router.push('/store')">Оптовикам</div>
      <div @click="$router.push('/composition')">Дистрибьютерам</div>
      <!-- <div @click="$router.push('/about')">О сайте</div> -->
    </div>
  </div>
</template>
<script>
export default {
  name: "nav-bar"
};
</script>
<style>
.navbar {
  position: fixed;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  height: 50px;
  /* background-color: #c5b79a; */
  background: -webkit-linear-gradient(180deg, #c5b79a 50%, #dbd6c0 100%);
  /* box-shadow: 2px 2px 4px gray; */
  display: flex;
  align-items: center;
  padding: 2rem 2.3rem 2rem 1rem;
  width: 100%;
  max-width: calc(100vw - 3.3rem);
  z-index: 1;
}
.navbar__btns {
  display: flex;
  flex-direction: row;
  gap: 20px;
}
.navbar div {
  cursor: pointer;
}
</style>
