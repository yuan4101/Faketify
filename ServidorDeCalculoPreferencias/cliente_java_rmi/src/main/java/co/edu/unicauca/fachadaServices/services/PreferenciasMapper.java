package co.edu.unicauca.fachadaServices.services;

import co.edu.unicauca.fachadaServices.DTO.PreferenciasDTORespuesta;
import co.edu.unicauca.fachadaServices.DTO.PreferenciaGeneroDTORespuesta;
import co.edu.unicauca.fachadaServices.DTO.PreferenciaArtistaDTORespuesta;
import co.edu.unicauca.fachadaServices.DTO.PreferenciaIdiomaDTORespuesta;
import co.edu.unicauca.grpc.Preferencias;

public class PreferenciasMapper {

    public static Preferencias.PreferenciaResponse convertirDTOaProtobuf(
            PreferenciasDTORespuesta preferencias) {

        Preferencias.PreferenciaResponse.Builder responseBuilder = 
            Preferencias.PreferenciaResponse.newBuilder();

        responseBuilder.setIdUsuario(preferencias.getIdUsuario());

        // Agregar géneros
        if (preferencias.getPreferenciasGeneros() != null) {
            for (PreferenciaGeneroDTORespuesta genero : preferencias.getPreferenciasGeneros()) {
                responseBuilder.addPreferenciasGeneros(
                    Preferencias.Genero.newBuilder()
                        .setNombreGenero(genero.getNombreGenero())
                        .setNumeroPreferencias(genero.getNumeroPreferencias())
                        .build()
                );
            }
        }

        // Agregar artistas
        if (preferencias.getPreferenciasArtistas() != null) {
            for (PreferenciaArtistaDTORespuesta artista : preferencias.getPreferenciasArtistas()) {
                responseBuilder.addPreferenciasArtistas(
                    Preferencias.Artista.newBuilder()
                        .setNombreArtista(artista.getNombreArtista())
                        .setNumeroPreferencias(artista.getNumeroPreferencias())
                        .build()
                );
            }
        }

        // Agregar idiomas
        if (preferencias.getPreferenciasIdiomas() != null) {
            for (PreferenciaIdiomaDTORespuesta idioma : preferencias.getPreferenciasIdiomas()) {
                responseBuilder.addPreferenciasIdiomas(
                    Preferencias.Idioma.newBuilder()
                        .setNombreIdioma(idioma.getNombreIdioma())
                        .setNumeroPreferencias(idioma.getNumeroPreferencias())
                        .build()
                );
            }
        }

        return responseBuilder.build();
    }
}
