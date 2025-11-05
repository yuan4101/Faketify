package co.edu.unicauca.main;

import java.io.IOException;
import co.edu.unicauca.fachadaServices.services.PreferenciasGrpcServer;

public class Main {
    public static void main(String[] args) throws IOException, InterruptedException {
        final PreferenciasGrpcServer server = new PreferenciasGrpcServer();
        server.start();
        server.blockUntilShutdown();
    }
}
