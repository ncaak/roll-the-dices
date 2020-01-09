package command

import (
	"strings"
)

const DEFAULT_HELP = "Especifica el comando para el que necesitas ayuda\n_Ejemplo:_\n'/ayuda *tira*'"

func getHelp(command string) (help string) {
	switch strings.TrimSpace(command) {
	case "tira":
		help = helpTira()
	case "v":
		help = helpV()
	case "dv":
		help = helpDv()
	case "t":
		help = helpT()
	case "agrupa":
		help = helpAgrupa()
	case "ayuda":
		help = helpAyuda()
	case "repite":
		help = helpRepite()
	default:
		help = DEFAULT_HELP
	}
	return
}

/*
 * Commands' help texts
 */
func helpTira() string {
	return `
Lanza el número de dados especificado, o un dado de 20 caras por defecto
Admite bonificadores y penalizadores que añadirá al total
La sintaxis de dados acepta los prefijos 'h' y 'l', siendo _h_ el numero de dados con valor más alto de entre los lanzados, y _l_ el más bajo
_Ejemplos_
- Tres dados de 6, uno de 10 y suma 3 al total
'/tira 3d6+1d10+3'
- Tira un dado de 20 y resta 2 al total
'/tira 1d20-2' o '/tira -2'
- Tira cuatro dados de 6 caras y escoge los 3 mayores
'/tira 3h4d6'
`
}

func helpV() string {
	return `
Este atajo hace una tirada de 'ventaja': lanza dos dados de 20 caras y devuelve el valor más alto
No hace falta proporcionar sintaxis de dados, solamente bonificaciones
_Ejemplo:_
- Tira dos dados de veinte y suma 3 al mayor de los resultados
'/v +3' o '/tira h2d20+3'
`
}

func helpDv() string {
	return `
Este atajo hace una tirada de 'ventaja': lanza dos dados de 20 caras y devuelve el valor más bajo
No hace falta proporcionar sintaxis de dados, solamente bonificaciones
_Ejemplo:_
- Tira dos dados de veinte y suma 3 al menor de los resultados
'/dv +3' o '/tira l2d20+3'
`
}

func helpT() string {
	return `
Despliega un teclado interactivo para hacer tiradas sencillas
`
}

func helpAgrupa() string {
	return `
Se usa igual que el comando *tira*, pero el resultado se distribuye mostrando subtotales
Si solo se lanzan grupos de dados y bonificadores, se muestra subtotal por cada grupo de dados y uno último con el subtotal de los bonificadores
Si se insertan etiquetas (_:texto_) los subtotales se harán por etiqueta
_Ejemplos_
- Tira un dado de 8, tres dados de 6 y suma dos bonificadores de 2
'/agrupa 1d8+3d6+2+2' Devolverá un subtotal para 1d8, otro para 3d6 y otro para +2+2
- Mismo ejemplo que antes pero con etiquetas
'/agrupa 1d8+2:espada +3d6+2:furtivo' Devolverá un subtotal para _espada_ (1d8+2) y otro para _furtivo_ (3d6+2)
`
}

func helpAyuda() string {
	return `
Este comando proporciona una breve ayuda sobre cada comando
`
}

func helpRepite() string {
	return `
Lanza una serie de la misma tirada devolviendo los resultados de cada uno
Actualmente el máximo número de repeticiones es 20
_Ejemplo:_
- Tira seis veces un dado de 20 con una bonificación de 3
'/repite 6 1d20+3' o '/tira 1d20+3' 6 veces
`
}
