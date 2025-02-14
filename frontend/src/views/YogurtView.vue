<template>
  <div class="yogurt">
    <div class="task-lists-container">
      <!-- Incomplete Tasks -->
      <div class="task-list">
        <h2 class="task-list-title">To Do</h2>
        <div class="task-container">
          <TaskComponent
            v-for="task in incompleteTasks"
            :key="task.task_id"
            :task="task"
            @task-updated="handleTaskUpdate"
            @task-deleted="handleTaskDelete"
            @open-edit-modal="openEditModal"
          />
        </div>
      </div>

      <!-- Vertical Separator -->
      <div class="separator"></div>

      <!-- Complete Tasks -->
      <div class="task-list">
        <h2 class="task-list-title">Completed</h2>
        <div class="task-container">
          <TaskComponent
            v-for="task in completedTasks"
            :key="task.task_id"
            :task="task"
            @task-updated="handleTaskUpdate"
            @task-deleted="handleTaskDelete"
            @open-edit-modal="openEditModal"
          />
        </div>
      </div>
    </div>
    <PlusButton @click="showCreateModal = true" />

    <!-- Create Modal -->
    <ModalComponent v-if="showCreateModal" @close="showCreateModal = false">
      <template v-slot:default>
        <form @submit.prevent="createTask">
          <div class="emoji-picker-container">
            <button type="button" class="emoji-button" @click="toggleEmojiPicker('create')">
              {{ newTask.emoji }}
            </button>
            <div v-if="showEmojiPicker === 'create'" class="emoji-list">
              <button 
                v-for="emoji in emojiList" 
                :key="emoji"
                type="button"
                class="emoji-option"
                @click="selectEmoji(emoji, 'create')"
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
                :value="'#' + newTask.color"
                @input="updateTaskColor($event, 'create')"
              >
              <span class="color-hex">#{{ newTask.color }}</span>
            </div>
          </div>

          <label>Title</label>
          <input v-model="newTask.title" required>
          
          <label>Priority (1-3)</label>
          <input type="number" min="1" max="3" v-model="newTask.priority" required>
          
          <label>Difficulty (1-3)</label>
          <input type="number" min="1" max="3" v-model="newTask.difficulty" required>
          
          <label>Size (1-3)</label>
          <input type="number" min="1" max="3" v-model="newTask.task_size" required>

          <button type="submit">Create Task</button>
        </form>
      </template>
    </ModalComponent>

    <!-- Edit Modal -->
    <ModalComponent v-if="showEditModal" @close="closeEditModal">
      <template v-slot:default>
        <form @submit.prevent="updateTask">
          <div class="emoji-picker-container">
            <button type="button" class="emoji-button" @click="toggleEmojiPicker('edit')">
              {{ selectedTask.emoji }}
            </button>
            <div v-if="showEmojiPicker === 'edit'" class="emoji-list">
              <button 
                v-for="emoji in emojiList" 
                :key="emoji"
                type="button"
                class="emoji-option"
                @click="selectEmoji(emoji, 'edit')"
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
                :value="'#' + selectedTask.color"
                @input="updateTaskColor($event, 'edit')"
              >
              <span class="color-hex">#{{ selectedTask.color }}</span>
            </div>
          </div>

          <label>Title</label>
          <input v-model="selectedTask.title" required>
          
          <label>Priority (1-3)</label>
          <input type="number" min="1" max="3" v-model="selectedTask.priority" required>
          
          <label>Difficulty (1-3)</label>
          <input type="number" min="1" max="3" v-model="selectedTask.difficulty" required>
          
          <label>Size (1-3)</label>
          <input type="number" min="1" max="3" v-model="selectedTask.task_size" required>

          <button type="submit">Save Changes</button>
        </form>
      </template>
    </ModalComponent>
  </div>
</template>

<script>
import TaskComponent from '../components/TaskComponent.vue'
import PlusButton from '../components/PlusButton.vue';
import ModalComponent from '../components/ModalComponent.vue';

