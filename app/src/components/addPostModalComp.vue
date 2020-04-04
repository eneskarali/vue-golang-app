<template>
  <div class="modal">
    <div class="modal-content">
      <h1>a new post</h1>
      <h3>write something and share it...</h3>
      <textarea
        type="text"
        name="postText"
        id="postText"
        class="postInput"
        autofocus
        placeholder="type here..."
      />
      <div class="button" v-on:click="shareOnClick">
        <p>
          <b>share</b>
        </p>
      </div>
      <div class="delButton" v-on:click="deleteOnClick">
        <p>
          <b>delete</b>
        </p>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "appPostModalComp",
  props: {
    msg: String
  },
  methods: {
    deleteOnClick() {
      var modal = document.getElementById("addPost");
      var addComp = document.getElementById("addPostComp");
      document.getElementById("postText").value = "";
      modal.style.display = "none";
      addComp.style.display = "none";
    },
    shareOnClick() {
      var posttext = document.getElementById("postText").value;
      console.log(localStorage.username);

      const params = {
        postBy: localStorage.username,
        posttext: posttext
      };
      axios
        .post(
          "http://localhost:8000/addpost",
          params,
          { withCredentials: true },
          {
            headers: {
              "content-type": "application/json"
            }
          }
        )
        .then(response => {
          if (response.status == 200) {
            var modal = document.getElementById("addPost");
            document.getElementById("postText").value = "";
            modal.style.display = "none";
            location.reload();
          }
        })
        .catch(err => {
          console.log("add post service unable:" + err);
        });
    }
  }
};
</script>

<style  scoped>
.modal {
  width: 50%;
  margin: auto;
}
.modal-content {
  background-color: #fefefe;
  margin: auto;
  padding: 20px;
  border: 1px solid #888;
  width: 100%;
  height: 380px;
  border-radius: 10px;
}
h3 {
  float: left;
  margin-top: 10px;
  margin-left: 8px;
}
.postInput {
  display: block;
  margin-top: 40px;
  height: 47%;
  width: 98%;
  cursor: pointer;
  overflow: auto;
  font-size: 16px;
  font: helvatica;
  border: 1px solid #8cd48f;
  border-radius: 10px;
  padding: 1%;
  font-family: Arial, Helvetica, sans-serif;
}

.postInput::placeholder {
  margin: 10px;
}

.postInput:focus {
  outline: none;
}

.button {
  color: #ffffff !important;
  background: #4caf50;
  margin-top: 12px;
  padding-top: 10px;
  padding-bottom: 10px;
  padding-right: 15px;
  padding-left: 15px;
  border: 4px solid #a0c2a1 !important;
  cursor: pointer;
  border-radius: 12px;
  font-size: 18px;
  color: black;
  font: bold;
  margin-left: 80%;
  margin-right: 00px;
}

.button:hover {
  color: #000000 !important;
  border: 4px solid #a0c2a1 !important;
  border-radius: 12px;
  transition: all 0.5s ease 0s;
}
.delButton {
  color: #ffffff !important;
  background: #8d1616;
  margin-top: 5px;
  padding: 10px;
  border: 4px solid #c74545 !important;
  transition: all 0.3s ease 0s;
  cursor: pointer;
  border-radius: 12px;
  font-size: 18px;
  color: black;
  font: bold;
  margin-left: 80%;
  margin-right: 0px;
}

.delButton:hover {
  color: #ffffff !important;
  border: 4px solid #000000 !important;
  border-radius: 12px;
  transition: all 0.5s ease 0s;
}

@media (max-width: 800px) {
  .modal {
    width: 80%;
    margin: 4%;
  }
  .delButton {
    margin-left: 70%;
  }
  .button {
    margin-left: 70%;
  }
}
</style>