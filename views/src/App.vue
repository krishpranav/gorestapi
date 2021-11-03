<template>
  <img alt="Vue logo" src="./assets/logo.png">

  <div>
    <div v-for="album in album" v-bind:key="album.id">
      <h2>{{ album.title }}</h2>
      <p>{{ album.body }}</p>
    </div>
  </div>

</template>

<script>
export default {
  data() {
    return {
      album: [],
    };
  },

  methods: {
    async getData() {
      try {
        const response = await this.$http.get(
          "http://jsonplaceholder.typicode.com/posts"
        );
        
        this.album = response.data;
      } catch (error) {
        console.log(error);
      }
    },
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
