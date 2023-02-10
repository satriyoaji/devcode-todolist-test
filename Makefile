DATABASE := "mysql://root:Password@tcp(localhost:3306)/todolist_challenge?parseTime=True&loc=Local"
#Database Migration DDL DML
migrate-up:
	migrate -path ./migrations -database $(DATABASE) -verbose up
migrate-down:
	migrate -path ./migrations -database $(DATABASE) -verbose down