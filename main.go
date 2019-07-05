/*
	Un CRUD completo de GoLang y MySQL
	@author parzibyte
*/
package main

import (
	"./base"
	"database/sql"                     // Interactuar con bases de datos
	"fmt"                              // Imprimir mensajes y esas cosas
	_ "github.com/go-sql-driver/mysql" // La librería que nos permite conectar a MySQL
)

//type Videojuego struct {
//	Nombre, Genero, Precio string
//	Id                     int
//}
// No se usa, esta en base.go
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
	fmt.Println("Bienvenido, Desea Instalar la base de datos? s/n")
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
	for eleccion != 5 {
		fmt.Print(menu)
		fmt.Scanln(&eleccion)
		switch eleccion {
		case 1:
			fmt.Println("Lista de opciones:" +
				"\n 1 Insertar en Tabla de Productos" +
				"\n 2 Insertar en Tabla de Generos" +
				"\n 3 Insertar en Tabla de Empleados" +
				"\n 4 Insertar en Tabla de Clientes" +
				"\n 5 Insertar en Tabla de Proveedores" +
				"\n 6 Insertar en Tabla de Pedidos" +
				"\n 7 Insertar en Tabla de Detalle de Pedidos")
			var opcion string
			fmt.Scanf("%s", &opcion)
			switch opcion {
			case "1":
				columnas := "Nombre,generos,valor,plataformas,proveedor_id,estrellas"
				valores := "%s,%s,%d,%s,%d,%d"
				estructura := base.Producto{}
				fmt.Println("Nombre:")
				fmt.Scanf("%s", &estructura.Nombre)
				fmt.Println("Generos:")
				fmt.Scanf("%s", &estructura.Generos)
				fmt.Println("Valor (Numero Entero)")
				fmt.Scanf("%d", &estructura.Valor)
				fmt.Println("Plataformas:")
				fmt.Scanf("%s", &estructura.Plataformas)
				fmt.Println("Proveedor_id (Numero Entero)")
				fmt.Scanf("%d", &estructura.Proveedor_id)
				fmt.Println("Estrellas (Numero Entero)")
				fmt.Scanf("%d", &estructura.Estrellas)
				base.Insertar_sql("productos", columnas, fmt.Sprintf(valores, estructura.Nombre, estructura.Generos, estructura.Valor, estructura.Plataformas, estructura.Proveedor_id, estructura.Estrellas))
				fmt.Println("Insertado con exito")
			case "2":
				columnas := "tipo"
				valores := "%s"
				estructura := base.Genero{}
				fmt.Println("Genero:")
				fmt.Scanf("%s", &estructura.Tipo)
				base.Insertar_sql("genero", columnas, fmt.Sprintf(valores, estructura.Tipo))
				fmt.Println("Insertado con exito")
			case "3":
				columnas := "Rut,Nombre,Sueldo,Area,Cargo,Direccion,Region"
				valores := "%s,%s,%d,%s,%s,%s,%s"
				estructura := base.Empleado{}
				fmt.Println("Rut:")
				fmt.Scanf("%s", &estructura.Rut)
				fmt.Println("Nombre:")
				fmt.Scanf("%s", &estructura.Nombre)
				fmt.Println("Sueldo (Numero Entero)")
				fmt.Scanf("%d", &estructura.Sueldo)
				fmt.Println("Area:")
				fmt.Scanf("%s", &estructura.Area)
				fmt.Println("Cargo")
				fmt.Scanf("%s", &estructura.Cargo)
				fmt.Println("Direccion")
				fmt.Scanf("%s", &estructura.Direccion)
				fmt.Println("Region")
				fmt.Scanf("%s", &estructura.Region)
				base.Insertar_sql("empleados", columnas, fmt.Sprintf(valores, estructura.Rut, estructura.Nombre, estructura.Sueldo, estructura.Area, estructura.Cargo, estructura.Direccion, estructura.Region))
				fmt.Println("Insertado con exito")
			case "4":
				columnas := "Rut,Nombre,Direccion,Region,telefono"
				valores := "%s,%s,%s,%s,%s"
				estructura := base.Clientes{}
				fmt.Println("Rut:")
				fmt.Scanf("%s", &estructura.Rut)
				fmt.Println("Nombre:")
				fmt.Scanf("%s", &estructura.Nombre)
				fmt.Println("Telefono:")
				fmt.Scanf("%s", &estructura.Telefono)
				fmt.Println("Direccion")
				fmt.Scanf("%s", &estructura.Direccion)
				fmt.Println("Region")
				fmt.Scanf("%s", &estructura.Region)
				base.Insertar_sql("clientes", columnas, fmt.Sprintf(valores, estructura.Rut, estructura.Nombre, estructura.Direccion, estructura.Region, estructura.Telefono))
				fmt.Println("Insertado con exito")
			case "5":
				columnas := "Nombre,Direccion,telefono"
				valores := "%s,%s,%s"
				estructura := base.Proveedores{}
				fmt.Println("Nombre:")
				fmt.Scanf("%s", &estructura.Nombre)
				fmt.Println("Telefono:")
				fmt.Scanf("%s", &estructura.Telefono)
				fmt.Println("Direccion")
				fmt.Scanf("%s", &estructura.Direccion)
				base.Insertar_sql("proveedores", columnas, fmt.Sprintf(valores, estructura.Nombre, estructura.Direccion, estructura.Telefono))
				fmt.Println("Insertado con exito")
			case "6":
				columnas := "Direccion,ClienteID,EmpleadoID,valor,DetalleID,MetodoPago"
				valores := "%s,%d,%d,%d,%d,%s"
				estructura := base.Pedidos{}
				fmt.Println("Direccion:")
				fmt.Scanf("%s", &estructura.Direccion)
				fmt.Println("Cliente ID (Numero Entero):")
				fmt.Scanf("%d", &estructura.ClienteID)
				fmt.Println("Empleado ID (Numero Entero)")
				fmt.Scanf("%d", &estructura.EmpleadoID)
				fmt.Println("Valor (Numero Entero):")
				fmt.Scanf("%d", &estructura.Valor)
				fmt.Println("Detalle ID (Numero Entero)")
				fmt.Scanf("%d", &estructura.DetalleID)
				fmt.Println("Metodo de Pago")
				fmt.Scanf("%s", &estructura.MetodoPago)
				base.Insertar_sql("pedidos", columnas, fmt.Sprintf(valores, estructura.Direccion, estructura.ClienteID, estructura.EmpleadoID, estructura.Valor, estructura.DetalleID, estructura.MetodoPago))
				fmt.Println("Insertado con exito")
			case "7":
				columnas := "id,productoID,cantidad"
				valores := "%d,%d,%d"
				estructura := base.Detalle_pedidos{}
				fmt.Println("Id (Numero Entero)")
				fmt.Scanf("%d", &estructura.Id)
				fmt.Println("Producto_id (Numero Entero)")
				fmt.Scanf("%d", &estructura.ProductoID)
				fmt.Println("Cantidad (Numero Entero)")
				fmt.Scanf("%d", &estructura.Cantidad)
				base.Insertar_sql("detalle_pedidos", columnas, fmt.Sprintf(valores, estructura.Id, estructura.ProductoID, estructura.Cantidad))
				fmt.Println("Insertado con exito")
			}
		case 2:
			fmt.Println("Lista de opciones:" +
				"\n 1 Mostrar Tabla de Productos" +
				"\n 2 Mostrar Tabla de Generos" +
				"\n 3 Mostrar Tabla de Empleados" +
				"\n 4 Mostrar Tabla de Clientes" +
				"\n 5 Mostrar Tabla de Proveedores" +
				"\n 6 Mostrar Tabla de Pedidos" +
				"\n 7 Mostrar Tabla de Detalle de Pedidos")
			var opcion, buscar string
			fmt.Scanf("%s", &opcion)
			switch opcion {
			case "1":
				fmt.Println("Desea buscar una ID en especifico?, por favor insertela, si no es el caso, Presione Enter\n")
				fmt.Scanf("%s", &buscar)
				productos, err := base.ObtenerProductos(buscar)
				if err != nil {
					fmt.Printf("Error obteniendo contactos: %v", err)
				} else {
					base.ImprimirProducto(productos)
				}
			case "2":
				fmt.Println("Desea buscar una ID en especifico?, por favor insertela, si no es el caso, Presione Enter\n")
				fmt.Scanf("%s", &buscar)
				Generos, err := base.ObtenerGenero(buscar)
				if err != nil {
					fmt.Printf("Error obteniendo contactos: %v", err)
				} else {
					base.ImprimirGeneros(Generos)
				}
			case "3":
				fmt.Println("Desea buscar una ID en especifico?, por favor insertela, si no es el caso, Presione Enter\n")
				fmt.Scanf("%s", &buscar)
				empleados, err := base.ObtenerEmpleados(buscar)
				if err != nil {
					fmt.Printf("Error obteniendo contactos: %v", err)
				} else {
					base.ImprimirEmpleados(empleados)
				}
			case "4":
				fmt.Println("Desea buscar una ID en especifico?, por favor insertela, si no es el caso, Presione Enter\n")
				fmt.Scanf("%s", &buscar)
				clientes, err := base.ObtenerClientes(buscar)
				if err != nil {
					fmt.Printf("Error obteniendo contactos: %v", err)
				} else {
					base.ImprimirClientes(clientes)
				}
			case "5":
				fmt.Println("Desea buscar una ID en especifico?, por favor insertela, si no es el caso, Presione Enter\n")
				fmt.Scanf("%s", &buscar)
				proveedores, err := base.ObtenerProveedores(buscar)
				if err != nil {
					fmt.Printf("Error obteniendo contactos: %v", err)
				} else {
					base.ImprimirProveedores(proveedores)
				}
			case "6":
				fmt.Println("Desea buscar una ID en especifico?, por favor insertela, si no es el caso, Presione Enter\n")
				fmt.Scanf("%s", &buscar)
				pedidos, err := base.ObtenerPedidos(buscar)
				if err != nil {
					fmt.Printf("Error obteniendo contactos: %v", err)
				} else {
					base.ImprimirPedidos(pedidos)
				}
			case "7":
				fmt.Println("Desea buscar una ID en especifico?, por favor insertela, si no es el caso, Presione Enter\n")
				fmt.Scanf("%s", &buscar)
				detalle_pedidos, err := base.ObtenerDetalle_Pedidos(buscar)
				if err != nil {
					fmt.Printf("Error obteniendo contactos: %v", err)
				} else {
					base.ImprimirDetalle_Pedidos(detalle_pedidos)
				}
			default:
				continue
			}

		case 3:
			fmt.Println("Lista de opciones:" +
				"\n 1 Actualizar la Tabla de Productos" +
				"\n 2 Actualizar la Tabla de Generos" +
				"\n 3 Actualizar la Tabla de Empleados" +
				"\n 4 Actualizar la Tabla de Clientes" +
				"\n 5 Actualizar la Tabla de Proveedores" +
				"\n 6 Actualizar la Tabla de Pedidos" +
				"\n 7 Actualizar la Tabla de Detalle de Pedidos")
			var opcion string
			fmt.Scanf("%s", &opcion)
			switch opcion {
			case "1":
				var id int
				estructura := base.Producto{}
				fmt.Println("Id a cambiar:")
				fmt.Scanf("%d", &id)
				fmt.Println("Nombre:")
				fmt.Scanf("%s", &estructura.Nombre)
				fmt.Println("Generos:")
				fmt.Scanf("%s", &estructura.Generos)
				fmt.Println("Valor (Numero Entero)")
				fmt.Scanf("%d", &estructura.Valor)
				fmt.Println("Plataformas:")
				fmt.Scanf("%s", &estructura.Plataformas)
				fmt.Println("Proveedor_id (Numero Entero)")
				fmt.Scanf("%d", &estructura.Proveedor_id)
				base.Actualizar_sqlp("productos", estructura.Nombre, estructura.Generos, estructura.Valor, estructura.Plataformas, estructura.Proveedor_id, id)
				fmt.Println("Cambiado con exito")

				fmt.Println("Cambiado con exito")
			case "2":
				var id int
				estructura := base.Genero{}
				fmt.Println("Id a cambiar:")
				fmt.Scanf("%d", &id)
				fmt.Println("Genero:")
				fmt.Scanf("%s", &estructura.Tipo)
				base.Actualizar_sqlg("genero", estructura.Tipo, id)
				fmt.Println("Cambiado con exito")
			case "3":
				var id int
				estructura := base.Empleado{}
				fmt.Println("Id a cambiar:")
				fmt.Scanf("%d", &id)
				fmt.Println("Rut:")
				fmt.Scanf("%s", &estructura.Rut)
				fmt.Println("Nombre:")
				fmt.Scanf("%s", &estructura.Nombre)
				fmt.Println("Sueldo (Numero Entero)")
				fmt.Scanf("%d", &estructura.Sueldo)
				fmt.Println("Area:")
				fmt.Scanf("%s", &estructura.Area)
				fmt.Println("Cargo")
				fmt.Scanf("%s", &estructura.Cargo)
				fmt.Println("Direccion")
				fmt.Scanf("%s", &estructura.Direccion)
				fmt.Println("Region")
				fmt.Scanf("%s", &estructura.Region)
				base.Actualizar_sqle("empleados", estructura.Rut, estructura.Nombre, estructura.Sueldo, estructura.Area, estructura.Cargo, estructura.Direccion, estructura.Region, id)
				fmt.Println("Cambiado con exito")
			case "4":
				estructura := base.Clientes{}
				var id int
				fmt.Println("Ingrese id a cambiar:")
				fmt.Scanf("%d", &id)
				fmt.Println("Rut:")
				fmt.Scanf("%s", &estructura.Rut)
				fmt.Println("Nombre:")
				fmt.Scanf("%s", &estructura.Nombre)
				fmt.Println("Telefono:")
				fmt.Scanf("%s", &estructura.Telefono)
				fmt.Println("Direccion")
				fmt.Scanf("%s", &estructura.Direccion)
				fmt.Println("Region")
				fmt.Scanf("%s", &estructura.Region)
				base.Actualizar_sqlc("clientes", estructura.Rut, estructura.Nombre, estructura.Direccion, estructura.Region, estructura.Telefono, id)
				fmt.Println("Cambiado con exito")
			case "5":
				var id int
				estructura := base.Proveedores{}
				fmt.Println("Ingrese id a cambiar:")
				fmt.Scanf("%d", &id)
				fmt.Println("Nombre:")
				fmt.Scanf("%s", &estructura.Nombre)
				fmt.Println("Telefono:")
				fmt.Scanf("%s", &estructura.Telefono)
				fmt.Println("Direccion")
				fmt.Scanf("%s", &estructura.Direccion)
				base.Actualizar_p("proveedores", estructura.Nombre, estructura.Direccion, estructura.Telefono, id)
				fmt.Println("Insertado con exito")
			case "6":
				var id int
				estructura := base.Pedidos{}
				fmt.Println("Ingrese id a cambiar:")
				fmt.Scanf("%d", &id)
				fmt.Println("Direccion:")
				fmt.Scanf("%s", &estructura.Direccion)
				fmt.Println("Cliente ID (Numero Entero):")
				fmt.Scanf("%d", &estructura.ClienteID)
				fmt.Println("Empleado ID (Numero Entero)")
				fmt.Scanf("%d", &estructura.EmpleadoID)
				fmt.Println("Valor (Numero Entero):")
				fmt.Scanf("%d", &estructura.Valor)
				fmt.Println("Detalle ID (Numero Entero)")
				fmt.Scanf("%d", &estructura.DetalleID)
				fmt.Println("Metodo de Pago")
				fmt.Scanf("%s", &estructura.MetodoPago)
				base.Actualizar_sqlpe("pedidos", estructura.Direccion, estructura.ClienteID, estructura.EmpleadoID, estructura.Valor, estructura.DetalleID, estructura.MetodoPago, id)
				fmt.Println("Cambiado con exito")
			case "7":
				estructura := base.Detalle_pedidos{}
				fmt.Println("Id a cambiar(Numero Entero)")
				fmt.Scanf("%d", &estructura.Id)
				fmt.Println("Producto_id (Numero Entero)")
				fmt.Scanf("%d", &estructura.ProductoID)
				fmt.Println("Cantidad (Numero Entero)")
				fmt.Scanf("%d", &estructura.Cantidad)
				base.Actualizar_sqlfi("detalle_pedidos", estructura.ProductoID, estructura.Cantidad, estructura.Id)
				fmt.Println("Cambiado con exito")
			}
		case 4:
			fmt.Println("Lista de opciones:" +
				"\n 1 Eliminar desde Tabla de Productos" +
				"\n 2 Eliminar desde Tabla de Generos" +
				"\n 3 Eliminar desde Tabla de Empleados" +
				"\n 4 Eliminar desde Tabla de Clientes" +
				"\n 5 Eliminar desde Tabla de Proveedores" +
				"\n 6 Eliminar desde Tabla de Pedidos" +
				"\n 7 Eliminar desde Tabla de Detalle de Pedidos")
			var opcion, buscar string
			fmt.Scanf("%s", &opcion)
			switch opcion {
			case "1":
				fmt.Println("Inserte la Id a eliminar")
				_, _ = fmt.Scanf("%s", &buscar)
				err := base.Eliminar_sql("productos", buscar)
				if err != nil {
					fmt.Println("Fallo al eliminar")
				} else {
					fmt.Println("Eliminado con exito")
				}
			case "2":
				fmt.Println("Inserte la Id a eliminar")
				_, _ = fmt.Scanf("%s", &buscar)
				err := base.Eliminar_sql("genero", buscar)
				if err != nil {
					fmt.Println("Fallo al eliminar")
				} else {
					fmt.Println("Eliminado con exito")
				}
			case "3":
				fmt.Println("Inserte la Id a eliminar")
				_, _ = fmt.Scanf("%s", &buscar)
				err := base.Eliminar_sql("empleados", buscar)
				if err != nil {
					fmt.Println("Fallo al eliminar")
				} else {
					fmt.Println("Eliminado con exito")
				}
			case "4":
				fmt.Println("Inserte la Id a eliminar")
				_, _ = fmt.Scanf("%s", &buscar)
				err := base.Eliminar_sql("clientes", buscar)
				if err != nil {
					fmt.Println("Fallo al eliminar")
				} else {
					fmt.Println("Eliminado con exito")
				}
			case "5":
				fmt.Println("Inserte la Id a eliminar")
				_, _ = fmt.Scanf("%s", &buscar)
				err := base.Eliminar_sql("proovedores", buscar)
				if err != nil {
					fmt.Println("Fallo al eliminar")
				} else {
					fmt.Println("Eliminado con exito")
				}
			case "6":
				fmt.Println("Inserte la Id a eliminar")
				_, _ = fmt.Scanf("%s", &buscar)
				err := base.Eliminar_sql("pedidos", buscar)
				if err != nil {
					fmt.Println("Fallo al eliminar")
				} else {
					fmt.Println("Eliminado con exito")
				}
			case "7":
				fmt.Println("Inserte la Id a eliminar")
				_, _ = fmt.Scanf("%s", &buscar)
				err := base.Eliminar_sql("detalle_pedidos", buscar)
				if err != nil {
					fmt.Println("Fallo al eliminar")
				} else {
					fmt.Println("Eliminado con exito")
				}
			default:
				continue
			}
		}
	}
}
