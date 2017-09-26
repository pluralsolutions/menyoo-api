# Install

....

# Live

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/a8aa8c2b6b9e9fc7dbec)


# API DOC

## All Products

GET - /restaurants/1/products

### Response Body

```json
{
   "id":1,
   "restaurant_id":1,
   "title":"Pizza",
   "description":"Pizza boa",
   "image":"http://image.com",
   "price_cents":2000,
   "ingredient_groups":[
      {
         "id":1,
         "title":"Queijo",
         "basic":false,
         "product_id":1,
         "ingredients":[
            {
               "id":2,
               "name":"Muçarela",
               "price_cents":50
            },
            {
               "id":1,
               "name":"Gorgonzola",
               "price_cents":140
            }
         ]
      }
   ]
}
```

## Show Product

GET - /restaurants/1/products/1

### Response Body

```json
[
   {
      "id":2,
      "restaurant_id":1,
      "title":"Pizza Presunto",
      "description":"Pizza de preseunto",
      "image":"http://image.com",
      "price_cents":3000,
      "ingredient_groups":null
   },
   {
      "id":1,
      "restaurant_id":1,
      "title":"Pizza",
      "description":"Pizza boa",
      "image":"http://image.com",
      "price_cents":2000,
      "ingredient_groups":null
   }
]
```

## Create order

POST - /orders

### Payload:

```json
{
   "id":3,
   "user_id":"qualquer user",
   "restaurant_id":1,
   "status":"requested",
   "products":[
      {
         "id":15,
         "product":{
            "id":2,
            "restaurant_id":1,
            "title":"Pizza Presunto",
            "description":"Pizza de preseunto",
            "image":"http://image.com",
            "price_cents":3000,
            "ingredient_groups":null
         },
         "product_id":2,
         "order_id":3,
         "quantity":46,
         "total_price_cents":138000,
         "ingredients":null
      },
      {
         "id":14,
         "product":{
            "id":1,
            "restaurant_id":1,
            "title":"Pizza",
            "description":"Pizza boa",
            "image":"http://image.com",
            "price_cents":2000,
            "ingredient_groups":null
         },
         "product_id":1,
         "order_id":3,
         "quantity":32,
         "total_price_cents":64000,
         "ingredients":null
      },
      {
         "id":16,
         "product":{
            "id":1,
            "restaurant_id":1,
            "title":"Pizza",
            "description":"Pizza boa",
            "image":"http://image.com",
            "price_cents":2000,
            "ingredient_groups":null
         },
         "product_id":1,
         "order_id":3,
         "quantity":5,
         "total_price_cents":10140,
         "ingredients":[
            {
               "id":1,
               "name":"Gorgonzola",
               "price_cents":140
            }
         ]
      }
   ]
}
```

### Response Body

```json
{
   "id":5,
   "user_id":"123das",
   "restaurant_id":1,
   "status":"requested",
   "products":[
      {
         "id":9,
         "product_id":1,
         "order_id":5,
         "quantity":2,
         "total_price_cents":4140,
         "ingredients":[
            {
               "id":1,
               "name":"Gorgonzola",
               "price_cents":140
            }
         ]
      },
      {
         "id":10,
         "product_id":2,
         "order_id":5,
         "quantity":1,
         "total_price_cents":3000,
         "ingredients":null
      }
   ]
}
```

## Update product's order quantity

PUT - /restaurant/{restaurant_id}/orders/{order_id}/products/{product_order_id}/quantity


### Headers

```
uid | {uid from FirebaseAuth}
```

### Payload

```
{
  "quantity": 12
}
```

### Response Body

