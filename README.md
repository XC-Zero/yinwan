# yinwan

## 开发测试工具

- Redis连接工具  [下载链接](https://gitee.com/qishibo/AnotherRedisDesktopManager/attach_files/934334/download/Another-Redis-Desktop-Manager.1.5.1.exe)
- Kafka连接工具   [下载链接](https://www.kafkatool.com/download2/offsetexplorer_64bit.exe)
- MongoDB连接工具  [下载链接](https://robomongo.org/)
- ElasticSearch 就用 Kibana 就好
- 其他关系型数据库就用 Goland 吧或 DataGrip 都差不多
- influxDB 就用它自己的网页可视化版吧
- Minio 也是用它网页管理就行
- 接口测试工具 可以用PostMan也可以用Goland自带的（`工具` -> `Http客户端` -> `在Http客户端创建请求` ）


## 数据库要求

- 设计要求
    - 全部使用逻辑外键，使用文档或注释说明关系（强制性）
    - 主键一律私有且使用 *int 类型且自增,主键的 json 别名一律为 id（强制性）
    - 不允许对数据库增加存储过程和触发器，所有数据处理、加工逻辑应全部在服务层（强制性）
    - 尽量不使用gorm中字段级别权限控制（建议性）
    - 一张表尽量不超过20个字段（建议性）
  

## 开发要求
  - 端口要求
    - 生产环境：只对外暴露唯一入口的端口
    - 测试环境：随意
  - GitHub要求
    - 使用自己名字分支进行开发
    - 以dev为主开发分支
    - 提交代码前一定要从dev拉取一遍代码