# 13-1 前言
資料庫是很重要的資料來源，不僅能放在本地端，更能從遠端供人存取。像是微軟 Azure、亞馬遜 AWS 也提供了雲端資料庫服務。  
這一章主要用意還是在展示 Go 語言是如何連接 SQL 資料庫。  

Go 語言是使用標準套件的 database/sql 做為 API，而 API 底下才會連接到資料庫需要的驅動程式。如、MySQL、Postgres、DB2、ODBC 等。大部分資料庫都有原生的 Go 語言驅動程式套件可下載；少數需要額外套件，如 C 語言實作的 SQLlute3 驅動程式，就需要安裝 GCC 工具（lesson 16）。  

之所以要用這種 API/驅動程式的架構，是為了以 Go 語言為統一抽象介面，使任何人無須了解資料庫的溝通細節就能操作他們。只需在一開始匯入正確的驅動程式和登入資料庫，在這之後的控制過程就是完全一樣的。  

若一開始就透過 database/sql 介面，那換不同資料庫只需更換驅動程式即可。


# 13-2 安裝 MySQL 資料庫
## 13-2-1 安裝 MySQL Server
```shell
sudo apt install -y mariadb-{server,client}
```

## 13-2-2 新增資料庫使用者 

`sudo mysql` 或 `mysql`

在終端機執行以上，會進入 MariaDB monitor 的提示字元：如下 
```
$ mysql
Welcome to the MariaDB monitor.  Commands end with ; or \g.
Your MariaDB connection id is 37
Server version: 10.11.6-MariaDB-0+deb12u1 Debian 12

Copyright (c) 2000, 2018, Oracle, MariaDB Corporation Ab and others.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

MariaDB [(none)]> use mysqldb;
Database changed
MariaDB [mysqldb]> ^DBye
```

1. 建立新使用者，在 MariaDB monitor 內修改：
```sql
create user 'jocelyn'@'localhost';
```

2. 賦予 user 完整操作權限：
```sql
grant all privileges on *.* to 'jocelyn'@'localhost';
```

3. 接著令設定生效：
```sql
flush privileges;
```
> 在這之後就可在主控台輸入 mysql 登入
> 關閉主控台：\q 或 ctrl+d

## 13-2-3 建立一個 MySQL 資料庫 
以上步驟只是建立了個 MySQL 伺服器而已，得新增一個資料庫（database），才能在裡面新增資料。

新增資料庫，透過 MariaDB 命令列客戶端來進行：
```sql
MariaDB [(none)]> create database mysqldb;
```

轉換成新增的資料庫：
```sql
MariaDB [(none)]> use mysqldb;
Database changed
MariaDB [mysqldb]> ^DBye
```

顯示所有資料庫：
```
MariaDB [mysql]> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| mysqldb            |
| performance_schema |
| sys                |
+--------------------+
5 rows in set (0.001 sec)
```

## 13-2-4 下載 Go 語言的 MySQL 驅動程式
