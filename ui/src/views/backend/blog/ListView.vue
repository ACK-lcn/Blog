<template>
  <div class="page">
    <!-- 页头 -->
    <a-breadcrumb>
      <a-breadcrumb-item>博客论坛管理</a-breadcrumb-item>
      <a-breadcrumb-item>文章管理</a-breadcrumb-item>
    </a-breadcrumb>
    <!-- 内容区 -->
    <!-- 表格操作区 -->
    <div class="table-op">
      <div>
        <a-button
          type="primary"
          @click="router.push({ name: 'BlogEdit' })"
          size="small"
          >创建文章</a-button
        >
      </div>
      <div>
        <a-input
          :style="{ width: '320px' }"
          @keyup.enter="queryBlogs"
          v-model="queryParams.keywords"
          placeholder="请输入文章标题, 敲回车键进行模糊搜索"
          allow-clear
        />
      </div>
    </div>
    <!-- 表格内容 -->
    <div>
      <!-- 使用后端分页,  自己适配 -->
      <a-table :data="blogs.items" :pagination="false">
        <template #columns>
          <a-table-column title="编号" data-index="id"></a-table-column>
          <a-table-column title="标题" data-index="title"></a-table-column>
          <a-table-column title="作者" data-index="author"></a-table-column>
          <a-table-column title="状态" data-index="status"></a-table-column>
          <!-- 使用dayjs来处理时间 -->
          <a-table-column title="状态">
            <template #cell="{ record }">
              {{ dayjs.unix(record.created_at).format("YYYY-MM-DD HH:mm") }}
            </template>
          </a-table-column>
          <a-table-column align="center" title="操作">
            <template #cell="{ record }">
              <a-space>
                <a-button
                  type="primary"
                  @click="
                    router.push({ name: 'BlogEdit', query: { id: record.id } })
                  "
                  >编辑</a-button
                >
                <a-button type="primary">发布</a-button>
                <a-popconfirm
                  content="是否确认删除该文章?"
                  type="error"
                  @ok="deleteBlog(record.id)"
                >
                  <a-button
                    status="danger"
                    :loading="deleteLoading === record.id"
                    >删除</a-button
                  >
                </a-popconfirm>
              </a-space>
            </template>
          </a-table-column>
        </template>
      </a-table>
      <!-- 适配后端分页 -->
      <div style="margin-top: 8px">
        <a-pagination
          :total="blogs.total"
          show-total
          show-jumper
          :page-size-options="[2, 10, 20, 30, 50]"
          show-page-size
          @page-size-change="onPageSizeChange"
          @change="onPageNumberChange"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import { onBeforeMount, ref } from "vue";
import { LIST_BLOG, DELETE_BLOG } from "../../../api/blog";
import { Message } from "@arco-design/web-vue";
import dayjs from "dayjs";
import { reactive } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();

// 查询文章列表，记录loadding状态
const blogs = ref({ total: 0, items: [] });
const queryLoading = ref(false);
const queryParams = reactive({
  page_size: 10,
  page_number: 1,
  keywords: "",
});
const queryBlogs = async () => {
  queryLoading.value = true;
  try {
    // { "total": 6,"items": [] }
    const resp = await LIST_BLOG(queryParams);
    blogs.value = resp;
  } finally {
    queryLoading.value = false;
  }
};

// 分页处理
const onPageSizeChange = (pageSize) => {
  queryParams.page_size = pageSize;
  queryParams.page_number = 1;
  queryBlogs();
};
const onPageNumberChange = (pageNumber) => {
  queryParams.page_number = pageNumber;
  queryBlogs();
};

// 当前处于删除中的文章的具体Id
const deleteLoading = ref(0);
const deleteBlog = async (id) => {
  try {
    deleteLoading.value = id;
    await DELETE_BLOG(id);

    // 删除成功后，需要刷新列表的数据
    queryBlogs();
  } catch (error) {
    Message.error(`删除失败: ${error.message}`);
  } finally {
    deleteLoading.value = 0;
  }
};

// 页面渲染之前，需要把数据提前准备好
onBeforeMount(async () => {
  await queryBlogs();
});
</script>

<style lang="css" scoped>
.table-op {
  height: 46px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

:deep(.arco-table .arco-spin) {
  height: unset;
}
</style>
