import client from "./client";

// 区分API请求, 可以全大写
// config?: AxiosRequestConfig<D>
// Get 请求, params?   url?a=1&b2
export var LIST_BLOG = (params) =>
  client.get("/api/blog/v1/blogs/", { params });

// 获取博客详情
export var GET_BLOG = (id, params) =>
  client.get(`/api/blog/v1/blogs/${id}`, { params });

// 创建博客
export var CREATE_BLOG = (data) => client.post("/api/blog/v1/blogs/", data);

// 更新博客
export var UPDATE_BLOG = (id, data) =>
  client.put(`/api/blog/v1/blogs/${id}`, data);

// 删除博客
export var DELETE_BLOG = (id) => client.delete(`/api/blog/v1/blogs/${id}`);
