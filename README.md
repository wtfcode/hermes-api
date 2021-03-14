# Inventory System Requirements

## El proyecto base del que nace este sistema es un ERP, un sistema con un nucleo altamente escalable, el actual está escrito en Java EE y cuenta con una versiòn desktop y otra web, se conecta a una base de datos IBM DB2, una muy antigua y su peso es cercano a 2TB.
## El aplicativo debe de soportar alta concurrencia y ser lo más rápido posible en cuanto a su ejecución y funcionamiento, la base de datos puede y debe cambiar a uno más actualizada y de ser necesario se pueden usar varias de ellas. Si se logra un funcionamiento en tiempo real mucho mejor, ya que el aplicativo actual ha presentando fallas por sobreescritura de los usuarios.

### Requerimientos
- Alta velocidad de ejecución
- Soporte de Concurrencia
- Login de Usuario con permisología (Administradors, Superusuarios, etc): Al ser un ERP se divide por departamentos, por lo que cada uno de ellos solo puede acceder a ciertas funcionalidades, al igual que la ubicación geografica de dicho departamento limita al usuario.
- Exportación de gráficos basados en los datos almacenados y de manera general o especifica, a gusto del usuario
- Registro de Facturas en sus distintas areas, cuentas por pagar (CxP), cuentas por cobrar (CxC), borradores, canceladas, divisas o moneda nacional, pago mixto, etc.
- Manejo de inventario de productos varios, por geolocalización, cliente, vendedor y propios. 
- Esta de más informar que el inventario puede aumentar o disminuir de acuerdo a las facturas emitidas.
- Tiempos de llegada de proximos cargamentos, tiempos de salida de exportaciones.


