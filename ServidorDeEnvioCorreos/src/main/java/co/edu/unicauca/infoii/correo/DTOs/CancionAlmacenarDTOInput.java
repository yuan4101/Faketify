package co.edu.unicauca.infoii.correo.DTOs;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Setter;
import com.fasterxml.jackson.annotation.JsonProperty;

@Getter
@Setter
@AllArgsConstructor
public class CancionAlmacenarDTOInput {
    
    @JsonProperty("titulo")
    private String titulo;
    
    @JsonProperty("artista")
    private String artista;
    
    @JsonProperty("genero")
    private String genero;
    
    @JsonProperty("año")
    private String año;
    
    @JsonProperty("idioma")
    private String idioma;
    
    @JsonProperty("duracion")
    private String duracion;
}
