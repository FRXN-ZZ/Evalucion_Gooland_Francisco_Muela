package main

import "fmt"


var nombresProductos []string
var subtotales []float64


func RegistrarVenta(nombre string, precio float64, cantidad int) {
	subtotal := precio * float64(cantidad)
	
	
	nombresProductos = append(nombresProductos, nombre)
	subtotales = append(subtotales, subtotal)
	
	fmt.Println("¡Venta registrada con éxito! Subtotal:", subtotal)
}


func MostrarEstadisticas() {
	if len(nombresProductos) == 0 {
		fmt.Println("No existen ventas registradas.")
		return
	}

	totalRecaudado := 0.0
	fmt.Println("\n--- ESTADÍSTICAS DE VENTAS ---")
	for i := 0; i < len(nombresProductos); i++ {
		fmt.Println("Producto:", nombresProductos[i], "| Subtotal:", subtotales[i])
		totalRecaudado += subtotales[i]
	}
	fmt.Println("Total recaudado:", totalRecaudado)
}


func main() {
	var opcion int

	for {
	
		fmt.Println("\n=== MENÚ ===")
		fmt.Println("1. Registrar una nueva venta")
		fmt.Println("2. Mostrar estadísticas")
		fmt.Println("3. Salir")
		fmt.Print("Elija una opción: ")
		
		fmt.Scan(&opcion)

		if opcion == 1 {
			
			fmt.Println("\n--- PRODUCTOS DISPONIBLES ---")
			fmt.Println("1. Arroz ($1.25)")
			fmt.Println("2. Leche ($0.95)")
			fmt.Println("3. Cacerola ($0.50)")
			
			var seleccion int
			fmt.Print("Seleccione el número del producto: ")
			fmt.Scan(&seleccion)

			var cantidad int
			fmt.Print("Ingrese la cantidad vendida: ")
			fmt.Scan(&cantidad)

			
			if seleccion == 1 {
				RegistrarVenta("Arroz", 1.25, cantidad)
			} else if seleccion == 2 {
				RegistrarVenta("Leche", 0.95, cantidad)
			} else if seleccion == 3 {
				RegistrarVenta("Cacerola", 0.50, cantidad)
			} else {
				fmt.Println("Número de producto inválido.")
			}

		} else if opcion == 2 {
			MostrarEstadisticas()

		} else if opcion == 3 {
			fmt.Println("Saliendo del programa...")
			break

		} else {
			fmt.Println("Opción no válida. Intente de nuevo.")
		}
	}
}