package co.edu.unicauca.fachadaServices.services;

import io.grpc.stub.StreamObserver;
import java.rmi.RemoteException;

import co.edu.unicauca.capaDeControladores.ControladorPreferenciasUsuariosInt;
import co.edu.unicauca.configuracion.servicios.ConexionRMIFactory;
import co.edu.unicauca.fachadaServices.DTO.PreferenciasDTORespuesta;
import co.edu.unicauca.grpc.Preferencias;
import co.edu.unicauca.grpc.PreferenciasServiceGrpc;

public class PreferenciasGrpcServiceImpl extends PreferenciasServiceGrpc.PreferenciasServiceImplBase {

    private ControladorPreferenciasUsuariosInt objRemoto;

    public PreferenciasGrpcServiceImpl() throws Exception {
        this.objRemoto = ConexionRMIFactory.crearConexionRMI();
    }

    @Override
    public void obtenerPreferenciasUsuario(
            Preferencias.PreferenciaRequest request,
            StreamObserver<Preferencias.PreferenciaResponse> responseObserver) {
        try {
            String nombreUsuario = request.getNombreUsuario();
            System.out.println("[gRPC] Solicitud recibida para: " + nombreUsuario);

            PreferenciasDTORespuesta preferencias = objRemoto.getPreferenciasUsuario(nombreUsuario);

            Preferencias.PreferenciaResponse response = 
                PreferenciasMapper.convertirDTOaProtobuf(preferencias);

            responseObserver.onNext(response);
            responseObserver.onCompleted();
            System.out.println("[gRPC] Respuesta enviada exitosamente");

        } catch (RemoteException e) {
            System.err.println("[gRPC] Error en RMI: " + e.getMessage());
            responseObserver.onError(e);

        } catch (Exception e) {
            System.err.println("[gRPC] Error inesperado: " + e.getMessage());
            e.printStackTrace();
            responseObserver.onError(e);
        }
    }
}
