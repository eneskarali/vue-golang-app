<template>
  <div class="header" >
    <img class = "logo" src="@/assets/login_icon.png" alt="Vue Logo" width="50" height="auto">
    <h2>DASHBOARD</h2>
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
  background: rgba(129, 218, 95, 0.63);
  color: #000000;
  width: 100% - 5px;
}

.content {
  padding: 10px;
}
h2{
    position: absolute;
    left: 50%;
    margin-left: -80px;
    margin-top: 10px;
}

.sticky {
  position: fixed;
  top: 0;
  width: 100%
}
.sticky + .content {
  padding-top: 102px;
}

.logo{
    float: left;
}

</style>
