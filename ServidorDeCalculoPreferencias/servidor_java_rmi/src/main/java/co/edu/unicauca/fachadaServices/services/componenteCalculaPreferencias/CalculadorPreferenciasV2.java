package co.edu.unicauca.fachadaServices.services.componenteCalculaPreferencias;

import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.stream.Collectors;

import co.edu.unicauca.fachadaServices.DTO.PreferenciaArtistaDTORespuesta;
import co.edu.unicauca.fachadaServices.DTO.PreferenciaGeneroDTORespuesta;
import co.edu.unicauca.fachadaServices.DTO.PreferenciaIdiomaDTORespuesta;
import co.edu.unicauca.fachadaServices.DTO.PreferenciasDTORespuesta;
import co.edu.unicauca.fachadaServices.DTO.ReproduccionesRealesDTOEntrada;

public class CalculadorPreferenciasV2 {

    public PreferenciasDTORespuesta calcular(String nombreUsuario,
            List<ReproduccionesRealesDTOEntrada> reproducciones) {

        // Filtrar reproducciones del usuario específico - CON VALIDACIONES
        List<ReproduccionesRealesDTOEntrada> reproduccionesUsuario = reproducciones.stream()
            .filter(Objects::nonNull)  // Filtrar nulos
            .filter(r -> r.getUsuario() != null && r.getUsuario().equals(nombreUsuario))
            .collect(Collectors.toList());

        System.out.println("  ✓ Reproducciones filtradas para " + nombreUsuario + ": " + 
            reproduccionesUsuario.size());

        // Contar géneros
        Map<String, Integer> contadorGeneros = new HashMap<>();
        reproduccionesUsuario.forEach(r -> {
            String genero = r.getGenero() != null ? r.getGenero() : "Desconocido";
            contadorGeneros.put(genero, contadorGeneros.getOrDefault(genero, 0) + 1);
        });

        // Contar artistas
        Map<String, Integer> contadorArtistas = new HashMap<>();
        reproduccionesUsuario.forEach(r -> {
            String artista = r.getArtista() != null ? r.getArtista() : "Desconocido";
            contadorArtistas.put(artista, contadorArtistas.getOrDefault(artista, 0) + 1);
        });

        // Contar idiomas - NUEVO
        Map<String, Integer> contadorIdiomas = new HashMap<>();
        reproduccionesUsuario.forEach(r -> {
            String idioma = r.getIdioma() != null ? r.getIdioma() : "Desconocido";
            contadorIdiomas.put(idioma, contadorIdiomas.getOrDefault(idioma, 0) + 1);
        });

        System.out.println("  ✓ Contadores: Géneros=" + contadorGeneros.size() + 
            ", Artistas=" + contadorArtistas.size() + 
            ", Idiomas=" + contadorIdiomas.size());

        // Convertir a DTOs de respuesta - géneros
        List<PreferenciaGeneroDTORespuesta> preferenciasGeneros = contadorGeneros.entrySet()
            .stream()
            .map(e -> new PreferenciaGeneroDTORespuesta(e.getKey(), e.getValue()))
            .sorted((a, b) -> b.getNumeroPreferencias().compareTo(a.getNumeroPreferencias()))
            .collect(Collectors.toList());

        // Convertir a DTOs de respuesta - artistas
        List<PreferenciaArtistaDTORespuesta> preferenciasArtistas = contadorArtistas.entrySet()
            .stream()
            .map(e -> new PreferenciaArtistaDTORespuesta(e.getKey(), e.getValue()))
            .sorted((a, b) -> b.getNumeroPreferencias().compareTo(a.getNumeroPreferencias()))
            .collect(Collectors.toList());

        // Convertir a DTOs de respuesta - idiomas - NUEVO
        List<PreferenciaIdiomaDTORespuesta> preferenciasIdiomas = contadorIdiomas.entrySet()
            .stream()
            .map(e -> new PreferenciaIdiomaDTORespuesta(e.getKey(), e.getValue()))
            .sorted((a, b) -> b.getNumeroPreferencias().compareTo(a.getNumeroPreferencias()))
            .collect(Collectors.toList());

        // Crear respuesta
        PreferenciasDTORespuesta respuesta = new PreferenciasDTORespuesta();
        respuesta.setIdUsuario(reproducciones.stream()
            .filter(Objects::nonNull)
            .filter(r -> r.getUsuario() != null && r.getUsuario().equals(nombreUsuario))
            .findFirst()
            .map(r -> r.getUsuario().hashCode())
            .orElse(0));

        respuesta.setPreferenciasGeneros(preferenciasGeneros);
        respuesta.setPreferenciasArtistas(preferenciasArtistas);
        respuesta.setPreferenciasIdiomas(preferenciasIdiomas);  // ← NUEVO

        return respuesta;
    }
}
