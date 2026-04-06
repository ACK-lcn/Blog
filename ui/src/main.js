import "@arco-design/web-vue/dist/arco.css";
// 尽量让第三方的样式放前面, 自定义的样式放后面，这样才能覆盖它
import "./assets/main.css";

import { createApp } from "vue";
import { createPinia } from "pinia";

import App from "./App.vue";
import router from "./router";

const app = createApp(App);

app.use(createPinia());
app.use(router);

// UI 组件
import ArcoVue from "@arco-design/web-vue";
import ArcoVueIcon from "@arco-design/web-vue/es/icon";
app.use(ArcoVue);
// 额外引入图标库
app.use(ArcoVueIcon);

app.mount("#app");
