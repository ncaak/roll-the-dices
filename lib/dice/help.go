package dice


const HELP = `
Comandos válidos para la versión actual:

- */tira* (_dados_) (_bonos_) (_texto_)
_dados_
(opcional, por defecto: 1d20)
Lanzará un número de dados *x* de *y* caras siguiendo el patrón xdy
_bonos_
(opcional)
Sumará o restará valores enteros siguiendo el patrón +x o -x
_texto_
(opcional)
Usará el texto introducido como prefijo del resultado de la tirada

- */v* (_bonos_) (_texto_)
Alias para tirada de ventaja: el valor más alto de dos dados de veinte caras.
Equivale a /tira h2d20

- */dv* (_bonos_) (_texto_)
Alias para tirada de desventaja: el valor más bajo de dos dados de veinte caras.
Equivale a /tira l2d20
`