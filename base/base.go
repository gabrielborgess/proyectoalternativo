package base

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/cheggaaa/pb.v1"
	_ "gopkg.in/cheggaaa/pb.v1"
	"strconv"
	"strings"
	"time"
)

type Detalle_pedidos struct {
	id, producto_id, cantidad int
}

type Pedidos struct {
	id, cliente_id, empleado_id, detalle_id, valor int
	direccion, metodo_pago                         string
}

type Proveedores struct {
	id                          int
	nombre, direccion, telefono string
}

type Clientes struct {
	id                                       int
	rut, nombre, direccion, region, telefono string
}

type Empleado struct {
	Id, Sueldo                                  int
	Rut, Nombre, Area, Cargo, Direccion, Region string
}

type Producto struct {
	Nombre, Plataformas, Generos       string
	Id, Valor, Proveedor_id, Estrellas int
}
type Genero struct {
	Id   int
	Tipo string
}

func Base() { //function principal donde llamamos la funcion de crear tablas etc
	count := 50
	bar := pb.StartNew(count)
	bar.ShowCounters = false
	bar.ShowElapsedTime = true
	bar.ShowFinalTime = false
	go creartablas(bar)
	time.Sleep(4 * time.Second)
}

func Droptable(nombre string) {
	Execdb("DROP TABLE IF EXISTS `" + nombre + "`")
}
func Createtable(nombre string, atributos string) {
	Execdb("CREATE TABLE " + nombre + " (" + atributos + ")")
}

