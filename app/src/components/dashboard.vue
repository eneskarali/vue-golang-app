<template>
  <div class="page">
    <headerComp class="header"></headerComp>
    <div class="content">
      <h1 class="title">welcome {{activeUserName}}</h1>
      <br />
      <div class="scrollList">
        <Scroll-view id="scrollViewParent">
          <template slot-scope="visibility">
            <div class="scrollViewItem" v-for="i in items" :key="i.date" :visibile="visibility">
              <div style="width:100%;height:100%" v-on:click="scrollItemClicked(i)">
                <div class="headInfo">
                  <img class="pp" src="@/assets/user_pp.png" alt="pp" width="40" height="40" />
                  <p class="name">
                    <b>{{i.name}} {{i.surname}}</b>
                  </p>
                </div>
                <br />
                <div class="Post">
                  <p class="postText">{{i.posttext}}</p>
                </div>
              </div>
            </div>
          </template>
        </Scroll-view>
      </div>
    </div>
    <div class="button" v-on:click="addClicked">
      <img class="add_icon" src="@/assets/add_icon.png" alt="add icon" width="34px" height="34px" />
      <p class="add_text">
        <b>NEW POST</b>
      </p>
    </div>
    <div id="addPost" class="addPostModal">
      <addPostModalComp id="addPostComp" style="display:none" />
      <postDetailComp
        :detailPostContent = this.detailPostContent
        id="postDetailComp"
        style="display:none"
      />
    </div>
  </div>
</template>

<script>
import headerComp from "./headerComp.vue";
import addPostModalComp from "./addPostModalComp.vue";
import postDetailComp from "./postDetailComp.vue";
import axios from "axios";

export default {
  name: "dashboard",
  props: {
    msg: String
  },
  components: {
    headerComp,
    addPostModalComp,
    postDetailComp
  },

  data() {
    return {
      page: 1,
      items: [],
      activeUserName: "",
      to: 10,
      from: 0,
      detailPostContent: {name:"", surname:"",posttext: "",date:"",postcount: ""}
      
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
          if (data) {
            this.items = this.items.concat(data);
          }
        })
        .catch(e => {
          console.log("get post method unable:" + e);
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
          this.from = document.getElementById(
            "scrollViewParent"
          ).childElementCount;
          this.fetchMore();
        }
      };
    },
    timeConverter(UNIX_timestamp) {
      var a = new Date(UNIX_timestamp * 1000);
      var months = [
        "Jan",
        "Feb",
        "Mar",
        "Apr",
        "May",
        "Jun",
        "Jul",
        "Aug",
        "Sep",
        "Oct",
        "Nov",
        "Dec"
      ];
      var year = a.getFullYear();
      var month = months[a.getMonth()];
      var date = a.getDate();
      var hour = a.getHours();
      var min = a.getMinutes();
      var time =
        ("0" + date).slice(-2) + " " + month + " " + year + " " + ("0" + hour).slice(-2) + ":" + ("0" + min).slice(-2);
      return time;
    },
    addClicked() {
      var addModal = document.getElementById("addPost");
      var addComp = document.getElementById("addPostComp");
      addModal.style.display = "block";
      addComp.style.display = "block";
    },

    scrollItemClicked(detail) {
      this.detailPostContent = JSON.parse(JSON.stringify(detail)) ;
      let date = this.timeConverter(this.detailPostContent.date)
      this.detailPostContent.date = date 
      var detailComp = document.getElementById("postDetailComp")
      var addModal = document.getElementById("addPost");
      detailComp.style.display = "block"
      addModal.style.display = "block"
    }
  },
  mounted() {
    this.scroll();
    window.onclick = function(event) {
      var modal = document.getElementById("addPost");
      var detailComp = document.getElementById("postDetailComp");
      var addComp = document.getElementById("addPostComp");
      if (event.target == modal) {
        modal.style.display = "none";
        detailComp.style.display = "none";
        addComp.style.display = "none";
      }
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
  min-height: 110px;
  overflow: auto;
  width: 50%;
  cursor: pointer;
  border-radius: 6px;
  display: flex;
  flex-wrap: wrap;
}
.scrollViewItem:hover {
  border: 1px solid #c5e2c6 !important;
  transition: all 0.3s ease 0s;
}

.pp {
  float: left;
  margin-top: 10px;
  margin-left: 15px;
  margin-right: 8px;
  display: block;
}
.name {
  margin-top: 20px;
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
  text-align: left;
  min-width: 100%;
}
.postText {
  display: block;
  text-align: left;
  margin-top: 10px;
  margin-left: 20px;
  margin-right: 20px;
  padding-bottom: 30px;
  word-break: break-word;
}

.addPostModal {
  display: none;
  position: fixed;
  z-index: 1;
  padding-top:12%;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  overflow: auto;
  background-color: rgb(0, 0, 0);
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
    padding-top: 30%;
  }
}
</style>