
# Receipt Processor

## Run
Clone the project

```bash
git clone git@github.com:joeattueyi/ja-receipts-processor
```

Change into project directory

```bash
cd ja-receipts-processor
```

build and run
```bash
docker build .
docker run -p 8080:8080 .
```

## API Documentation

### Process receipts
```http
  POST /receipts/process

  { "id": "7fb1377b-b223-49d9-a31a-5a02701dd310"}
```

### Get points for a receipt
```http
GET /receipts/{id}/points

{ "points": 32 }
```