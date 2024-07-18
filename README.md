# Project 01 - Player profile

## Description

El proyecto consiste en una API REST que expone una base de datos que maneja 3 entidades:

- User
- PlayerProfile
- Achievement

Un Usuario puede tener varios PlayerProfiles y un PlayerProfile puede tener varios Achievements.

La relación sería la siguiente:	

```mermaid
---
title: DB Players
---

erDiagram
    USER {
        int user_id PK
        string username
        string password
        string email
        int age
    }

    PLAYERPROFILE {
        int player_profile_id PK
        string nickname
        string avatar
        int level
        int experience
        int points
        int user_id FK
    }

    ACHIEVEMENT {
        int achievement_id PK
        string name
        string description
        int player_profile_id FK
    }

    USER ||--o{ PLAYERPROFILE : owns
    PLAYERPROFILE ||--o{ ACHIEVEMENT : unlocks
```

## Tecnologías

- Golang 1.22.1
- Gin Gonic
- Gorm
- PostgreSQL
- Docker
- Docker Compose
- Makefile
- GitHub Actions
- Terraform
- AWS
- Swagger
- SonarCloud
- Synk

La aplicación consta de una API con los respectivos endpoints para cada entidad, como base de datos se utiliza PostgreSQL y se utiliza Docker Compose para levantar la base de datos y la aplicación. Se despliega mediante un Pipeline de GitHub Actions a un entorno de AWS utilizando Terraform.


## Diagrama de Clases

```mermaid
classDiagram
    class User {
        +int ID
        +string Username
        +string Password
        +string Email
        +int Age
        +[]PlayerProfile Profiles
        +CreateUser() error
        +GetUser(id int) (User, error)
        +UpdateUser() error
        +DeleteUser() error
    }

    class PlayerProfile {
        +int ID
        +string Nickname
        +string Avatar
        +int Level
        +int Experience
        +int Points
        +int UserID
        +User User
        +[]Achievement Achievements
        +CreateProfile() error
        +GetProfile(id int) (PlayerProfile, error)
        +UpdateProfile() error
        +DeleteProfile() error
    }

    class Achievement {
        +int ID
        +string Name
        +string Description
        +int PlayerProfileID
        +PlayerProfile PlayerProfile
        +CreateAchievement() error
        +GetAchievement(id int) (Achievement, error)
        +UpdateAchievement() error
        +DeleteAchievement() error
    }

    User "1" --o "*" PlayerProfile : tiene
    PlayerProfile "1" --o "*" Achievement : desbloquea
```

## Endpoints

### User

- **GET /users**: Devuelve todos los usuarios.
- **GET /users/{id}**: Devuelve un usuario por su id.
- **POST /users**: Crea un usuario.
- **PUT /users/{id}**: Actualiza un usuario por su id.

### PlayerProfile

- **GET /player-profiles**: Devuelve todos los perfiles de jugador.
- **GET /player-profiles/{id}**: Devuelve un perfil de jugador por su id.
- **POST /player-profiles**: Crea un perfil de jugador.
- **PUT /player-profiles/{id}**: Actualiza un perfil de jugador por su id.

### Achievement

- **GET /achievements**: Devuelve todos los logros.
- **GET /achievements/{id}**: Devuelve un logro por su id.
- **POST /achievements**: Crea un logro.
- **PUT /achievements/{id}**: Actualiza un logro por su id.

```mermaid
sequenceDiagram
    actor Cliente
    participant Auth
    participant User
    participant PlayerProfile
    participant Achievement

    Cliente->>Auth: POST /auth/login
    Auth-->>Cliente: JWT Token

    Cliente->>User: GET /api/v1/users
    User-->>Cliente: Lista paginada de usuarios

    Cliente->>User: POST /api/v1/users
    User-->>Cliente: Usuario creado

    Cliente->>PlayerProfile: GET /api/v1/player-profiles
    PlayerProfile-->>Cliente: Lista paginada de perfiles de jugador

    Cliente->>PlayerProfile: POST /api/v1/player-profiles
    PlayerProfile-->>Cliente: Perfil de jugador creado

    Cliente->>Achievement: GET /api/v1/achievements
    Achievement-->>Cliente: Lista paginada de logros

    Cliente->>Achievement: POST /api/v1/achievements
    Achievement-->>Cliente: Logro creado

    Cliente->>Auth: POST /auth/logout
    Auth-->>Cliente: Sesión cerrada
```
## Infraestructura de AWS

- **VPC**: Para la configuración de la red virtual privada interna de la aplicación.
- **EC2**: Para la creación de la instancia virtual que alojará la aplicación.
- **RDS**: Para la creación de la base de datos PostgreSQL.
- **Security Groups**: Para la configuración de los puertos de entrada y salida de la aplicación.
- **IAM**: Para la creación de los roles y permisos necesarios para la aplicación.
- **S3**: Para el almacenamiento de los archivos de configuración de Terraform.


## Diagrama de Arquitectura


```mermaid
graph TB
    subgraph "GitHub"
        A[Repositorio de Código]
        B[GitHub Actions]
    end

    subgraph "AWS Cloud"
        subgraph "VPC"
            C[EC2 Instance]
            D[RDS PostgreSQL]
            E[Security Groups]
        end
        F[IAM Roles]
        G[S3 Bucket]
    end

    subgraph "Developer Environment"
        H[Docker Compose]
        I[Local PostgreSQL]
        J[API Application]
    end

    A -->|Trigger| B
    B -->|Deploy| C
    B -->|Configure| F
    B -->|Store State| G
    C -->|Connect| D
    C -->|Uses| E
    C -->|Access| F
    H -->|Local Development| I
    H -->|Local Development| J
    J -->|Deployed to| C
```

