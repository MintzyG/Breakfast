<template>
  <button class="plus-button" @click="$emit('click')">+</button>
</template>

<script>
export default {
  name: 'PlusButton',
};
</script>

<style scoped>
.plus-button {
  position: fixed;
  bottom: 20px;
  right: 20px;
  width: 50px;
  height: 50px;
  background: #3498db;
  color: white;
  font-size: 24px;
  border: none;
  border-radius: 50%;
  cursor: pointer;
}
.plus-button:hover {
  background: #2980b9;
}
</style>

