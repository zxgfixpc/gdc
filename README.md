# gdc
基于gin的脚手架（gin-dev-cli）

`项目生成命令:`
    

---

## 笔记

---

### Mysql安装配置
1. 启动     1. /opt/homebrew/opt/mysql/bin/mysqld_safe --datadir\=/opt/homebrew/var/mysql &
2. 创建数据库：
   1. 主:CREATE DATABASE mydatabase;
   2. 读1： CREATE DATABASE mydatabase_slave1;
   3. 读2： CREATE DATABASE mydatabase_slave2;
   4. 表：
      create table student(
        id int(10) primary key auto_increment,
        name varchar(15) not null,
        age int(10) not null
      );
   5. 数据：insert into student (name, age) values ('zhao', 1);   // 主1，读2，3
3. 添加一个访问用户：CREATE USER 'zdu'@'%' IDENTIFIED BY 'zdu123456';
4. 添加权限：GRANT ALL PRIVILEGES ON mydatabase.* TO 'zdu'@'%';  // 授权3个数据库
5. 刷新权限：FLUSH PRIVILEGES;


    
    