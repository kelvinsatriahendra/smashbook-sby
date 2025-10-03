<script setup lang="ts">
import { ref } from 'vue';
import { useToast } from 'vue-toastification';
import { RouterLink } from 'vue-router';

const email = ref('');
const toast = useToast();

const handleForgotPassword = async () => {
  if (!email.value) {
    toast.error('Silakan masukkan alamat email Anda.');
    return;
  }

  // NOTE: Logika untuk mengirim email akan kita buat di backend nanti.
  // Untuk sekarang, kita hanya akan menampilkan pesan sukses.
  toast.success(`Jika email ${email.value} terdaftar, link reset akan dikirimkan.`);
  
  // Di aplikasi nyata, di sinilah kita akan memanggil API backend:
  // await api.post('/forgot-password', { email: email.value });
};
</script>

<template>
  <v-container class="fill-height">
    <v-row align="center" justify="center">
      <v-col cols="12" sm="8" md="4">
        <v-sheet elevation="12" rounded="lg" class="pa-8">
          <h2 class="text-center mb-4">Lupa Password</h2>
          <p class="text-center mb-6">
            Masukkan alamat email Anda, dan kami akan mengirimkan link untuk mereset password Anda.
          </p>
          <v-form @submit.prevent="handleForgotPassword">
            <v-text-field
              v-model="email"
              label="Email"
              type="email"
              prepend-inner-icon="mdi-email-outline"
              required
              class="mb-4"
            ></v-text-field>

            <v-btn type="submit" block color="primary" size="large">Kirim Link Reset</v-btn>
          </v-form>
           <p class="text-center mt-4">
            Kembali ke <RouterLink to="/login">Login</RouterLink>
          </p>
        </v-sheet>
      </v-col>
    </v-row>
  </v-container>
</template>