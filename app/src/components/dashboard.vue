<template>
  <div class="page">
    <headerComp></headerComp>
  </div>
</template>

<script>
import headerComp from './headerComp.vue'
export default {
  name: "dashboard",
  props: {
    msg: String
  },
  components:{
    headerComp
  }
};
</script>