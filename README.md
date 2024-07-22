# Project 01 - Player Profile

## Description

The project consists of a REST API that exposes a database managing 3 entities:

- User
- PlayerProfile
- Achievement

A User can have multiple PlayerProfiles, and a PlayerProfile can have multiple Achievements.

The relationship would be as follows:

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

## Technologies

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

The application consists of an API with respective endpoints for each entity. PostgreSQL is used as the database, and Docker Compose is used to set up both the database and the application. Deployment is handled via a GitHub Actions pipeline to an AWS environment using Terraform.

## Class Diagram

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

    User "1" --o "*" PlayerProfile : has
    PlayerProfile "1" --o "*" Achievement : unlocks
```

## Endpoints

### User

- **GET /users**: Returns all users.
- **GET /users/{id}**: Returns a user by their id.
- **POST /users**: Creates a user.
- **PUT /users/{id}**: Updates a user by their id.

### PlayerProfile

- **GET /player-profiles**: Returns all player profiles.
- **GET /player-profiles/{id}**: Returns a player profile by its id.
- **POST /player-profiles**: Creates a player profile.
- **PUT /player-profiles/{id}**: Updates a player profile by its id.

### Achievement

- **GET /achievements**: Returns all achievements.
- **GET /achievements/{id}**: Returns an achievement by its id.
- **POST /achievements**: Creates an achievement.
- **PUT /achievements/{id}**: Updates an achievement by its id.

```mermaid
sequenceDiagram
    actor Client
    participant Auth
    participant User
    participant PlayerProfile
    participant Achievement

    Client->>Auth: POST /auth/login
    Auth-->>Client: JWT Token

    Client->>User: GET /api/v1/users
    User-->>Client: Paginated list of users

    Client->>User: POST /api/v1/users
    User-->>Client: User created

    Client->>PlayerProfile: GET /api/v1/player-profiles
    PlayerProfile-->>Client: Paginated list of player profiles

    Client->>PlayerProfile: POST /api/v1/player-profiles
    PlayerProfile-->>Client: Player profile created

    Client->>Achievement: GET /api/v1/achievements
    Achievement-->>Client: Paginated list of achievements

    Client->>Achievement: POST /api/v1/achievements
    Achievement-->>Client: Achievement created

    Client->>Auth: POST /auth/logout
    Auth-->>Client: Session closed
```

## AWS Infrastructure

- **VPC**: For the internal virtual private network configuration of the application.
- **EC2**: For the creation of the virtual instance hosting the application.
- **RDS**: For the creation of the PostgreSQL database.
- **Security Groups**: For configuring the application’s inbound and outbound ports.
- **IAM**: For creating the necessary roles and permissions for the application.
- **S3**: For storing Terraform configuration files.

## Architecture Diagram

```mermaid
graph TB
    subgraph "GitHub"
        A[Code Repository]
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