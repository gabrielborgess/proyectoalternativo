/*
	Un CRUD completo de GoLang y MySQL
	@author parzibyte
*/
package main

import (
	"./base"
	"bufio"                            // Leer líneas incluso si tienen espacios
	"database/sql"                     // Interactuar con bases de datos
	"fmt"                              // Imprimir mensajes y esas cosas
	_ "github.com/go-sql-driver/mysql" // La librería que nos permite conectar a MySQL
	"os"                               // El búfer, para leer desde la terminal con os.Stdin
)

type Videojuego struct {
	Nombre, Genero, Precio string
	Id                     int
}

func obtenerBaseDeDatos() (db *sql.DB, e error) {
	usuario := "admin_admin"
	pass := "ganzo10."
	host := "tcp(158.69.60.190:3306)"
	nombreBaseDeDatos := "admin_proyecto"
	// Debe tener la forma usuario:contraseña@protocolo(host:puerto)/nombreBaseDeDatos
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	var setup string
	fmt.Println("Bienvenido, Desea Instalar la base de datos? S/n")
	fmt.Scanf("%s", &setup)
	if setup == "s" {
		base.Base()
	}
	menu := `¿Qué deseas hacer?
[1] -- Insertar
[2] -- Mostrar
[3] -- Actualizar
[4] -- Eliminar
[5] -- Salir
----->	`
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
			err := insertar(c)
			if err != nil {
				fmt.Printf("Error insertando: %v", err)
			} else {
				fmt.Println("Insertado correctamente")
			}
		case 2:
			contactos, err := obtenerContactos()
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
			err := actualizar(c)
			if err != nil {
				fmt.Printf("Error actualizando: %v", err)
			} else {
				fmt.Println("Actualizado correctamente")
			}
		case 4:
			fmt.Println("Ingresa el ID del contacto que deseas eliminar:")
			fmt.Scanln(&c.Id)
			err := eliminar(c)
			if err != nil {
				fmt.Printf("Error eliminando: %v", err)
			} else {
				fmt.Println("Eliminado correctamente")
			}
		}
	}
}

func eliminar(c Videojuego) error {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	sentenciaPreparada, err := db.Prepare("DELETE FROM administrador WHERE id = ?")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()

	_, err = sentenciaPreparada.Exec(c.Id)
	if err != nil {
		return err
	}
	return nil
}

func insertar(c Videojuego) (e error) {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	// Preparamos para prevenir inyecciones SQL
	sentenciaPreparada, err := db.Prepare("INSERT INTO administrador (nombre, genero, precio) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()
	// Ejecutar sentencia, un valor por cada '?'
	_, err = sentenciaPreparada.Exec(c.Nombre, c.Genero, c.Precio)
	if err != nil {
		return err
	}
	return nil
}

func obtenerContactos() ([]Videojuego, error) {
	contactos := []Videojuego{}
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	filas, err := db.Query("SELECT id, nombre, genero, precio FROM administrador")

	if err != nil {
		return nil, err
	}
	// Si llegamos aquí, significa que no ocurrió ningún error
	defer filas.Close()

	// Aquí vamos a "mapear" lo que traiga la consulta en el while de más abajo
	var c Videojuego

	// Recorrer todas las filas, en un "while"
	for filas.Next() {
		err = filas.Scan(&c.Id, &c.Nombre, &c.Genero, &c.Precio)
		// Al escanear puede haber un error
		if err != nil {
			return nil, err
		}
		// Y si no, entonces agregamos lo leído al arreglo
		contactos = append(contactos, c)
	}
	// Vacío o no, regresamos el arreglo de contactos
	return contactos, nil
}

func actualizar(c Videojuego) error {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	sentenciaPreparada, err := db.Prepare("UPDATE administrador SET nombre = ?, genero = ?, precio = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()
	// Pasar argumentos en el mismo orden que la consulta
	_, err = sentenciaPreparada.Exec(c.Nombre, c.Genero, c.Precio, c.Id)
	return err // Ya sea nil o sea un error, lo manejaremos desde donde hacemos la llamada
}
