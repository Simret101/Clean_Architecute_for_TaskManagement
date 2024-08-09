# Task Manager Refactored Architecture Documentation

## 1. Overview

The Task Manager project is designed as a modular web application using Go. It utilizes the Gin framework for HTTP routing and middleware, and incorporates JWT for authentication and password hashing for security. The architecture follows a clean separation of concerns to ensure maintainability and scalability.

## 2. Folder Structure

```plaintext
task-manager/
├── Delivery/
│   ├── main.go
│   ├── controllers/
│   │   └── controller.go
│   └── routers/
│       └── router.go
├── Domain/
│   └── domain.go
├── Infrastructure/
│   ├── auth_middleWare.go
│   ├── jwt_service.go
│   ├── password_service.go
├── Repositories/
│   ├── task_repository.go
│   └── user_repository.go
├── Usecases/
│   ├── task_usecases.go
│   └── user_usecases.go
└── tests/
    ├── mock_usecases.go                // Mock implementations for use cases
    ├── task_controller_test.go         // Unit tests for task controller
    ├── user_controller_test.go         // Unit tests for user controller
    ├── jwt_service_test.go             // Unit tests for JWT service
    ├── password_service_test.go        // Unit tests for password service
    ├── task_usecases_test.go            // Unit tests for task use cases
    └── user_usecases_test.go            // Unit tests for user use cases
```

## 3. Components

### 3.1. Delivery Layer

- **Purpose**: Handles HTTP requests and responses, manages routing and controllers.
- **Files**:
  - `main.go`: Entry point of the application. Initializes the application and sets up routes.
  - `controllers/`: Contains controllers which process HTTP requests, interact with use cases, and return responses.
  - `routers/`: Manages routing setup and middleware.

### 3.2. Domain Layer

- **Purpose**: Defines core business entities and their interactions.
- **Files**:
  - `domain.go`: Contains definitions for domain entities like `Task` and `User`.

### 3.3. Infrastructure Layer

- **Purpose**: Provides implementations for security, authentication, and other services.
- **Files**:
  - `auth_middleWare.go`: Implements middleware for JWT authentication.
  - `jwt_service.go`: Manages JWT creation and validation.
  - `password_service.go`: Handles password hashing and comparison.

### 3.4. Repositories Layer

- **Purpose**: Manages data persistence and retrieval from storage.
- **Files**:
  - `task_repository.go`: Implements CRUD operations for tasks.
  - `user_repository.go`: Implements CRUD operations for users.

### 3.5. Usecases Layer

- **Purpose**: Contains business logic and interacts with repositories to fulfill use cases.
- **Files**:
  - `task_usecases.go`: Manages task-related business logic.
  - `user_usecases.go`: Manages user-related business logic.

### 3.6. Tests

- **Purpose**: Contains unit and integration tests for different components.
- **Files**:
  - `mock_usecases.go`: Contains mock implementations for use cases to facilitate testing.
  - `task_controller_test.go`: Tests for task controller endpoints.
  - `user_controller_test.go`: Tests for user controller endpoints.
  - `jwt_service_test.go`: Tests for JWT service.
  - `password_service_test.go`: Tests for password service.
  - `task_usecases_test.go`: Tests for task use cases.
  - `user_usecases_test.go`: Tests for user use cases.

## 4. Design Decisions

1. **Separation of Concerns**:
   - **Reasoning**: By dividing the application into distinct layers (Delivery, Domain, Infrastructure, Repositories, Usecases), the architecture ensures that each layer has a single responsibility, which simplifies maintenance and testing.

2. **Use of Gin Framework**:
   - **Reasoning**: Gin provides a fast HTTP web framework with minimal overhead, making it suitable for high-performance applications.

3. **JWT for Authentication**:
   - **Reasoning**: JWT is used for stateless authentication, which simplifies scalability and security.

4. **Password Hashing**:
   - **Reasoning**: Secure password storage and comparison are implemented using hashing to protect user credentials.

5. **Unit Testing and Mocks**:
   - **Reasoning**: Mock implementations of use cases are used in tests to isolate and verify individual components without dependency on external systems.

## 5. Guidelines for Future Development

1. **Adding New Features**:
   - **Follow the Layered Architecture**: Ensure that new features are added in the appropriate layer (e.g., business logic in Usecases, routing in Controllers).
   - **Write Tests**: Add corresponding tests for new features to maintain code quality and functionality.

2. **Refactoring**:
   - **Maintain Consistency**: Refactor code while adhering to the architectural boundaries and naming conventions.
   - **Document Changes**: Update documentation to reflect any changes in the architecture or design.

3. **Error Handling**:
   - **Centralize Error Handling**: Use centralized error handling mechanisms in the Delivery layer to manage errors uniformly.

4. **Security**:
   - **Review Security Practices**: Regularly review and update security practices, especially in authentication and password management.

5. **Performance**:
   - **Profile and Optimize**: Use profiling tools to identify performance bottlenecks and optimize code as needed.

6. **Code Reviews**:
   - **Peer Reviews**: Implement a peer review process to ensure code quality and adherence to architectural guidelines.

## 6. Conclusion

This documentation provides a comprehensive overview of the refactored architecture of the Task Manager project. By adhering to the guidelines and design decisions outlined, future development can be streamlined, ensuring maintainability, scalability, and robustness of the application.
POStMAN DOCUMENTATION: https://documenter.getpostman.com/view/37289771/2sA3rzKsPp