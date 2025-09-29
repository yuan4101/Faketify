# Faketify - Sistema de Streaming Musical

Servicio de streaming musical que permite explorar géneros, buscar canciones y reproducir audio en tiempo real via gRPC.

## 🎵 Características

- **Catálogo Musical**: 10 géneros y 20 canciones de ejemplo
- **Búsqueda**: Por título y filtrado por género  
- **Streaming**: Reproducción en tiempo real con fragmentos de 64KB
- **Interfaz**: Consola interactiva con colores y formatos
- **Protocolo**: Comunicación gRPC para alta eficiencia

## 🏗️ Arquitectura

### Servicios gRPC

#### SongService
- `GetGenres()` - Lista todos los géneros musicales
- `GetSongsByGenre()` - Canciones filtradas por género
- `GetSong()` - Metadatos de canción específica

#### AudioService  
- `GetStreamingSong()` - Stream de audio en tiempo real

### Componentes Principales

- **Servidor**: Maneja catálogo y streaming
- **Cliente**: Interfaz de usuario y reproducción
- **Common**: Utilidades y modelos compartidos

## 🚀 Uso

1. **Iniciar Servidor de Streaming**:
```bash
cd ServidorDeStreaming/Views/
go run streamingServer.go
```

2. **Iniciar Servidor de Canciones**:
```bash
cd ServidorDeCanciones/Views/
go run songsServer.go
```

3. **Ejecutar Cliente**:
```bash
cd Cliente/main/
go run client.go
```

3. **Navegación**:
   - Ver géneros musicales
   - Explorar canciones por género
   - Reproducir canciones en streaming
   - Controlar reproducción en tiempo real

## 🔧 Tecnologías

- **Go** - Lenguaje principal
- **gRPC** - Comunicación entre servicios
- **Protocol Buffers** - Serialización de datos
- **BEEP** - Reproducción de audio

## 📊 Protocolo de Streaming

- Fragmentación en chunks de 64KB
- Transmisión secuencial via gRPC streams
- Sincronización con canales Go
- Decodificación MP3 en tiempo real

## 🎨 Interfaz

Menús interactivos con:
- Colores ANSI para mejor legibilidad
- Navegación jerárquica intuitiva
- Controles de reproducción en tiempo real
- Feedback visual del estado del sistema