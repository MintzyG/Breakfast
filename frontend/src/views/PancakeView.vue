<template>
  <div class="pancake">
    <div class="notes-container">
      <NoteComponent 
        v-for="note in notes" 
        :key="note.note_id" 
        :note="note" 
        @open-modal="handleNoteClick"
      />
    </div>

    <PlusButton @click="showCreateModal = true" />

    <!-- Create Note Modal -->
    <ModalComponent v-if="showCreateModal" @close="showCreateModal = false">
      <template v-slot:default>
        <form @submit.prevent="createNote">
          <div class="emoji-picker-container">
            <EmojiPicker
              v-model="newNote.emoji"
              @updateemoji="updateNoteEmoji($event, 'new')"
            />
          </div>

          <div class="color-picker-container">
            <label>Color</label>
            <ColorPicker
              v-model="newNote.color"
              @updatecolor="updateNoteColor($event, 'new')"
            />
          </div>

          <label>Title</label>
          <input v-model="newNote.title" required>

          <label>Content</label>
          <textarea v-model="newNote.content" required></textarea>

          <button type="submit">Create Note</button>
        </form>
      </template>
    </ModalComponent>

    <!-- Edit Note Modal -->
    <ModalComponent v-if="showEditModal" @close="closeEditModal">
      <template v-slot:default>
        <form @submit.prevent="updateNote">
          <div class="emoji-picker-container">
            <EmojiPicker
              v-model="selectedNote.emoji"
              @updateemoji="updateNoteEmoji($event, 'edit')"
            />
          </div>

          <div class="color-picker-container">
            <label>Color</label>
            <ColorPicker
              v-model="selectedNote.color"
              @updatecolor="updateNoteColor($event, 'edit')"
            />
          </div>

          <label>Title</label>
          <input v-model="selectedNote.title" required>

          <label>Content</label>
          <textarea v-model="selectedNote.content" required></textarea>

          <div class="modal-actions">
            <button type="submit">Save Changes</button>
            <button type="button" @click="showDeleteConfirm = true" class="delete-btn">
              Delete Note
            </button>
          </div>
        </form>
      </template>
    </ModalComponent>

    <!-- Delete Confirmation Modal -->
    <ModalComponent v-if="showDeleteConfirm" @close="showDeleteConfirm = false">
      <template v-slot:default>
        <h3>Are you sure you want to delete this note?</h3>
        <div class="confirm-buttons">
          <button @click="deleteNote" class="confirm-delete">Yes, Delete</button>
          <button @click="showDeleteConfirm = false">Cancel</button>
        </div>
      </template>
    </ModalComponent>
  </div>
</template>

<script>
import NoteComponent from '../components/NoteComponent.vue';
import PlusButton from '../components/PlusButton.vue';
import ModalComponent from '../components/ModalComponent.vue';
import ColorPicker from '../components/ColorPicker.vue';
import EmojiPicker from '../components/EmojiPicker.vue';

export default {
  name: 'PancakeView',
  components: {
    NoteComponent,
    PlusButton,
    ModalComponent,
    ColorPicker,
    EmojiPicker
  },
  data() {
    return {
      showCreateModal: false,
      showEditModal: false,
      showDeleteConfirm: false,
      showEmojiPicker: false, // new
      colorInput: '#4CAF50',
      notes: [],
      newNote: { 
        title: '', 
        content: '',
        emoji: '📝',
        color: '4CAF50'
      },
      selectedNote: null,
      emojiList: ['📝', '📌', '⭐', '❤️', '🎯', '✨', '💡', '📚'] // new
    };
  },
  mounted() {
    this.fetchNotes();
  },
  methods: {
    async fetchNotes() {
      try {
        const response = await this.$api.get('/pancake', {
          headers: { Authorization: `Bearer ${localStorage.getItem('authToken')}` }
        });
        this.notes = response.data;
      } catch (error) {
        console.error('Error fetching notes:', error);
      }
    },

    handleNoteClick(note) {
      this.selectedNote = { ...note };
      this.colorInput = '#' + note.color;
      this.showEditModal = true;
    },

    closeEditModal() {
      this.showEditModal = false;
      this.selectedNote = null;
    },

    updateNoteEmoji(emoji, type) {
      if (type === 'new') {
        this.newNote.emoji = emoji;
      } else {
        this.selectedNote.emoji = emoji;
      }
    },

    updateNoteColor(color, type) {
      if (type === 'new') {
        this.newNote.color = color;
      } else {
        this.selectedNote.color = color;
      }
    },

    async updateNote() {
      try {
        await this.$api.patch(`/pancake/${this.selectedNote.note_id}`, this.selectedNote, {
          headers: { Authorization: `Bearer ${localStorage.getItem('authToken')}` }
        });
        await this.fetchNotes();
        this.closeEditModal();
      } catch (error) {
        console.error('Error updating note:', error);
      }
    },

    async deleteNote() {
      try {
        await this.$api.delete(`/pancake/${this.selectedNote.note_id}`, {
          headers: { Authorization: `Bearer ${localStorage.getItem('authToken')}` }
        });
        await this.fetchNotes();
        this.showDeleteConfirm = false;
        this.closeEditModal();
      } catch (error) {
        console.error('Error deleting note:', error);
      }
    },

    async createNote() {
      try {
        await this.$api.post('/pancake', this.newNote, {
          headers: { Authorization: `Bearer ${localStorage.getItem('authToken')}` }
        });
        await this.fetchNotes();
        this.newNote = { 
          title: '', 
          content: '', 
          emoji: '📝',
          color: '4CAF50' // Reset to default green
        };
        this.colorInput = '#4CAF50'; // Reset color input
        this.showCreateModal = false;
      } catch (error) {
        console.error('Error creating note:', error);
      }
    },
  }
};
</script>

<style scoped>
/* Original styles preserved */
.pancake {
  display: flex;
  flex-direction: column;
  height: 100%;
  padding: 20px;
}

.notes-container {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

form {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

/* New additions for modals only */
.modal-actions {
  margin-top: 20px;
  display: flex;
  gap: 10px;
  justify-content: flex-end;
}

.delete-btn {
  background: #ff4444;
  color: white;
  padding: 8px 16px;
  border-radius: 4px;
}

.confirm-buttons {
  margin-top: 20px;
  display: flex;
  gap: 10px;
  justify-content: flex-end;
}

.confirm-delete {
  background: #ff4444;
  color: white;
  padding: 8px 16px;
  border-radius: 4px;
}

/* Ensure modal content has reasonable sizing */
.modal-content {
  min-width: 400px;
  max-width: 600px;
}

.color-picker-container {
  display: flex;
  flex-direction: column;
}
</style>
