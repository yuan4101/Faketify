module localClient/grpc-client

go 1.24.5

require (
	github.com/faiface/beep v1.1.0
	google.golang.org/grpc v1.76.0
	google.golang.org/protobuf v1.36.10
	localServer/grpc-songServer v0.0.0
	localServer/grpc-streamingServer v0.0.0
)

require (
	github.com/hajimehoshi/go-mp3 v0.3.0 // indirect
	github.com/hajimehoshi/oto v0.7.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/exp v0.0.0-20190306152737-a1d7652674e8 // indirect
	golang.org/x/image v0.0.0-20190227222117-0694c2d4d067 // indirect
	golang.org/x/mobile v0.0.0-20190415191353-3e0bab5405d6 // indirect
	golang.org/x/net v0.42.0 // indirect
	golang.org/x/sys v0.34.0 // indirect
	golang.org/x/text v0.27.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250804133106-a7a43d27e69b // indirect
)

replace localServer/grpc-songServer => ../ServidorDeCanciones

replace localServer/grpc-streamingServer => ../ServidorDeStreaming
