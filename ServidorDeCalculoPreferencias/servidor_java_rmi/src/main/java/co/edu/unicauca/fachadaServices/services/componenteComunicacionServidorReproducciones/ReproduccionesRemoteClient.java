package co.edu.unicauca.fachadaServices.services.componenteComunicacionServidorReproducciones;

import co.edu.unicauca.fachadaServices.DTO.ReproduccionesDTOEntrada;
import feign.Headers;
import feign.RequestLine;
import java.util.List;

public interface ReproduccionesRemoteClient {
    @RequestLine("GET /playback/listarReproducciones")
    @Headers("Accept: application/json")
    List<ReproduccionesDTOEntrada> obtenerReproducciones();
}
