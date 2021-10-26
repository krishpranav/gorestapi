<template>
  <div>
    <div id="app" v-for="albums in albums" v-bind:key="albums.id" >
      <!-- <img alt="Vue logo" src="./assets/logo.png"> -->
      <HelloWorld msg="Album Details"/>
      <h2>{{ albums.title }}</h2>
      <p>{{ albums.body }}</p>
      <!-- <h2>{{ aboutapi.title }} </h2> -->
    </div>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'

export default {
  name: 'App',

  components: {
    HelloWorld
  },

  data() {
    return {
      albums: [],
    };
  },

  methods: {
    async getAlbum() {
      try {
        const response = await this.$http.get(
          "http://localhost:8080/albums"
        );
        this.albums = response.data;
      } catch (error) {
        console.log('OOPS');
        console.log(error);
      }
    },
    
    async getAboutApi() {
      try {
        const aboutapi = await this.$http.get(
          "http://localhost:8080/about"
        );
        this.aboutapi = aboutapi.data;  
      } catch (error) {
        console.log('OOPS');
        console.log('error');
      }
    }

  },

  created() {
    this.getAlbum();
    this.getAboutApi();
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
