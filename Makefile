#ABOUT POSTGRES
DB_DRIVE = postgresql
DB_USER = postgres
DB_PASS = qwerty
DB_IP = localhost
DB_PORT = 5432
DB_NAME = postgres
#extra options
DB_OPT = sslmode=disable


#ABOUT PATHS
PATH_TO_SCHEMA = ./schema


MigrUP:
	migrate -path $(PATH_TO_SCHEMA) -database "$(DB_DRIVE)://$(DB_USER):$(DB_PASS)@$(DB_IP):$(DB_PORT)/$(DB_NAME)?$(DB_OPT)" -verbose up

MigrDown: MigrUP
	migrate -path $(PATH_TO_SCHEMA) -database "$(DB_DRIVE)://$(DB_USER):$(DB_PASS)@$(DB_IP):$(DB_PORT)/$(DB_NAME)?$(DB_OPT)" -verbose up
