Introducción a goroutines y channels (ejemplos)

Objetivo
- Entender por qué se produce el `deadlock` en el ejercicio original y aprender patrones seguros con channels.

Puntos clave
- Canales sin buffer (unbuffered) sincronizan inmediatamente el envío y la recepción: la operación de envío bloqueará hasta que exista una goroutine receptora esperando, y la recepción bloqueará hasta que exista un emisor.
- Canales con buffer permiten enviar hasta N mensajes sin receptor inmediato. Si envías más que el buffer, el envío bloqueará.
- Cerrar un canal (`close(ch)`) es una forma habitual de señalar "no habrá más valores"; todas las recepciones posteriores devuelven el valor cero sin bloquear y un range sobre el canal terminará.
- Un `send` entrega el valor a un único receptor (no lo distribuye a todos). Si hay múltiples receptores y sólo un send, sólo uno recibirá y los demás seguirán esperando.

Ejemplos incluidos
- `examples/unbuffered_fix.go`: muestra usar `close(ch)` para señalizar y evitar bucles de recepciones que causen bloqueo.
- `examples/buffered_channel.go`: muestra cómo un canal con buffer permite acumular mensajes antes de que haya receptor.
- `examples/close_and_range.go`: muestra el patrón `for range ch` y cerrar el canal cuando se termina de producir.
- `examples/multi_receivers.go`: muestra comportamiento con múltiples receivers y cómo `close` puede ayudar a que todos terminen.

Consejos rápidos
- Si necesitas enviar una "señal" (sin dato) y solo esperas un receptor, un canal sin buffer con un solo envío o usar `close` son suficientes.
- Para mensajes múltiples o desacoplar emisores/receptores, considera canales con buffer.
- Para notificaciones de termino a múltiples goroutines, usa `close` (o un `sync.WaitGroup` combinándolo con canales cerrados si necesitas más control).
