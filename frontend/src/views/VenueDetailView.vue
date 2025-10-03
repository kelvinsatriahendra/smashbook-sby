<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import axios from 'axios' // Kita masih butuh axios di sini
import { useAuthStore } from '@/stores/auth'
import { useToast } from 'vue-toastification'

interface Venue { ID: number; name: string; address: string; description: string; image_url: string; }
interface Court { ID: number; name: string; floor_type: string; }

const route = useRoute()
const authStore = useAuthStore()
const toast = useToast()
const venue = ref<Venue | null>(null)
const courts = ref<Court[]>([])
const isLoading = ref(true)
const error = ref<string | null>(null)

const bookingDate = ref(new Date())
const bookingTime = ref('')

const handleBooking = async (courtId: number) => {
  if (!authStore.isAuthenticated) {
    toast.warning('Anda harus login terlebih dahulu untuk membuat booking!');
    return;
  }
  if (!bookingDate.value || !bookingTime.value) {
    toast.error('Silakan pilih tanggal dan jam terlebih dahulu.');
    return;
  }
  try {
    const date = bookingDate.value.toISOString().split('T')[0];
    const startTime = new Date(`${date}T${bookingTime.value}:00.000Z`);
    const endTime = new Date(startTime.getTime() + 60 * 60 * 1000);
    
    const bookingData = {
      start_time: startTime.toISOString(),
      end_time: endTime.toISOString(),
    };
    const headers = { Authorization: `Bearer ${authStore.token}` };
    
    await axios.post(`http://localhost:8000/api/courts/${courtId}/bookings`, bookingData, { headers });
    toast.success(`Booking untuk ${bookingDate.value.toLocaleDateString('id-ID')} jam ${bookingTime.value} berhasil!`);
  } catch (err: any) {
    if (err.response && err.response.status === 409) {
      toast.error('Jadwal Bentrok! Silakan pilih waktu lain.');
    } else {
      toast.error('Gagal membuat booking.');
    }
    console.error(err);
  }
}

const fetchVenueData = async () => {
  const venueId = route.params.id;
  try {
    const venueResponse = await axios.get(`http://localhost:8000/api/venues/${venueId}`)
    venue.value = venueResponse.data.data
    const courtsResponse = await axios.get(`http://localhost:8000/api/venues/${venueId}/courts`)
    courts.value = courtsResponse.data.data
  } catch (err) {
    error.value = 'Gagal memuat data GOR atau lapangan.'
    console.error(err)
  } finally {
    isLoading.value = false
  }
}

onMounted(() => {
  fetchVenueData()
})
</script>

<template>
  <div>
    <div v-if="isLoading">
      <v-progress-linear indeterminate color="primary"></v-progress-linear>
    </div>
    <v-alert type="error" v-else-if="error">
      {{ error }}
    </v-alert>
    <div v-else-if="venue">
      <v-card>
        <v-img
          v-if="venue.image_url"
          :src="venue.image_url"
          height="300px"
          cover
        ></v-img>
        <div v-else class="bg-grey-lighten-2" style="height: 300px;"></div>

        <v-card-title class="text-h4 d-flex justify-space-between align-center">
          {{ venue.name }}
          <v-btn 
            v-if="authStore.isOwner" 
            :to="{ name: 'edit-venue', params: { id: venue.ID } }" 
            color="grey" 
            icon="mdi-pencil"
            variant="text"
            title="Edit / Upload Gambar"
          ></v-btn>
        </v-card-title>
        <v-card-subtitle>{{ venue.address }}</v-card-subtitle>
        <v-card-text class="mt-4">{{ venue.description }}</v-card-text>
      </v-card>

      <v-card class="mt-6">
        <v-card-title>Booking Lapangan</v-card-title>
        <v-card-text>
          <v-row>
            <v-col cols="12" md="6">
              <v-date-picker v-model="bookingDate" title="Pilih Tanggal"></v-date-picker>
            </v-col>
            <v-col cols="12" md="6">
              <v-select
                v-model="bookingTime"
                label="Pilih Jam Mulai"
                :items="Array.from({length: 15}, (_, i) => `${(i + 8).toString().padStart(2, '0')}:00`)"
                prepend-inner-icon="mdi-clock-time-four-outline"
              ></v-select>
              
              <v-list v-if="courts.length > 0">
                <v-list-subheader>PILIH LAPANGAN</v-list-subheader>
                <v-list-item
                  v-for="court in courts"
                  :key="court.ID"
                  :title="court.name"
                  :subtitle="`Lantai: ${court.floor_type}`"
                >
                  <template v-slot:append>
                    <v-btn color="primary" @click="handleBooking(court.ID)">Book</v-btn>
                  </template>
                </v-list-item>
              </v-list>
              <p v-else>Belum ada lapangan tersedia di GOR ini.</p>
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>
    </div>
  </div>
</template>