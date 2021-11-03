import { createApp } from 'vue'
import App from './App.vue'
import axios from 'axios';
import Vue from 'vue';

Vue.prototype.$http = axios;

createApp(App).mount('#app')
