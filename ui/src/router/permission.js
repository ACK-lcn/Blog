import { state } from "@/stores/app";

export var beforeEach = async (to) => {
  if (to.fullPath.startsWith("/backend")) {
    if (!state.value.is_login) {
      // 针对去往backend的页面, 才需要登录
      // 跳转到登录页面, next就是router push方法
      // 例如：用户访问的目标页面是: CommentList,
      // 跳转到登录页面 登录成功后(携带上目标页面的路由名称), 需要重定向到目标页面去(push)
      // 直接Return路由
      return {
        name: "Login",
        query: {
          redirectName: to.name,
        },
      };
    }
  }
};
