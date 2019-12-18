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
	Id, ProductoID, Cantidad int
}

type Pedidos struct {
	Id, ClienteID, EmpleadoID, DetalleID, Valor int
	Direccion, MetodoPago                       string
}

type Proveedores struct {
	Id                          int
	Nombre, Direccion, Telefono string
}

type Clientes struct {
	Id                                       int
	Rut, Nombre, Direccion, Region, Telefono string
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
	count := 14
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
	atributos := []string{"Id integer NOT NULL AUTO_INCREMENT, Rut varchar(255), Nombre varchar(255), Direccion varchar(255), Region varchar(255),Telefono varchar(255),PRIMARY KEY (`Id`)",
		"Id integer NOT NULL AUTO_INCREMENT, Rut varchar(255), Nombre varchar(255), Sueldo integer, Area varchar(255), Cargo varchar(255), Direccion varchar(255), Region varchar(255),PRIMARY KEY (`Id`)",
		"Id integer NOT NULL AUTO_INCREMENT, Direccion varchar(255), ClienteID integer, EmpleadoID integer, Valor integer, DetalleID integer, MetodoPago varchar(255),PRIMARY KEY (`Id`)",
		"Id integer, ProductoID integer, Cantidad integer",
		"Id integer NOT NULL AUTO_INCREMENT, Nombre varchar(255),generos varchar(255), Valor integer, plataformas varchar(255), proveedor_id integer, estrellas integer,PRIMARY KEY (`Id`)",
		"Id integer NOT NULL AUTO_INCREMENT, tipo varchar(255),PRIMARY KEY (`Id`)",
		"Id integer NOT NULL AUTO_INCREMENT, Nombre varchar(255), Direccion varchar(255), Telefono varchar(255),PRIMARY KEY (`Id`)"}
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
	pass := "contraseña"
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
func Insertar_sql(tabla string, columnas string, valores string) {
	//
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
func Eliminar_sql(tabla string, id string) error {
	db, err := ObtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()
	sentenciaPreparada, err := db.Prepare(fmt.Sprintf("DELETE FROM `%s` WHERE Id=?", tabla))
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
		filas, err = db.Query(fmt.Sprintf("SELECT * FROM productos where Id=%s", buscarid))
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
		filas, err = db.Query(fmt.Sprintf("SELECT * FROM genero where Id=%s", buscarid))
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
		filas, err = db.Query(fmt.Sprintf("SELECT * FROM empleados where Id=%s", buscarid))
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
		filas, err = db.Query(fmt.Sprintf("SELECT * FROM clientes where Id=%s", buscarid))
	} else {
		filas, err = db.Query("SELECT * FROM clientes")
	}

	if err != nil {
		return nil, err
	}
	defer filas.Close()
	var cliente Clientes
	for filas.Next() {
		err = filas.Scan(&cliente.Id, &cliente.Rut, &cliente.Nombre, &cliente.Direccion, &cliente.Region, &cliente.Telefono)
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
		fmt.Printf("Id: %d\n", genero.Id)
		fmt.Printf("Rut: %s\n", genero.Rut)
		fmt.Printf("Nombre: %s\n", genero.Nombre)
		fmt.Printf("Direccion: %s\n", genero.Direccion)
		fmt.Printf("Region: %s\n", genero.Region)
		fmt.Printf("Telefono: %s\n", genero.Telefono)
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
		filas, err = db.Query(fmt.Sprintf("SELECT * FROM proveedores where Id=%s", buscarid))
	} else {
		filas, err = db.Query("SELECT * FROM proveedores")
	}

	if err != nil {
		return nil, err
	}
	defer filas.Close()
	var proveedor Proveedores
	for filas.Next() {
		err = filas.Scan(&proveedor.Id, &proveedor.Nombre, &proveedor.Direccion, &proveedor.Telefono)
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
		fmt.Printf("Id: %d\n", Proveedor.Id)
		fmt.Printf("Nombre: %s\n", Proveedor.Nombre)
		fmt.Printf("Direccion: %s\n", Proveedor.Direccion)
		fmt.Printf("Telefono: %s\n", Proveedor.Telefono)
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
		filas, err = db.Query(fmt.Sprintf("SELECT * FROM pedidos where Id=%s", buscarid))
	} else {
		filas, err = db.Query("SELECT * FROM pedidos")
	}

	if err != nil {
		return nil, err
	}
	defer filas.Close()
	var pedido Pedidos
	for filas.Next() {
		err = filas.Scan(&pedido.Id, &pedido.Direccion, &pedido.ClienteID, &pedido.EmpleadoID, &pedido.Valor, &pedido.DetalleID, &pedido.MetodoPago)
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
		fmt.Printf("Id: %d\n", pedido.Id)
		fmt.Printf("Direccion: %s\n", pedido.Direccion)
		fmt.Printf("Cliente ID: %d\n", pedido.ClienteID)
		fmt.Printf("pedido ID: %d\n", pedido.EmpleadoID)
		fmt.Printf("Valor: %d\n", pedido.Valor)
		fmt.Printf("Detalle ID: %d\n", pedido.DetalleID)
		fmt.Printf("Metodo de Pago: %s\n", pedido.MetodoPago)
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
		filas, err = db.Query(fmt.Sprintf("SELECT * FROM detalle_pedidos where Id=%s", buscarid))
	} else {
		filas, err = db.Query("SELECT * FROM detalle_pedidos")
	}

	if err != nil {
		return nil, err
	}
	defer filas.Close()
	var detallePedidos Detalle_pedidos
	for filas.Next() {
		err = filas.Scan(&detallePedidos.Id, &detallePedidos.ProductoID, &detallePedidos.Cantidad)
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
		fmt.Printf("Id: %d\n", detallePedidos.Id)
		fmt.Printf("Producto ID: %d\n", detallePedidos.ProductoID)
		fmt.Printf("Cantidad: %d\n", detallePedidos.Cantidad)
	}
}

func Actualizar_sqlp(tabla string, nombre string, generos string, valor int, plataforma string, proveedor int, id int) {
	//

	fmt.Println(fmt.Sprintf("UPDATE `%s` SET `Nombre`='%s',`generos`='%s',`valor`='%d',`plataformas`='%s',`proveedor_id`='%d' WHERE `id`='%d';", tabla, nombre, generos, valor, plataforma, proveedor, id))
	Execdb(fmt.Sprintf("UPDATE `%s` SET `Nombre`='%s',`generos`='%s',`valor`='%d',`plataformas`='%s',`proveedor_id`='%d' WHERE `id`='%d';", tabla, nombre, generos, valor, plataforma, proveedor, id))
	return
}

func Actualizar_sqlg(tabla string, genero string, id int) {
	//

	fmt.Println(fmt.Sprintf("UPDATE `%s` SET `Tipo`='%s' WHERE `id`='%d';", tabla, genero, id))
	Execdb(fmt.Sprintf("UPDATE `%s` SET `Tipo`='%s' WHERE `id`='%d';", tabla, genero, id))
	return
}

func Actualizar_sqle(tabla string, Rut string, Nombre string, Sueldo int, Area string, Cargo string, Direccion string, Region string, id int) { //%s,%s,%d,%s,%s,%s,%s %d
	//

	fmt.Println(fmt.Sprintf("UPDATE `%s` SET `Rut`='%s',`Nombre`='%s',`Sueldo`='%d',`Area`='%s',`Cargo`='%s',`Direccion`='%s',`Region`='%s' WHERE `id`='%d;'", tabla, Rut, Nombre, Sueldo, Area, Cargo, Direccion, Region, id))
	Execdb(fmt.Sprintf("UPDATE `%s` SET `Rut`='%s',`Nombre`='%s',`Sueldo`='%d',`Area`='%s',`Cargo`='%s',`Direccion`='%s',`Region`='%s' WHERE `id`='%d';", tabla, Rut, Nombre, Sueldo, Area, Cargo, Direccion, Region, id))
	return
}

func Actualizar_sqlc(tabla string, Rut string, Nombre string, Direccion string, Region string, telefono string, id int) { //%s,%s,%s,%s,%s,%d
	//

	fmt.Println(fmt.Sprintf("UPDATE `%s` SET `Rut`='%s',`Nombre`='%s',`Direccion`='%s',`Region`='%s',`telefono`='%s' WHERE `id`='%d';", tabla, Rut, Nombre, Direccion, Region, telefono, id))
	Execdb(fmt.Sprintf("UPDATE `%s` SET `Rut`='%s',`Nombre`='%s',`Direccion`='%s',`Region`='%s',`telefono`='%s' WHERE `id`='%d';", tabla, Rut, Nombre, Direccion, Region, telefono, id))
	return
}

func Actualizar_p(tabla string, Nombre string, Direccion string, telefono string, id int) { //"%s,%s,%s"
	//

	fmt.Println(fmt.Sprintf("UPDATE `%s` SET `Nombre`='%s',`Direccion`='%s',`telefono`='%s' WHERE `id`='%d';", tabla, Nombre, Direccion, telefono, id))
	Execdb(fmt.Sprintf("UPDATE `%s` SET `Nombre`='%s',`Direccion`='%s',`telefono`='%s' WHERE `id`='%d';", tabla, Nombre, Direccion, telefono, id))
	return
}

func Actualizar_sqlpe(tabla string, Direccion string, cliente_id int, empleado_id int, valor int, detalle_id int, metodo_pago string, id int) { //"%s,%d,%d,%d,%d,%s,%d"
	//

	fmt.Println(fmt.Sprintf("UPDATE `%s` SET `Direccion`='%s',`cliente_id`='%d',`empleado_id`='%d',`valor`='%d',`detalle_id`='%d',`metodo_pago`='%s' WHERE `id`='%d';", tabla, Direccion, cliente_id, empleado_id, valor, detalle_id, metodo_pago, id))
	Execdb(fmt.Sprintf("UPDATE `%s` SET `Direccion`='%s',`cliente_id`='%d',`empleado_id`='%d',`valor`='%d',`detalle_id`='%d',`metodo_pago`='%s' WHERE `id`='%d';", tabla, Direccion, cliente_id, empleado_id, valor, detalle_id, metodo_pago, id))
	return
}

func Actualizar_sqlfi(tabla string, id int, producto_id int, cantidad int) { //"%d,%d,%d,%d"
	//

	fmt.Println(fmt.Sprintf("UPDATE `%s` SET `producto_id`='%d',`cantidad`='%d' WHERE `id`='%d';", tabla, producto_id, cantidad, id))
	Execdb(fmt.Sprintf("UPDATE `%s` SET `producto_id`='%d',`cantidad`='%d' WHERE `id`='%d';", tabla, producto_id, cantidad, id))
	return
}
