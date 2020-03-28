<template>
  <div class="page">
    <img class="logo" src="@/assets/login_icon.png" alt="Vue Logo" width="200" height="200" />
    <br />
    <div class="login">
      <h1>Sing In</h1>
      <ul>
        <li>
          <p>User Name:</p>
          <input type="text" />
        </li>
        <li>
          <p>Password:</p>
          <input type="password" />
        </li>
        <li>
          <div class="button" v-on:click="change">
            <b>Sing In</b>
          </div>
        </li>
      </ul>
    </div>
    <div class="footerInfo">
      <p>Bu proje Enes Karali tarafından, Kartaca "Çekirdekten Yetişenler Programı" için geliştirilmiştir.</p>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "loginPage",
  props: {
    msg: String
  },
  methods: {
    change() {
      const params = {
        username: "user1",
        password: "password1"
      };
      axios
        .post(
          "http://localhost:8000/signin",
          params,
          { withCredentials: true },
          {
            headers: {
              "content-type": "application/json"
            }
          }
        )
        .then(response => {
          console.log("OK!" + response);
          this.$router.push("/");
        })
        .catch(err => {
          console.log("sing in service unable:" + err);
        });
    }
  }
};
</script>

<style scoped>
.login p {
  float: left;
  font-size: 22px;
  margin: 2px;
  color: black;
  font: normal;
}
h1 {
  margin-top: 10px;
  margin-bottom: 15px;
  color: black;
}
ul {
  list-style-type: none;
  padding: 0;
  margin: 0;
  border-style: solid;
  border-color: #4caf50;
  border-block-width: 20px;
  border-inline-width: 15px;
  border-radius: 10px;
  background-color: #4caf50;
}
li {
  margin: 15px 10px;
}
input {
  height: 35px;
  width: 98%;
  font-size: 20px;
}
.login {
  margin-top: 50px;
  text-align: center;
  width: 550px;
  display: inline-block;
}
.footerInfo {
  position: fixed;
  bottom: 5px;
  width: 100%;
}
.footerInfo p {
  color: #235825;
  font-size: 16px;
}

.page {
  text-align: center;
}

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
}

.button:hover {
  color: #4caf50 !important;
  border-radius: 50px;
  transition: all 0.5s ease 0s;
}

.logo {
  margin-top: 50px;
}

@media (max-width: 800px) {
  .login {
    width: 350px;
  }

  .footerInfo p {
    color: #235825;
    margin: 0;
    font-size: 12px;
  }
}
</style>