<template>
  <div class="container">
    <button
      type="button"
      class="emoji-btn"
      @click="toggleEmojiPicker"
    >
      {{ selectedEmoji || '😀' }}
    </button>
    <div v-if="showEmojiPicker" class="emoji-lst">
      <button 
        v-for="emoji in emojiList"
        :key="emoji"
        type="button"
        class="emoji-opt"
        @click="updateSelectedEmoji(emoji)"
      >
        {{ emoji }}
      </button>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    modelValue: String,
  },
  data() {
    return {
      showEmojiPicker: false,
      emojiList: ['😀', '😅', '😂', '😍', '🥳', '🤔', '😎', '😢', '😡', '👍', '👎'],
      selectedEmoji: null,
    };
  },
  watch: {
    modelValue: {
      immediate: true,
      handler(newValue) {
        this.selectedEmoji = newValue;
      },
    },
  },
  methods: {
    toggleEmojiPicker() {
      this.showEmojiPicker = !this.showEmojiPicker;
    },
    updateSelectedEmoji(emoji) {
      this.selectedEmoji = emoji;
      this.$emit('updateemoji', emoji);
      this.showEmojiPicker = false;
    },
    GetSelectedEmoji() {
      return this.selectedEmoji;
    },
  },
};
</script>

<style scoped>
.container {
  position: relative;
  display: inline-block;
}

.emoji-btn {
  font-size: 1.5rem;
  padding: 0.5rem;
  border: none;
  background: rgba(255, 255, 255, 0.9);
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  cursor: pointer;
  transition: background-color 0.2s;
}

.emoji-btn:hover {
  background-color: #f0f0f0;
}

.emoji-lst {
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  gap: 0.5rem;
  padding: 0.5rem;
  position: absolute;
  background: white;
  border: 1px solid #ddd;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  z-index: 10;
}

.emoji-opt {
  font-size: 1.5rem;
  padding: 0.5rem;
  border: none;
  background: none;
  cursor: pointer;
  border-radius: 4px;
  transition: background-color 0.2s;
}

.emoji-option:hover {
  background-color: #f0f0f0;
}
</style>

