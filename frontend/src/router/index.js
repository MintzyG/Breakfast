import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '../views/LoginView.vue'
import SignupView from '../views/SignupView.vue'
import DashboardView from '../views/DashboardView.vue'
import DashboardHome from '../views/DashboardHome.vue'
import PancakeView from '../views/PancakeView.vue'
import YogurtView from '../views/YogurtView.vue'
import EspressoView from '../views/EspressoView.vue'
import Profile from '../views/Profile.vue'
import Settings from '../views/Settings.vue'

import axios from 'axios';
import * as jwtDecode from 'jwt-decode';
const jwtSecret = import.meta.env.VITE_JWT_SECRET;

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/signup',
      name: 'signup',
      component: SignupView,
      meta: { freeAccess: true },
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView,
      meta: { freeAccess: true },
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: DashboardView,
      children: [
        {
          path: '',
          name: 'dashboard-home',
          component: DashboardHome,
        },
        {
          path: 'pancake',
          name: 'pancake',
          component: PancakeView,
        },
        {
          path: 'yogurt',
          name: 'yogurt',
          component: YogurtView,
        },
        {
          path: 'espresso',
          name: 'espresso',
          component: EspressoView,
        },
        {
          path: 'profile',
          name: 'profile',
          component: Profile,
        },
        {
          path: 'settings',
          name: 'settings',
          component: Settings,
        },
      ],
    },
  ],
})

router.beforeEach((to, from, next) => {
  const authToken = localStorage.getItem('authToken');

  if (!to.meta.freeAccess && !authToken) {
    next({ name: 'login' });
  } else if (!to.meta.freeAccess && authToken) {
    verifyJWT()
      .then(() => {
        next();
      })
      .catch(() => {
        next({ name: 'login' });
      });
  } else {
    next();
  }
});

async function verifyJWT() {
  const authToken = localStorage.getItem('authToken');

  if (!authToken) {
    console.error('No auth token found');
    return;
  }

  try {
    const response = await axios.post('http://localhost:8080/verify-jwt', {}, {
      headers: {
        'Authorization': `Bearer ${authToken}`,
      },
    });

    if (response.status === 200) {
      console.log('JWT is valid');
    } else {
      console.error('JWT verification failed:', response.data);
      this.$router.push({ name: 'login' });
    }
  } catch (error) {
    if (error.response && error.response.status === 401) {
      console.error('Invalid JWT: ', response.data);
      this.$router.push({ name: 'login' });
    } else {
      console.error('Error verifying JWT:', error);
      this.$router.push({ name: 'login' });
    }
  }
}

export default router
