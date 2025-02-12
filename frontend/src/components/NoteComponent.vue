<template>
  <div class="note" @click="openModal">
    <div class="title">{{ note.title }}</div>
    <div class="content">{{ note.content }}</div>
  </div>
</template>

<script>
export default {
  name: 'NoteComponent',
  props: {
    note: Object,
  },
  methods: {
    openModal() {
      this.$emit('open-modal', this.note);
    }
  }
};
</script>

<style scoped>
.note {
  width: 200px;
  height: 200px;
  background: #f9f9f9;
  border-radius: 10px;
  box-shadow: 2px 2px 5px rgba(0, 0, 0, 0.1);
  display: flex;
  flex-direction: column;
  cursor: pointer;
  overflow: hidden;
  position: relative; /* Added for proper title positioning */
}

.title {
  position: absolute; /* Position the title absolutely */
  top: 0;
  left: 0;
  right: 0; /* Stretch across the full width */
  background: #4CAF50;
  color: white;
  padding: 10px;
  font-weight: bold;
  border-radius: 10px 10px 0 0;
  text-overflow: ellipsis;
  white-space: nowrap;
  overflow: hidden;
}

.content {
  margin-top: 40px; /* Add space for the title */
  padding: 10px;
  flex-grow: 1;
  overflow-y: auto; /* Allow scrolling if content is too long */
  word-wrap: break-word; /* Enable word wrapping */
  white-space: pre-wrap; /* Preserve whitespace and wrapping */
}
</style>
