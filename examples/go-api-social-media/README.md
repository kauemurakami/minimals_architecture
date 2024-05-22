## Driver Postgress
```shell 
go get github.com/jackc/pgx  
go get github.com/jackc/pgx/v5/pgconn@v5.5.5
```  


## Create .env
```go get github.com/joho/godotenv```  
Create in root folder create a file ```.env``` and add environment variables:  
```
DB_HOST=127.0.0.1
DB_USER=myuser
DB_PASS=mypass
DB_NAME=api-social-media
DB_PORT=postgresport normally 5432
API_PORT=3000
DB_SSL=disable
DB_TZ=America/Sao_Paulo your timezone
```

## Connection with db
Create a database with pgAdmin or in terminal  
In your root folder create ```/app/core/db/db_config.go``` and add this code  
```go
package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func SetupDB() {
	connStr := "host=DB_HOST port=DB_PORTS user=DB_USER password=DB_PASS dbname=DB_NAME sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Verify the connection is alive.
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	log.Println("Database connection established")
}
```
By creating this way we will be able to recover the connection whenever we want with db.DB  

## Add extension uuid-ossp
In pgadmin -> your-db -> extendions -> create -> uuid-ossp  
OR in migrations  
```sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    nick VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    pass VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```
Using uuid-ossp to suport uuids used with package ```go get github.com/google/uuid```  

## Migration
```shell 
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```  

##### Create migration files up down
```shell
migrate create -ext sql -dir ./app/core/db/migrations create_table_users
```  

##### Migrate up to db
```shell
migrate -database postgres://DB_USER:DB_PASS@DB_HOST:DB_PORT/DB_NAME?sslmode=disable -path ./app/core/db/migrations up 
```  
##### Migrate down to db
```shell
migrate -database postgres://DB_USER:DB_PASS@DB_HOST:DB_PORT/DB_NAME?sslmode=disable -path ./app/core/db/migrations down 
```

...
...
go get github.com/dgrijalva/jwt-go
...
go get golang.org/x/crypto/bcrypt 
...