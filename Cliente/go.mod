module cliente.local/grpc-cliente

go 1.24.5

// Servidor de canciones
require servidor.local/grpc-servidorCanciones v0.0.0
replace servidor.local/grpc-servidorCanciones => ../servidorDeCanciones

// Servidor de streaming
require servidor.local/grpc-servidorStreaming v0.0.0
replace servidor.local/grpc-servidorStreaming => ../servidorDeStreaming