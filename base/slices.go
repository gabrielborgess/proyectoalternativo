package base

import (
	"fmt"
	"strings"
)

// Util para borrar un indice exacto en un slice con puntero (agrega "&" tipo: &slice)
func RemoveSlice(a *[]string, i int) {
	*a = append((*a)[:i], (*a)[i+1:]...)
}

func ExampleTest() {
	a := []string{"holi", "jajasaludos"}
	RemoveSlice(&a, 0)
	AppendSliceBySplit(&a, "uf,jaja,mira,che,como,te,bailo,con,el,ejemplo", ",")
	fmt.Println("Con el primer elemento eliminado:", a)
	RemoveSlice(&a, 5)
	fmt.Println("Despues de eliminar el 6to elemento:", a)
}

// Util para agregar elementos extras a un slice con un separador (agrega "&" tipo: &slice)
// Tambien sirve por ejemplo cuando tienes un slice, no quieres hacer otro con strings.split y solo quieres agregar cosas
func AppendSliceBySplit(slice *[]string, element string, separator string) {
	element_slice := strings.Split(element, separator)
	for i := 0; i < len(element_slice); i++ {
		*slice = append(*slice, element_slice[i])
	}
}
