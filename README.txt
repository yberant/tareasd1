NOTA: ESTE MISMO README ESTARÁ PRESENTE EN TODAS LAS MÁQUINAS 

Nombre equipo: "pero ella me levantó"
Integrantes:
- Diego Valderas
- Yoel Berant

Distribución de "servicios" por máquina virtual
- Cliente: root@dist90
- Logística: root@dist89
- Camión: root@dist91
- Financiero: root@dist92

Para echar a andar cada servicio, se debe ejecutar el comando "make run" desde el directorio
en el que se encuentra su MakeFile (el cual es el mismo directorio en el que se
encuentra este readme) en su respectiva máquina virtual

Para hacer andar el sistema:

1- PRIMERO se debe echar a andar el servicio de Logística, una vez ejecutado
el makefile se imprimirá por consola la dirección de IP de donde está corriendo
Logística. Se sugiere antotar esta dirección o "copiarla".

Una vez se empiece a ejecutar logística se creará en la carpeta "logística" 
(en la máquina que este corriendo obviamente) el archivo "logs.csv", donde se guardará el registro de todos los pedidos echos por clientes.

IMPORTANTE: Los números de puertos serán definidos por el usuario:
Desde logística se pedirá ingresar (uno por uno) el número de puerto
en el que se escuchará al camión normal y los dos números de puerto en los
que se escuchará a cada camión de retail. Cuando un camión se conecte con logística, esta le informará al camión si es de tipo normal o retail según el número de puerto en el que esté conectado.

Una vez ingresados los 3 números de puerto desde los cuales logística escuchará a los camiones
se podrá ingresar a logística dos comandos:
 - "quit": se mata a la ejecución de logística
 - "listen"

Para que se pueda establecer una conexión etre logística y un cliente se debe ingresar "listen". Cada vez que se ingrese "listen" se pedirá ingresar un número de puerto en el que se escuchará a un cliente.

Si se pretende escuchar a más de un cliente, se debe ingresar "listen" y otro número de puerto.

Demás está decir que todos los números de puerto ingresados obviamente deben estar disponibles y no se deben repetir. Además, no se debe usar el 5672, pues ese está reservado para la conexión de rabbitMQ con el servicio financiero

Una vez ingresados los números de puerto, será posible conectar a logística camiones y clientes

2- Financiero
Ejecute el servicio financiero, ingrese la dirección IP de logística (el número de puerto ya está seleccionado: 5672) y el servicio escuchará a la cola de RabbitMq. 

Cada vez que el servicio financiero reciba información de un paquete recibido o no recibido calculará la ganancia/pérdida y los costos (según el número de reintentos) asignados a la entrega, actualizando el archivo "finanzas.csv", localizado en la carpeta "financiero", en la máquina del servicio.

3- Para ingresar a un camión, se debe ejecutar un servicio de camión. Por pantalla se pedirá ingresar la dirección de IP de logística y uno de los números de puerto en el que se está escuchando a un camión.
Estos números de puerto debieron ser ya ingresados a logística previamente (ver paso 1). Según el número de puerto se le asignará al camión su tipo ("normal" o "retail") y su ID ("CamionA"(normal), "CamionB"(retail) o "CamionC"(retail))

Luego de que se establesca exitosamente la conexión entre un camión y logística, se creará un archivo csv  en la carpeta "camion" de nombre "registro<ID Camion>.csv" (obviamente en la máquina donde se encuentra el servicio de camión) con sus registros.

Se pedirá desde la consola del camión que se ingrese su tiempo de espera (cuanto se demora en intentar entregar un paquete y cuanto espera a un segundo paquete en logística)

4- Para ingresar un cliente se debe ejecutar un servicio de cliente

Primero se pedirá ingresar la dirección IP del servicio de logística y un número de puerto desde el cual logistica esta escuchando a un cliente, que previamente debió haber sido ingresado a logística con el comando "listen" a logística (ver paso 1). 

Luego se pedirá el tipo de cliente, para eso debe ingresarse por consola las opciones "0", "1" y "2":
 * "0" (Pyme): si se elige esta opción, el servicio realizará pedidos de tipo pyme, desde el archivo pymes.csv.
 * "1" (Retail) si se elige esta opción, el servicio realizará pedidos de tipo retail, desde el archivo retail.csv.
 * "2" (Seguidor) si se elige esta opción, el servicio ofrecerá la opción de solicitar estados de productos, ingresando por consola sus codigos de seguimiento.
 
 ...



