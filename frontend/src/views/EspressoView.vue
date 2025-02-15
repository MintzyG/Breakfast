<template>
  <div class="espresso-view">
    <div class="content-container">
      <!-- Timer Section -->
      <div class="timer-section">
        <FlowmodoroComponent
          @session-completed="handleSessionComplete"
        />
      </div>

      <!-- Sessions Section -->
      <div class="sessions-section">
        <h2 class="sessions-title">Past Sessions</h2>
        <div class="sessions-container">
          <FlowmodoroEntry
            v-for="session in sessions"
            :key="session.session_id"
            :session="session"
            @delete-session="showDeleteConfirmation"
            @edit-session="showEditModal"
          />
        </div>
      </div>
    </div>

    <!-- Delete Confirmation Modal -->
    <ModalComponent v-if="showDeleteModal" @close="closeDeleteModal">
      <template v-slot:default>
        <h3>Delete Session</h3>
        <p>Are you sure you want to delete this session?</p>
        <div class="modal-actions">
          <button @click="confirmDeleteSession" class="delete-btn">Delete</button>
          <button @click="closeDeleteModal" class="cancel-btn">Cancel</button>
        </div>
      </template>
    </ModalComponent>

    <!-- Edit Modal -->
    <ModalComponent v-if="showEditModal" @close="closeEditModal">
      <template v-slot:default>
        <h3>Edit Session</h3>
        <form @submit.prevent="updateSession" class="edit-form">
          <div class="emoji-picker-container">
            <button type="button" class="emoji-button" @click="toggleEmojiPicker">
              {{ selectedSession?.emoji }}
            </button>
            <div v-if="showEmojiPicker" class="emoji-list">
              <button 
                v-for="emoji in emojiList" 
                :key="emoji"
                type="button"
                class="emoji-option"
                @click="selectEmoji(emoji)"
              >
                {{ emoji }}
              </button>
            </div>
          </div>

          <div class="color-picker-container">
            <label>Color</label>
            <div class="color-input-wrapper">
              <input 
                type="color" 
                :value="selectedSession ? '#' + selectedSession.color : '#4CAF50'"
                @input="updateSessionColor"
              >
              <span class="color-hex">#{{ selectedSession?.color }}</span>
            </div>
          </div>

          <div class="form-group">
            <label>Session Name</label>
            <input 
              v-model="selectedSession.session_name" 
              required
              class="form-input"
            >
          </div>

          <div class="modal-actions">
            <button type="submit" class="save-btn">Save Changes</button>
            <button type="button" @click="closeEditModal" class="cancel-btn">Cancel</button>
          </div>
        </form>
      </template>
    </ModalComponent>
  </div>
</template>

<script>
import FlowmodoroComponent from '../components/FlowmodoroComponent.vue';
import FlowmodoroEntry from '../components/FlowmodoroEntry.vue';
import ModalComponent from '../components/ModalComponent.vue';

export default {
  name: 'EspressoView',
  components: {
    FlowmodoroComponent,
    FlowmodoroEntry,
    ModalComponent
  },
  data() {
    return {
      sessions: [],
      showDeleteModal: false,
      showEditModal: false,
      showEmojiPicker: false,
      sessionToDelete: null,
      selectedSession: null,
      emojiList: ['⏱️', '🎯', '💪', '🧠', '📚', '💻', '🎨', '🎵', '✍️', '🔬']
    };
  },
  mounted() {
    this.fetchSessions();
  },
  methods: {
    async fetchSessions() {
      try {
        const response = await this.$api.get('/espresso', {
          headers: { Authorization: `Bearer ${localStorage.getItem('authToken')}` }
        });
        this.sessions = response.data;
      } catch (error) {
        console.error('Error fetching sessions:', error);
      }
    },
    async handleSessionComplete(session) {
      try {
        await this.$api.post('/espresso', session, {
          headers: { Authorization: `Bearer ${localStorage.getItem('authToken')}` }
        });
        await this.fetchSessions();
      } catch (error) {
        console.error('Error saving session:', error);
      }
    },
    showDeleteConfirmation(session) {
      this.sessionToDelete = session;
      this.showDeleteModal = true;
    },
    closeDeleteModal() {
      this.showDeleteModal = false;
      this.sessionToDelete = null;
    },
    async confirmDeleteSession() {
      try {
        await this.$api.delete(`/espresso/${this.sessionToDelete.session_id}`, {
          headers: { Authorization: `Bearer ${localStorage.getItem('authToken')}` }
        });
        await this.fetchSessions();
        this.closeDeleteModal();
      } catch (error) {
        console.error('Error deleting session:', error);
      }
    },
    showEditModal(session) {
      this.selectedSession = { ...session };
      this.showEditModal = true;
    },
    closeEditModal() {
      this.showEditModal = false;
      this.selectedSession = null;
      this.showEmojiPicker = false;
    },
    toggleEmojiPicker() {
      this.showEmojiPicker = !this.showEmojiPicker;
    },
    selectEmoji(emoji) {
      if (this.selectedSession) {
        this.selectedSession.emoji = emoji;
      }
      this.showEmojiPicker = false;
    },
    updateSessionColor(event) {
      if (this.selectedSession) {
        this.selectedSession.color = event.target.value.substring(1);
      }
    },
    async updateSession() {
      if (!this.selectedSession) return;

      try {
        const updatedSession = {
          session_name: this.selectedSession.session_name,
          emoji: this.selectedSession.emoji,
          color: this.selectedSession.color
        };

        await this.$api.patch(`/espresso/${this.selectedSession.session_id}`, updatedSession, {
          headers: { Authorization: `Bearer ${localStorage.getItem('authToken')}` }
        });

        await this.fetchSessions();
        this.closeEditModal();
      } catch (error) {
        console.error('Error updating session:', error);
      }
    }
  }
};
</script>

