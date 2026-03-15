import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'


// UI组件
import ArcoVue from '@arco-design/web-vue';
import '@arco-design/web-vue/dist/arco.css';
// 额外引入图标库
import ArcoVueIcon from '@arco-design/web-vue/es/icon';
app.use(ArcoVue);
app.use(ArcoVueIcon);


createApp(App).use(router).mount('#app')
