package co.edu.unicauca.grpc;

import io.grpc.Server;

import io.grpc.ServerBuilder;

import java.io.IOException;

public class PreferenciasGrpcServer {

    private static final int PORT = 50052;

    private Server server;

    public void start() throws IOException {
        server = ServerBuilder.forPort(PORT)
            .addService(new PreferenciasGrpcServiceImpl())
            .build()
            .start();

        System.out.println("✓ Servidor gRPC de Preferencias escuchando en puerto " + PORT);

        Runtime.getRuntime().addShutdownHook(new Thread(() -> {
            System.err.println("Apagando servidor gRPC");
            PreferenciasGrpcServer.this.stop();
        }));
    }

    public void stop() {
        if (server != null) {
            server.shutdown();
        }
    }

    public void blockUntilShutdown() throws InterruptedException {
        if (server != null) {
            server.awaitTermination();
        }
    }

    public static void main(String[] args) throws IOException, InterruptedException {
        final PreferenciasGrpcServer server = new PreferenciasGrpcServer();
        server.start();
        server.blockUntilShutdown();
    }
}
