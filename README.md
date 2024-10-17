# worker-go

Proyecto de prueba para procesar mensajes de Kafka y poblar base de datos Mongo.

Para levantar kafka y mongo db ejecutar: 

```
  docker-compose up -d
```

Para levantar el worker y la api externa ejecutar

```
  make worker-api
```

Para enviar mensajes ejecutar

```
  docker-compose exec kafka kafka-console-producer --topic test-topic --bootstrap-server localhost:9092
```

Enviar mensaje de ejemplo:
```
  { "orderId": "order-1234", "clientId": "670f1e74c2f2ce81e7ec441d", "products": [{ "productId": "670f0de272bfc343993ddc5f", "name": "Producto 1", "price": 3990}]}
```

Se pueden modificar distintos valores como el clientId o productId para generar errores y verificar reintentos.


Para crear clientes en la DB

```
  curl --location 'http://localhost:8080/clients' \
  --header 'Content-Type: application/json' \
  --data '{
      "name": "prueba",
      "active": true
  }'
```

Para crear productos

```
  curl --location 'http://localhost:8080/products' \
  --header 'Content-Type: application/json' \
  --data '{
      "name": "Producto 1",
      "description": "Producto de prueba",
      "price": 3990
  }'
```

Para dar de baja Kafka y la MongoDB

```
  docker-compose down
```
