<template>
  <aside class="sidebar">
    <nav>
      <ul>
        <li><router-link to="/dashboard">Dashboard</router-link></li>
        <li><router-link to="/dashboard/profile">Profile</router-link></li>
        <li><router-link to="/dashboard/settings">Settings</router-link></li>
        <li><router-link to="/dashboard/pancake">Pancake</router-link></li>
        <li><router-link to="/dashboard/yogurt">Yogurt</router-link></li>
      </ul>
    </nav>
  </aside>
</template>

<script>
export default {
  name: 'Sidebar',
};
</script>

<style scoped>
.sidebar {
  width: 200px;
  height: calc(100vh - 60px); /* Ensures sidebar stays below the header */
  background-color: #2c3e50;
  position: fixed;
  top: 60px; /* Positions below the header */
  left: 0;
  padding-top: 20px; /* Pushes the content down */
}

.sidebar nav ul {
  list-style: none;
  padding: 0;
  margin: 0;
  padding-top: 15px; /* Extra space for the buttons */
}

.sidebar nav ul li {
  margin-bottom: 10px;
}

.sidebar nav ul li a {
  display: block;
  padding: 12px 16px;
  background-color: #34495e;
  color: #ecf0f1;
  text-decoration: none;
  font-size: 16px;
  border-radius: 4px;
  transition: background 0.2s ease-in-out;
}

.sidebar nav ul li a:hover {
  background-color: #3b5877;
}
</style>

