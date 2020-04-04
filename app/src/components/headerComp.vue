<template>
  <div class="header">
    <img class="logo" src="@/assets/login_icon.png" alt="Vue Logo" width="50" height="auto" />
    <h1>timeline</h1>
    <div class="button" v-on:click="logoutOnClick">
      <img
        class="logout_icon"
        src="@/assets/logout_icon.png"
        alt="add icon"
        width="30px"
        height="30px"
      />
      <p class="logout_text">
        <b>logout</b>
      </p>
    </div>
  </div>
</template>

<script>
export default {
  name: "headerComp",
  props: {
    msg: String
  },
  methods: {
    logoutOnClick() {
      this.$router.push("/login");
      document.cookie =
        "token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
    }
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
.button {
  color: #ffffff !important;
  background: #c74545;
  margin-top: 0px;
  padding-top: 2px;
  padding-bottom: 10px;
  padding-right: 15px;
  padding-left: 15px;
  border: 4px solid #c74545 !important;
  cursor: pointer;
  border-radius: 12px;
  font-size: 18px;
  color: black;
  font: bold;
  position: fixed;
  float: left;
  right: 15px;
  width: 100px;
  max-height: 30px;
}

.button:hover {
  color: #ffffff !important;
  border: 4px solid #000000 !important;
  border-radius: 12px;
  transition: all 0.5s ease 0s;
}
.logout_icon {
  display: block;
  float: left;
  margin-top: 2px;
  padding-top: 0;
  margin-right: 13px;
}
.logout_text {
  margin-top: 8px;
}
@media (max-width: 800px) {
  .button {
    padding-top: 5px;
    width: 20px;
  }
  .logout_icon {
    width: 20px;
    height: 20px;
    margin-right: 6px;
  }
  .logout_text {
    margin-top: 2px;
    display: none;
  }
}
</style>