```json
{
   "id":2,
   "user_id":"0piEoeeFuhOJgcdXIWR5xTCAkux2",
   "restaurant_id":1,
   "status":"requested",
   "products":[
      {
         "id":3,
         "product":{
            "id":2,
            "restaurant_id":1,
            "title":"Pizza Presunto",
            "description":"Pizza de preseunto",
            "image":"http://image.com",
            "price_cents":3000,
            "ingredient_groups":null
         },
         "product_id":2,
         "order_id":2,
         "quantity":4,
         "total_price_cents":12000,
         "ingredients":null,
         "evaluation":{
            "id":0,
            "user_id":"",
            "product_order_id":0,
            "score":0
         }
      },
      {
         "id":4,
         "product":{
            "id":1,
            "restaurant_id":1,
            "title":"Pizza",
            "description":"Pizza boa",
            "image":"http://image.com",
            "price_cents":2000,
            "ingredient_groups":null
         },
         "product_id":1,
         "order_id":2,
         "quantity":3,
         "total_price_cents":6190,
         "ingredients":[
            {
               "id":2,
               "name":"Muçarela",
               "price_cents":50
            },
            {
               "id":1,
               "name":"Gorgonzola",
               "price_cents":140
            }
         ],
         "evaluation":{
            "id":0,
            "user_id":"",
            "product_order_id":0,
            "score":0
         }
      }
   ]
}
```

## Retrieve products order from user

GET - /users/me/restaurants/1/orders/4


### Headers

```
uid | {uid from FirebaseAuth}
```

### Response Body

```json
{
   "id":4,
   "user_id":"souza",
   "restaurant_id":1,
   "status":"requested",
   "products":[
      {
         "id":6,
         "product_id":1,
         "order_id":4,
         "quantity":12,
         "total_price_cents":24140,
         "ingredients":[
            {
               "id":1,
               "name":"Gorgonzola",
               "price_cents":140
            }
         ]
      },
      {
         "id":7,
         "product_id":1,
         "order_id":4,
         "quantity":3,
         "total_price_cents":6190,
         "ingredients":[
            {
               "id":2,
               "name":"Muçarela",
               "price_cents":50
            },
            {
               "id":1,
               "name":"Gorgonzola",
               "price_cents":140
            }
         ]
      }
   ]
}
```

## Place an Order

POST - users/me/restaurants/1/orders/3/place

### Headers

```
uid | {uid from FirebaseAuth}
```

### Response Body

```json
{
  "id": 3,
  "user_id": "carlos",
  "restaurant_id": 1,
  "status": "paid",
  "products": null
}
```

## Evaluations

POST - users/me/restaurants/1/products/3/evaluations

### Headers

```
uid | {uid from FirebaseAuth}
```

### Response Body

```json
{
  "id": 8,
  "user_id": "souza",
  "product_id": 1,
  "score": 30
}
```

## All products by user from restaurant

GET - /users/me/restaurants/{restaurant_id}/products

### Headers
```
uid | {uid from FirebaseAuth}
```

### Response Body
```json
{
   "products":[
      {
         "id":12,
         "product":{
            "id":1,
            "restaurant_id":1,
            "title":"Pizza",
            "description":"Pizza boa",
            "image":"http://image.com",
            "price_cents":2000,
            "ingredient_groups":null
         },
         "product_id":1,
         "order_id":2,
         "quantity":61,
         "total_price_cents":122000,
         "ingredients":null,
         "evaluation":{
            "id":0,
            "user_id":"",
            "product_order_id":0,
            "score":0
         }
      }
   ]
}
```


## Current Order

GET - /users/me/restaurants/{restaurant_id}/current_order

### Headers
```
uid | {uid from FirebaseAuth}
```

### Response Body
```json
{
   "id":3,
   "user_id":"IxqhAVmgtkh7COqhfdm2PLPLtxG3",
   "restaurant_id":1,
   "status":"requested",
   "products":[
      {
         "id":5,
         "product":{
            "id":2,
            "restaurant_id":1,
            "title":"Pizza Presunto",
            "description":"Pizza de preseunto",
            "image":"http://image.com",
            "price_cents":3000,
            "ingredient_groups":null
         },
         "product_id":2,
         "order_id":3,
         "quantity":1,
         "total_price_cents":3500,
         "ingredients":[
            {
               "id":3,
               "name":"Ingrediente adicional",
               "price_cents":500
            }
         ],
         "evaluation":{
            "id":0,
            "user_id":"",
            "product_order_id":0,
            "score":0
         }
      }
   ]
}
```
