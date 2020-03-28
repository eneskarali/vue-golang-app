<template>
  <div class="page">
    <headerComp class="header"></headerComp>
    <div class="content">
      <Scroll-view>
        <template slot-scope="visibility">
          <img v-for="i in items" :src="i.url" :key="i.id" :visibile="visibility" />
        </template>
      </Scroll-view>
    </div>
    <div class="button">
      <img class="add_icon" src="@/assets/add_icon.png" alt="add icon" width="34px" height="34px" />
      <p class="add_text">
        <b>ADD</b>
      </p>
    </div>
  </div>
</template>

<script>
import headerComp from "./headerComp.vue";
import axios from "axios";

export default {
  name: "dashboard",
  props: {
    msg: String
  },
  components: {
    headerComp
  },
  data() {
    return {
      page: 1,
      items: [],
      loading: true
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
      axios
        .get(`https://jsonplaceholder.typicode.com/albums/${this.page}/photos`)
        .then(({ data }) => {
          this.loading = true;
          this.items = this.items.concat(data.slice(1, 6));
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
    }
  },
  mounted() {
    this.scroll();
  }
};
</script>

<style scoped>
.button {
  color: #ffffff !important;
  background: #4caf50;
  margin-bottom: 0;
  padding: 15px;
  border: 4px solid #000000 !important;
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
img {
  display: block;
  margin-left: auto;
  margin-right: auto;
}

.header {
  position: fixed;
  top: 0px;
  width: 100%;
  height: 0;
}
.content {
  margin-top: 80px;
}
</style>