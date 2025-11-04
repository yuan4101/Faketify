package co.edu.unicauca.infoii.correo.componenteRecibirMensajes;

import org.springframework.stereotype.Service;
import co.edu.unicauca.infoii.correo.DTOs.CancionAlmacenarDTOInput;
import co.edu.unicauca.infoii.correo.commons.Simulacion;
import org.springframework.amqp.rabbit.annotation.RabbitListener;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;

@Service
public class MessageConsumer {

    private String[] fraseMotivadoras = {
        "¡La música es la poesía del aire! 🎵",
        "Cada nota es una puerta hacia la emoción. 🎶",
        "La música nos conecta a todos. 🎼",
        "Vive la música, vive la vida. 🎸",
        "La mejor medicina es la música. 💫",
        "Donde hay música, hay esperanza. ✨",
        "Cada canción cuenta una historia. 📖",
        "La música es el lenguaje universal. 🌍"
    };

    @RabbitListener(queues = "notificaciones_canciones")
    public void receiveMessage(CancionAlmacenarDTOInput objCancion) {
        System.out.println("\n📨 ========== SERVIDOR DE CORREOS ==========");
        System.out.println("Datos de la canción recibidos:");
        System.out.println("  Título: " + objCancion.getTitulo());
        System.out.println("  Artista: " + objCancion.getArtista());
        System.out.println("  Género: " + objCancion.getGenero());
        
        // Simular envío de correo
        System.out.println("\n⏳ Enviando correo electrónico...");
        Simulacion.simular(3000, "Procesando");
        
        // Obtener fecha y hora actual
        LocalDateTime ahora = LocalDateTime.now();
        DateTimeFormatter formatter = DateTimeFormatter.ofPattern("dd/MM/yyyy HH:mm:ss");
        String fechaHora = ahora.format(formatter);
        
        // Obtener frase motivadora aleatoria
        String frase = fraseMotivadoras[(int) (Math.random() * fraseMotivadoras.length)];
        
        // Mostrar contenido del correo
        System.out.println("\n📧 ========== CONTENIDO DEL CORREO ==========");
        System.out.println("Para: admin@faketify.com");
        System.out.println("Asunto: 🎵 Nueva Canción Registrada - " + objCancion.getTitulo());
        System.out.println("─────────────────────────────────────────");
        System.out.println("¡Hola Faketify!");
        System.out.println("\nUna nueva canción ha sido registrada en el servidor.");
        System.out.println("Aquí están los detalles:\n");
        System.out.println("📋 METADATOS DE LA CANCIÓN:");
        System.out.println("  • Título: " + objCancion.getTitulo());
        System.out.println("  • Artista: " + objCancion.getArtista());
        System.out.println("  • Género: " + objCancion.getGenero());
        System.out.println("  • Año: " + objCancion.getAño());
        System.out.println("  • Idioma: " + objCancion.getIdioma());
        System.out.println("  • Duración: " + objCancion.getDuracion());
        System.out.println("\n📅 Fecha y Hora del Registro: " + fechaHora);
        System.out.println("\n💡 Frase Motivadora: " + frase);
        System.out.println("\n¡Sigue disfrutando de la música! 🎶");
        System.out.println("─────────────────────────────────────────");
        
        System.out.println("\n✅ Correo enviado exitosamente");
        System.out.println("🎵 =========================================\n");
    }
}
