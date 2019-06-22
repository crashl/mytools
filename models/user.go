package models

type User struct {
	Id   int
	Name string
	Pwd  string
}
/*
create table user(
id int not null auto_increment primary key,
name varchar(50) not null,
pwd varchar(100) not null
);
*/