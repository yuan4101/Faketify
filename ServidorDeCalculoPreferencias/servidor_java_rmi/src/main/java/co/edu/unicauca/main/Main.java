package co.edu.unicauca.main;

import co.edu.unicauca.capaDeControladores.ControladorPreferenciasUsuariosIml;
import co.edu.unicauca.configuracion.lector.LectorPropiedadesConfig;
import co.edu.unicauca.configuracion.servicios.ServidorDeObjetos;

public class Main {
    public static void main(String[] args) {
        int puertoNS = Integer.parseInt(LectorPropiedadesConfig.get("ns.port"));
        String direccionIPNS = LectorPropiedadesConfig.get("ns.host");

        ServidorDeObjetos.arrancarNS(direccionIPNS, puertoNS);
        ControladorPreferenciasUsuariosIml objControladorPreferencias = ServidorDeObjetos.crearObjetoRemoto();
        String identificadorObjetoRemoto = "objControladorPreferenciasUsuarios";
        ServidorDeObjetos.registrarObjetoRemoto(objControladorPreferencias, direccionIPNS, puertoNS, identificadorObjetoRemoto);
        
        System.out.printf("\n\t\t----- SERVIDOR CALCULO DE PREFERENCIAS (Java RMI) [%s:%s] -----\n",direccionIPNS, puertoNS);
        System.out.println("-> Objeto remoto registrado como: " + identificadorObjetoRemoto);
    }
}
