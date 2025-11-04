package co.edu.unicauca.fachadaServices.services;

import java.util.List;
import co.edu.unicauca.fachadaServices.DTO.PreferenciasDTORespuesta;
import co.edu.unicauca.fachadaServices.DTO.ReproduccionesRealesDTOEntrada;
import co.edu.unicauca.fachadaServices.services.componenteCalculaPreferencias.CalculadorPreferenciasV2;
import co.edu.unicauca.fachadaServices.services.componenteComunicacionServidorReproducciones.ComunicacionServidorReproducciones;

public class PreferenciasServiceImpl implements IPreferenciasService {

    private ComunicacionServidorReproducciones comunicacionServidorReproducciones;
    private CalculadorPreferenciasV2 calculadorPreferencias;

    public PreferenciasServiceImpl() {
        this.comunicacionServidorReproducciones = new ComunicacionServidorReproducciones();
        this.calculadorPreferencias = new CalculadorPreferenciasV2();
    }

    @Override
    public PreferenciasDTORespuesta getPreferenciasUsuario(String nombreUsuario) {
        try {
            // Obtener todas las reproducciones
            List<ReproduccionesRealesDTOEntrada> reproducciones = 
                comunicacionServidorReproducciones.obtenerReproduccionesRemotas();
            
            if (reproducciones == null || reproducciones.isEmpty()) {
                System.out.println("No se encontraron reproducciones para el usuario: " + nombreUsuario);
                return new PreferenciasDTORespuesta();
            }
            
            // Calcular preferencias
            return calculadorPreferencias.calcular(nombreUsuario, reproducciones);
            
        } catch (Exception e) {
            System.err.println("Error al obtener preferencias: " + e.getMessage());
            e.printStackTrace();
            return new PreferenciasDTORespuesta();
        }
    }
}
