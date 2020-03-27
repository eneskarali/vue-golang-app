<template>
  <div class="page">
    <headerComp></headerComp>
    <div class="button">
      <img class="add_icon" src="@/assets/add_icon.png" alt="add icon" width="30px" height="30px" />
      <p class="add_text">
        <b>ADD</b>
      </p>
    </div>
  </div>
</template>

<script>
import headerComp from "./headerComp.vue";
export default {
  name: "dashboard",
  props: {
    msg: String
  },
  components: {
    headerComp
  }
};
</script>

<style scoped>
.button {
  color: #ffffff !important;
  background: #000000;
  margin-bottom: 0;
  padding: 15px;
  border: 4px solid #494949 !important;
  display: inline-block;
  transition: all 0.3s ease 0s;
  cursor: pointer;
  border-radius: 50px;
  font-size: 18px;
  margin: 2px;
  color: black;
  font: bold;
  position: fixed;
  bottom: 4%;
  right: 6%;
}

.button:hover {
  color: #4caf50 !important;
  border-radius: 50px;
  transition: all 0.5s ease 0s;
}

.add_icon {
  display: block;
  float: left;
  margin-right: 10px;
}

.add_text {
  margin-top: 5px;
  padding-right: 5px;
  font-size: 18px;
  display: inline-block;
}
</style>