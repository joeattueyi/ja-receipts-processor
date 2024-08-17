
# Receipt Processor

## Run via Docker
Clone the project

```bash
git clone https://github.com/joeattueyi/ja-receipts-processor.git
```

Change into project directory

```bash
cd ja-receipts-processor
```

build and run
```bash
docker build -t joeattueyi/ja-receipts-processor
docker run -p 8080:8080 joeattueyi/ja-receipts-processor
```

## Run locally
If you already have Go installed you can do the following

Clone the project

```bash
git clone https://github.com/joeattueyi/ja-receipts-processor.git
```

Change into project directory

```bash
cd ja-receipts-processor
```

Run
```bash
go run .
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