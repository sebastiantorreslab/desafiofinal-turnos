# Sistema de Reserva de Turnos API 🦷

Esta API RESTful gestiona la reserva de turnos para una clínica odontológica. Permite la administración completa de odontólogos y pacientes, así como la asignación y gestión de turnos.

## Características 🌟

- CRUD completo para odontólogos y pacientes.
- Registro y gestión de turnos con odontólogos.
- Seguridad en las operaciones críticas a través de middleware.
- Documentación interactiva de la API con Swagger.

## Entidades 📑

### Dentistas

- **POST** `/dentists`: Agregar un nuevo dentista.
- **GET** `/dentists/{id}`: Obtener datos de un dentista por ID.
- **PUT** `/dentists/{id}`: Actualizar los datos de un dentista.
- **PATCH** `/dentists/{id}`: Actualizar datos específicos de un dentista.
- **DELETE** `/dentists/{id}`: Eliminar un dentista.

### Pacientes

- **POST** `/patients`: Agregar un nuevo paciente.
- **GET** `/patients/{id}`: Obtener datos de un paciente por ID.
- **PUT** `/patients/{id}`: Actualizar los datos de un paciente.
- **PATCH** `/patients/{id}`: Actualizar datos específicos de un paciente.
- **DELETE** `/patients/{id}`: Eliminar un paciente.

### Turnos

- **POST** `/shifts`: Agregar un nuevo turno.
- **GET** `/shifts/{id}`: Obtener un turno por ID.
- **PUT** `/shifts/{id}`: Actualizar un turno.
- **PATCH** `/shifts/{id}`: Actualizar datos específicos de un turno.
- **DELETE** `/shifts/{id}`: Eliminar un turno.
- **POST** `/shifts/patient`: Agregar un turno utilizando DNI del paciente y matrícula del dentista.
- **GET** `/shifts/patient?dni={dni}`: Obtener detalle de turnos por DNI del paciente.

## Seguridad 🔒

La seguridad se maneja con middleware para validar la autenticidad en operaciones de creación, actualización y eliminación.

## Documentación 📖

Accede a la documentación interactiva generada con Swagger para ver todos los detalles de los endpoints y modelos de datos.

## Requerimientos Técnicos 🛠️

La aplicación está estructurada en la siguiente arquitectura de capas:

- **Capa de Entidades de Negocio**: Define las estructuras y modelos de datos.
- **Capa de Acceso a Datos (Repository)**: Abstracción para la interacción con la base de datos.
- **Capa de Acceso a Datos (Base de datos)**: Puede ser una base de datos relacional (H2, MySQL) o no relacional (MongoDB).
- **Capa Service**: Contiene la lógica de negocio.
- **Capa Handler**: Define los controladores de la API que interactúan con las capas service y repository.

---

Desarrollado por [Sebastian Torres y Sebastian Alfonso](https://github.com/sebastiantorreslab/desafiofinal-turnos/tree/main).


# desafiofinal-turnos

- development commits over **develop branch.**
- **main branch** for final delivery.
