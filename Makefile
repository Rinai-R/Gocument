USER_PATH := /home/rinai/PROJECTS/Gocument/Server/User
DOC_PATH := /home/rinai/PROJECTS/Gocument/Server/Document
AUTH_PATH := /home/rinai/PROJECTS/Gocument/Server/Auth
API_PATH := /home/rinai/PROJECTS/Gocument/Server/Api

start-user:
	cd $(USER_PATH) && go run main.go &

start-doc:
	cd $(DOC_PATH) && go run main.go &

start-auth:
	cd $(AUTH_PATH) && go run main.go &

start-api:
	cd $(API_PATH) && go run main.go &

start-all: start-user start-doc start-auth start-api

stop-user:
	pkill -f "$(USER_PATH)/main.go"

stop-doc:
	pkill -f "$(DOC_PATH)/main.go"

stop-auth:
	pkill -f "$(AUTH_PATH)/main.go"

stop-api:
	pkill -f "$(API_PATH)/main.go"

stop-all: stop-user stop-doc stop-auth stop-api
