import MainPage from "@/views/MainPage";   
import AboutPage from "@/views/AboutPage";   
// import PostsPage from "@/pages/PostsPage";
// import About from "@/pages/About";
// import PostIdPage from "@/pages/PostIdPage";
// import PostsPageWithStore from "@/pages/PostsPageWithStore";
// import PostsPageCompositionAPI from "@/pages/PostsPageCompositionAPI";
import { createRouter, createWebHistory } from "vue-router";

const routes = [
  {
    path: "/",
    component: MainPage,
  },
  {
    path: "/about",
    component: AboutPage,
  }
];

const router = createRouter({
    routes,
    history: createWebHistory(process.env.BASE_URL)
})

export default router;
