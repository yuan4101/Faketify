package co.edu.unicauca.fachadaServices.services.componenteCalculaPreferencias;

import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

import co.edu.unicauca.fachadaServices.DTO.PreferenciaArtistaDTORespuesta;
import co.edu.unicauca.fachadaServices.DTO.PreferenciaGeneroDTORespuesta;
import co.edu.unicauca.fachadaServices.DTO.PreferenciaIdiomaDTORespuesta;
import co.edu.unicauca.fachadaServices.DTO.PreferenciasDTORespuesta;

public class PreferenciasMapper {

    public static List<PreferenciaGeneroDTORespuesta> mapGeneros(Map<String, Integer> contadores) {
        return contadores.entrySet()
            .stream()
            .map(e -> new PreferenciaGeneroDTORespuesta(e.getKey(), e.getValue()))
            .sorted((a, b) -> b.getNumeroPreferencias().compareTo(a.getNumeroPreferencias()))
            .collect(Collectors.toList());
    }

    public static List<PreferenciaArtistaDTORespuesta> mapArtistas(Map<String, Integer> contadores) {
        return contadores.entrySet()
            .stream()
            .map(e -> new PreferenciaArtistaDTORespuesta(e.getKey(), e.getValue()))
            .sorted((a, b) -> b.getNumeroPreferencias().compareTo(a.getNumeroPreferencias()))
            .collect(Collectors.toList());
    }

    public static List<PreferenciaIdiomaDTORespuesta> mapIdiomas(Map<String, Integer> contadores) {
        return contadores.entrySet()
            .stream()
            .map(e -> new PreferenciaIdiomaDTORespuesta(e.getKey(), e.getValue()))
            .sorted((a, b) -> b.getNumeroPreferencias().compareTo(a.getNumeroPreferencias()))
            .collect(Collectors.toList());
    }

    public static PreferenciasDTORespuesta construirRespuesta(
            String idUsuario,
            Map<String, Integer> contadorGeneros,
            Map<String, Integer> contadorArtistas,
            Map<String, Integer> contadorIdiomas) {
        
        PreferenciasDTORespuesta respuesta = new PreferenciasDTORespuesta();
        respuesta.setIdUsuario(idUsuario);
        respuesta.setPreferenciasGeneros(mapGeneros(contadorGeneros));
        respuesta.setPreferenciasArtistas(mapArtistas(contadorArtistas));
        respuesta.setPreferenciasIdiomas(mapIdiomas(contadorIdiomas));
        
        return respuesta;
    }
}
