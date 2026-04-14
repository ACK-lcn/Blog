<template>
  <div class="page">
    <a-breadcrumb class="detail-breadcrumb">
      <a-breadcrumb-item>
        <a-link @click="router.push({ name: 'FrontendBlogs' })">首页</a-link>
      </a-breadcrumb-item>
      <a-breadcrumb-item>文章详情</a-breadcrumb-item>
    </a-breadcrumb>
    <a-spin :loading="queryLoading" style="width: 100%">
      <div v-if="blog.id" class="blog-detail-card">
        <div class="article-title">{{ blog.title }}</div>
        <div class="blog-meta">
          <span>作者：{{ blog.author || "匿名作者" }}</span>
          <span>发布时间：{{ formatTime(blog.published_at || blog.created_at) }}</span>
          <span>最后更新：{{ formatTime(blog.updated_at) }}</span>
        </div>
        <div v-if="blog.summary" class="article-summary">{{ blog.summary }}</div>
        <MdPreview :model-value="blog.content || ''" preview-theme="smart-blue" />
      </div>
      <a-empty v-else-if="!queryLoading" description="文章不存在或已删除" />
    </a-spin>
  </div>
</template>

<script setup>
import { onBeforeMount, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import dayjs from "dayjs";
import { Message } from "@arco-design/web-vue";
import { MdPreview } from "md-editor-v3";
import "md-editor-v3/lib/preview.css";
import { GET_BLOG } from "@/api/blog";

const route = useRoute();
const router = useRouter();

const queryLoading = ref(false);
const blog = ref({
  id: 0,
  title: "",
  author: "",
  summary: "",
  content: "",
  created_at: 0,
  updated_at: 0,
  published_at: 0,
});

const formatTime = (unixTime) => {
  if (!unixTime) return "-";
  return dayjs.unix(unixTime).format("YYYY-MM-DD HH:mm");
};

const queryDetail = async () => {
  const blogId = route.params.id;
  if (!blogId) {
    return;
  }
  queryLoading.value = true;
  try {
    const resp = await GET_BLOG(blogId);
    blog.value = resp;
  } catch (error) {
    Message.error(`获取文章失败: ${error.message}`);
  } finally {
    queryLoading.value = false;
  }
};

onBeforeMount(() => {
  queryDetail();
});
</script>

<style lang="css" scoped>
.detail-breadcrumb {
  margin-bottom: 12px;
}

.blog-detail-card {
  border: 1px solid rgb(229, 230, 235);
  border-radius: 8px;
  padding: 20px;
  background: #fff;
}

.article-title {
  font-size: 30px;
  line-height: 1.35;
  color: rgb(29, 33, 41);
  font-weight: 700;
  margin-bottom: 12px;
}

.blog-meta {
  margin-bottom: 12px;
  color: rgb(134, 144, 156);
  display: flex;
  gap: 16px;
  font-size: 13px;
  flex-wrap: wrap;
}

.article-summary {
  padding: 10px 12px;
  margin-bottom: 12px;
  border-left: 3px solid rgb(64, 128, 255);
  background: rgb(245, 247, 255);
  color: rgb(78, 89, 105);
}
</style>
