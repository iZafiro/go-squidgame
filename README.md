# go-squidgame

# instalacion

Deberia estar todo instalado en las máquinas. En caso de que no funcione RabbitMQ, ejectuar "systemctl start rabbitmq-server" en dist178 y dist179.

# uso

Para cada maquina se debe usar el comando "make [programa]" con los programas correspondientes desde la carpeta go-squidgame.

En dist177 ejecutar "make player" desde en una consola, y en otra consola ejecutar "make datanodeone".

En dist178 ejecutar "make leader" desde en una consola, y en otra consola ejecutar "make datanodetwo".

En dist178 ejecutar "make pool" desde en una consola, y en otra consola ejecutar "make datanodethree".

En dist177 ejecutar "make namenode".