<template>
  <div class="modal">
    <div class="modal-content">
      <div class="headInfo">
        <img class="pp" src="@/assets/user_pp.png" alt="pp" width="60" height="60" />
        <h2 class="name">
          <b>{{this.detailPostContent.name}} {{this.detailPostContent.surname}}</b>
        </h2>
        <li class="count">post count: {{this.detailPostContent.postcount}}</li>
      </div>
      <br />
      <div class="post">
        <p class="postText">{{this.detailPostContent.posttext}}</p>
      </div>
      <p class="date">{{this.detailPostContent.date}}</p>
    </div>
  </div>
</template>

<script>
export default {
  name: "postDetailComp",
  props: {
    msg: String,
    detailPostContent: {}
  }
};
</script>

<style  scoped>
.modal {
  width: 50%;
  margin: auto;
  border: 2px solid #85ac86;
  border-radius: 11px;
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
.pp {
  float: left;
}

.name {
  float: left;
  margin-top: 15px;
  margin-left: 10px;
}

.count {
  float: left;
  margin-top: 20px;
  margin-left: 20px;
  color: #888;
}
.date {
  margin-top: 10px;
  float: right;
  margin-right: 15px;
}
.post {
  height: 75%;
  width: 100%;
  overflow: auto;
}
.postText {
  margin-top: 20px;
  margin-left: 10px;
  margin-right: 10px;
  text-align: left;
  word-break: break-word;
}

@media (max-width: 800px) {
  .modal {
    width: 80%;
    margin: 4%;
  }
}
</style>