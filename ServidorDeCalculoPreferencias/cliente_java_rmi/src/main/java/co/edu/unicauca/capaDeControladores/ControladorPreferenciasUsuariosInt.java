package co.edu.unicauca.capaDeControladores;

import java.rmi.Remote;
import java.rmi.RemoteException;
import co.edu.unicauca.fachadaServices.DTO.PreferenciasDTORespuesta;

//Hereda de la clase remore, lo cual lo convierte en interfaz remota
public interface ControladorPreferenciasUsuariosInt extends Remote {
    //Definicion primer metodo remoto
    public PreferenciasDTORespuesta getPreferenciasUsuario(String id) throws RemoteException;
    //cada definicion del metodo debe especificar que puede lanzar la excepcion java.rmi.remoteException
}