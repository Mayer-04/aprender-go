package main

import (
	"fmt"
	"time"
)

/*
* Paquete time en Go
El paquete "time" proporciona funcionalidades para manejar fechas, horas, duraciones y temporizadores.

- time.Now(): Obtiene la fecha y hora actual.
- time.Add(d time.Duration): Suma una duración a una fecha y hora.
- time.Sub(t time.Time): Calcula la diferencia entre dos instancias de tiempo.
- time.Format(layout string): Formatea la fecha y hora a un string.
- time.ParseDuration(s string): Convierte un string a un tipo `time.Duration``.
- time.Sleep(d time.Duration): Pausa la ejecución durante una duración específica.
- Evitar llamadas a time.After en funciones repetidas como bucles o controladores HTTP.

* Una fecha límite se refiere a un punto específico en el tiempo determinado con uno de los siguientes:
- time.Duration desde ahora (por ejemplo, en 250 ms). Es un alias para int64.
- time.Time (por ejemplo, 2023-02-07 00:00:00 UTC).

* Constantes

- time.Local: Representa la zona horaria configurada en la máquina donde se ejecuta el código
- time.UTC: Representa el tiempo universal coordinado
- time.Nanosecond: 1
- time.Microsecond: 1000 * Nanosecond
- time.Millisecond: 1000 * Microsecond
- time.Second: 1000 * Millisecond (estándar)
- time.Minute: 60 * Second (10 * time.Minute)
- time.Hour: 60 * Minute (24 * time.Hour)
- time.RFC3339: Es el estándar para JSON, APIs REST y fechas en internet. Ej: 2006-01-02T15:04:05Z07:00
- time.Kitchen: Mostrar horas sencillas en interfaces de usuario. Ej: 3:04PM
- Meses: time.January, time.February, ..., time.December
- Días: time.Sunday, time.Monday, ..., time.Saturday
*/

