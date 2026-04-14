<template>
  <div>
    <TopBar></TopBar>
    <div class="frontend-main">
      <div class="frontend-container">
        <!-- 需要显示子路由的页面, 必须要添加RouteView -->
        <router-view></router-view>
      </div>
    </div>
  </div>
</template>

<script setup>
import TopBar from "../../components/TopBar.vue";
</script>

<style lang="css" scoped>
.frontend-main {
  width: 100%;
  min-height: calc(100vh - 45px);
  background: rgb(247, 248, 250);
  display: flex;
  justify-content: center;
}

.frontend-container {
  width: 100%;
  max-width: 1080px;
  min-height: calc(100vh - 45px);
  background: #fff;
  border-left: 1px solid rgb(229, 230, 235);
  border-right: 1px solid rgb(229, 230, 235);
}
</style>
