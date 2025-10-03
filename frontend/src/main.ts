import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'

// Impor Vuetify
import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import '@mdi/font/css/materialdesignicons.css' // Impor ikon

// Impor Toastification (masih kita pakai)
import Toast from 'vue-toastification'
import 'vue-toastification/dist/index.css'

// Buat instance Vuetify
const vuetify = createVuetify({
  components,
  directives,
})

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(Toast)
app.use(vuetify) // Gunakan Vuetify

app.mount('#app')