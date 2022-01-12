import { createWebHashHistory, createRouter } from "vue-router";
import Landing from "@/pages/Landing.vue";


const router = createRouter({
    history: createWebHashHistory(),
    routes: [
        {
            path: "/",
            name: "landing",
            component: Landing,
        },
    ],
});

export default router;