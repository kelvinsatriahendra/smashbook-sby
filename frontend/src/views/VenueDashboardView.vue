<script setup lang="ts">
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useAuthStore } from '@/stores/auth';
import { useToast } from 'vue-toastification';

interface User { name: string; }
interface Court { name: string; }
interface Booking {
  ID: number;
  court: Court;
  user: User;
  start_time: string;
  end_time: string;
}

// KODE YANG BENAR SEPERTI INI:
const bookings = ref<Booking[]>([]); 
const isLoading = ref(true);
const error = ref<string | null>(null);
const authStore = useAuthStore();
const toast = useToast();

const fetchOwnerBookings = async () => {
  try {
    const headers = { Authorization: `Bearer ${authStore.token}` };
    const response = await axios.get('http://localhost:8000/api/my-venue/bookings', { headers });
    bookings.value = response.data.data;
  } catch (err: any) {
    error.value = err.response?.data?.message || "Gagal mengambil data booking.";
    toast.error(error.value);
  } finally {
    isLoading.value = false;
  }
};

const formatDate = (dateString: string) => {
  const options: Intl.DateTimeFormatOptions = { day: '2-digit', month: 'short', year: 'numeric', hour: '2-digit', minute: '2-digit' };
  return new Date(dateString).toLocaleDateString('id-ID', options);
}

onMounted(() => {
  fetchOwnerBookings();
});
</script>

<template>
  <div>
    <h2 class="text-h4 mb-4">Dashboard GOR Saya</h2>
    <p class="text-subtitle-1 mb-6">Daftar semua booking yang masuk ke GOR Anda.</p>
    
    <div v-if="isLoading">
      <p aria-busy="true">Memuat data booking...</p>
    </div>
    <v-alert type="error" v-else-if="error && bookings.length === 0">
      {{ error }}
    </v-alert>
    <v-card v-else-if="bookings.length > 0">
      <v-table>
        <thead>
          <tr>
            <th class="text-left">Nama Pemesan</th>
            <th class="text-left">Nama Lapangan</th>
            <th class="text-left">Waktu Mulai</th>
            <th class="text-left">Waktu Selesai</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="booking in bookings" :key="booking.ID">
            <td>{{ booking.user.name }}</td>
            <td>{{ booking.court.name }}</td>
            <td>{{ formatDate(booking.start_time) }}</td>
            <td>{{ formatDate(booking.end_time) }}</td>
          </tr>
        </tbody>
      </v-table>
    </v-card>
    <v-alert type="info" v-else>
      <p>Belum ada booking yang masuk ke GOR Anda.</p>
    </v-alert>
  </div>
</template>