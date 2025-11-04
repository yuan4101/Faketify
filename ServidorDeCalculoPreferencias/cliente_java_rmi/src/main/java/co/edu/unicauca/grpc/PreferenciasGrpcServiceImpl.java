package co.edu.unicauca.grpc;

import io.grpc.stub.StreamObserver;

import java.rmi.RemoteException;

import co.edu.unicauca.capaDeControladores.ControladorPreferenciasUsuariosInt;

import co.edu.unicauca.configuracion.servicios.ClienteDeObjetos;

import co.edu.unicauca.configuracion.lector.LectorPropiedadesConfig;

import co.edu.unicauca.fachadaServices.DTO.PreferenciasDTORespuesta;

import co.edu.unicauca.fachadaServices.DTO.PreferenciaGeneroDTORespuesta;

import co.edu.unicauca.fachadaServices.DTO.PreferenciaArtistaDTORespuesta;

import co.edu.unicauca.fachadaServices.DTO.PreferenciaIdiomaDTORespuesta;

public class PreferenciasGrpcServiceImpl extends PreferenciasServiceGrpc.PreferenciasServiceImplBase {

    private ControladorPreferenciasUsuariosInt objRemoto;

    public PreferenciasGrpcServiceImpl() {
        try {
            int puertoNS = Integer.parseInt(LectorPropiedadesConfig.get("ns.port"));
            String direccionIPNS = LectorPropiedadesConfig.get("ns.host");
            String identificadorObjetoRemoto = "objControladorPreferenciasUsuarios";

            this.objRemoto = ClienteDeObjetos.obtenerObjetoRemoto(
                direccionIPNS,
                puertoNS,
                identificadorObjetoRemoto
            );
            System.out.println("✓ gRPC conectado exitosamente con servidor RMI");
        } catch (Exception e) {
            System.err.println("✗ Error conectando con servidor RMI: " + e.getMessage());
            e.printStackTrace();
        }
    }

    @Override
    public void obtenerPreferenciasUsuario(
            Preferencias.PreferenciaRequest request,
            StreamObserver<Preferencias.PreferenciaResponse> responseObserver) {
        try {
            String nombreUsuario = request.getNombreUsuario();
            System.out.println("[gRPC] Solicitud recibida para: " + nombreUsuario);

            PreferenciasDTORespuesta preferencias = objRemoto.getPreferenciasUsuario(nombreUsuario);

            Preferencias.PreferenciaResponse.Builder responseBuilder =
                Preferencias.PreferenciaResponse.newBuilder();

            responseBuilder.setIdUsuario(preferencias.getIdUsuario());

            // Agregar géneros
            if (preferencias.getPreferenciasGeneros() != null) {
                for (PreferenciaGeneroDTORespuesta genero : preferencias.getPreferenciasGeneros()) {
                    responseBuilder.addPreferenciasGeneros(
                        Preferencias.Genero.newBuilder()
                            .setNombreGenero(genero.getNombreGenero())
                            .setNumeroPreferencias(genero.getNumeroPreferencias())
                            .build()
                    );
                }
            }

            // Agregar artistas
            if (preferencias.getPreferenciasArtistas() != null) {
                for (PreferenciaArtistaDTORespuesta artista : preferencias.getPreferenciasArtistas()) {
                    responseBuilder.addPreferenciasArtistas(
                        Preferencias.Artista.newBuilder()
                            .setNombreArtista(artista.getNombreArtista())
                            .setNumeroPreferencias(artista.getNumeroPreferencias())
                            .build()
                    );
                }
            }

            // NUEVO: Agregar idiomas
            if (preferencias.getPreferenciasIdiomas() != null) {
                for (PreferenciaIdiomaDTORespuesta idioma : preferencias.getPreferenciasIdiomas()) {
                    responseBuilder.addPreferenciasIdiomas(
                        Preferencias.Idioma.newBuilder()
                            .setNombreIdioma(idioma.getNombreIdioma())
                            .setNumeroPreferencias(idioma.getNumeroPreferencias())
                            .build()
                    );
                }
            }

            responseObserver.onNext(responseBuilder.build());
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
