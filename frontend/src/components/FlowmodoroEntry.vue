<template>
  <div class="session-entry">
    <div class="session-content">
      <div class="session-left" :style="{ background: '#' + session.color || '#4CAF50' }">
        <span class="session-emoji">{{ session.emoji }}</span>
      </div>
      <div class="session-right">
        <div class="session-header">
          <span class="session-name">{{ session.session_name }}</span>
          <div class="session-actions">
            <span class="session-duration">{{ formatDuration(session.duration) }}</span>
            <button class="action-btn edit-btn" @click="$emit('edit-session', session)">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M17 3a2.828 2.828 0 1 1 4 4L7.5 20.5 2 22l1.5-5.5L17 3z"></path>
              </svg>
            </button>
            <button class="action-btn delete-btn" @click="$emit('delete-session', session)">
              🗑️
            </button>
          </div>
        </div>
        <div class="session-stats">
          <span class="stat">
            <span class="stat-icon">⏸️</span>
            {{ formatDuration(session.break_time) }}
          </span>
          <span class="stat">
            <span class="stat-icon">⚡</span>
            {{ session.distractions }}
          </span>
        </div>
        <div class="session-time">
          <span class="stat-icon">📅</span>
          {{ formatSessionTime(session.focus_start, session.focus_end) }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'FlowmodoroEntry',
  props: {
    session: {
      type: Object,
      required: true,
      validator: (obj) => {
        return [
          'emoji',
          'session_name',
          'duration',
          'break_time',
          'distractions',
          'focus_start',
          'break_factor'
        ].every(prop => prop in obj);
      }
    }
  },
  emits: ['delete-session', 'edit-session'],
  methods: {
    formatDuration(seconds) {
      const hours = Math.floor(seconds / 3600);
      const minutes = Math.floor((seconds % 3600) / 60);
      const remainingSeconds = seconds % 60;
      return hours > 0 
        ? `${hours}h:${minutes.toString().padStart(2, '0')}m:${remainingSeconds.toString().padStart(2, '0')}s`
        : `${minutes}m:${remainingSeconds.toString().padStart(2, '0')}s`;
    },
    formatSessionTime(start, end) {
      const startDate = new Date(start);
      const endDate = new Date(end);

      const formatDate = (date) => {
        return `${date.getDate()}/${date.getMonth() + 1}/${date.getFullYear()}`;
      };

      const formatTime = (date) => {
        return date.toLocaleTimeString('en-US', { 
          hour12: false,
          hour: '2-digit',
          minute: '2-digit',
          second: '2-digit'
        });
      };

      const startDateStr = formatDate(startDate);
      const endDateStr = formatDate(endDate);

      return startDateStr === endDateStr
        ? `${startDateStr} ${formatTime(startDate)} - ${formatTime(endDate)}`
        : `${startDateStr} ${formatTime(startDate)} - ${endDateStr} ${formatTime(endDate)}`;
    }
  }
};
</script>

<style scoped>
.session-entry {
  background: white;
  border-radius: 8px;
  margin-bottom: 1rem;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
  transition: transform 0.2s, box-shadow 0.2s;
}

.session-entry:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.session-content {
  display: grid;
  grid-template-columns: 120px 1fr;
  min-height: 100px;
  border-radius: 8px;
  overflow: hidden;
}

.session-left {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 1rem;
}

.session-emoji {
  font-size: 2rem;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 3rem;
  height: 3rem;
  background: rgba(255, 255, 255, 0.9);
  border-radius: 8px;
}

.session-right {
  padding: 1rem;
  background: white;
}

.session-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 0.5rem;
}

.session-name {
  font-weight: 600;
  color: #2c3e50;
}

.session-duration {
  color: #666;
  font-size: 0.9rem;
}

.session-stats {
  display: flex;
  gap: 1rem;
  align-items: center;
  color: #666;
  font-size: 0.9rem;
}

.stat {
  display: flex;
  align-items: center;
  gap: 0.25rem;
}

.stat-icon {
  font-size: 1rem;
}

.session-time {
  margin-top: 0.5rem;
  color: #666;
  font-size: 0.9rem;
  display: flex;
  align-items: center;
  gap: 0.25rem;
}

.session-actions {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.action-btn {
  background: none;
  border: none;
  padding: 0.25rem;
  cursor: pointer;
  opacity: 0.6;
  transition: opacity 0.2s, background-color 0.2s;
  border-radius: 4px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.action-btn:hover {
  opacity: 1;
}

.edit-btn:hover {
  background: #e5e7eb;
}

.delete-btn:hover {
  background: #fee2e2;
}

@media (max-width: 640px) {
  .session-content {
    grid-template-columns: 80px 1fr;
  }
  
  .session-stats {
    flex-wrap: wrap;
    gap: 0.5rem;
  }
  
  .session-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 0.25rem;
  }
}
</style>
