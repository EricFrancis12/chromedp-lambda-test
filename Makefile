build_serverless:
	env GOOS=linux go build -o bootstrap main.go

deploy_serverless: build_serverless
	serverless deploy --stage prod
