curl --request POST \
  --url http://localhost:3000/api/v1/product \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "product1",
	"price": 10.99,
	"stock": 10
  }'
