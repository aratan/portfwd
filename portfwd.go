package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
)

// handleConnection se encarga de gestionar una conexión entrante y redirigirla.
// Inicia dos goroutines para copiar datos en ambas direcciones simultáneamente.
func handleConnection(clientConn net.Conn, remoteAddr string) {
	// Aseguramos que la conexión del cliente se cierre al finalizar la función.
	defer clientConn.Close()

	log.Printf("[+] Conexión aceptada de %s", clientConn.RemoteAddr())

	// Intentamos conectar al destino remoto.
	remoteConn, err := net.Dial("tcp", remoteAddr)
	if err != nil {
		log.Printf("[!] No se pudo conectar al destino remoto %s: %v", remoteAddr, err)
		return
	}
	// Aseguramos que la conexión remota se cierre al finalizar la función.
	defer remoteConn.Close()

	log.Printf("[+] Redirigiendo %s <--> %s", clientConn.RemoteAddr(), remoteAddr)

	// Creamos un canal para señalizar cuando la copia de datos ha terminado.
	done := make(chan struct{})

	// Goroutine para copiar datos del cliente al remoto.
	go func() {
		// io.Copy es una función bloqueante que copia datos hasta EOF o un error.
		io.Copy(remoteConn, clientConn)
		done <- struct{}{} // Señaliza que ha terminado.
	}()

	// Goroutine para copiar datos del remoto al cliente.
	go func() {
		io.Copy(clientConn, remoteConn)
		done <- struct{}{} // Señaliza que ha terminado.
	}()

	// Esperamos a que cualquiera de las dos goroutines termine.
	// Al terminar una, la otra se desbloqueará (por cierre de socket) y terminará también.
	// Esperar a ambas señales asegura que todo el trabajo ha finalizado.
	<-done
	<-done

	log.Printf("[-] Conexión cerrada con %s", clientConn.RemoteAddr())
}

func main() {
	// Definimos los flags para la línea de comandos.
	// flag.Int devuelve un puntero a un entero.
	localPort := flag.Int("lport", 8001, "Puerto local en el que escuchar.")
	remoteHost := flag.String("rhost", "127.0.0.1", "Host remoto al que conectar.")
	remotePort := flag.Int("rport", 80, "Puerto remoto al que conectar.")

	// Parseamos los argumentos de la línea de comandos.
	flag.Parse()

	// Construimos las direcciones a partir de los flags.
	localAddr := fmt.Sprintf("0.0.0.0:%d", *localPort)
	remoteAddr := fmt.Sprintf("%s:%d", *remoteHost, *remotePort)

	// Iniciamos el listener en la dirección y puerto local especificados.
	listener, err := net.Listen("tcp", localAddr)
	if err != nil {
		// Si no podemos iniciar el listener, el programa no puede continuar.
		log.Fatalf("[!] Error al iniciar el listener en %s: %v", localAddr, err)
	}
	// Nos aseguramos de que el listener se cierre cuando main() termine.
	defer listener.Close()

	log.Printf("[+] Escuchando en %s, redirigiendo a %s", localAddr, remoteAddr)

	// Bucle infinito para aceptar nuevas conexiones.
	for {
		// Accept() es una llamada bloqueante que espera por una nueva conexión.
		conn, err := listener.Accept()
		if err != nil {
			// Si hay un error al aceptar, lo registramos y continuamos esperando.
			log.Printf("[!] Error al aceptar una conexión: %v", err)
			continue
		}

		// Para cada conexión, iniciamos una nueva goroutine para manejarla.
		// Esto permite que el bucle principal vuelva a esperar por más conexiones inmediatamente.
		go handleConnection(conn, remoteAddr)
	}
}
