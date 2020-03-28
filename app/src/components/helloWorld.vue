
<template>
  <div class="hello">
    <headerComp></headerComp>
    <img class="logo" src="@/assets/hello.png" alt="Vue Logo" width="200" height="200" />
    <h3>hi!</h3>
  </div>
</template>

<script>
import headerComp from "./headerComp";

export default {
  name: "helloWorld",

  props: {
    msg: String
  },

  components: {
    headerComp
  }
};
</script>

<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
.logo {
  display: block;
  margin-left: auto;
  margin-right: auto;
  margin-top: 20px;
  width: auto;
}
</style>
