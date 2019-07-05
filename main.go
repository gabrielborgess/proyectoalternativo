package main

import (
	"./base"                           // El búfer, para leer desde la terminal con os.Stdin
	"bufio"                            // Leer líneas incluso si tienen espacios
	"fmt"                              // Imprimir mensajes y esas cosas
	_ "github.com/go-sql-driver/mysql" // La librería que nos permite conectar a MySQL
	"os"
)

type Videojuego struct {
	Nombre, Genero, Precio string
	Id                     int
}

func main() {
	menu := `¿Qué deseas hacer?
[1] -- Insertar
[2] -- Mostrar
[3] -- Actualizar
[4] -- Eliminar
[5] -- Salir
----->	`
	base.ObtenerBaseDeDatos()
	var eleccion int
	var c Videojuego

	for eleccion != 5 {
		fmt.Print(menu)
		fmt.Scanln(&eleccion)
		scanner := bufio.NewScanner(os.Stdin)
		switch eleccion {
		case 1:
			fmt.Println("Ingresa el nombre:")
			if scanner.Scan() {
				c.Nombre = scanner.Text()
			}
			fmt.Println("Genero:")
			if scanner.Scan() {
				c.Genero = scanner.Text()
			}
			fmt.Println("Precio:")
			if scanner.Scan() {
				c.Precio = scanner.Text()
			}
			err := base.Insertar(c)
			if err != nil {
				fmt.Printf("Error insertando: %v", err)
			} else {
				fmt.Println("Insertado correctamente")
			}
		case 2:
			contactos, err := base.ObtenerContactos()
			if err != nil {
				fmt.Printf("Error obteniendo contactos: %v", err)
			} else {
				for _, contacto := range contactos {
					fmt.Println("====================")
					fmt.Printf("Id: %d\n", contacto.Id)
					fmt.Printf("Nombre: %s\n", contacto.Nombre)
					fmt.Printf("Genero: %s\n", contacto.Genero)
					fmt.Printf("Precio: %s\n", contacto.Precio)
				}
			}
		case 3:
			fmt.Println("Ingresa el id:")
			fmt.Scanln(&c.Id)
			fmt.Println("Ingresa el nuevo nombre:")
			if scanner.Scan() {
				c.Nombre = scanner.Text()
			}
			fmt.Println("Ingresa la nueva dirección:")
			if scanner.Scan() {
				c.Genero = scanner.Text()
			}
			fmt.Println("Ingresa el nuevo correo electrónico:")
			if scanner.Scan() {
				c.Precio = scanner.Text()
			}
			err := base.Actualizar(c)
			if err != nil {
				fmt.Printf("Error actualizando: %v", err)
			} else {
				fmt.Println("Actualizado correctamente")
			}
		case 4:
			fmt.Println("Ingresa el ID del contacto que deseas eliminar:")
			fmt.Scanln(&c.Id)
			err := base.Eliminar(c)
			if err != nil {
				fmt.Printf("Error eliminando: %v", err)
			} else {
				fmt.Println("Eliminado correctamente")
			}
		}
	}
}
