<template>
    <div class="border p-4 rounded">
      <h3 class="text-lg font-bold">{{ category.name }}</h3>
    </div>
  </template>
  
  <script setup>
  defineProps(["category"]);
  </script>
  