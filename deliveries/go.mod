module dominos.com/deliveries

go 1.15

replace github.com/dominos/logs => ../deliveries

require (
	github.com/dominos/logs v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-kit/kit v0.10.0
	github.com/gorilla/mux v1.7.3
	github.com/joho/godotenv v1.3.0
	github.com/lib/pq v1.9.0
)
