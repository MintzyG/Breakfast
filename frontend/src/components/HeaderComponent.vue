<template>
  <header class="header">
    <h1>Welcome to the Dashboard</h1>
  </header>
</template>

<script>
export default {
  name: 'Header',
};
</script>

<style scoped>
.header {
  background-color: #34495e;
  color: white;
  padding: 15px 20px;
  text-align: center;
  width: 100%;
  position: fixed;
  top: 0;
  left: 0;
  height: 60px; /* Adjust as needed */
  z-index: 1000;
}
</style>

