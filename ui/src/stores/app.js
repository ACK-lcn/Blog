// 保持我们当前的应用程序的状态
import { useStorage } from "@vueuse/core";

export const state = useStorage(
  "blog",
  {
    is_login: false,
    token: {},
    menu: {
      selectedKeys: [],
      openKeys: [],
    },
  },
  localStorage,
  { mergeDefaults: true }
);
