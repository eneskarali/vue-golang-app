<template>
  <div class="page">
    <headerComp class="header"></headerComp>
    <div class="content">
      <h1 class="title">welcome {{activeUserName}}</h1>
      <br />
      <div class="scrollList">
        <Scroll-view>
          <template slot-scope="visibility">
            <div
              class="scrollViewItem"
              v-for="i in items"
              :key="i.date"
              :visibile="visibility"
            >
              <div  class="headInfo">
                <img class="pp" src="@/assets/user_pp.png" alt="pp" width="40" height="40" />
                <p class="name">
                  <b>{{i.name}} {{i.surname}}</b>
                </p>
              </div>
              <br />
              <div class="Post">
                <br />
                <p
                  class="postText"
                >{{i.postText}}</p>
              </div>
            </div>
          </template>
        </Scroll-view>
      </div>
    </div>
    <div class="button" v-on:click="addClicked">
      <img class="add_icon" src="@/assets/add_icon.png" alt="add icon" width="34px" height="34px" />
      <p class="add_text">
        <b>ADD POST</b>
      </p>
    </div>
    <div id="addPost" class="addPostModal">
      <addPostModalComp />
    </div>
  </div>
</template>

<script>
import headerComp from "./headerComp.vue";
import addPostModalComp from "./addPostModalComp.vue";
import axios from "axios";

export default {
  name: "dashboard",
  props: {
    msg: String
  },
  components: {
    headerComp,
    addPostModalComp
  },

  data() {
    return {
      page: 1,
      items: [],
      loading: true,
      activeUserName: "",
      to: 10,
      from: 0
    };
  },
  watch: {
    page: {
      immediate: true,
      handler: function() {
        this.fetchMore();
      }
    }
  },
  methods: {
    fetchMore() {
      this.loading = true;
      const params = {
        from: this.from,
        to: this.to
      };
      axios
        .post(
          "http://localhost:8000/get10post",
          params,
          { withCredentials: true },
          {
            headers: {
              "content-type": "application/json"
            }
          }
        )
        .then(({ data }) => {
          this.loading = true;
          console.log(data)
          this.items = this.items.concat(data.slice(0, 7));
        })
        .catch(e => {
          this.loading = false;
          console.log(e);
        });
    },
    scroll() {
      window.onscroll = () => {
        let bottomOfWindow =
          Math.max(
            window.pageYOffset,
            document.documentElement.scrollTop,
            document.body.scrollTop
          ) +
            window.innerHeight ===
          document.documentElement.offsetHeight;

        if (bottomOfWindow) {
          this.fetchMore();
        }
      };
    },
    addClicked() {
      var addModal = document.getElementById("addPost");
      addModal.style.display = "block";
    }
  },
  mounted() {
    this.scroll();
    window.onclick = function(event) {
      var modal = document.getElementById("addPost");
      if (event.target == modal) modal.style.display = "none";
    };

    if (localStorage.name) {
      this.activeUserName = localStorage.name.toLowerCase().trim();
    }
  }
};
</script>

<style scoped>
.button {
  color: #ffffff !important;
  background: #4caf50;
  margin-bottom: 0;
  padding: 12px;
  border: 4px solid #000000 !important;
  display: inline-block;
  transition: all 0.3s ease 0s;
  cursor: pointer;
  border-radius: 50px;
  font-size: 16px;
  margin: 2px;
  color: black;
  font: bold;
  position: fixed;
  bottom: 4%;
  right: 6%;
}

.button:hover {
  color: #000000 !important;
  border: 4px solid #a0c2a1 !important;
  border-radius: 50px;
  transition: all 0.5s ease 0s;
}

.add_icon {
  display: block;
  float: left;
  margin-right: 10px;
}

.add_text {
  margin-top: 7px;
  padding-right: 5px;
  font-size: 18px;
  display: inline-block;
}

.header {
  position: fixed;
  top: 0px;
  width: 100%;
  height: 0;
}
.content {
  margin-top: 85px;
}
.scrollViewItem {
  border: 1px solid #85ac86;
  margin-top: 20px;
  margin-left: auto;
  margin-right: auto;
  min-height: 180px;
  overflow: auto;
  width: 50%;
  cursor: pointer;
  border-radius: 6px;
}
.scrollViewItem:hover {
  border: 1px solid #c5e2c6 !important;
  transition: all 0.3s ease 0s;
}

.pp {
  float: left;
  margin: 6px;
  display: block;
}
.name {
  margin-top: 15px;
  padding-left: 3px;
  font-size: 18px;
  float: left;
}
.title {
  float: left;
  margin-left: 2%;
}

.scrollList {
  margin-top: 35px;
}
.Post {
  float: left;
}

.postText {
  display: block;
  text-align: left;
  margin-top: 0px;
  margin-left: 40px;
  margin-right: 35px;
  padding-bottom: 35px;
}
.addPostModal {
  display: none; /* Hidden by default */
  position: fixed; /* Stay in place */
  z-index: 1; /* Sit on top */
  padding-top: 13%; /* Location of the box */
  left: 0;
  top: 0;
  width: 100%; /* Full width */
  height: 100%; /* Full height */
  overflow: auto; /* Enable scroll if needed */
  background-color: rgb(0, 0, 0); /* Fallback color */
  background-color: rgba(17, 17, 17, 0.5);
}
.close {
  color: #aaaaaa;
  float: right;
  font-size: 28px;
  font-weight: bold;
}

.close:hover,
.close:focus {
  color: #000;
  text-decoration: none;
  cursor: pointer;
}

@media (max-width: 800px) {
  .scrollViewItem {
    width: 96%;
  }
  .button {
    bottom: 2%;
    right: 1%;
  }
  .addPostModal {
    padding-top: 45%;
  }
}
</style>