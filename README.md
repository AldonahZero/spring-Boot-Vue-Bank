# 项目架构
![架构图](https://github.com/AldonahZero/spring-Boot-Vue-Bank/blob/master/%E5%BE%AE%E6%9C%8D%E5%8A%A1%E6%9E%B6%E6%9E%84%E5%9B%BE.jpg "架构图")


扫码加微信，备注技术新潮流.

![me](https://github.com/AldonahZero/spring-Boot-Vue-Bank/blob/master/Master.png "me")

## 项目介绍

我,请始皇[打钱]是一个前后端分离的工具人系统，项目采用 SpringBoot+Vue 开发，项目加入常见的企业级应用所涉及到的技术点，例如 Redis、RabbitMQ 等(主要是多用用工具多踩踩坑)。


- 项目地址：[https://github.com/AldonahZero/spring-Boot-Vue-Bank](https://github.com/AldonahZero/spring-Boot-Vue-Bank) 
- 作者邮箱：[aldnoah.usa@foxmail.com](aldnoah.usa@foxmail.com) 

### 项目技术栈
#### 后端技术-Go

| 技术                 | 说明                | 官网                                                 |
| -------------------- | ------------------- | ---------------------------------------------------- |
| gin           | 框架        | https://github.com/gin-gonic/gin             |
| MongoDB       | 认证和授权框架      | https://www.mongodb.com        |
| Swagger       | 自动化文档生成工具      | https://swagger.io/      |
| grom       | ORM框架      | github.com/jinzhu/gorm      |
#### 后端技术-Java

| 技术                 | 说明                | 官网                                                 |
| -------------------- | ------------------- | ---------------------------------------------------- |
| SpringBoot           | 容器+MVC框架        | https://spring.io/projects/spring-boot               |
| SpringSecurity       | 认证和授权框架      | https://spring.io/projects/spring-security           |
| MyBatis              | ORM框架             | http://www.mybatis.org/mybatis-3/zh/index.html       |
| MyBatisGenerator     | 数据层代码生成      | http://www.mybatis.org/generator/index.html          |
| PageHelper           | MyBatis物理分页插件 | http://git.oschina.net/free/Mybatis_PageHelper       |
| Swagger-UI           | 文档生产工具        | https://github.com/swagger-api/swagger-ui            |
| Hibernator-Validator | 验证框架            | http://hibernate.org/validator                       |
| Elasticsearch        | 搜索引擎            | https://github.com/elastic/elasticsearch             |
| Kafka	               | 消息队列            | https://www.rabbitmq.com/                            |
| Redis                | 分布式缓存          | https://redis.io/                                    |
| MongoDb              | NoSql数据库         | https://www.mongodb.com                              |
| Docker               | 应用容器引擎        | https://www.docker.com                               |
| Druid                | 数据库连接池        | https://github.com/alibaba/druid                     |
| OSS                  | 对象存储            | https://github.com/aliyun/aliyun-oss-java-sdk        |
| MinIO                | 对象存储            | https://github.com/minio/minio                       |
| JWT                  | JWT登录支持         | https://github.com/jwtk/jjwt                         |
| LogStash             | 日志收集工具        | https://github.com/logstash/logstash-logback-encoder |
| Lombok               | 简化对象封装工具    | https://github.com/rzwitserloot/lombok               |

#### 前端技术

| 技术       | 说明                  | 官网                                   |
| ---------- | --------------------- | -------------------------------------- |
| Vue        | 前端框架              | https://vuejs.org/                     |
| Vue-router | 路由框架              | https://router.vuejs.org/              |
| Vuex       | 全局状态管理框架      | https://vuex.vuejs.org/                |
| Element    | 前端UI框架            | https://element.eleme.io               |
| Axios      | 前端HTTP框架          | https://github.com/axios/axios         |
| v-charts   | 基于Echarts的图表框架 | https://v-charts.js.org/               |
| Js-cookie  | cookie管理工具        | https://github.com/js-cookie/js-cookie |
| nprogress  | 进度条控件            | https://github.com/rstacruz/nprogress  |


## 快速本地部署

1. clone 项目到本地 `https://github.com/AldonahZero/spring-Boot-Vue-Bank.git`
2. 数据库脚本放在项目根目录下，在 MySQL 中执行数据库脚本，导入数据库，并修改项目中关于数据的配置（resources 目录下的 application.properties 文件中）
3. 提前准备好 Redis，在 项目的 application.properties 文件中，将 Redis 配置改为自己的
4. 提前准备好 MQ，在项目的 application.properties 文件中将 MQ 的配置改为自己的
5. 在 IntelliJ IDEA 中打开 bdemo 项目，启动
6. npm 运行前端项目

**OK，至此，服务端就启动成功了，此时我们直接在地址栏输入 `http://localhost:8080` 即可访问我们的项目，如果要做二次开发，请继续看第七、八步。**

7. 进入到vuehr目录中，在命令行依次输入如下命令：

```
# 安装依赖
npm install

# 在 localhost:8080 启动项目
npm run serve
```

由于我在 vuehr 项目中已经配置了端口转发，将数据转发到 Spring Boot 上，因此项目启动之后，在浏览器中输入 `http://localhost:8080` 就可以访问我们的前端项目了，所有的请求通过端口转发将数据传到 Spring Boot 中（注意此时不要关闭 Sprin gBoot 项目）。

8. 最后可以用 WebStorm 等工具打开 vuehr 项目，继续开发，开发完成后，当项目要上线时，依然进入到 vuehr 目录，然后执行如下命令：

```
npm run build
```

该命令执行成功之后，vuehr 目录下生成一个 dist 文件夹，将该文件夹中的两个文件 static 和 index.html 拷贝到 Spring Boot 项目中 resources/static/ 目录下，然后就可以像第 6 步那样直接访问了（关于前后端分离部署，大家也可以参考这个[使用 Nginx 部署前后端分离项目，解决跨域问题](https://mp.weixin.qq.com/s/C7PIck3SIPPTcA3NX3ELoQ)）。


**步骤 7 中需要大家对 NodeJS、NPM 等有一定的使用经验，不熟悉的小伙伴可以先自行搜索学习下，推荐 [Vue 官方教程](https://cn.vuejs.org/v2/guide/)。**

## 参考

- [vue-chat](https://github.com/microzz/vue-chat)

## License

    Copyright 2020 Aldno

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

        http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
 