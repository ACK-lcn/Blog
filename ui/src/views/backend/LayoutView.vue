<template>
  <div>
    <!-- 顶部导航组件 -->
    <!-- 顶部 -->
    <TopBar></TopBar>
    <!-- 导航与内容 -->
    <div class="layout-content">
      <div class="menu">
        <a-menu
          :style="{ width: '220px', height: '100%' }"
          show-collapse-button
          breakpoint="xl"
          :open-keys="state.menu.openKeys"
          :selected-keys="state.menu.selectedKeys"
          @menuItemClick="onMenuItemClick"
          @subMenuClick="onSubMenuClick"
        >
          <a-sub-menu key="BackendBlogs">
            <template #icon><icon-apps></icon-apps></template>
            <template #title>博客管理</template>
            <a-menu-item key="BackendBlogs">文章列表</a-menu-item>
          </a-sub-menu>
          <a-sub-menu key="CommentList">
            <template #icon><icon-bug></icon-bug></template>
            <template #title>评论管理</template>
            <a-menu-item key="CommentList">评论列表</a-menu-item>
          </a-sub-menu>
        </a-menu>
      </div>
      <!-- 内容区支持滚动条 -->
      <div style="overflow: auto; width: 100%; height: calc(100vh - 45px)">
        <router-view></router-view>
      </div>
    </div>
  </div>
</template>

<script setup>
import TopBar from "../../components/TopBar.vue";
import { useRouter } from "vue-router";
import { state } from "@/stores/app";

const router = useRouter();

// 选中的菜单项 key 数组
const onMenuItemClick = (key) => {
  state.value.menu.selectedKeys = [key];
  router.push({ name: key });
  // 保持选择状态
};

// 展开的子菜单 key 数组
const onSubMenuClick = (key) => {
  state.value.menu.openKeys = [key];
};

// 查询文章的列表
</script>

<style lang="css" scoped>
.menu {
  border-right: 1px solid rgb(229, 230, 235);
}

.layout-content {
  display: flex;
  width: 100%;
  height: calc(100vh - 45px);
}

:deep(.arco-scrollbar) {
  width: 100%;
}
</style>
