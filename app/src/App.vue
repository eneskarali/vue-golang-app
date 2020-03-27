<template>
  <div id="app">
    <router-view class="rooter" />
  </div>
</template>

<script>
export default {
  name: "App"
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
.router {
  margin: 0;
  padding: 0;
}

* {
  margin: 0;
}
</style>
