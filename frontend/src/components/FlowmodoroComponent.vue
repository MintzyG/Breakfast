<template>
  <div class="timer-card">
    <div class="settings-section" v-if="!activeSession">
      <div class="settings-header" :style="{ background: `#${selectedColor}` }">
        <div class="emoji-section">
          <EmojiPicker
            v-model="selectedEmoji"
            @updateEmoji="updateEmoji"
          />
        </div>
        <div class="color-picker">
          <ColorPicker
            v-model="selectedColor"
            @updatecolor="updateColor"
          />
        </div>
      </div>
    </div>

    <!-- Timer Display -->
    <div class="timer-display">
      <h1>{{ isBreak ? formatTime(breakTimeRemaining) : formatTime(timeElapsed) }}</h1>
      <div class="timer-status" v-if="activeSession">
        {{ isBreak ? 'Break Time' : 'Focus Time' }}
        <span v-if="!isBreak" class="break-available">
          Break available: {{ formatTime(availableBreakTime) }}
        </span>
      </div>
    </div>

    <!-- Controls -->
    <div class="timer-controls">
      <button 
        class="control-btn"
        :class="{ 'active': activeSession && !isBreak }"
        @click="toggleTimer"
        v-if="!isBreak"
      >
        {{ !activeSession ? 'Start' : 'Pause' }}
      </button>

      <button 
        v-if="activeSession && !isBreak"
        class="control-btn break-btn"
        @click="startBreak"
      >
        Break
      </button>

      <button 
        v-if="isBreak"
        class="control-btn resume-btn"
        @click="resumeFromBreak"
      >
        Skip Break
      </button>

      <button 
        v-if="activeSession"
        class="control-btn distraction-btn"
        @click="addDistraction"
      >
        Distraction ({{ distractions }})
      </button>

      <button 
        v-if="activeSession"
        class="control-btn stop-btn"
        @click="stopSession"
      >
        Stop
      </button>
    </div>
  </div>
</template>

<script>
import ColorPicker from './ColorPicker.vue';
import EmojiPicker from './EmojiPicker.vue';
export default {
  name: 'FlowmodoroComponent',
  components: { ColorPicker, EmojiPicker },
  props: {
    initialSettings: {
      type: Object,
      default: () => ({})
    }
  },
  emits: ['session-completed'],
  data() {
    const savedSession = localStorage.getItem('currentSession');
    const defaultData = {
      activeSession: false,
      isBreak: false,
      timeElapsed: 0,
      breakTime: 0,
      breakTimeRemaining: 0,
      breakTimeUsed: 0,
      distractions: 0,
      selectedEmoji: '☕',
      selectedColor: '4CAF50',
      showEmojiPicker: false,
      emojiList: ['☕', '📚', '💻', '🎨', '✍️', '🎵', '💪', '🧘', '🎯', '⚡'],
      sessionStartTime: null,
      breakStartTime: null
    };

    return savedSession ? { ...defaultData, ...JSON.parse(savedSession) } : defaultData;
  },
  computed: {
    availableBreakTime() {
      // Ensure we're working with numbers and have valid values
      const elapsed = Number(this.timeElapsed) || 0;
      const factor = Number(this.breakFactor) || 5;
      const used = Number(this.breakTimeUsed) || 0;

      // Calculate total earned break time
      const totalEarned = Math.floor(elapsed / factor);

      // Return the difference (with minimum of 0)
      return Math.max(0, totalEarned - used);
    }
  },
  watch: {
    initialSettings: {
      handler(newSettings) {
        if (newSettings.emoji) this.selectedEmoji = newSettings.emoji;
        if (newSettings.color) this.selectedColor = newSettings.color;
        if (newSettings.breakFactor) this.breakFactor = newSettings.breakFactor;
      },
      immediate: true
    },
    timeElapsed() {
      this.saveToLocalStorage();
    },
    breakTimeUsed() {
      this.saveToLocalStorage();
    },
    distractions() {
      this.saveToLocalStorage();
    }
  },
  beforeUnmount() {
    this.clearTimer();
  },
  methods: {
    saveToLocalStorage() {
      if (this.activeSession) {
        const sessionData = {
          activeSession: this.activeSession,
          isBreak: this.isBreak,
          timeElapsed: this.timeElapsed,
          breakTime: this.breakTime,
          breakTimeRemaining: this.breakTimeRemaining,
          breakTimeUsed: this.breakTimeUsed,
          distractions: this.distractions,
          selectedEmoji: this.selectedEmoji,
          selectedColor: this.selectedColor,
          sessionStartTime: this.sessionStartTime,
          breakStartTime: this.breakStartTime
        };
        localStorage.setItem('currentSession', JSON.stringify(sessionData));
      } else {
        localStorage.removeItem('currentSession');
      }
    },
    toggleTimer() {
      if (!this.activeSession) {
        this.startSession();
      } else {
        this.pauseSession();
      }
    },
    startSession() {
      this.activeSession = true;
      this.sessionStartTime = new Date();
      this.startTimer();
      this.saveToLocalStorage();
    },
    startTimer() {
      this.timerInterval = setInterval(() => {
        if (this.isBreak) {
          this.breakTimeRemaining--;
          if (this.breakTimeRemaining <= 0) {
            this.resumeFromBreak();
          }
        } else {
          this.timeElapsed++;
        }
      }, 1000);
    },
    pauseSession() {
      this.clearTimer();
      this.activeSession = false;
    },
    startBreak() {
      if (this.availableBreakTime <= 0) return;

      this.isBreak = true;
      this.breakStartTime = new Date();
      this.breakTimeRemaining = this.availableBreakTime;
      this.clearTimer();
      this.startTimer();
    },

    resumeFromBreak() {
      this.isBreak = false;
      if (this.breakStartTime) {
        const breakDuration = Math.floor((new Date() - this.breakStartTime) / 1000);
        this.breakTime += breakDuration;
        this.breakTimeUsed += breakDuration; // Track used break time
        this.breakStartTime = null;
      }
      this.clearTimer();
      this.startTimer();
    },
    stopSession() {
      if (this.isBreak) {
        this.resumeFromBreak();
      }
      this.clearTimer();

      const session = {
        emoji: this.selectedEmoji,
        color: this.selectedColor,
        session_name: `Focus Session`,
        duration: this.timeElapsed,
        break_time: this.breakTimeUsed,
        distractions: this.distractions,
        focus_start: this.sessionStartTime,
        focus_end: new Date()
      };

      this.$emit('session-completed', session);
      this.resetSession();
      localStorage.removeItem('currentSession');
    },
    clearTimer() {
      if (this.timerInterval) {
        clearInterval(this.timerInterval);
        this.timerInterval = null;
      }
    },
    resetSession() {
      this.activeSession = false;
      this.isBreak = false;
      this.timeElapsed = 0;
      this.breakTime = 0;
      this.breakTimeRemaining = 0;
      this.breakTimeUsed = 0; // Reset used break time
      this.distractions = 0;
      this.sessionStartTime = null;
      this.breakStartTime = null;
      this.selectedEmoji = '☕';
      this.selectedColor = '#4CAF50';
      this.breakFactor = 5;
    },
    addDistraction() {
      this.distractions++;
    },
    updateEmoji(emoji) {
      this.selectedEmoji = emoji;
    },
    updateColor(color) {
      this.selectedColor = color;
    },    formatTime(seconds) {
      if (!seconds || isNaN(seconds)) {
        return '0:00';
      }

      const hours = Math.floor(seconds / 3600);
      const minutes = Math.floor((seconds % 3600) / 60);
      const remainingSeconds = Math.floor(seconds % 60);

      if (hours > 0) {
        return `${hours}:${minutes.toString().padStart(2, '0')}:${remainingSeconds.toString().padStart(2, '0')}`;
      }
      return `${minutes}:${remainingSeconds.toString().padStart(2, '0')}`;
    }
  }
};
</script>

