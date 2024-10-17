worker: 
	go run cmd/worker/main.go

api:
	go run cmd/api/main.go

worker-api:
	make worker & make api