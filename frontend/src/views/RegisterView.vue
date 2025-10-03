<script setup lang="ts">
import { ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'
import { useToast } from 'vue-toastification'

const name = ref('')
const email = ref('')
const password = ref('')
const router = useRouter()
const toast = useToast()

const handleRegister = async () => {
  try {
    await axios.post('http://localhost:8000/api/register', {
      name: name.value,
      email: email.value,
      password: password.value
    });

    toast.success('Registrasi Berhasil! Silakan login.');
    router.push('/login');

  } catch (error) {
    toast.error('Registrasi Gagal. Email mungkin sudah digunakan.');
    console.error('Registration failed:', error);
  }
}
</script>

<template>
  <v-container class="fill-height">
    <v-row align="center" justify="center">
      <v-col cols="12" sm="8" md="4">
        <v-sheet elevation="12" rounded="lg" class="pa-8">
          <h2 class="text-center mb-6">Register</h2>
          <v-form @submit.prevent="handleRegister">
            <v-text-field
              v-model="name"
              label="Nama Lengkap"
              prepend-inner-icon="mdi-account-outline"
              required
              class="mb-4"
            ></v-text-field>

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
              class="mb-4"
            ></v-text-field>

            <v-btn type="submit" block color="primary" size="large">Register</v-btn>
          </v-form>
        </v-sheet>
      </v-col>
    </v-row>
  </v-container>
</template>

<style scoped>
/* Style lama bisa dihapus */
</style>