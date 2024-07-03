<template>
  <div class="footer">
    <div class="container">
      <div class="row">
        <div class="col-3">
          <div class="footer-item" @click="$router.push('/about')">О нас</div>
          <div class="footer-item" @click="$router.push('/store')">
            Оптовикам
          </div>
          <div class="footer-item" @click="$router.push('/composition')">
            Дистрибьютерам
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  name: "MyFooter",
};
</script>
<style>
.footer {
  padding: 20px;
  background-color: black;
  text-align: left;
}
.footer-item {
  padding: 15px 20px;
}
</style>
