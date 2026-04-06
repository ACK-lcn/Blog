import client from "./client";

// 颁发Token
export var LOGIN = (data) => {
  return client({
    url: "/api/blog/v1/tokens/",
    method: "post",
    data: data,
  });
};
