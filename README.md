# dynamodb-stream-pipe

Read dynamodb stream and pass to standard input of specified command.

# usags

Example of execution wc -l

Use with AccessKey/SecretKey
```bash
AWS_REGION=us-east-1 AWS_ACCESS_KEY_ID=xxxx AWS_SECRET_ACCESS_KEY=xxxx \
go run cmd/pipe/main.go --table table_name --command=wc --args=-l
```

Use with DynamoDB Local
```bash
AWS_REGION=us-east-1 AWS_ACCESS_KEY_ID=xxxx AWS_SECRET_ACCESS_KEY=xxxx \
go run cmd/pipe/main.go --table table_name --command=wc --args=-l --endpoint=http://localhost:8883
```
