<template>
  <div class="note" @click="openModal">
    <div class="title">
      <span class="note-emoji">{{ note.emoji }}</span>
      {{ note.title }}
    </div>
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
  position: relative;
}

.title {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  background: #4CAF50;
  color: white;
  padding: 10px;
  font-weight: bold;
  border-radius: 10px 10px 0 0;
  text-overflow: ellipsis;
  white-space: nowrap;
  overflow: hidden;
  display: flex;
  align-items: center;
}

.note-emoji {
  margin-right: 8px;
  font-size: 1.2em;
}

.content {
  margin-top: 40px;
  padding: 10px;
  flex-grow: 1;
  overflow-y: auto;
  word-wrap: break-word;
  white-space: pre-wrap;
}

.note-emoji {
  margin-right: 8px;
}
</style>
