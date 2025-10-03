<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()
</script>

<template>
  <v-app>
    <v-app-bar app color="primary" dark>
      <v-toolbar-title>
        <RouterLink to="/" class="white-link">SmashBook SBY</RouterLink>
      </v-toolbar-title>
      <v-spacer></v-spacer>

      <v-btn to="/" text>Home</v-btn>

      <v-btn v-if="!authStore.isAuthenticated" to="/login" text>Login</v-btn>

      <v-btn v-if="authStore.isAuthenticated" to="/my-bookings" text>Booking Saya</v-btn>

      <v-btn v-if="authStore.isOwner" to="/dashboard" text>Dashboard</v-btn>

      <v-btn v-if="authStore.isAuthenticated" @click="authStore.logout()" text>Logout</v-btn>
    </v-app-bar>

    <v-main>
      <v-container>
        <RouterView />
      </v-container>
    </v-main>

    <v-footer app padless>
      <v-col class="text-center" cols="12">
        {{ new Date().getFullYear() }} — <strong>SmashBook SBY</strong>
      </v-col>
    </v-footer>
  </v-app>
</template>

<style scoped>
.white-link {
  color: white;
  text-decoration: none;
}
</style>