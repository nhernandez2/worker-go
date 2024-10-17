# worker-go

Levantar kafka: docker-compose up -d
Verificar que esten corriendo: docker-compose ps

Crear topico:
docker-compose exec kafka kafka-topics --create --topic test-topic --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1

Iniciar productor:
docker-compose exec kafka kafka-console-producer --topic test-topic --bootstrap-server localhost:9092

Mensaje de ejemplo: 

{ "orderId": "order-1234", "clientId": "670f1e74c2f2ce81e7ec441d", "products": [{ "productId": "670f0de272bfc343993ddc5f", "name": "Producto 1", "price": 3990}]}


Iniciar consumidor
docker-compose exec kafka kafka-console-consumer --topic test-topic --bootstrap-server localhost:9092 --from-beginning

Apagar kafka
docker-compose down
