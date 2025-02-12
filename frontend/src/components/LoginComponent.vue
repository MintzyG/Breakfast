<template>
  <form @submit.prevent="submitForm">
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
      email: '',
      password: ''
    };
  },
  methods: {
    async submitForm() {
      try {
        const response = await this.$api.post('/login', {
          email: this.email,
          password: this.password
        });
        localStorage.setItem('authToken', response.data.token);
        const token = response.data.token;
        console.log('Raw Token:', token);
        console.log('Type:', typeof token);
        const authToken = localStorage.getItem('authToken');
        console.log('Retrieved Token:', authToken);
        console.log('Type of Retrieved Token:', typeof authToken);
        this.$router.push({ name: 'dashboard' });
      } catch (error) {
        console.error('Error logging in:', error);
      }
    }
  }
};
</script>