func main() {
	// Obtener la fecha y hora actual en el que se ejecuta el programa
	now := time.Now()
	fmt.Println("Fecha y hora actual:", now) // 2026-01-08 23:15:36.010874459 -0500 -05 m=+0.000017307

	// Sumar 2 horas a la fecha y hora actual
	// Si la hora actual es 5 pm y le sumamos 2 horas, la hora resultante es 7 pm
	future := now.Add(2 * time.Hour)
	fmt.Println("Fecha y hora después de 2 horas:", future) // 2026-01-09 01:15:36.010874459 -0500 -05 m=+7200.000017307

	// Calcular la diferencia de tiempo entre dos instancias
	difference := future.Sub(now)
	fmt.Println("Diferencia entre future y now:", difference) // 2h0m0s

	// Formatear la fecha y hora actual en un string
	formattedTime := now.Format("2006-01-02 15:04:05")
	fmt.Println("Fecha y hora formateada:", formattedTime) // 2026-01-08 23:15:36

	// Parsear un string a un tipo time.Duration
	duration, err := time.ParseDuration("2h45m")
	if err != nil {
		panic(err) // Manejar el error adecuadamente en un contexto real
	}
	fmt.Println("Duración parseada:", duration) // Duración parseada: 2h45m0s

	// Obtener el número de horas de una duración
	hours := duration.Hours()
	fmt.Printf("Horas en la duración: %.2f\n", hours) // Output: 2.75

	// Obtener el número de minutos de una duración
	minutes := duration.Minutes()
	fmt.Printf("Minutos en la duración: %.2f\n", minutes) // Output: 165.00

	// Obtener el número de segundos de una duración
	seconds := duration.Seconds()
	fmt.Printf("Segundos en la duración: %.2f\n", seconds) // Output: 9900.00

	// Pausar la ejecución del programa por 2 segundos
	// Una duración negativa o cero hace que `Sleep()` regrese inmediatamente
	fmt.Println("Esperando 2 segundos...")
	time.Sleep(2 * time.Second)
	fmt.Println("Continuando la ejecución después de la pausa.")

	// Crear una fecha específica (1 de enero de 2024 a las 10:00 AM)
	specificDate := time.Date(2024, time.January, 1, 10, 0, 0, 0, time.UTC)
	fmt.Println("Fecha específica:", specificDate) // Fecha específica: 2024-01-01 10:00:00 +0000 UTC

	// Comparar dos fechas after (después) o before (antes)
	if future.After(specificDate) {
		fmt.Println("El tiempo futuro es después de la fecha específica.") // Se imprime este Println
	} else if future.Before(specificDate) {
		fmt.Println("El tiempo futuro es antes de la fecha específica.")
	} else {
		fmt.Println("El tiempo futuro es igual a la fecha específica.")
	}

	// Usar un temporizador para realizar una acción después de un tiempo
	timer := time.NewTimer(3 * time.Second)
	fmt.Println("Temporizador de 3 segundos iniciado...")
	<-timer.C
	fmt.Println("Temporizador finalizado.")

	// time.Since() regresa el tiempo 'transcurrido' desde una fecha de inicio
	// Es un alias de time.Now().Sub(t).
	elapsed := time.Since(now)
	fmt.Printf("¡Listo! Tomó %v\n", elapsed) // ¡Listo! Tomó 5.0185089s

	_ = time.Unix(0, 0)

	// convierte una cadena de texto (string) que representa una fecha y hora en un valor del tipo time.Time,
	// siempre que esa cadena siga un formato específico. Espera dos argumentos:
	// 1. layout: El formato en el que esperas la fecha/hora
	// 2. value: La cadena de texto con la fecha/hora que quieres convertir
	t, _ := time.Parse("2006-01-02 15:04", "2025-05-27 18:30")
	fmt.Println("Parse()", t) // Output: 2025-05-27 18:30:00 +0000 UTC

	// Ejemplo 1: Fecha simple
	t2, _ := time.Parse("2006-01-02", "2025-05-27")
	fmt.Println(t2) // Output: 2025-05-27 00:00:00 +0000 UTC

	// Ejemplo 3: Fecha con hora y segundos
	t3, _ := time.Parse("02/01/2006 15:04:05", "27/05/2025 14:05:30")
	fmt.Println(t3) // Output: 2025-05-27 14:05:30 +0000 UTC

	//* time.Now().Format() Convierte la hora actual a string con formato personalizado.

	// Ejemplo 1: Fecha actual en formato "año-mes-día"
	fmt.Println(time.Now().Format("2006-01-02")) // Output: 2026-01-08

	// Ejemplo 2: Fecha y hora con minutos
	fmt.Println(time.Now().Format("02/01/2006 15:04")) // Output: 08/01/2026 23:15

	// Ejemplo 3: Solo hora con AM/PM
	fmt.Println(time.Now().Format("03:04 PM")) // Output: 11:15 PM

	// Extraer hora, minuto y segundo individualmente
	h, m, s := now.Clock()
	// Reloj Desglosado: 23 horas, 19 minutos, 36 segundos
	fmt.Printf("Reloj Desglosado: %02d horas, %02d minutos, %02d segundos\n", h, m, s)

	// Extraer componentes de fecha
	year, month, day := now.Date()
	// Calendario: Día 8 del mes January de 2026
	fmt.Printf("Calendario: Día %d del mes %v de %d\n", day, month, year)

	// Zona Horaria Local
	// Usamos time.Local para decirle a Go que interprete la hora con la zona del sistema
	tLocal, err := time.ParseInLocation("2006-01-02 15:04", "2025-05-27 18:30", time.Local)
	if err != nil {
		fmt.Println("Error parseando:", err)
	} else {
		fmt.Println("Parse Local:", tLocal) // Output: 2025-05-27 18:30:00 -0500 -05
	}

	// Constantes
	fmt.Println(time.UTC)     // Output: UTC
	fmt.Println(time.Local)   // Output: Local
	fmt.Println(time.RFC3339) // Output: 2006-01-02T15:04:05Z07:00
	fmt.Println(time.Kitchen) // Output: 3:04PM

	// 1. Obtengo la FECHA
	// 2. La muevo a la ZONA (Objeto time.UTC)
	// 3. Le aplico el FORMATO (String time.Kitchen)
	now2 := time.Now().In(time.UTC).Format(time.Kitchen)
	fmt.Println("now2", now2) // Output: now2 4:36AM
}
