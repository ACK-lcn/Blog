import { createRouter, createWebHistory } from "vue-router";
import { beforeEach } from "./permission.js";
// import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    // 登录页
    {
      path: "/login",
      name: "Login",
      component: () => import("../views/login/LoginView.vue"),
    },
    // 前台
    {
      path: "/frontend",
      name: "FrontendLayout",
      component: () => import("../views/frontend/LayoutView.vue"),
      children: [
        {
          // /blogs
          // blogs --> /frontend/blogs
          path: "blogs",
          name: "FrontendBlogs",
          component: () => import("../views/frontend/blog/ListView.vue"),
        },
      ],
    },
    // 后台
    {
      path: "/backend",
      name: "BackendLayout",
      component: () => import("../views/backend/LayoutView.vue"),
      redirect: { name: "BackendBlogs" },
      children: [
        {
          // /blogs
          // blogs --> /frontend/blogs
          path: "blogs",
          name: "BackendBlogs",
          component: () => import("../views/backend/blog/ListView.vue"),
        },
        {
          // /blogs
          // blogs --> /frontend/blogs
          path: "blogs_edit",
          name: "BlogEdit",
          component: () => import("../views/backend/blog/EditView.vue"),
        },
        {
          // /blogs
          // blogs --> /frontend/blogs
          path: "comments",
          name: "CommentList",
          component: () => import("../views/backend/comment/ListPage.vue"),
        },
      ],
    },
  ],
});

// 补充路由的页面导航守卫
router.beforeEach(beforeEach);

export default router;
