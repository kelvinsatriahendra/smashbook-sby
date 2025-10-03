import { createRouter, createWebHistory } from 'vue-router'
import VenuesView from '../views/VenuesView.vue'
import LoginView from '../views/LoginView.vue'
import RegisterView from '../views/RegisterView.vue'
import ForgotPasswordView from '../views/ForgotPasswordView.vue'
import ResetPasswordView from '../views/ResetPasswordView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/', name: 'venues', component: VenuesView },
    { path: '/venues/:id', name: 'venue-detail', component: () => import('../views/VenueDetailView.vue') },
    { path: '/login', name: 'login', component: LoginView },
    { path: '/register', name: 'register', component: RegisterView },
    { path: '/my-bookings', name: 'my-bookings', component: () => import('../views/MyBookingsView.vue') },
    { path: '/dashboard', name: 'dashboard', component: () => import('../views/VenueDashboardView.vue') },
    { path: '/forgot-password', name: 'forgot-password', component: ForgotPasswordView },
    { path: '/reset-password', name: 'reset-password', component: ResetPasswordView },
    // PERBAIKAN DI SINI
    { path: '/add-venue', name: 'add-venue', component: () => import('../views/View.vue') }, 
    { 
      path: '/venues/:id/edit', 
      name: 'edit-venue', 
      component: () => import('../views/EditVenueView.vue') 
    }
  ]
})

export default router