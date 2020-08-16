# Prueba concepto go con docker y gestor de dependencias

Se crea un fichero de go con un endpoint para pruebas. Se controlan las dependencias y se configura un Dockerfile para que gestione dos imágenes. En una se realizará la construccion y otra será una imágene con el tamaño mas reducido posible en la que irá arrancado el enpoint y expuestos los puertos.

```http
GET localhost:8080?name=test
```

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `name` | `string` | Your name |
