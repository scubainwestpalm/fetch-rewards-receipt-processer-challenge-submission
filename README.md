# __*Fetch Rewards Receipt Processor Challenge Submission*__

Hey Everyone,

This is the documentation for my solution I created to provide the functionality for the API as described in the [challenge repo](https://github.com/fetch-rewards/receipt-processor-challenge?tab=readme-ov-file).
Please keep in mind that I actually am experienced mostly with **NodeJS** and not in **Go**, so I made all of this in **Go** as it is your _preferred language_ for the solution, as well as to demonstrate that I can learn new languages and utilize them quickly.

## Project Structure

Overall I tried to keep the project as simplistic as possible since I wanted to write concise and human readable code. I did use two libraries in order to make things simple and speed up development, those being [gin](https://github.com/gin-gonic/gin) and [uuid](https://github.com/google/uuid).

> [!TIP]
> Logic in `./main.go` and `./router/*.go`
>

My solution runs a http server via **gin** on port __8080__ using the default host during development(localhost). It has two paths as described in the [challenge repo](https://github.com/fetch-rewards/receipt-processor-challenge?tab=readme-ov-file), `POST /receipts/process` for __processing__ a receipt(calculating its Points and returning a __uuid__) and `GET /receipts/:id/points` for getting the points of a receipt by its __uuid__.


#### Directory

> [!NOTE]
> Documentation for the packages and their methods are [here](DOCS.md) in a separate markdown to keep the _README_ from getting too large.  
>

```
./
    main.go
    go.mod
    go.sum
    .gitignore
    README.md
    DOCS.md
    ./models
        item.go
        receipt.go
    ./router
        receipts.go
        router.go
    ./utils
        points_helpers.go
```

## Application Flow:


#### Processing Receipt and Getting UUID

User makes `POST` request to `/receipts/process` with the request body as a valid receipt JSON 
```json
{
    "retailer": "Walgreens",
    "purchaseDate": "2022-01-02",
    "purchaseTime": "08:13",
    "total": "2.65",
    "items": [
        {"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
        {"shortDescription": "Dasani", "price": "1.40"}
    ]
}
```

__*Then*__

API responds with new _uuid_ for processed receipt and stores the points on the server.

```json
{ "id": "a4566f0b-c7de-4b58-948e-7467993abd64" }
```

#### Fetching Points for Processed Receipt by UUID

User makes `GET` request to `/receipts/:id/points` with _uuid_ returned from previous `POST` request.
> Example Path: `/receipts/a4566f0b-c7de-4b58-948e-7467993abd64/points`
>

__*Then*__

API responds with points for processed receipt under specified _uuid_.

```json
{ "points": 21 }
```