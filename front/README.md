# 博客后台


npm run dev后打开页面报错 处理方法
Cannot assign to read only property 'exports' of object BaseClient.js
安装
npm install --save-dev @babel/plugin-transform-modules-commonjs
在项目根目录新增.babelrc文件，并在文件中加入

{
  "plugins": ["@babel/plugin-transform-modules-commonjs"]
}