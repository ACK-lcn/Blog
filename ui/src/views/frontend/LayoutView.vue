<template>
  <div>
    <TopBar></TopBar>
    <div>博客论坛前台页</div>
    <!-- 需要显示子路由的页面, 必须要添加RouteView -->
    <router-view></router-view>
  </div>
</template>

<script setup>
import TopBar from "../../components/TopBar.vue";
</script>

<style lang="css" scoped></style>
