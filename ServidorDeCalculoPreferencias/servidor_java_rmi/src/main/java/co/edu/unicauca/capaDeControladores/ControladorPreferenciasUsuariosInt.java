package co.edu.unicauca.capaDeControladores;

import java.rmi.Remote;
import java.rmi.RemoteException;
import co.edu.unicauca.fachadaServices.DTO.PreferenciasDTORespuesta;

public interface ControladorPreferenciasUsuariosInt extends Remote {
    
    // Método que recibe el nombre del usuario como parámetro
    public PreferenciasDTORespuesta getPreferenciasUsuario(String nombreUsuario) throws RemoteException;
}
