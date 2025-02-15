<template>
  <div class="container">
    <button
      type="button"
      class="color-btn"
      @click="toggleColorPicker"
    >
      <div class="preview" :style="{ backgroundColor: `#${modelValue || defaultColor}` }"></div>
      <span class="color-hex">#{{ modelValue || defaultColor }}</span>
    </button>
    <div v-if="showColorPicker" class="color-picker">
      <input
        type="color"
        :value="`#${modelValue || defaultColor}`"
        @input="updateSelectedColor"
        class="color-input"
      >
      <span class="color-hex">#{{ modelValue || defaultColor }}</span>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    modelValue: String,
    defaultColor: {
      type: String,
      default: '4CAF50'
    }
  },
  data() {
    return {
      showColorPicker: false
    };
  },
  methods: {
    toggleColorPicker() {
      this.showColorPicker = !this.showColorPicker;
    },
    updateSelectedColor(event) {
      const color = event.target.value.substring(1).toUpperCase();
      this.$emit('update:modelValue', color);
      this.$emit('updatecolor', color);
    },
  },
};
</script>

<style scoped>
.container {
  position: relative;
  display: inline-block;
}

.color-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 0.5rem;
  border: 1px solid #ddd;
  background: white;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.2s;
}

.preview {
  width: 24px;
  height: 24px;
  border-radius: 4px;
}

.color-picker {
  position: absolute;
  top: 100%;
  left: 0;
  margin-top: 0.5rem;
  padding: 0.5rem;
  background: white;
  border: 1px solid #ddd;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  display: flex;
  align-items: center;
  gap: 8px;
  z-index: 10;
}

.color-input {
  width: 50px;
  height: 30px;
  padding: 0;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.color-hex {
  font-family: monospace;
  font-size: 0.9em;
  color: #666;
}
</style>
