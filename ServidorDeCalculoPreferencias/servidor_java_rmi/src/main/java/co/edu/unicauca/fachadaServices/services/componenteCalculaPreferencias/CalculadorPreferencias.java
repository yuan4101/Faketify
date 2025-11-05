package co.edu.unicauca.fachadaServices.services.componenteCalculaPreferencias;

import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.stream.Collectors;

import co.edu.unicauca.fachadaServices.DTO.PreferenciasDTORespuesta;
import co.edu.unicauca.fachadaServices.DTO.ReproduccionesDTOEntrada;

public class CalculadorPreferencias {

    public PreferenciasDTORespuesta calcular(String nombreUsuario,
            List<ReproduccionesDTOEntrada> reproducciones) {

        List<ReproduccionesDTOEntrada> reproduccionesUsuario = reproducciones.stream()
            .filter(Objects::nonNull)
            .filter(r -> r.getUsuario() != null && r.getUsuario().equals(nombreUsuario))
            .collect(Collectors.toList());

        System.out.println("  - Reproducciones filtradas para " + nombreUsuario + ": " + 
            reproduccionesUsuario.size());

        Map<String, Integer> contadorGeneros = contarGeneros(reproduccionesUsuario);
        Map<String, Integer> contadorArtistas = contarArtistas(reproduccionesUsuario);
        Map<String, Integer> contadorIdiomas = contarIdiomas(reproduccionesUsuario);

        System.out.println("  - Contadores: Géneros: " + contadorGeneros.size() + 
            ", Artistas: " + contadorArtistas.size() + 
            ", Idiomas: " + contadorIdiomas.size());

        String idUsuario = obtenerIdUsuario(reproducciones, nombreUsuario);

        return PreferenciasMapper.construirRespuesta(
            idUsuario,
            contadorGeneros,
            contadorArtistas,
            contadorIdiomas
        );
    }

    private Map<String, Integer> contarGeneros(
            List<ReproduccionesDTOEntrada> reproducciones) {
        
        Map<String, Integer> contador = new HashMap<>();
        reproducciones.forEach(r -> {
            String genero = r.getGenero() != null ? r.getGenero() : "Desconocido";
            contador.put(genero, contador.getOrDefault(genero, 0) + 1);
        });
        return contador;
    }

    private Map<String, Integer> contarArtistas(
            List<ReproduccionesDTOEntrada> reproducciones) {
        
        Map<String, Integer> contador = new HashMap<>();
        reproducciones.forEach(r -> {
            String artista = r.getArtista() != null ? r.getArtista() : "Desconocido";
            contador.put(artista, contador.getOrDefault(artista, 0) + 1);
        });
        return contador;
    }

    private Map<String, Integer> contarIdiomas(
            List<ReproduccionesDTOEntrada> reproducciones) {
        
        Map<String, Integer> contador = new HashMap<>();
        reproducciones.forEach(r -> {
            String idioma = r.getIdioma() != null ? r.getIdioma() : "Desconocido";
            contador.put(idioma, contador.getOrDefault(idioma, 0) + 1);
        });
        return contador;
    }

    private String obtenerIdUsuario(
            List<ReproduccionesDTOEntrada> reproducciones,
            String nombreUsuario) {
        
        return reproducciones.stream()
            .filter(Objects::nonNull)
            .filter(r -> r.getUsuario() != null && r.getUsuario().equals(nombreUsuario))
            .findFirst()
            .map(r -> String.valueOf(r.getUsuario().hashCode()))
            .orElse("0");
    }
}
