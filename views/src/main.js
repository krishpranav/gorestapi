import { createApp } from 'vue'
import App from './App.vue'
import axios from 'axios';

Vue.prototype.$http = axios;

createApp(App).mount('#app')
