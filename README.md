# simple-ping-api
This is a simple ping API to demonstrate the basics of API.

## Steps to run DB using docker

- Install docker on system.
- Run ```docker run -p 2000:3306 --name some-mysql -e MYSQL_ROOT_PASSWORD=password -d mysql```
- Run  ```docker ps``` to check if container is running.

## Steps to connect to DB container 

- Run ```docker exec -it some-mysql /bin/bash```
- Run ```mysql -u root -p ``` and enter password.
- This will open mysql for you run query to create database and tables.

```CREATE DATABASE test```

```CREATE TABLE employee(id int AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255), age int)```

## Steps to run the app

- Firstly download dependecies, run ```go get ./...```
- Finally run ```go run main.go```
