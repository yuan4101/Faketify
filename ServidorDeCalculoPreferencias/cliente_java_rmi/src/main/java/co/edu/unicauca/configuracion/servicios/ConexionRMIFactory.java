package co.edu.unicauca.configuracion.servicios;

import co.edu.unicauca.capaDeControladores.ControladorPreferenciasUsuariosInt;
import co.edu.unicauca.configuracion.lector.LectorPropiedadesConfig;

public class ConexionRMIFactory {

    public static ControladorPreferenciasUsuariosInt crearConexionRMI() throws Exception {
        try {
            int puertoNS = Integer.parseInt(LectorPropiedadesConfig.get("ns.port"));
            String direccionIPNS = LectorPropiedadesConfig.get("ns.host");
            String identificadorObjetoRemoto = "objControladorPreferenciasUsuarios";

            ControladorPreferenciasUsuariosInt objRemoto = ClienteDeObjetos.obtenerObjetoRemoto(
                direccionIPNS,
                puertoNS,
                identificadorObjetoRemoto
            );

            if (objRemoto == null) {
                throw new Exception("No se pudo obtener la referencia RMI del objeto remoto");
            }

            System.out.println("-> gRPC conectado exitosamente con servidor RMI");
            return objRemoto;

        } catch (Exception e) {
            System.err.println("-> Error conectando con servidor RMI: " + e.getMessage());
            e.printStackTrace();
            throw new Exception("No se pudo inicializar la conexión RMI", e);
        }
    }
}
