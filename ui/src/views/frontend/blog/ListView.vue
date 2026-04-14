<template>
  <div class="page">
    <div class="hero">
      <div class="hero-title">欢迎来到博客论坛空间</div>
      <div class="hero-desc">这里展示已发布的文章内容，支持关键词搜索和文章详情阅读。</div>
      <a-input-search
        v-model="queryParams.keywords"
        placeholder="输入标题关键字搜索文章"
        search-button
        allow-clear
        @search="queryBlogs"
      />
    </div>

    <div class="list-wrap">
      <a-spin :loading="queryLoading" style="width: 100%">
        <div class="blog-grid">
          <a-card
            v-for="item in blogs.items"
            :key="item.id"
            hoverable
            class="blog-card"
            @click="toDetail(item.id)"
          >
            <template #title>
              <div class="blog-title">{{ item.title }}</div>
            </template>
            <div class="blog-summary">{{ getPreviewText(item) }}</div>
            <div class="blog-meta">
              <span>{{ item.author || "匿名作者" }}</span>
              <span>{{ dayjs.unix(item.published_at || item.created_at).format("YYYY-MM-DD HH:mm") }}</span>
            </div>
            <a-button type="text" size="mini" @click.stop="toDetail(item.id)">
              阅读全文
            </a-button>
          </a-card>
        </div>
      </a-spin>
    </div>
    <a-empty v-if="!queryLoading && blogs.total === 0" description="暂无文章" />
    <a-pagination
      v-if="blogs.total > 0"
      :total="blogs.total"
      :current="queryParams.page_number"
      :page-size="queryParams.page_size"
      :page-size-options="[5, 10, 20, 30, 50]"
      show-total
      show-jumper
      show-page-size
      @change="onPageNumberChange"
      @page-size-change="onPageSizeChange"
    />
  </div>
</template>

<script setup>
import { onBeforeMount, reactive, ref } from "vue";
import { useRouter } from "vue-router";
import dayjs from "dayjs";
import { LIST_BLOG } from "@/api/blog";

const router = useRouter();

const blogs = ref({ total: 0, items: [] });
const queryLoading = ref(false);
const queryParams = reactive({
  page_size: 10,
  page_number: 1,
  keywords: "",
  status: "published",
});

const queryBlogs = async () => {
  queryLoading.value = true;
  try {
    const resp = await LIST_BLOG(queryParams);
    blogs.value = resp;
  } finally {
    queryLoading.value = false;
  }
};

const onPageSizeChange = (pageSize) => {
  queryParams.page_size = pageSize;
  queryParams.page_number = 1;
  queryBlogs();
};

const onPageNumberChange = (pageNumber) => {
  queryParams.page_number = pageNumber;
  queryBlogs();
};

const toDetail = (id) => {
  router.push({ name: "FrontendBlogDetail", params: { id } });
};

const removeMarkdown = (text) => {
  if (!text) return "";
  return text
    .replace(/```[\s\S]*?```/g, "")
    .replace(/`([^`]+)`/g, "$1")
    .replace(/!\[[^\]]*\]\([^)]+\)/g, "")
    .replace(/\[([^\]]+)\]\([^)]+\)/g, "$1")
    .replace(/[#>*-]/g, "")
    .replace(/\s+/g, " ")
    .trim();
};

const getPreviewText = (item) => {
  const text = item.summary || removeMarkdown(item.content) || "暂无摘要";
  return text.length > 120 ? `${text.slice(0, 120)}...` : text;
};

onBeforeMount(() => {
  queryBlogs();
});
</script>

<style lang="css" scoped>
.hero {
  margin-bottom: 16px;
  padding: 20px;
  border: 1px solid rgb(229, 230, 235);
  border-radius: 8px;
  background: linear-gradient(135deg, #f5f7ff 0%, #f9fbff 100%);
}

.hero-title {
  font-size: 24px;
  font-weight: 600;
  color: rgb(29, 33, 41);
}

.hero-desc {
  margin: 8px 0 14px;
  color: rgb(78, 89, 105);
}

.list-wrap {
  margin: 8px 0 12px;
  min-height: 460px;
}

.blog-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 12px;
}

.blog-card {
  cursor: pointer;
}

.blog-title {
  font-size: 18px;
  font-weight: 600;
  color: rgb(29, 33, 41);
}

.blog-summary {
  color: rgb(78, 89, 105);
  line-height: 1.7;
  min-height: 52px;
}

.blog-meta {
  margin-top: 10px;
  color: rgb(134, 144, 156);
  font-size: 12px;
  display: flex;
  gap: 16px;
}
</style>
