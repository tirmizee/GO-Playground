curl -XPOST -H "Content-type: application/json" -d '{
"productName": "SSSS",
"productType" : "sw333"
}' 'http://localhost:8080/product'

curl -XPOST -H "Content-type: application/json" -d '{
"ids": [1,2,3,4,5,6]
}' 'http://localhost:8080/customer'
