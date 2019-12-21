.PHONY:compile
.PHONY:run
.PHONY:ship

compile:
	GOOS=linux go build -o myapp
run: compile
	/usr/local/Cellar/aws-sam-cli/0.38.0/bin/sam local invoke "myapp" -e event.json
ship: compile
	zip hello_lambda.zip myapp
