# Install

....


# API DOC

## Products by restaurant

GET - /restaurants/1/products


## Product

GET - /restaurants/1/products/1

## Create order

POST - /orders
#### Payload:
```
{
   "user_id":"123das",
   "restaurant_id":1,
   "products":[
      {
         "product_id":1,
         "quantity":2,
         "ingredients":[
            {
               "id":1
            },
            {
               "id":5
            }
         ]
      },
      {
         "product_id":2,
         "quantity":1
      }
   ]
}
```

#### Response Body:
```
{
   "ID":5,
   "user_id":"123das",
   "restaurant_id":1,
   "status":"requested",
   "products":[
      {
         "ID":9,
         "product_id":1,
         "order_id":5,
         "quantity":2,
         "total_price_cents":4140,
         "ingredients":[
            {
               "ID":1,
               "IngredientGroup":{
                  "id":0,
                  "title":"",
                  "basic":false,
                  "product_id":0,
                  "ingredients":null
               },
               "name":"Gorgonzola",
               "price_cents":140
            }
         ]
      },
      {
         "ID":10,
         "product_id":2,
         "order_id":5,
         "quantity":1,
         "total_price_cents":3000,
         "ingredients":null
      }
   ]
}
```

## Add item into order

PUT - /orders

#### Payload:

```
{
  "user_id":"123das",
  "restaurant_id":1,
  "products_order_id": 1
}
```

## Remove item from order
DELETE - /orders

#### Payload:
```
{
  "user_id":"123das",
  "restaurant_id":1,
  "products_order_id": 1
}
```