# AWS Lambda function - Go example

## Getting Started

- Install Docker — Docker is used to run the AWS Lambda locally
- Install aws-sam-local — We will use this tool to be able to run Lambda locally on our machine during the development process.
- Install Go packages

```
go get -u github.com/aws/aws-lambda-go/events
go get -u github.com/aws/aws-lambda-go/lambda
go get -u github.com/tidwall/gjson 
```

# To Run

```
cd $GOPATH/src/hello_lambda
make run
```

The above lines will change the directory to the application, compile and run the app. Note that on the very first run Docker will download the image to be able to run the application using aws-sam-local, this may take a couple of minutes. Subsequent calls will be much faster.

# Using with AWS API Gateway

If you have the api gateway setup, you can call your Lambda function from a terminal window using:

**Important: The token has to match whatever you put in your template.yaml file. The token you see below is an example. Make sure to change it in your template.yaml file, and your go code**

```
curl -d '{"param1":"my param value","token":"35c760f4-b3dc-4657-b4f3-2c6566d4f42e"}' -H "Content-Type: application/json" -X POST  https://{url of api gateway}/V1/{endpoint}
```