func Execdb(query string) { // Usa la funcion que cree yo para hacer las query, saldrá mejor y mas facil
	db, err := ObtenerBaseDeDatos()
	if err != nil {
		fmt.Printf("error al conectar")
		return
	}
	defer db.Close()
	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func Atributos() ([]string, []string) {
	tablas := []string{"clientes",
		"empleados",
		"pedidos",
		"detalle_pedidos",
		"productos",
		"genero",
		"proveedores"}
	atributos := []string{"id integer NOT NULL AUTO_INCREMENT, Rut varchar(255), Nombre varchar(255), Direccion varchar(255), Region varchar(255),telefono varchar(255),PRIMARY KEY (`id`)",
		"id integer NOT NULL AUTO_INCREMENT, Rut varchar(255), Nombre varchar(255), Sueldo integer, Area varchar(255), Cargo varchar(255), Direccion varchar(255), Region varchar(255),PRIMARY KEY (`id`)",
		"id integer NOT NULL AUTO_INCREMENT, Direccion varchar(255), cliente_id integer, empleado_id integer, valor integer, detalle_id integer, metodo_pago varchar(255),PRIMARY KEY (`id`)",
		"id integer, producto_id integer, cantidad integer",
		"id integer NOT NULL AUTO_INCREMENT, Nombre varchar(255),generos varchar(255), valor integer, plataformas varchar(255), proveedor_id integer, estrellas integer,PRIMARY KEY (`id`)",
		"id integer NOT NULL AUTO_INCREMENT, tipo varchar(255),PRIMARY KEY (`id`)",
		"id integer NOT NULL AUTO_INCREMENT, Nombre varchar(255), Direccion varchar(255), telefono varchar(255),PRIMARY KEY (`id`)"}
	return tablas, atributos
}

func Show_Struct() {
	tablas, atributos := Atributos()
	for i := range tablas {
		fmt.Printf("\nTabla\t\tAtributos\n%s\t\t%s\n", tablas[i], atributos[i])
	}
}

func creartablas(a *pb.ProgressBar) { //una funcion aparte encargada solo de crear tablas
	tablas, atributos := Atributos()
	go func() {
		for i := range tablas {
			go Droptable(tablas[i])
			a.Increment()
		}
	}()
	time.Sleep(1 * time.Second)
	go func() {
		for i := range atributos {
			go Createtable(tablas[i], atributos[i])
			a.Increment()
		}
	}()
	time.Sleep(100 * time.Millisecond)
	a.FinishPrint("Programa Cargado\n")
	time.Sleep(100 * time.Millisecond)

}
func ObtenerBaseDeDatos() (db *sql.DB, e error) {
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

// Los valores se separan por coma, la tabla se declara, se utiliza una db ya creada para evitar las sobreconexiones a base y se confirma el cierre o no al final de ejecución
func Insertar_sql(tabla string, columnas string, valores string, db *sql.DB, cerrar bool) {
	//
	if cerrar {
		defer db.Close()
	}
	preparado := ""
	iterar := strings.Split(valores, ",")
	for i := range iterar {
		_, err := strconv.Atoi(iterar[i])
		if err != nil {
			if i+1 == len(iterar) {
				preparado += "\"" + iterar[i] + "\""
			} else {
				preparado += "\"" + iterar[i] + "\","
			}
		} else {
			if i+1 == len(iterar) {
				preparado += iterar[i]
			} else {
				preparado += iterar[i] + ","
			}
		}

	}
	fmt.Println(fmt.Sprintf("Sql ejecutada:\nINSERT INTO `%s` (%s) VALUES(%s);", tabla, columnas, preparado))
	Execdb(fmt.Sprintf("INSERT INTO `%s` (%s) VALUES(%s);", tabla, columnas, preparado))
	return
}
func eliminar_sql(tabla string, id string) error {
	db, err := ObtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()
	sentenciaPreparada, err := db.Prepare(fmt.Sprint("DELETE FROM `%s` WHERE id=%s", tabla, id))
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()

	_, err = sentenciaPreparada.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
func ObtenerProductos(buscarid string) ([]Producto, error) {
	productos := []Producto{}
	db, err := ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var filas *sql.Rows
	if buscarid != "" {
		filas, err = db.Query(fmt.Sprintf("SELECT * FROM productos where id=%s", buscarid))
	} else {
		filas, err = db.Query("SELECT * FROM productos")
	}

	if err != nil {
		return nil, err
	}
	defer filas.Close()
	var producto Producto
	for filas.Next() {
		err = filas.Scan(&producto.Id, &producto.Nombre, &producto.Generos, &producto.Valor, &producto.Plataformas, &producto.Proveedor_id, &producto.Estrellas)
		if err != nil {
			return nil, err
		}
		productos = append(productos, producto)
	}
	return productos, nil
}

func ImprimirProducto(productos []Producto) {
	for _, Producto := range productos {
		fmt.Println("====================")
		fmt.Printf("Id: %d\n", Producto.Id)
		fmt.Printf("Nombre: %s\n", Producto.Nombre)
		fmt.Printf("Generos: %s\n", Producto.Generos)
		fmt.Printf("Precio: %d\n", Producto.Valor)
		fmt.Printf("Plataformas: %s\n", Producto.Plataformas)
		fmt.Printf("Proveedor ID: %d\n", Producto.Proveedor_id)
		fmt.Printf("Estrellas: %d\n", Producto.Estrellas)
	}
}
func ObtenerGenero(buscarid string) ([]Genero, error) {
	generos := []Genero{}
	db, err := ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var filas *sql.Rows
	if buscarid != "" {
		filas, err = db.Query(fmt.Sprintf("SELECT * FROM genero where id=%s", buscarid))
	} else {
		filas, err = db.Query("SELECT * FROM genero")
	}

	if err != nil {
		return nil, err
	}
	defer filas.Close()
	var genero Genero
	for filas.Next() {
		err = filas.Scan(&genero.Id, &genero.Tipo)
		if err != nil {
			return nil, err
		}
		generos = append(generos, genero)
	}
	return generos, nil
}

func ImprimirGeneros(generos []Genero) {
	fmt.Println(">ID<======>Genero<")
	for _, genero := range generos {
		fmt.Printf("%d \t \t %s\n", genero.Id, genero.Tipo)
	}
}

func ObtenerEmpleados(buscarid string) ([]Empleado, error) {
	empleados := []Empleado{}
	db, err := ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var filas *sql.Rows
	if buscarid != "" {
		filas, err = db.Query(fmt.Sprintf("SELECT * FROM empleados where id=%s", buscarid))
	} else {
		filas, err = db.Query("SELECT * FROM empleados")
	}

	if err != nil {
		return nil, err
	}
	defer filas.Close()
	var empleado Empleado
	for filas.Next() {
		err = filas.Scan(&empleado.Id, &empleado.Rut, &empleado.Nombre, &empleado.Sueldo, &empleado.Area, &empleado.Cargo, &empleado.Direccion, &empleado.Region)
		if err != nil {
			return nil, err
		}
		empleados = append(empleados, empleado)
	}
	return empleados, nil
}

func ImprimirEmpleados(generos []Empleado) {
	for _, Empleado := range generos {
		fmt.Println("====================")
		fmt.Printf("Id: %d\n", Empleado.Id)
		fmt.Printf("Rut: %s\n", Empleado.Rut)
		fmt.Printf("Nombre: %s\n", Empleado.Nombre)
		fmt.Printf("Sueldo: %d\n", Empleado.Sueldo)
		fmt.Printf("Area: %s\n", Empleado.Area)
		fmt.Printf("Cargo: %s\n", Empleado.Cargo)
		fmt.Printf("Direccion: %s\n", Empleado.Direccion)
		fmt.Printf("Region: %s\n", Empleado.Region)
	}
}

func ObtenerClientes(buscarid string) ([]Clientes, error) {
	clientes := []Clientes{}
	db, err := ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var filas *sql.Rows
	if buscarid != "" {
		filas, err = db.Query(fmt.Sprintf("SELECT * FROM clientes where id=%s", buscarid))
	} else {
		filas, err = db.Query("SELECT * FROM clientes")
	}

	if err != nil {
		return nil, err
	}
	defer filas.Close()
	var cliente Clientes
	for filas.Next() {
		err = filas.Scan(&cliente.id, &cliente.rut, &cliente.nombre, &cliente.direccion, &cliente.region, &cliente.telefono)
		if err != nil {
			return nil, err
		}
		clientes = append(clientes, cliente)
	}
	return clientes, nil
}

func ImprimirClientes(generos []Clientes) {
	for _, genero := range generos {
		fmt.Println("====================")
		fmt.Printf("Id: %d\n", genero.id)
		fmt.Printf("Rut: %s\n", genero.rut)
		fmt.Printf("Nombre: %s\n", genero.nombre)
		fmt.Printf("Direccion: %s\n", genero.direccion)
		fmt.Printf("Region: %s\n", genero.region)
		fmt.Printf("Telefono: %s\n", genero.telefono)
	}
}
func ObtenerProveedores(buscarid string) ([]Proveedores, error) {
	proveedores := []Proveedores{}
	db, err := ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var filas *sql.Rows
	if buscarid != "" {
		filas, err = db.Query(fmt.Sprintf("SELECT * FROM proveedores where id=%s", buscarid))
	} else {
		filas, err = db.Query("SELECT * FROM proveedores")
	}

	if err != nil {
		return nil, err
	}
	defer filas.Close()
	var proveedor Proveedores
	for filas.Next() {
		err = filas.Scan(&proveedor.id, &proveedor.nombre, &proveedor.direccion, &proveedor.telefono)
		if err != nil {
			return nil, err
		}
		proveedores = append(proveedores, proveedor)
	}
	return proveedores, nil
}

func ImprimirProveedores(proveedores []Proveedores) {
	for _, Proveedor := range proveedores {
		fmt.Println("====================")
		fmt.Printf("Id: %d\n", Proveedor.id)
		fmt.Printf("Nombre: %s\n", Proveedor.nombre)
		fmt.Printf("Direccion: %s\n", Proveedor.direccion)
		fmt.Printf("Telefono: %s\n", Proveedor.telefono)
	}
}

func ObtenerPedidos(buscarid string) ([]Pedidos, error) {
	pedidos := []Pedidos{}
	db, err := ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var filas *sql.Rows
	if buscarid != "" {
		filas, err = db.Query(fmt.Sprintf("SELECT * FROM pedido where id=%s", buscarid))
	} else {
		filas, err = db.Query("SELECT * FROM pedido")
	}

	if err != nil {
		return nil, err
	}
	defer filas.Close()
	var pedido Pedidos
	for filas.Next() {
		err = filas.Scan(&pedido.id, &pedido.direccion, &pedido.cliente_id, &pedido.empleado_id, &pedido.valor, &pedido.detalle_id, &pedido.metodo_pago)
		if err != nil {
			return nil, err
		}
		pedidos = append(pedidos, pedido)
	}
	return pedidos, nil
}

func ImprimirPedidos(pedidos []Pedidos) {
	for _, pedido := range pedidos {
		fmt.Println("====================")
		fmt.Printf("Id: %d\n", pedido.id)
		fmt.Printf("Direccion: %s\n", pedido.direccion)
		fmt.Printf("Cliente ID: %d\n", pedido.cliente_id)
		fmt.Printf("pedido ID: %d\n", pedido.empleado_id)
		fmt.Printf("Valor: %d\n", pedido.valor)
		fmt.Printf("Detalle ID: %d\n", pedido.detalle_id)
		fmt.Printf("Metodo de Pago: %s\n", pedido.metodo_pago)
	}
}
func ObtenerDetalle_Pedidos(buscarid string) ([]Detalle_pedidos, error) {
	detalle_pedidos := []Detalle_pedidos{}
	db, err := ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var filas *sql.Rows
	if buscarid != "" {
		filas, err = db.Query(fmt.Sprintf("SELECT * FROM detalle_pedidos where id=%s", buscarid))
	} else {
		filas, err = db.Query("SELECT * FROM detalle_pedidos")
	}

	if err != nil {
		return nil, err
	}
	defer filas.Close()
	var detallePedidos Detalle_pedidos
	for filas.Next() {
		err = filas.Scan(&detallePedidos.id, &detallePedidos.producto_id, &detallePedidos.cantidad)
		if err != nil {
			return nil, err
		}
		detalle_pedidos = append(detalle_pedidos, detallePedidos)
	}
	return detalle_pedidos, nil
}

func ImprimirDetalle_Pedidos(detalle_pedidos []Detalle_pedidos) {
	for _, detallePedidos := range detalle_pedidos {
		fmt.Println("====================")
		fmt.Printf("Id: %d\n", detallePedidos.id)
		fmt.Printf("Producto ID: %d\n", detallePedidos.producto_id)
		fmt.Printf("cantidad: %d\n", detallePedidos.cantidad)
	}
}
