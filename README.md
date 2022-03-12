# simple-websockets-chat-app

## Building and Deploying

### Compilation

```bash
make clean build
```

### Deploying

```bash
make deploy
```

## Local development

### Start containers

```bash
docker compose up -d
```

### Accessing dynamdb admin

http://localhost:8001/

### Invoke function locally

```bash
sam local invoke ConnectFunction -e testdata/event_connection.json
```

## Using wscat for Testing

```bash
ENDPOINT=`aws cloudformation describe-stacks --stack-name simple-websockets-chat-app --query "Stacks[0].Outputs[?OutputKey=='WebSocketEndpoint'].OutputValue" --output text`
```

```bash
wscat -c "${ENDPOINT}?room=chat:1"
```

```json
{"message": "test", "room": "chat:1", "data": "some message."}
```

## Deleting

```bash
sam delete --config-file samconfig.toml
````
