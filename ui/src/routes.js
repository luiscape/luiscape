import { createWebHashHistory, createRouter } from "vue-router";
import Landing from "@/pages/Landing.vue";
import Post from "@/pages/Post.vue"


const router = createRouter({
    history: createWebHashHistory(),
    routes: [
        {
            path: "/",
            name: "landing",
            component: Landing,
        },
        {
            path: "/post/:section/:path",
            name: "post",
            component: Post,
        }
    ],
});

export default router;