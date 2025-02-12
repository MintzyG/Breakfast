<template>
  <form @submit.prevent="submitForm">
    <div>
      <label for="name">Name</label>
      <input id="name" v-model="name" type="text" />
    </div>
    <div>
      <label for="email">Email</label>
      <input id="email" v-model="email" type="email" />
    </div>
    <div>
      <label for="password">Password</label>
      <input id="password" v-model="password" type="password" />
    </div>
    <button type="submit">Login</button>
  </form>
</template>

<script>
export default {
  data() {
    return {
      name: '',
      email: '',
      password: ''
    };
  },
  methods: {
    async submitForm() {
      try {
        const response = await this.$api.post('/register', {
          name: this.name,
          email: this.email,
          password: this.password
        });
        localStorage.setItem('authToken', response.data);
        this.$router.push({ name: 'dashboard' });
      } catch (error) {
        console.error('Error logging in:', error);
      }
    }
  }
};
</script>

