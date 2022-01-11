<template>
  <el-container>
    <el-col :xs="18" :sm="18" :md="10" :lg="10" :xl="6" class="content">
      <h1>Luis Capelo</h1>
      <p>
        I am interested in creating products that enhance human abilities,
        including cognition and creativity.
      </p>
      <p>
        I was previously the co-founder and CTO of Grid.ai; alumni of
        <a href="https://glossier.com" target="_blank">Glossier</a>,
        <a href="https://forbes.com" target="_blank">Forbes Media</a>,
        <a href="https://flowminder.org" target="_blank">Flowminder</a>, and the
        <a href="https://data.humdata.org/" target="_blank">United Nations</a>.
      </p>
    </el-col>
  </el-container>
</template>

<script>
export default {
  name: "Landing",
};
</script>


<style scoped>
.content {
  /* top right bottom left */
  margin: 0 0 0 12%;
}
</style>
