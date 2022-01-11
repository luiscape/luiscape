<template>
  <el-container>
    <el-col :xs="24" :sm="24" :md="12" :lg="12" :xl="12" class="content">
      <h2>Posts</h2>
      <div>
        <el-row v-for="post in posts" :key="post.id" class="posts">
          <div v-if="post.Type == 'post'">
            <br />
            <!-- todo: create function with router to load md -->
            <a :href="post.Path">
              {{ post.Path }}
            </a>
            <br />
            <span class="timestamp">
              {{ moment(post.CreationTime).format("MMMM DD, YYYY") }}
            </span>
          </div>
        </el-row>
      </div>
    </el-col>
  </el-container>
</template>

<script>
import moment from "moment";

export default {
  name: "Posts",
  data() {
    return {
      posts: [],
    };
  },
  created: function () {
    this.moment = moment;
  },
  mounted() {
    this.fetchEssaysDatabase();
  },
  methods: {
    fetchEssaysDatabase: function () {
      var self = this;
      this.axios
        .get("/posts/database.json")
        .then(function (response) {
          self.posts = response.data.Entries;
          console.log(self.posts);
        })
        .catch(function (error) {
          console.log(error);
        });
    },
  },
};
</script>


<style scoped>
.content {
  /* top right bottom left */
  margin: 0 0 0 12%;
}
h2 {
  margin-bottom: 0px;
}
.timestamp {
  font-size: 14px;
}
</style>
