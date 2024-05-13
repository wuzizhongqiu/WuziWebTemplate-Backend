# GoWeb 后端开发万用模板

## 模板特点

### 使用框架 & 插件

* Kratos 作为 web 开发框架（同时支持 grpc）

* JWT 作为认证工具（后续可能会改成 Session + Redis 的形式维持登录态）

* Wire 作为依赖注入工具

* Validate 作为参数校验工具

* Protoc-gen-go 作为代码生成工具，加快开发效率

* Gen 作为更友好 & 更安全 GORM 代码生成工具

    

### 数据存储

* MySQL 数据库
* Redis 内存数据库

后续考虑

* 腾讯云 COS 对象存储

* Elasticsearch 搜索引擎
* kafka 消息队列



### 业务特性

* protocol 代码生成（通过写 proto 文件形式完成代码框架开发/项目配置/自定义错误码/参数校验 等等）
* 封装通用响应类 & 全局日志
* openapi 接口文档

后续

* 使用 Session Redis 实现分布式登录



### 业务功能

* 用户登录、注册、检索、权限管理（后续改成 Session Redis 的分布式登录会添加登出功能）
* 帖子创建（后续会支持更多功能）

后续

* 完善用户相关功能
* 完善帖子相关功能



### 单元测试

* 暂无



### 架构设计

* 参考 Kratos 官方描述：[Go工程化 - Project Layout 最佳实践 | Kratos (go-kratos.dev)](https://go-kratos.dev/blog/go-project-layout/)



## 快速上手

1）下载依赖

```shell
依赖略多，先偷个懒
```

2）启动项目

```shell
kratos run
```



## 数据库表设计

```sql
-- 用户表
create table if not exists users 
(
    id bigint(32) unsigned not null auto_increment comment '主键',
    create_time timestamp not null default current_timestamp comment '创建时间',
    update_time timestamp not null default current_timestamp on update current_timestamp comment '更新时间',
    delete_at timestamp null default null comment '逻辑删除标记',
    
    account varchar(256) not null comment '账号',
    password varchar(512) not null comment '密码',
    username varchar(256) not null default '新用户' comment '用户昵称',
    avatar_url varchar(1024) not null default '' comment '用户头像',
    gender tinyint not null default '1' comment '性别:1男;2女;3隐藏',
    user_info varchar(1024) not null default '' comment '个人简介',
    
    user_role tinyint(4) not null default '10' comment '身份:10用户;20管理者;30封禁',
    
    primary key (`id`)
) engine=innodb default charset=utf8mb4 comment='用户表';
    
    
-- 帖子表
create table if not exists post
(
    id         bigint auto_increment comment 'id' primary key,
    create_time datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    delete_at   tinyint  default 0                 not null comment '是否删除',

    title      varchar(512)                   not null comment '标题',
    content    text                           not null comment '内容',
    tags       varchar(1024)                  not null comment '标签列表（json 数组）',
    thumb_num   int      default 0                 not null comment '点赞数',
    favour_num  int      default 0                 not null comment '收藏数',
    user_id     bigint                             not null comment '创建用户 id',

    index idx_userId (user_id)
) engine=innodb default charset=utf8mb4 comment '帖子';


-- 帖子点赞表
create table if not exists post_thumb
(
    id         bigint auto_increment comment 'id' primary key,
    create_time datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',

    post_id     bigint                             not null comment '帖子 id',
    user_id     bigint                             not null comment '创建用户 id',

    index idx_postId (post_id),
    index idx_userId (user_id)
) engine=innodb default charset=utf8mb4 comment '帖子点赞';


-- 帖子收藏表
create table if not exists post_favour
(
    id         bigint auto_increment comment 'id' primary key,
    create_time datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',

    post_id     bigint                             not null comment '帖子 id',
    user_id     bigint                             not null comment '创建用户 id',

    index idx_postId (post_id),
    index idx_userId (user_id)
) engine=innodb default charset=utf8mb4 comment '帖子收藏';
```









