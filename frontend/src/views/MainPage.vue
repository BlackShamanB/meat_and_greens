<template>
  <div>
    <div class="page__header">
      <h1>Магазин</h1>
    </div>
    <div class="container">
      <product-listing
        name="классика"
        :product-list="classicProducts"
      ></product-listing>
      <product-listing
        name="сало"
        :product-list="classicProducts"
      ></product-listing>
      <product-listing
        name="котлеты"
        :product-list="classicProducts"
      ></product-listing>
      <product-listing
        name="блины"
        :product-list="classicProducts"
      ></product-listing>
    </div>
  </div>
</template>
<script>
import ProductListing from "@/components/ProductListing.vue";
export default {
  data() {
    return {
      classicProducts: [
        {
          img: "1",
          name: "Пельмени",
          price: "100",
          color: "#a72a31",
        },
        {
          img: "2",
          name: "Манты",
          price: "134",
          color: "#4f972f",
        },
        {
          img: "3",
          name: "Сало",
          price: "876",
          color: "#cebda0",
        },
        {
          img: "4",
          name: "Сало",
          price: "876",
          color: "#cf9e9d",
        },
      ],
    };
  },
  components: {
    ProductListing,
  },
};
</script>
<style>
.page__header {
  padding: 7.7rem 2rem 0;
  margin-bottom: 30px;
  background: -webkit-linear-gradient(180deg, #c5b79a 50%, #dbd6c0 100%);
}
.page__header h1 {
  font-size: 56px;
}
.product__cards {
  display: flex;
}
</style>
