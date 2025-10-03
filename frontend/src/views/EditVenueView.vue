<script setup lang="ts">
import { ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useToast } from 'vue-toastification';
import api from '@/services/apiService';

const route = useRoute();
const router = useRouter();
const toast = useToast();
const selectedFile = ref<File[]>([]);
const isLoading = ref(false);

const handleImageUpload = async () => {
  if (selectedFile.value.length === 0) {
    toast.error('Silakan pilih sebuah gambar terlebih dahulu.');
    return;
  }

  isLoading.value = true;
  const venueId = route.params.id;
  const formData = new FormData();
  // PERBAIKAN DI SINI
  formData.append('image', selectedFile.value[0]); 

  try {
    await api.post(`/venues/${venueId}/upload-image`, formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });

    toast.success('Gambar berhasil di-upload!');
    router.push({ name: 'venue-detail', params: { id: venueId } });
  } catch (error) {
    toast.error('Gagal meng-upload gambar.');
    console.error(error);
  } finally {
    isLoading.value = false;
  }
};
</script>

<template>
  <v-container>
    <v-sheet max-width="600" class="mx-auto pa-8" elevation="12" rounded="lg">
      <h2 class="text-h4 mb-6">Upload Gambar GOR</h2>
      <v-form @submit.prevent="handleImageUpload">
        <v-file-input
          v-model="selectedFile"
          label="Pilih Gambar GOR"
          accept="image/*"
          prepend-icon="mdi-camera"
          show-size
        ></v-file-input>

        <v-btn 
          type="submit" 
          block 
          color="primary" 
          size="large" 
          class="mt-4"
          :loading="isLoading"
          :disabled="isLoading"
        >
          Upload
        </v-btn>
      </v-form>
    </v-sheet>
  </v-container>
</template>