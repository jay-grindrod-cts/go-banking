# Go Banking

Simple CLI Application to practice Go.

To connect to a Database (you will need to set one up called `go_banking`):

```bash
export DBUSER=dbusername
export DBPASS=dbpassword
export DBHOST=dbhost
```

Your database should have the following `users` table set up:

```mysql
+----------+--------------+------+-----+---------+----------------+
| Field    | Type         | Null | Key | Default | Extra          |
+----------+--------------+------+-----+---------+----------------+
| id       | int(11)      | NO   | PRI | NULL    | auto_increment |
| name     | varchar(128) | NO   |     | NULL    |                |
| email    | varchar(128) | NO   |     | NULL    |                |
| password | varchar(128) | NO   |     | NULL    |                |
| balance  | int(11)      | YES  |     | NULL    |                |
+----------+--------------+------+-----+---------+----------------+
```

To run the application:

`go run .`
