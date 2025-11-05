package co.edu.unicauca.fachadaServices.services;

import io.grpc.Server;
import io.grpc.ServerBuilder;
import java.io.IOException;

public class PreferenciasGrpcServer {

    private static final int PORT = 50052;
    private Server server;

    public void start() throws IOException {
        try {
            PreferenciasGrpcServiceImpl servicio = new PreferenciasGrpcServiceImpl();
            
            server = ServerBuilder.forPort(PORT)
                .addService(servicio)
                .build()
                .start();

            System.out.printf("\n\t\t----- SERVIDOR MEDIADOR (Java gRPC <-> Java RMI) [%d] -----\n", PORT);

            Runtime.getRuntime().addShutdownHook(new Thread(() -> {
                System.err.println("Apagando servidor gRPC");
                PreferenciasGrpcServer.this.stop();
            }));
            
        } catch (Exception e) {
            System.err.println("-> Error iniciando servidor gRPC: " + e.getMessage());
            e.printStackTrace();
            throw new IOException("No se pudo iniciar el servidor gRPC", e);
        }
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
}