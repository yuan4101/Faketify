package co.edu.unicauca.fachadaServices.DTO;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class ReproduccionesRealesDTOEntrada {

    @JsonProperty("Usuario")
    private String usuario;

    @JsonProperty("Genero")
    private String genero;

    @JsonProperty("Artista")
    private String artista;

    @JsonProperty("Titulo")
    private String titulo;

    @JsonProperty("Cliente")
    private String cliente;

    @JsonProperty("Idioma")
    private String idioma;

    @JsonProperty("FechaHora")
    private String fechaHora;
}