export default {
  name: 'YogurtView',
  components: { TaskComponent, PlusButton, ModalComponent },
  data() {
    return {
      showCreateModal: false,
      showEditModal: false,
      showEmojiPicker: null, // 'create' or 'edit' or null
      tasks: [],
      newTask: { 
        title: '',
        completed: false,
        priority: 1,
        difficulty: 1,
        task_size: 1,
        emoji: '✅',
        color: '4CAF50'
      },
      selectedTask: null,
      emojiList: ['✅', '📌', '⭐', '❤️', '🎯', '✨', '💡', '📚', '🔥', '⚡']
    };
  },
  computed: {
    sortedTasks() {
      return [...this.tasks].sort((a, b) => {
        const getRadixValue = (task) => {
          return (task.priority * 100) + (task.task_size * 10) + task.difficulty;
        };
        const aValue = getRadixValue(a);
        const bValue = getRadixValue(b);
        return bValue - aValue;
      });
    },
    incompleteTasks() {
      return this.sortedTasks.filter(task => !task.completed);
    },
    completedTasks() {
      return this.sortedTasks.filter(task => task.completed);
    }
  },
  mounted() {
    this.fetchTasks();
  },
  methods: {
    async fetchTasks() {
      try {
        const response = await this.$api.get('/yogurt', {
          headers: { Authorization: `Bearer ${localStorage.getItem('authToken')}` }
        });
        this.tasks = response.data;
      } catch (error) {
        console.error('Error fetching tasks:', error);
      }
    },
    toggleEmojiPicker(mode) {
      this.showEmojiPicker = this.showEmojiPicker === mode ? null : mode;
    },
    selectEmoji(emoji, mode) {
      if (mode === 'create') {
        this.newTask.emoji = emoji;
      } else {
        this.selectedTask.emoji = emoji;
      }
      this.showEmojiPicker = null;
    },
    updateTaskColor(event, mode) {
      const color = event.target.value.substring(1);
      if (mode === 'create') {
        this.newTask.color = color;
      } else {
        this.selectedTask.color = color;
      }
    },
    openEditModal(task) {
      this.selectedTask = { ...task };
      this.showEditModal = true;
    },
    closeEditModal() {
      this.showEditModal = false;
      this.selectedTask = null;
      this.showEmojiPicker = null;
    },
    handleTaskUpdate(updatedTask) {
      const index = this.tasks.findIndex(t => t.task_id === updatedTask.task_id);
      if (index !== -1) {
        this.tasks.splice(index, 1, updatedTask);
      }
    },
    handleTaskDelete(taskId) {
      const index = this.tasks.findIndex(t => t.task_id === taskId);
      if (index !== -1) {
        this.tasks.splice(index, 1);
      }
    },
    async createTask() {
      try {
        const response = await this.$api.post('/yogurt', this.newTask, {
          headers: { Authorization: `Bearer ${localStorage.getItem('authToken')}` }
        });
        this.tasks.push(response.data);
        this.newTask = {
          title: '',
          completed: false,
          priority: 1,
          difficulty: 1,
          task_size: 1,
          emoji: '✅',
          color: '4CAF50'
        };
        this.showCreateModal = false;
      } catch (error) {
        console.error('Error creating task:', error);
      }
    },
    async updateTask() {
      try {
        await this.$api.patch(`/yogurt/${this.selectedTask.task_id}`, this.selectedTask, {
          headers: { Authorization: `Bearer ${localStorage.getItem('authToken')}` }
        });
        this.handleTaskUpdate(this.selectedTask);
        this.closeEditModal();
      } catch (error) {
        console.error('Error updating task:', error);
      }
    },
  }
};
</script>

<style scoped>
.yogurt {
  display: flex;
  flex-direction: column;
  align-items: center;
  height: 100%;
  padding: 20px;
}

.task-lists-container {
  display: flex;
  width: 100%;
  gap: 20px;
  margin-bottom: 20px;
  height: calc(100vh - 100px);
  overflow: hidden;
}

.task-list {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  min-width: 0;
}

.task-list-title {
  font-size: 1.5rem;
  font-weight: 600;
  margin-bottom: 20px;
  color: #333;
}

.task-container {
  width: 100%;
  overflow-y: auto;
  padding: 0 10px;
}

.separator {
  width: 1px;
  background-color: #e0e0e0;
  height: 100%;
}

/* Ensure task containers can scroll independently */
.task-container::-webkit-scrollbar {
  width: 8px;
}

.task-container::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 4px;
}

.task-container::-webkit-scrollbar-thumb {
  background: #888;
  border-radius: 4px;
}

.task-container::-webkit-scrollbar-thumb:hover {
  background: #555;
}

/* Emoji Picker Styles */
.emoji-picker-container {
  position: relative;
  margin-bottom: 15px;
}

.emoji-button {
  font-size: 1.5em;
  padding: 5px 15px;
  border: 1px solid #ddd;
  border-radius: 4px;
  background: white;
  cursor: pointer;
}

.emoji-list {
  position: absolute;
  top: 100%;
  left: 0;
  background: white;
  border: 1px solid #ddd;
  border-radius: 4px;
  padding: 10px;
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  gap: 5px;
  z-index: 1000;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.emoji-option {
  font-size: 1.5em;
  padding: 5px;
  border: none;
  background: none;
  cursor: pointer;
  border-radius: 4px;
}

.emoji-option:hover {
  background: #f5f5f5;
}

/* Color Picker Styles */
.color-picker-container {
  margin-bottom: 15px;
}

.color-input-wrapper {
  display: flex;
  align-items: center;
  gap: 10px;
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
  font-size: 0.9em;
}

/* Form Styles */
form {
  display: flex;
  flex-direction: column;
  gap: 10px;
  min-width: 300px;
}

label {
  font-weight: 500;
  margin-top: 5px;
}

input {
  padding: 8px;
  border: 1px solid #ddd;
  border-radius: 4px;
}

button[type="submit"] {
  background: #4CAF50;
  color: white;
  padding: 10px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  margin-top: 10px;
}

button[type="submit"]:hover {
  background: #45a049;
}
</style>
