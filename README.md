# Bookstore REST API with MySQL, GORM, JWT and Fiber framework

### Run Mysql via vagrant and docker ( Optional )
```bash
vagrant up
vagrant ssh
docker container run --name mysql -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root#123 mysql
```