<template>
    <button class="btn">
      <slot></slot>
    </button>
  </template>
  <script>
  export default {
    name: "my-button",
  };
  </script>
  <style>
  .btn {
    padding: 10px 15px;
    background-color: none;
    border: 1px solid teal;
    border-radius: 5px;
  }
  </style>
  