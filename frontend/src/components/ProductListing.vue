<template lang="">
  <div class="row mb-60">
    <h2 class="product-listing-title">{{ name }}</h2>
    <div class="col-4" v-for="product in productList" :key="product.name">
      <product-card :product="product" :data-ref="product.name"></product-card>
    </div>
  </div>
</template>
<script>
// import axios from 'axios'
import ProductCard from "@/components/ProductCard.vue";
export default {
  data() {
    return {
      // productList: []
    };
  },
  components: { ProductCard },
  props: {
    name: String,
    productList: Array,
  },
  mounted() {
    // axios.get('http://127.0.0.1:3000/api/products').then(response => {
    //   this.productList = response.data
    // })
  }
};
</script>
<style>
.product-listing-title {
  font-size: 32px;
}
.mb-60 {
  margin-bottom: 60px;
}
</style>
