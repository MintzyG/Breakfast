<template>
  <div class="task">
    <div class="task-content" :style="{ '--task-color': '#' + task.color }">
      <div class="task-left">
        <div class="custom-checkbox">
          <div 
            class="checkbox-circle" 
            :class="{ 'checked': task.completed }"
            @click="toggleCompleted"
          >
            <svg 
              v-if="task.completed"
              xmlns="http://www.w3.org/2000/svg" 
              viewBox="0 0 24 24" 
              fill="none" 
              stroke="white" 
              stroke-width="2" 
              stroke-linecap="round" 
              stroke-linejoin="round"
              class="checkmark"
            >
              <polyline points="20 6 9 17 4 12"></polyline>
            </svg>
          </div>
        </div>
      </div>
      <div class="task-right">
        <div class="title">
          <span class="task-emoji">{{ task.emoji }}</span>
          {{ task.title }}
        </div>
        <div class="stats">
          <span class="stat" :class="getStatClass(task.priority)">P{{ task.priority }}</span>
          <span class="stat" :class="getStatClass(task.task_size)">S{{ task.task_size }}</span>
          <span class="stat" :class="getStatClass(task.difficulty)">D{{ task.difficulty }}</span>
          <button class="action-button" @click.stop="$emit('open-edit-modal', task)">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M17 3a2.828 2.828 0 1 1 4 4L7.5 20.5 2 22l1.5-5.5L17 3z"></path>
            </svg>
          </button>
          <button class="action-button" @click.stop="showDeleteConfirm = true">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M3 6h18"></path>
              <path d="M19 6v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6"></path>
              <path d="M8 6V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2"></path>
            </svg>
          </button>
        </div>
      </div>
    </div>

    <!-- Delete Confirmation Modal -->
    <ModalComponent v-if="showDeleteConfirm" @close="showDeleteConfirm = false">
      <template v-slot:default>
        <h3>Are you sure you want to delete this task?</h3>
        <div class="confirm-buttons">
          <button @click="confirmDelete" class="confirm-delete">Yes, Delete</button>
          <button @click="showDeleteConfirm = false">Cancel</button>
        </div>
      </template>
    </ModalComponent>
  </div>
</template>

<script>
import ModalComponent from './ModalComponent.vue';

export default {
  name: 'TaskComponent',
  components: { ModalComponent },
  props: {
    task: Object,
  },
  data() {
    return {
      showDeleteConfirm: false,
    };
  },
  methods: {
    getStatClass(value) {
      switch (value) {
        case 3: return 'stat-red';
        case 2: return 'stat-yellow';
        case 1: return 'stat-green';
        default: return '';
      }
    },
    async toggleCompleted() {
      try {
        const updatedTask = {
          ...this.task,
          completed: !this.task.completed
        };
        await this.$api.patch(`/yogurt/${this.task.task_id}`, updatedTask, {
          headers: { Authorization: `Bearer ${localStorage.getItem('authToken')}` }
        });
        this.$emit('task-updated', updatedTask);
      } catch (error) {
        console.error('Error updating task:', error);
        this.task.completed = !this.task.completed;
      }
    },
    async confirmDelete() {
      try {
        await this.$api.delete(`/yogurt/${this.task.task_id}`, {
          headers: { Authorization: `Bearer ${localStorage.getItem('authToken')}` }
        });
        this.$emit('task-deleted', this.task.task_id);
        this.showDeleteConfirm = false;
      } catch (error) {
        console.error('Error deleting task:', error);
      }
    }
  }
};
</script>

<style scoped>
.task {
  width: 600px;
  height: 50px;
  background: #f9f9f9;
  border-radius: 10px;
  box-shadow: 2px 2px 5px rgba(0, 0, 0, 0.1);
  margin-bottom: 10px;
}

.task-content {
  display: flex;
  height: 100%;
  --task-color: #4CAF50;
}

.task-left {
  background-color: var(--task-color);
  padding: 0 15px;
  display: flex;
  align-items: center;
  border-radius: 10px 0 0 10px;
}

.task-right {
  flex: 1;
  display: flex;
  align-items: center;
  padding: 0 15px;
  justify-content: space-between;
}

.completed {
  display: flex;
  align-items: center;
}

.completed input[type="checkbox"] {
  width: 18px;
  height: 18px;
  cursor: pointer;
}

.title {
  flex: 1;
  display: flex;
  align-items: center;
  font-size: 1em;
  margin-right: 15px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.task-emoji {
  margin-right: 8px;
}

.stats {
  display: flex;
  gap: 10px;
  align-items: center;
}

.stat {
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 0.9em;
  font-weight: 500;
  color: white;
}

.stat-red { background: #ff4444; }
.stat-yellow { background: #ffbb33; color: #333; }
.stat-green { background: #00C851; }

.action-button {
  background: none;
  border: none;
  cursor: pointer;
  padding: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #666;
  border-radius: 50%;
  transition: all 0.2s ease;
}

.action-button:hover {
  background: #f0f0f0;
  color: #333;
}

.action-button:last-child:hover {
  color: #ff4444;
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

/* Modal Form Styles */
form {
  display: flex;
  flex-direction: column;
  gap: 10px;
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
  border: none;
  cursor: pointer;
}

.custom-checkbox {
  display: flex;
  align-items: center;
  justify-content: center;
}

.checkbox-circle {
  width: 20px;
  height: 20px;
  border: 2px solid white;
  border-radius: 50%;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: transparent;
  transition: background-color 0.2s ease;
}

.checkbox-circle.checked {
  background-color: var(--task-color);
  border-color: var(--task-color);
}

.checkmark {
  width: 14px;
  height: 14px;
  stroke: white;
}

</style>

