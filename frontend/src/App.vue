<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import Header from './components/Header.vue'
import Checker from './components/Checker.vue'
</script>

<template>
  <div class="container">
    <Header />
    <Checker />
  </div>
</template>

<style scoped>
.container {
  height: 100vh;
  width: 100vw;
  position: fixed;
  /* display: flex; */
  top: 0;
  left: 0;
}
</style>
