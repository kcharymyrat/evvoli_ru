# E-Commerce Web Application

This project is a multilingual e-commerce web application and API built using Go. It supports English, Russian, and Turkmen languages and uses PostgreSQL as the database. The project is structured to follow best practices and is designed for scalability and maintainability.

## Project Structure

```plaintext
project_root/
├── api/
│   ├── config/
│   │   └── config.go
│   ├── handler/
│   │   ├── product_handler.go
│   │   ├── user_handler.go
│   │   └── order_handler.go
│   ├── model/
│   │   ├── product.go
│   │   ├── user.go
│   │   └── order.go
│   ├── repository/
│   │   ├── product_repository.go
│   │   ├── user_repository.go
│   │   └── order_repository.go
│   ├── middleware/
│   │   └── auth_middleware.go
│   ├── service/
│   │   └── product_service.go
│   ├── util/
│   │   └── response_util.go
│   └── route.go
├── cmd/
│   ├── api/
│   │   └── main.go
│   └── web/
│       └── main.go
├── internal/
│   ├── helper/
│   │   └── string_helper.go
│   ├── log/
│   │   ├── api.log
│   │   └── web.log
│   ├── mail/
│   │   ├── mailer/
│   │   │   ├── order_mailer.go
│   │   │   └── user_mailer.go
│   │   ├── template/
│   │   │   ├── order_confirmation.html
│   │   │   └── user_welcome.html
│   ├── migration/
│   │   └── 20220501_create_tables.sql
│   ├── validator/
│   │   ├── model/
│   │   │   └── product_validator.go
│   │   └── util/
│   │       └── validation_util.go
│   ├── util/
│   │   └── common_util.go
│   └── deploy/
│       └── deploy.sh
├── web/
│   ├── config/
│   │   └── config.go
│   ├── asset/
│   │   ├── css/
│   │   ├── js/
│   │   └── image/
│   ├── controller/
│   │   ├── product_controller.go
│   │   └── user_controller.go
│   ├── layout/
│   │   └── main_layout.html
│   ├── model/
│   │   └── view_model.go
│   ├── template/
│   │   ├── product_template.html
│   │   └── user_template.html
│   └── util/
│       └── template_util.go
├── go.mod
└── Makefile
```

## Setup

### Prerequisites

- Go (1.22 or later)
- PostgreSQL
- Make

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/project_root.git
    cd project_root
    ```

2. Install dependencies:

    ```sh
    go mod download
    ```

3. Set up the database:
    - Create a PostgreSQL database.
    - Apply migrations:

        ```sh
        psql -U yourusername -d yourdatabase -f internal/migration/20220501_create_tables.sql
        ```

4. Configure environment variables:
    - Create a `.env` file in the root directory and add the necessary environment variables:

        ```env
        DB_HOST=localhost
        DB_PORT=5432
        DB_USER=yourusername
        DB_PASSWORD=yourpassword
        DB_NAME=yourdatabase
        TEMPLATE_PATH=web/template
        STATIC_PATH=web/asset
        ```

### Running the Application

- To run the API:

    ```sh
    cd cmd/api
    go run main.go
    ```

- To run the web application:

    ```sh
    cd cmd/web
    go run main.go
    ```

## Project Components

### API

- **config/**: Contains configuration settings for the API.
- **handler/**: Contains HTTP handlers for API endpoints.
- **model/**: Contains data models for the API.
- **repository/**: Contains database interaction logic for the API.
- **middleware/**: Contains middleware functions for the API.
- **service/**: Contains business logic services for the API.
- **util/**: Contains utility functions for the API.
- **route.go**: Defines the API routes.

### Web

- **config/**: Contains configuration settings for the web application.
- **asset/**: Contains static assets such as CSS, JavaScript, and images.
- **controller/**: Contains HTTP handlers for web pages.
- **layout/**: Contains layout templates for the web pages.
- **model/**: Contains view models for the web application.
- **template/**: Contains HTML templates for the web pages.
- **util/**: Contains utility functions for the web application.

### Internal

- **helper/**: Contains helper functions.
- **log/**: Contains log files.
- **mail/**: Contains mailer logic and templates.
- **migration/**: Contains database migration scripts.
- **validator/**: Contains validation logic for models.
- **util/**: Contains common utility functions.
- **deploy/**: Contains deployment scripts.

## Contributing

Contributions are welcome! Please fork the repository and create a pull request with your changes. Ensure your code adheres to the project's coding standards and includes appropriate tests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For questions or suggestions, please open an issue on GitHub or contact the project maintainer at [your.email@example.com](mailto:your.email@example.com).
