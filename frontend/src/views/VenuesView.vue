<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { RouterLink } from 'vue-router'
import api from '@/services/apiService';
import { useAuthStore } from '@/stores/auth'

interface Venue {
  ID: number;
  name: string;
  address: string;
  description: string;
  image_url: string;
}

const venues = ref<Venue[]>([])
const isLoading = ref(true)
const error = ref<string | null>(null)
const authStore = useAuthStore()

const fetchVenues = async () => {
  try {
    const response = await api.get('/venues')
    venues.value = response.data.data
  } catch (err) {
    error.value = 'Gagal memuat data. Pastikan server backend berjalan.'
    console.error(err)
  } finally {
    isLoading.value = false
  }
}

onMounted(() => {
  fetchVenues()
})
</script>

<template>
  <div>
    <div class="d-flex justify-space-between align-center mb-6">
      <h2 class="text-h4">Daftar GOR Badminton</h2>
      <v-btn
        v-if="authStore.isAuthenticated"
        to="/add-venue"
        color="primary"
        prepend-icon="mdi-plus-circle-outline"
      >
        Tambah GOR
      </v-btn>
    </div>

    <div v-if="isLoading">
      <v-progress-linear indeterminate color="primary"></v-progress-linear>
      <p class="text-center mt-4">Memuat data GOR...</p>
    </div>
    
    <v-alert type="error" v-else-if="error" class="mb-4">
      {{ error }}
    </v-alert>

    <v-row v-else-if="venues.length > 0">
      <v-col
        v-for="venue in venues"
        :key="venue.ID"
        cols="12"
        sm="6"
        md="4"
      >
        <v-card
          class="mx-auto"
          hover
          :to="{ name: 'venue-detail', params: { id: venue.ID } }"
        >
          <v-img
            v-if="venue.image_url"
            :src="venue.image_url"
            height="200px"
            cover
            class="align-end"
          >
            <v-card-title class="text-white">{{ venue.name }}</v-card-title>
          </v-img>
          
          <v-card-item v-else>
            <v-card-title>{{ venue.name }}</v-card-title>
          </v-card-item>

          <v-card-subtitle class="pt-4 pb-2">{{ venue.address }}</v-card-subtitle>
        </v-card>
      </v-col>
    </v-row>

    <v-alert type="info" v-else>
      Belum ada data GOR yang tersedia.
    </v-alert>
  </div>
</template>