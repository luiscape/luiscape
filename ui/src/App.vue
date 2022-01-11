<template>
  <div class="main">
    <Landing/>
    <Posts />
  </div>
</template>

<script>
import Landing from "./components/Landing.vue";
import Posts from "./components/Posts.vue";

export default {
  components: {
    Landing,
    Posts,
  },
};
</script>

<style>
* {
  font-weight: 400;
  font-family: 'Work Sans', sans-serif;
}
h1 {
  font-family: 'Lora', serif;  
  font-weight: 700;
  font-size: 60px;
}
h2 {
  font-family: 'Work Sans', serif;  
  font-weight: 800;
  font-size: 26px;
}
.main {
  margin-top: 10%;
}
</style>
