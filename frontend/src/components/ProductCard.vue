<template>
  <div class="productCard" :style="{ backgroundColor: product.color }">
    <img class="product_img" :src="getHref" alt="" />
    <div class="product_name">{{ product.name }}</div>
    <div class="product_price">Цена {{ product.price }}₽</div>
  </div>
</template>
<script>
export default {
  data() {
    return {};
  },
  props: {
    product: Object,
  },
  computed: {
    getHref() {
      return `${this.product.img}.png`;
    },
  },
};
</script>
<style>
.productCard {
  display: flex;
  flex-direction: column;
  padding: 3.4rem;
  cursor: pointer;
  text-align: left;
  transition: all 0.2s ease;
}
.productCard:not(:hover) {
  background-color: unset !important;
}
.product_img {
  transition: all 0.2s ease;
  transform: translateY(0rem) scale(1);
}
.productCard:hover .product_img {
  transform: translateY(-1.7rem) scale(1);
}
.product_name {
  font-size: 16px;
  line-height: 20px;
  font-weight: 600;
  margin-bottom: 8px;
}
</style>
