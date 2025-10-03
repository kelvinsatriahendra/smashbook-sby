import { defineStore } from 'pinia'
import axios from 'axios'
import { ref } from 'vue'

// Definisikan interface untuk User
interface User {
  id: number;
  name: string;
  email: string;
  role: string;
}

export const useAuthStore = defineStore('auth', () => {
  const token = ref(localStorage.getItem('token') || null)
  // Simpan seluruh objek user, bukan null lagi
  const user = ref<User | null>(JSON.parse(localStorage.getItem('user') || 'null'))

  const isAuthenticated = ref(!!token.value)
  // Buat computed property untuk mengecek role
  const isOwner = ref(user.value?.role === 'venue_owner')

  async function login(email: string, password: string) {
    try {
      const response = await axios.post('http://localhost:8000/api/login', {
        email: email,
        password: password
      });

      const responseToken = response.data.data.token
      const responseUser = response.data.data.user

      token.value = responseToken
      user.value = responseUser
      isAuthenticated.value = true
      isOwner.value = responseUser.role === 'venue_owner'

      // Simpan di localStorage
      localStorage.setItem('token', responseToken)
      localStorage.setItem('user', JSON.stringify(responseUser))

      return true
    } catch (error) {
      console.error('Login failed:', error)
      return false
    }
  }

  function logout() {
    token.value = null
    user.value = null
    isAuthenticated.value = false
    isOwner.value = false
    
    // Hapus dari localStorage
    localStorage.removeItem('token')
    localStorage.removeItem('user')

    window.location.href = '/login';
  }

  return { token, user, isAuthenticated, isOwner, login, logout }
})