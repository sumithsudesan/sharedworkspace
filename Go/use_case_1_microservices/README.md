**Overview**

This repository contains the source code for multiple microservices related to order management and authentication. Below is a breakdown of each component:

**order_service**

- **Features**:
  - **HTTP Service**: Accepts JSON data for order management operations.
  - **Endpoints**:
    - Place Order: Allows users to place orders.
    - Get Order List: Retrieves a list of orders associated with the user.
    - Delete Order: Deletes an order based on order ID.
  - **Authentication**: Validates users using JWT tokens.
  - **Integration**: Communicates with `order_database_service` via gRPC to store and retrieve order data.
  - **Authorization**: Interacts with `auth_service` for user authorization.

**order_database_service**

- **Features**:
  - **Database Service**: Manages order data.
  - **Protocol**: Uses gRPC for communication.
  - **Database Support**: Utilizes either MongoDB or PostgreSQL based on configuration.


**auth_service**

- **Features**:
  - **Authentication Service**: Provides registration and authentication functionalities.
  - **Protocol**: Utilizes gRPC for communication.
  - **Data Management**: Relies on `auth_database_service` to manage user data.

**auth_database_service**

- **Features**:
  - **Database Service**: Manages user data.
  - **Protocol**: Implements gRPC for communication.
  - **Database Support**: Supports MySQL or PostgreSQL based on configuration.

**Service Interactions**

**User** <---jSOn---> **order_service** <---auth proto---> **auth_service** <--auth proto--> **auth_database_service**

**Description**:
Users interact with `order_service` via JSON payloads. `order_service` communicates with `auth_service` using a protocol buffer (proto) defined for authentication. `auth_service` in turn interacts with `auth_database_service` using another proto for user data management.

**Details**:
- **User <-> order_service**: Users interact with `order_service` through HTTP JSON requests, which allows them to place orders, retrieve order information, and perform other order-related operations.
- **order_service <-> auth_service**: `order_service` sends authentication-related requests to `auth_service` using a protocol buffer (`auth.proto`), which includes operations like validating JWT tokens or verifying user permissions.
- **auth_service <-> auth_database_service**: `auth_service` interacts with `auth_database_service` using another protocol buffer (`auth.proto`), handling operations such as user authentication, registration, and user data management.
- **order_service <-> order_database_service**: `order_service` uses gRPC to communicate with `order_database_service`, sending requests defined in `order.proto` to store, retrieve, update, or delete order information in MongoDB or PostgreSQL, depending on the configuration.

**Directory Structure**

**order_service/**
├── cmd/
│   ├── main.go
├── build/
│   ├── Dockerfile.order-service
├── deploy/
│   ├── order-service.yaml
│   ├── configmap.yaml
├── data-definitions/<submodule>
│   ├── auth/
│   │    ├──auth.proto
│   ├── order/
│   │    ├──order.proto
│   ├── user/
│   │    ├──user.proto
└── pkg/order/
│   ├── authserviceclient.go
│   ├── configuration.go
│   └── model.go
│   └── orderdatabasegrpcclient.go
│   └── server.go

**order_database_service/**
├── cmd/
│   ├── main.go
├── build/
│   ├── Dockerfile.order-db-service
├── deploy/
│   ├── order-database-service.yaml
│   ├── mongodb.yaml
│   ├── configmap.yaml
├── data-definitions/<submodule>
│   ├── auth/
│   │    ├──auth.proto
│   ├── order/
│   │    ├──order.proto
│   ├── user/
│   │    ├──user.proto
└── pkg/database/
│   ├── databaseservice.go
│   ├── configuration.go
│   └── mongodatabase.go
│   └── orderrepository.go
│   └── postgressdatabase.go
│   └── server.go

**auth_service/**
├── cmd/
│   ├── main.go
├── build/
│   ├── Dockerfile.auth-service
├── deploy/
│   ├── order-service.yaml
│   ├── configmap.yaml
├── data-definitions/<submodule>
│   ├── auth/
│   │    ├──auth.proto
│   ├── order/
│   │    ├──order.proto
│   ├── user/
│   │    ├──user.proto
└── pkg/auth/
│   ├── authserviceclient.go
│   ├── authservice.go
│   └── configuration.go
│   └── model.go
│   └── server.go

**auth_database_service/**
├── cmd/
│   ├── main.go
├── build/
│   ├── Dockerfile.auth-db-service
├── deploy/
│   ├── order-service.yaml
│   ├── configmap.yaml
├── data-definitions/<submodule>
│   ├── auth/
│   │    ├──auth.proto
│   ├── order/
│   │    ├──order.proto
│   ├── user/
│   │    ├──user.proto
└── pkg/database/
│   ├── authdbserviceserver.go
│   ├── mysqldatabase.go
│   └── configuration.go
│   └── userrepository.go
│   └── server.go

**data-definitions/**
├── auth/
│    ├──auth.proto
├── order/
│    ├──order.proto
├── user/
│    ├──user.proto

**data-definitions**:
Organize the shared .proto files in a separate repository.

In `auth_service` and `auth_data_service` repositories, add the proto-definitions repository as a submodule.

Navigate to the root of your service repository (`auth_database_service`) and run:

```bash
git submodule add https://github.com/your-org/data-definitions.git data-definitions
protoc -I data-definitions/ --go_out=. --go-grpc_out=. data-definitions/auth/auth.proto
protoc -I data-definitions/ --go_out=. --go-grpc_out=. data-definitions/user/user.proto
