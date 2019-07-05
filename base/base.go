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
	tablas := []string{"clientes_nombre", "clientes_rut", "clientes_direccion", "empleados_nombre", "empleados_sueldo",
		"empleados_rut", "empleados_area", "empleados_direccion", "empleados_cargo", "area", "pedidos_direccion",
		"pedidos_empleado", "pedidos_cliente", "pedidos_detalle", "detalle_pedidos",
		"productos_nombre", "productos_valor", "productos_proveedor",
		"productos_genero", "productos_plataforma", "productos_estrellas",
		"plataforma", "proveedores", "genero", "metodo_pago"}
	atributos := []string{"id integer, nombre varchar(150)", "id integer, rut integer, dv integer",
		"id integer, direccion varchar(255), region varchar(150), telefono varchar(12)", "id integer, nombre varchar(255)",
		"id integer, sueldo integer", "id integer, rut integer, dv integer", "id integer, area_id integer",
		"id integer, direccion varchar(255), region varchar(150)", "id integer, cargo varchar(255)", "id integer, tipo varchar(255)",
		"id integer, direccion varchar(255)", "id integer, empleado_id integer", "id integer, cliente_id integer",
		"id integer, valor integer, detalle_id integer", "id integer, producto_id integer, cantidad integer", "id integer, nombre varchar(255)",
		"id integer, valor integer", "id integer, proveedor_id integer", "id integer, genero_id integer",
		"id integer, plataforma_id integer", "id integer, estrellas integer", "id integer, plataforma varchar(25)",
		"id integer, nombre varchar(255), direccion varchar(255), telefono varchar(12)", "id integer, genero varchar(32)", "id integer, nombre varchar(32)"}
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

type Juego struct {
	Nombre string
	Genero string
	Precio int
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
