<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/login_icon.png" width="200" height="200">
    <loginPage/>
  </div>
</template>

<script>
import loginPage from './components/loginPage.vue'

export default {
  name: 'App',
  components: {
    loginPage
  }
}
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
