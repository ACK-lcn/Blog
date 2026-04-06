// 封装一个统一的全局的http的客户端

import axios from "axios";
import { Message } from "@arco-design/web-vue";

var instance = axios.create({
  baseURL: "",
  // 超时时间
  timeout: 5000,
  // 后端Gin 使用的Bind函数, 而非BindJson, 补充请求Data是哪个格式
  headers: { "Content-Type": "application/json" },
});

// 这里是请求拦截器
// instance.interceptors.request

// 通过响应拦截器统一处理异常
instance.interceptors.response.use(
  (resp) => {
    // 返回处理后的数据, {code: 0, data: {}} , resp.data
    return resp.data;
  },
  (error) => {
    let msg = error.message;
    // 处理自定义异常
    // 根据code做自己的业务逻辑(token过期)
    if (error.response.data && error.response.data.message) {
      // 通用逻辑处理
      msg = error.response.data.message;

      // 自定义业务逻辑处理:
      switch (error.response.data.code) {
        // Token过期, 跳转到Login页面
        case 5001:
          console.log(error.response.data);
          window.location.assign("/login");
          break;

        default:
          break;
      }

      // 是否要注入Error, 业务页面需要拿到异常
      // 只需要获取页面异常
      return Promise.reject(error.response.data);
    }

    // 直接把异常信息提示处理
    Message.error(`${msg}`);
  }
);

export default instance;
