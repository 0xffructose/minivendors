<template>
  <main>
    <router-view />
  </main>
</template>

<script setup lang="ts">
</script>

<style scoped>
nav {
  padding: 1rem;
  margin-bottom: 1rem;
  background-color: #f8f9fa;
  border-radius: 8px;
}
nav a {
  font-weight: bold;
  color: #2c3e50;
  text-decoration: none;
  margin: 0 10px;
}
nav a.router-link-exact-active {
  color: #42b983; /* Aktif olan sayfanın link rengi yeşil olur */
}
</style>