<template>
  <div class="login-page">
    <div class="login-form">
      <a-form :model="loginForm" @submit="handleSubmit">
        <!-- 定义里面那个字段 -->
        <a-form-item hide-label>
          <span class="login-title">欢迎登录博客论坛系统</span>
        </a-form-item>
        <a-form-item
          hide-label
          field="username"
          :rules="[{ required: true, message: '请输入用户名' }]"
        >
          <a-input
            v-model="loginForm.username"
            placeholder="请输入用户名"
            allow-clear
          >
            <template #prefix>
              <icon-user />
            </template>
          </a-input>
        </a-form-item>
        <a-form-item
          hide-label
          field="password"
          :rules="[
            { required: true, message: '请输入密码' },
            { minLength: 6, message: '密码至少6位' },
          ]"
        >
          <a-input-password
            v-model="loginForm.password"
            placeholder="请输入密码"
          >
            <template #prefix>
              <icon-lock />
            </template>
          </a-input-password>
        </a-form-item>
        <a-form-item hide-label>
          <a-button html-type="submit" style="width: 100%" type="primary"
            >登录</a-button
          >
        </a-form-item>
      </a-form>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { LOGIN } from "../../api/token";
import { useRouter } from "vue-router";
import { state } from "../../stores/app";

// 路由对象
const router = useRouter();

// 定义表单数据
// 表单对应的响应式数据, 提交给后端的数据
// 按照Vblog login api 来设计
const loginForm = ref({
  username: "",
  password: "",
});
// 表单数据提交函数
const handleSubmit = async (data) => {
  // 表单校验成功, 才与后端交互
  if (data.errors !== undefined) {
    return;
  }

  try {
    const resp = await LOGIN(loginForm.value);
    // 保存登录状态
    state.value.is_login = true;
    state.value.token = resp;

    // 需要跳转 后台页面, Vue router的 Router对象
    // 通过 vue router库来获取一个router对象
    // name 路由的名称
    router.push({ name: "BackendBlogs" });
  } catch (error) {
    // 就是Promise 的Reject 的error
    console.log(error);
  }
};
</script>

<style lang="css" scoped>
.login-page {
  height: 100vh;
  width: 100vw;
  display: flex;
  align-items: center;
  justify-content: center;
}

.login-form {
  display: flex;
  align-items: center;
  flex-direction: column;
  height: 400px;
  width: 400px;
}

.login-title {
  display: flex;
  width: 100%;
  font-size: 16px;
  font-weight: 500;
  justify-content: center;
}
</style>
../../api/token
