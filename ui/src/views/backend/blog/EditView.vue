<template>
  <div>
    <a-page-header :title="title" @back="router.go(-1)"> </a-page-header>
    <!-- 页面内容区 -->
    <div class="page">
      <a-form ref="formRef" layout="vertical" :model="blog">
        <a-form-item
          field="title"
          label="标题"
          :rules="[{ required: true, message: '请输入用标题' }]"
        >
          <a-input v-model="blog.title" placeholder="请输入文章标题" />
        </a-form-item>
        <a-form-item
          field="summary"
          label="摘要"
          :rules="[{ required: true, message: '请输入摘要' }]"
        >
          <a-textarea v-model="blog.summary" placeholder="请输入文章摘要" />
        </a-form-item>
        <a-form-item field="content" label="内容">
          <!-- 选择一个markdown编辑器 -->
          <MdEditor
            style="height: calc(100vh - 340px)"
            v-model="blog.content"
            @onSave="onSave"
          />
        </a-form-item>
      </a-form>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from "vue-router";
import { ref, onBeforeMount } from "vue";
import { MdEditor } from "md-editor-v3";
import "md-editor-v3/lib/style.css";
import { CREATE_BLOG, UPDATE_BLOG, GET_BLOG } from "@/api/blog";
import { Message } from "@arco-design/web-vue";

const router = useRouter();
const title = ref("创建文章");
// 创建文章
const blog = ref({
  id: 0,
  title: "",
  summary: "",
  content: "",
});

// 来绑定名叫formRef 元素(组件)
const formRef = ref(null);
const onSave = async () => {
  // 手动触发校验与提交
  // 1. 校验表单
  const validate = await formRef.value.validate();

  const payload = blog.value;
  // 2. 保持内容
  if (!validate) {
    console.log(blog.value);
    // 判断是否有id
    if (blog.value.id) {
      // 更新
      try {
        const resp = await UPDATE_BLOG(payload.id, payload);
        blog.value = resp;
        Message.success(`更新成功`);
      } catch (error) {
        Message.error(`更新失败: ${error.message}`);
      }
    } else {
      // 创建
      try {
        const resp = await CREATE_BLOG(payload);
        blog.value = resp;
        Message.success(`保持成功`);
      } catch (error) {
        Message.error(`保持失败: ${error.message}`);
      }
    }
  }
};

// 补充编辑逻辑
const filledUpateData = async () => {
  const blogId = router.currentRoute.value.query.id;
  if (blogId) {
    title.value = "更新文章";
    try {
      const resp = await GET_BLOG(blogId);
      blog.value = resp;
    } catch (error) {
      Message.error(`保持失败: ${error.message}`);
    }
  }
};

onBeforeMount(() => {
  filledUpateData();
});
</script>

<style lang="css" scoped></style>