<style scoped>
.espresso-view {
  min-height: 100vh;
  background-color: #f8f9fa;
  padding: 2rem;
}

.content-container {
  max-width: 1200px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 2rem;
}

.timer-section {
  position: sticky;
  top: 2rem;
  height: fit-content;
}

.sessions-section {
  min-height: 100%;
}

.sessions-title {
  font-size: 1.5rem;
  font-weight: 600;
  color: #2c3e50;
  margin-bottom: 1.5rem;
}

.sessions-container {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 1rem;
  margin-top: 1.5rem;
}

.delete-btn {
  padding: 0.5rem 1rem;
  background-color: #ef4444;
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  transition: background-color 0.2s;
}

.delete-btn:hover {
  background-color: #dc2626;
}

.cancel-btn {
  padding: 0.5rem 1rem;
  background-color: #e5e7eb;
  color: #374151;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  transition: background-color 0.2s;
}

.cancel-btn:hover {
  background-color: #d1d5db;
}

/* Form Styles */
form {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  min-width: 300px;
}

label {
  font-weight: 500;
  color: #374151;
}

input {
  padding: 0.5rem;
  border: 1px solid #d1d5db;
  border-radius: 6px;
  font-size: 1rem;
}

input:focus {
  outline: none;
  border-color: #60a5fa;
  box-shadow: 0 0 0 2px rgba(96, 165, 250, 0.2);
}

/* Emoji Picker Styles */
.emoji-picker-container {
  position: relative;
  margin-bottom: 1rem;
}

.emoji-button {
  font-size: 1.5rem;
  padding: 0.5rem 1rem;
  border: 1px solid #d1d5db;
  border-radius: 6px;
  background: white;
  cursor: pointer;
  transition: border-color 0.2s;
}

.emoji-button:hover {
  border-color: #9ca3af;
}

.emoji-list {
  position: absolute;
  top: 100%;
  left: 0;
  background: white;
  border: 1px solid #d1d5db;
  border-radius: 6px;
  padding: 0.5rem;
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  gap: 0.5rem;
  z-index: 1000;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
}

.emoji-option {
  font-size: 1.5rem;
  padding: 0.5rem;
  border: none;
  background: none;
  cursor: pointer;
  border-radius: 4px;
  transition: background-color 0.2s;
}

.emoji-option:hover {
  background-color: #f3f4f6;
}

/* Color Picker Styles */
.color-picker-container {
  margin-bottom: 1rem;
}

.color-input-wrapper {
  display: flex;
  align-items: center;
  gap: 1rem;
}

input[type="color"] {
  width: 50px;
  height: 30px;
  padding: 0;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.color-hex {
  font-family: monospace;
  font-size: 0.9rem;
  color: #4b5563;
}

button[type="submit"] {
  padding: 0.5rem 1rem;
  background-color: #10b981;
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-weight: 500;
  transition: background-color 0.2s;
}

button[type="submit"]:hover {
  background-color: #059669;
}

/* Responsive Design */
@media (max-width: 1024px) {
  .content-container {
    grid-template-columns: 1fr;
  }

  .timer-section {
    position: static;
    margin-bottom: 2rem;
  }
}

@media (max-width: 640px) {
  .espresso-view {
    padding: 1rem;
  }

  form {
    min-width: auto;
  }
}
</style>
