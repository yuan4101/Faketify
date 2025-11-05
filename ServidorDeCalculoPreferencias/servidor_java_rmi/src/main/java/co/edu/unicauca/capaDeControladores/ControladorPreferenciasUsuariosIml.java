package co.edu.unicauca.capaDeControladores;

import java.rmi.RemoteException;
import java.rmi.server.UnicastRemoteObject;
import co.edu.unicauca.fachadaServices.DTO.PreferenciasDTORespuesta;
import co.edu.unicauca.fachadaServices.services.IPreferenciasService;

public class ControladorPreferenciasUsuariosIml extends UnicastRemoteObject 
        implements ControladorPreferenciasUsuariosInt {
    
    private static final long serialVersionUID = 1L;
    private IPreferenciasService servicioFachadaPreferencias;

    public ControladorPreferenciasUsuariosIml(IPreferenciasService servicioFachadaPreferencias) 
            throws RemoteException {
        super();
        this.servicioFachadaPreferencias = servicioFachadaPreferencias;
    }

    @Override
    public PreferenciasDTORespuesta getPreferenciasUsuario(String nombreUsuario) throws RemoteException {
        System.out.println("\nRMI: Solicitud de preferencias para usuario: " + nombreUsuario);
        return this.servicioFachadaPreferencias.getPreferenciasUsuario(nombreUsuario);
    }
}
