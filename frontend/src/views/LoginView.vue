<script setup lang="ts">
import { ref } from 'vue'
import { useAuthStore } from '@/stores/auth'
import { useRouter, RouterLink } from 'vue-router'
import { useToast } from 'vue-toastification'

const email = ref('')
const password = ref('')
const authStore = useAuthStore()
const router = useRouter()
const toast = useToast()

const handleLogin = async () => {
  const success = await authStore.login(email.value, password.value)
  if (success) {
    toast.success('Login Berhasil!')
    router.push('/')
  } else {
    toast.error('Email atau Password salah!')
  }
}
</script>

<template>
  <v-container class="fill-height">
    <v-row align="center" justify="center">
      <v-col cols="12" sm="8" md="4">
        <v-sheet elevation="12" rounded="lg" class="pa-8">
          <h2 class="text-center mb-6">Login</h2>
          <v-form @submit.prevent="handleLogin">
            <v-text-field
              v-model="email"
              label="Email"
              type="email"
              prepend-inner-icon="mdi-email-outline"
              required
              class="mb-4"
            ></v-text-field>

            <v-text-field
              v-model="password"
              label="Password"
              type="password"
              prepend-inner-icon="mdi-lock-outline"
              required
            ></v-text-field>
            
            <div class="text-right mb-4">
              <RouterLink to="/forgot-password" class="text-caption">Lupa Password?</RouterLink>
            </div>

            <v-btn type="submit" block color="primary" size="large">Login</v-btn>
          </v-form>
          <p class="text-center mt-4">
            Belum punya akun? <RouterLink to="/register">Daftar di sini</RouterLink>
          </p>
        </v-sheet>
      </v-col>
    </v-row>
  </v-container>
</template>

<style scoped>
a {
  text-decoration: none;
  color: #1a73e8; /* Warna link standar Vuetify */
}
a:hover {
  text-decoration: underline;
}
</style>