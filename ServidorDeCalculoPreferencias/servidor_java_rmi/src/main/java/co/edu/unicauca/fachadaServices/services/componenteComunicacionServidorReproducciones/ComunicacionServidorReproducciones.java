package co.edu.unicauca.fachadaServices.services.componenteComunicacionServidorReproducciones;

import co.edu.unicauca.fachadaServices.DTO.ReproduccionesDTOEntrada;
import feign.Feign;
import feign.jackson.JacksonDecoder;
import java.util.ArrayList;
import java.util.List;

public class ComunicacionServidorReproducciones {
    private static final String BASE_URL = "http://localhost:5000";
    private final ReproduccionesRemoteClient client;

    public ComunicacionServidorReproducciones() {
        this.client = Feign.builder()
                .decoder(new JacksonDecoder())
                .target(ReproduccionesRemoteClient.class, BASE_URL);
    }

    public List<ReproduccionesDTOEntrada> obtenerReproduccionesRemotas() {
        try {
            List<ReproduccionesDTOEntrada> reproducciones = client.obtenerReproducciones();
            return reproducciones != null ? reproducciones : new ArrayList<>();
        } catch (Exception e) {
            System.err.println("Error al obtener reproducciones: " + e.getMessage());
            e.printStackTrace();
            return new ArrayList<>();
        }
    }
}
