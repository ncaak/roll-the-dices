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

- */agrupa* (_dados_) (_bonos_) (_etiqueta_) (_texto_)
Resuelve tiradas igual que el comando *tira* pero devuelve el resultado distribuido por bloques.
_etiqueta_
(opcional)
Agrupa la tirada de dados anterior a la propia etiqueta mostrando un subtotal. Se pueden incluir más de una etiqueta, cada bloque irá acompañado de la tirada y su subtotal correspondiente. En este modo, no se mostrará total de la tirada. Sigue el patrón :texto

- */v* (_bonos_) (_texto_)
Alias para tirada de ventaja: el valor más alto de dos dados de veinte caras.
Equivale a /tira h2d20

- */dv* (_bonos_) (_texto_)
Alias para tirada de desventaja: el valor más bajo de dos dados de veinte caras.
Equivale a /tira l2d20

- */t*
Despliega una botonera con las opciones más comunes de tiradas simples.
Al pulsar uno de los botones resolverá la tirada correspondiente.
`