<style scoped>
.timer-card {
  background: white;
  border-radius: 12px;
  padding: 2rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  max-width: 400px;
  margin: 0 auto;
}

.emoji-section {
  text-align: center;
  margin-bottom: 1.5rem;
}

.emoji-picker {
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  gap: 0.5rem;
  margin-top: 1rem;
}

.emoji-btn {
  font-size: 1.5rem;
  padding: 0.5rem;
  border: none;
  background: none;
  cursor: pointer;
  border-radius: 4px;
  transition: background-color 0.2s;
}

.emoji-btn:hover {
  background-color: #f0f0f0;
}

.timer-display {
  text-align: center;
  margin-bottom: 2rem;
}

.timer-display h1 {
  font-size: 3.5rem;
  font-weight: bold;
  margin-bottom: 0.5rem;
}

.timer-status {
  color: #666;
  font-size: 1.1rem;
}

.break-available {
  font-size: 0.9rem;
  margin-left: 0.5rem;
}

.timer-controls {
  display: flex;
  flex-wrap: wrap;
  gap: 0.75rem;
  justify-content: center;
}

.control-btn {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 9999px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
  min-width: 120px;
}

.control-btn:not(.active) {
  background-color: #4CAF50;
  color: white;
}

.control-btn:not(.active):hover {
  background-color: #45a049;
}

.control-btn.active {
  background-color: #f44336;
  color: white;
}

.control-btn.active:hover {
  background-color: #da190b;
}

.break-btn {
  background-color: #2196F3;
  color: white;
}

.break-btn:hover {
  background-color: #1976D2;
}

.resume-btn {
  background-color: #4CAF50;
  color: white;
}

.resume-btn:hover {
  background-color: #45a049;
}

.distraction-btn {
  background-color: #FF9800;
  color: white;
}

.distraction-btn:hover {
  background-color: #F57C00;
}

.stop-btn {
  background-color: #f44336;
  color: white;
}

.stop-btn:hover {
  background-color: #da190b;
}

.settings-section {
  margin-bottom: 1.5rem;
}

.settings-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1.5rem;
  border-radius: 8px 8px 0 0;
  margin: -2rem -2rem 1rem -2rem;
}

.emoji-section {
  position: relative;
  margin-bottom: 0;
}

.emoji-display {
  font-size: 2.5rem;
  cursor: pointer;
  user-select: none;
  background: rgba(255, 255, 255, 0.9);
  padding: 0.5rem;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.settings-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.break-factor-setting {
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-bottom: 1rem;
}

.break-factor-input {
  width: 60px;
  padding: 0.5rem;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.color-picker {
  display: flex;
  align-items: center;
}

.color-input {
  width: 40px;
  height: 40px;
  padding: 0;
  border: 3px solid white;
  border-radius: 4px;
  cursor: pointer;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}
</style>
