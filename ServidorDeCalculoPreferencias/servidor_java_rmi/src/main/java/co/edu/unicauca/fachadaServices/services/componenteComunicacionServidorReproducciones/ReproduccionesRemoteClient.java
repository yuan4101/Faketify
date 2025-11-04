package co.edu.unicauca.fachadaServices.services.componenteComunicacionServidorReproducciones;

import co.edu.unicauca.fachadaServices.DTO.ReproduccionesRealesDTOEntrada;
import feign.Headers;
import feign.RequestLine;
import java.util.List;

public interface ReproduccionesRemoteClient {
    @RequestLine("GET /tendencias/listarReproducciones")
    @Headers("Accept: application/json")
    List<ReproduccionesRealesDTOEntrada> obtenerReproducciones();
}
