<script setup lang="ts">
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useAuthStore } from '@/stores/auth';

interface Booking {
  ID: number;
  court_id: number;
  start_time: string;
  end_time: string;
}

const bookings = ref<Booking[]>([]);
const isLoading = ref(true);
const authStore = useAuthStore();

const fetchMyBookings = async () => {
  try {
    const headers = { Authorization: `Bearer ${authStore.token}` };
    const response = await axios.get('http://localhost:8000/api/my-bookings', { headers });
    bookings.value = response.data.data;
  } catch (error) {
    console.error("Failed to fetch bookings:", error);
  } finally {
    isLoading.value = false;
  }
};

onMounted(() => {
  fetchMyBookings();
});
</script>

<template>
  <div class="my-bookings">
    <h2>Riwayat Booking Saya</h2>

    <div v-if="isLoading">
      <p aria-busy="true">Memuat riwayat booking...</p>
    </div>

    <div v-else-if="bookings.length > 0">
      <article v-for="booking in bookings" :key="booking.ID">
        <p><strong>Lapangan ID:</strong> {{ booking.court_id }}</p>
        <p><strong>Mulai:</strong> {{ new Date(booking.start_time).toLocaleString('id-ID') }}</p>
        <p><strong>Selesai:</strong> {{ new Date(booking.end_time).toLocaleString('id-ID') }}</p>
      </article>
    </div>
    <div v-else>
      <p>Anda belum memiliki riwayat booking.</p>
    </div>
  </div>
</template>

<style scoped>
/* Style tidak berubah */
</style>