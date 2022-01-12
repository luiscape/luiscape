<template>
  <div class="main">
    <Bio />
    <Posts />
  </div>
</template>

<script>
import Bio from "@/components/Bio.vue";
import Posts from "@/components/Posts.vue";

export default {
  components: {
    Bio,
    Posts,
  },
};
</script>

<style scoped>
.main {
  margin-top: 10%;
}
</style>
