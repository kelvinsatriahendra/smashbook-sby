<script setup lang="ts">
import { ref } from 'vue';
import { useAuthStore } from '@/stores/auth';
import axios from 'axios';
import { useToast } from 'vue-toastification';
import { useRouter } from 'vue-router';

// --- State untuk menampung data dari form ---
const name = ref('');
const address = ref('');
const description = ref('');
const imageFile = ref<File[]>([]); // State untuk menampung file gambar

// --- Inisialisasi hooks ---
const authStore = useAuthStore();
const toast = useToast();
const router = useRouter();

// --- Fungsi untuk menangani submit form ---
const handleAddVenue = async () => {
  if (imageFile.value.length === 0) {
    toast.error("Silakan pilih gambar untuk GOR.");
    return;
  }

  // Kita gunakan FormData untuk mengirim file dan teks bersamaan
  const formData = new FormData();
  formData.append('name', name.value);
  formData.append('address', address.value);
  formData.append('description', description.value);
  formData.append('image', imageFile.value[0]); // Ambil file pertama

  try {
    const headers = {
      'Content-Type': 'multipart/form-data',
      Authorization: `Bearer ${authStore.token}`,
    };

    // Kita akan buat endpoint baru ini di backend nanti
    await axios.post('http://localhost:8000/api/venues', formData, { headers });

    toast.success('GOR baru berhasil ditambahkan!');
    router.push('/'); // Kembali ke halaman utama setelah sukses
  } catch (error) {
    toast.error('Gagal menambahkan GOR.');
    console.error(error);
  }
};
</script>

<template>
  <v-container>
    <v-sheet max-width="600" class="mx-auto pa-8" elevation="12" rounded="lg">
      <h2>Tambah GOR Baru</h2>
      <v-form @submit.prevent="handleAddVenue">
        <v-text-field
          v-model="name"
          label="Nama GOR"
          prepend-inner-icon="mdi-stadium-variant"
          required
          class="mb-4"
        ></v-text-field>

        <v-text-field
          v-model="address"
          label="Alamat"
          prepend-inner-icon="mdi-map-marker"
          required
          class="mb-4"
        ></v-text-field>

        <v-textarea
          v-model="description"
          label="Deskripsi"
          prepend-inner-icon="mdi-information-outline"
          required
          class="mb-4"
        ></v-textarea>

        <v-file-input
          v-model="imageFile"
          label="Foto GOR"
          accept="image/*"
          prepend-icon="mdi-camera"
          required
          class="mb-4"
        ></v-file-input>

        <v-btn type="submit" color="primary" block>Tambah GOR</v-btn>
      </v-form>
    </v-sheet>
  </v-container>
</template>

<style scoped>
/* Anda bisa menambahkan styling custom di sini jika perlu */
h2 {
  margin-bottom: 1.5rem;
  text-align: center;
}
</style>