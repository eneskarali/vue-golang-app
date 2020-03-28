<template>
  <div class="header">
    <img class="logo" src="@/assets/login_icon.png" alt="Vue Logo" width="50" height="auto" />
    <h1>timeline</h1>
  </div>
</template>

<script>
export default {
  name: "headerComp",
  props: {
    msg: String
  }
};
</script>

<style  scoped>
.header {
  padding: 5px 5px;
  padding-bottom: 55px;
  background: rgba(89, 177, 54, 0.88);
  color: #000000;
  width: 100% - 5px;
}
h1 {
  position: absolute;
  left: 50%;
  margin-left: -80px;
  margin-top: 7px;
}

.logo {
  float: left;
}
</style>
