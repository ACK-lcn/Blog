<template>
  <div class="top-bar">
    <div style="margin-left: 12px">博客论坛系统</div>
    <div>
      <!-- 处理后台才显示 -->
      <a-button
        type="text"
        v-if="pageType() === 'backend'"
        @click="router.push({ name: 'FrontendLayout' })"
      >
        前台
      </a-button>
      <!-- 处于前台 才显示, 并且已经登录才显示后台 -->
      <a-button
        type="text"
        v-if="pageType() === 'frontend' && state.is_login"
        @click="router.push({ name: 'BackendLayout' })"
      >
        后台
      </a-button>
      <a-button type="text" v-if="!state.is_login" @click="login">
        登录
      </a-button>
      <a-button type="text" @click="logout" v-if="state.is_login">
        <span style="margin-right: 12px">退出</span>
        <icon-export />
      </a-button>
    </div>
  </div>
</template>

<script setup>
import { state } from "../stores/app";
import { useRouter } from "vue-router";

const router = useRouter();
const logout = () => {
  state.value.is_login = false;
  // 调用后端的Token销毁接口(留给你们自己实现)

  // 重新跳转到登录页面进行重新登录
  router.push({ name: "Login" });
};
const login = () => {
  router.push({ name: "Login" });
};

// 判断当前处于前台还是后台(url)
const pageType = () => {
  if (router.currentRoute.value.fullPath.startsWith("/frontend")) {
    return "frontend";
  } else {
    return "backend";
  }
};
pageType();
</script>

<style lang="css" scoped>
.top-bar {
  height: 45px;
  border-bottom: 1px solid rgb(229, 230, 235);
  display: flex;
  align-items: center;
  justify-content: space-between;
}

/* 调整第三方组件库的样式, 使用css :deep函数 */
:deep(.arco-btn) {
  height: 100%;
}
</style>
