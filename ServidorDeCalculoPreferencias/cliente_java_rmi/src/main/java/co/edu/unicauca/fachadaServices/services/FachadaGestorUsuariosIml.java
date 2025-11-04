
package co.edu.unicauca.fachadaServices.services;

import java.rmi.RemoteException;

import co.edu.unicauca.capaDeControladores.ControladorPreferenciasUsuariosInt;
import co.edu.unicauca.fachadaServices.DTO.PreferenciasDTORespuesta;

public class FachadaGestorUsuariosIml 
{
    private final ControladorPreferenciasUsuariosInt objRemoto;

    public FachadaGestorUsuariosIml(ControladorPreferenciasUsuariosInt objRemoto) {
       
        this.objRemoto = objRemoto;
    }  
    
     public PreferenciasDTORespuesta getReferencias(String usuario) throws RemoteException {
        return this.objRemoto.getPreferenciasUsuario(usuario);
    }  
}

