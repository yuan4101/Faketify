package co.edu.unicauca.grpc;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.50.0)",
    comments = "Source: preferencias.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class PreferenciasServiceGrpc {

  private PreferenciasServiceGrpc() {}

  public static final String SERVICE_NAME = "co.edu.unicauca.grpc.PreferenciasService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<co.edu.unicauca.grpc.Preferencias.PreferenciaRequest,
      co.edu.unicauca.grpc.Preferencias.PreferenciaResponse> getObtenerPreferenciasUsuarioMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ObtenerPreferenciasUsuario",
      requestType = co.edu.unicauca.grpc.Preferencias.PreferenciaRequest.class,
      responseType = co.edu.unicauca.grpc.Preferencias.PreferenciaResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<co.edu.unicauca.grpc.Preferencias.PreferenciaRequest,
      co.edu.unicauca.grpc.Preferencias.PreferenciaResponse> getObtenerPreferenciasUsuarioMethod() {
    io.grpc.MethodDescriptor<co.edu.unicauca.grpc.Preferencias.PreferenciaRequest, co.edu.unicauca.grpc.Preferencias.PreferenciaResponse> getObtenerPreferenciasUsuarioMethod;
    if ((getObtenerPreferenciasUsuarioMethod = PreferenciasServiceGrpc.getObtenerPreferenciasUsuarioMethod) == null) {
      synchronized (PreferenciasServiceGrpc.class) {
        if ((getObtenerPreferenciasUsuarioMethod = PreferenciasServiceGrpc.getObtenerPreferenciasUsuarioMethod) == null) {
          PreferenciasServiceGrpc.getObtenerPreferenciasUsuarioMethod = getObtenerPreferenciasUsuarioMethod =
              io.grpc.MethodDescriptor.<co.edu.unicauca.grpc.Preferencias.PreferenciaRequest, co.edu.unicauca.grpc.Preferencias.PreferenciaResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ObtenerPreferenciasUsuario"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  co.edu.unicauca.grpc.Preferencias.PreferenciaRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  co.edu.unicauca.grpc.Preferencias.PreferenciaResponse.getDefaultInstance()))
              .setSchemaDescriptor(new PreferenciasServiceMethodDescriptorSupplier("ObtenerPreferenciasUsuario"))
              .build();
        }
      }
    }
    return getObtenerPreferenciasUsuarioMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static PreferenciasServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PreferenciasServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PreferenciasServiceStub>() {
        @java.lang.Override
        public PreferenciasServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PreferenciasServiceStub(channel, callOptions);
        }
      };
    return PreferenciasServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static PreferenciasServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PreferenciasServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PreferenciasServiceBlockingStub>() {
        @java.lang.Override
        public PreferenciasServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PreferenciasServiceBlockingStub(channel, callOptions);
        }
      };
    return PreferenciasServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static PreferenciasServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PreferenciasServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PreferenciasServiceFutureStub>() {
        @java.lang.Override
        public PreferenciasServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PreferenciasServiceFutureStub(channel, callOptions);
        }
      };
    return PreferenciasServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class PreferenciasServiceImplBase implements io.grpc.BindableService {

    /**
     */
    public void obtenerPreferenciasUsuario(co.edu.unicauca.grpc.Preferencias.PreferenciaRequest request,
        io.grpc.stub.StreamObserver<co.edu.unicauca.grpc.Preferencias.PreferenciaResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getObtenerPreferenciasUsuarioMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getObtenerPreferenciasUsuarioMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                co.edu.unicauca.grpc.Preferencias.PreferenciaRequest,
                co.edu.unicauca.grpc.Preferencias.PreferenciaResponse>(
                  this, METHODID_OBTENER_PREFERENCIAS_USUARIO)))
          .build();
    }
  }

  /**
   */
  public static final class PreferenciasServiceStub extends io.grpc.stub.AbstractAsyncStub<PreferenciasServiceStub> {
    private PreferenciasServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PreferenciasServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PreferenciasServiceStub(channel, callOptions);
    }

    /**
     */
    public void obtenerPreferenciasUsuario(co.edu.unicauca.grpc.Preferencias.PreferenciaRequest request,
        io.grpc.stub.StreamObserver<co.edu.unicauca.grpc.Preferencias.PreferenciaResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getObtenerPreferenciasUsuarioMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class PreferenciasServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<PreferenciasServiceBlockingStub> {
    private PreferenciasServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PreferenciasServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PreferenciasServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public co.edu.unicauca.grpc.Preferencias.PreferenciaResponse obtenerPreferenciasUsuario(co.edu.unicauca.grpc.Preferencias.PreferenciaRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getObtenerPreferenciasUsuarioMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class PreferenciasServiceFutureStub extends io.grpc.stub.AbstractFutureStub<PreferenciasServiceFutureStub> {
    private PreferenciasServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PreferenciasServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PreferenciasServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<co.edu.unicauca.grpc.Preferencias.PreferenciaResponse> obtenerPreferenciasUsuario(
        co.edu.unicauca.grpc.Preferencias.PreferenciaRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getObtenerPreferenciasUsuarioMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_OBTENER_PREFERENCIAS_USUARIO = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final PreferenciasServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(PreferenciasServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_OBTENER_PREFERENCIAS_USUARIO:
          serviceImpl.obtenerPreferenciasUsuario((co.edu.unicauca.grpc.Preferencias.PreferenciaRequest) request,
              (io.grpc.stub.StreamObserver<co.edu.unicauca.grpc.Preferencias.PreferenciaResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class PreferenciasServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    PreferenciasServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return co.edu.unicauca.grpc.Preferencias.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("PreferenciasService");
    }
  }

  private static final class PreferenciasServiceFileDescriptorSupplier
      extends PreferenciasServiceBaseDescriptorSupplier {
    PreferenciasServiceFileDescriptorSupplier() {}
  }

  private static final class PreferenciasServiceMethodDescriptorSupplier
      extends PreferenciasServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    PreferenciasServiceMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (PreferenciasServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new PreferenciasServiceFileDescriptorSupplier())
              .addMethod(getObtenerPreferenciasUsuarioMethod())
              .build();
        }
      }
    }
    return result;
  }
}
