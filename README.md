# gdc
基于gin的脚手架（gin-dev-cli）

`项目生成命令:`
    

---

## 笔记

---

### Mysql安装配置
1. Mac默认启动Mysql     ```/opt/homebrew/opt/mysql/bin/mysqld_safe --datadir\=/opt/homebrew/var/mysql &```
2. 创建数据库：```CREATE DATABASE mydatabase;```
3. 添加一个访问用户：```CREATE USER 'zdu'@'%' IDENTIFIED BY 'zdu123456';```
4. 添加权限：```GRANT ALL PRIVILEGES ON mydatabase.* TO 'zdu'@'%';``` 
5. 刷新权限：```FLUSH PRIVILEGES;```


    
    