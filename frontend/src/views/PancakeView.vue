<template>
  <div class="pancake">
    <h2>Pancake Notes</h2>

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

export default {
  name: 'PancakeView',
  components: { NoteComponent, PlusButton, ModalComponent },
  data() {
    return {
      showCreateModal: false,
      showEditModal: false,
      showDeleteConfirm: false,
      notes: [],
      newNote: { title: '', content: '' },
      selectedNote: null
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
      this.showEditModal = true;
    },

    closeEditModal() {
      this.showEditModal = false;
      this.selectedNote = null;
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
        this.newNote = { title: '', content: '' };
        this.showCreateModal = false;
      } catch (error) {
        console.error('Error creating note:', error);
      }
    }
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
</style>
