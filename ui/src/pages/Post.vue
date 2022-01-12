<template>
  <el-container>
    <el-col :xs="24" :sm="24" :md="12" :lg="12" :xl="12" class="content">
      <div v-html="parsedMarkdown"></div>
    </el-col>
  </el-container>
</template>

<script>
import { marked } from "marked";

export default {
  name: "Post",
  data() {
    return {
      markdown: "test",
    };
  },
  computed: {
    parsedMarkdown() {
      return marked(this.markdown);
    },
  },
  mounted() {
      this.fetchMarkdownData()
  },
  methods: {
    fetchMarkdownData: function () {
      var self = this;
      this.axios
        .get(`posts/${self.$route.params.section}/${self.$route.params.path}`)
        .then(function (response) {
          self.markdown = response.data;
        })
        .catch(function (error) {
          console.log(error);
        });
    },
  },
};
</script>


<style scoped>
@media screen and (max-width : 480px)
{
  h1
  {
    font-size: 10px;
  }
}
@media screen and (max-width : 1204px)
{
  h1
  {
    font-size: 60px;
  }
}
.content {
  /* top right bottom left */
  margin: 0 0 0 12%;
}
</style>
