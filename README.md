# Receipt Processor

A simple web service for processing receipts and calculating points based on predefined rules.

## Endpoints

### Process Receipts

- **Path**: `/receipts/process`
- **Method**: `POST`
- **Payload**: Receipt JSON
- **Response**: JSON containing an ID for the receipt

### Get Points

- **Path**: `/receipts/{id}/points`
- **Method**: `GET`
- **Response**: JSON containing the number of points awarded

## Rules for Calculating Points

- One point for every alphanumeric character in the retailer name.
- 50 points if the total is a round dollar amount with no cents.
- 25 points if the total is a multiple of 0.25.
- 5 points for every two items on the receipt.
- If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
- 6 points if the day in the purchase date is odd.
- 10 points if the time of purchase is after 2:00pm and before 4:00pm.

## Setup and Installation

1. Ensure Go is installed on machine.

2. Clone the repository:

   ```bash
   git clone https://github.com/Manasvi-Shahare/Receipt-Processor.git
   cd Receipt-Processor

3. Install dependencies:
   
   ```bash
   go mod tidy

## Build and Run the Application

1. Build the Application:

   ```bash
   make build

2. Run the Application:

   ```bash
   make run


## Testing

To run the tests:

```bash
make test
```

## Example Usage

1. Process a Receipt

   ```bash
   curl -X POST http://localhost:8080/receipts/process \
   -H "Content-Type: application/json" \
   -d '{
       "retailer": "Walgreens",
       "purchaseDate": "2022-01-02",
       "purchaseTime": "08:13",
       "total": "2.65",
       "items": [
           {"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
           {"shortDescription": "Dasani", "price": "1.40"}
       ]
   }'
    ```
   
2. Get Points

   ```bash
   curl http://localhost:8080/receipts/<ID>/points
   ```
   
   Replace <ID> with the ID returned from the /receipts/process endpoint.
