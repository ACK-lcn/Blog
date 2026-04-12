import client from "./client";

// 颁发Token
export var LOGIN = (data) => {
  return client({
    url: "/api/blog/v1/tokens/",
    method: "post",
    data: data,
  });
};

// 销毁 Token（后端 DELETE /v1/tokens/，Body: access_token, refresh_token）
export var LOGOUT = (data) => {
  return client({
    url: "/api/blog/v1/tokens/",
    method: "delete",
    data: data,
  });
};
