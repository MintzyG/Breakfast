<template>
  <div class="dashboard">
    <Header />
    <div class="main-area">
      <Sidebar />
      <div class="content">
        <router-view />
      </div>
    </div>
  </div>
</template>

<script>
import Sidebar from './SidebarComponent.vue';
import Header from './HeaderComponent.vue';

export default {
  name: 'Dashboard',
  components: {
    Sidebar,
    Header,
  },
};
</script>

<style scoped>
.dashboard {
  display: flex;
  flex-direction: column;
  height: 100vh;
}

.main-area {
  display: flex;
  flex: 1;
  margin-top: 60px;
}

.content {
  flex: 1;
  padding: 20px;
  margin-left: 200px;
  overflow-y: auto;
}
</style>

