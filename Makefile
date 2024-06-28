M_UP:
	migrate -path ./migrations -database 'postgres://postgres:root@localhost:5432/dbname?sslmode=disable' up 

M_DOWN:
	migrate -path ./migrations -database 'postgres://postgres:root@localhost:5432/dbname?sslmode=disable' down

create:
	migrate create -ext sql -dir migrations -seq