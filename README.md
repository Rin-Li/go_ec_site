# EC Site Backend

This is a backend system for an e-commerce (EC) site, built with Go using the Gin framework. It provides APIs for user management, product management, orders, and payments.

## Features

- **User Management**: Registration, login, profile updates, and avatar upload.
- **Product Management**: Product listing, details, searching, and category management.
- **Order Management**: Order creation, listing, and payment processing.
- **Address Management**: CRUD operations for user addresses.
- **Favorites Management**: Add, view, and remove favorite products.
- **Carousel**: Fetch homepage carousel images.
- **Seckill (Flash Sale)**: Special purchase events for limited-time offers.

## Technologies Used

- **Go**: Backend development.
- **Gin**: Web framework for API handling.
- **JWT**: Authentication and authorization.
- **MySQL**: Database for storing user and product data.
- **Docker**: Deployment and environment setup.
- **Redis**: Caching for improved performance.

## Installation

### Prerequisites

- Go 1.18+
- MySQL
- Redis
- Docker (Optional, for containerized deployment)

### Setup

1. Clone the repository:

   ```
   git clone https://github.com/your-repo/go-ec-site.git
   cd go-ec-site
   ```

2. Install dependencies:

   ```
   go mod tidy
   ```

3. Configure the database:

   - Update `config.yaml` with MySQL and Redis connection details.

4. Run the application:

   ```
   go run main.go
   ```

5. Access API via `http://localhost:8080/api/v1`

## API Endpoints

### Public Routes

| Method | Endpoint                | Description             |
| ------ | ----------------------- | ----------------------- |
| POST   | `/api/v1/user/register` | User registration       |
| POST   | `/api/v1/user/login`    | User login              |
| GET    | `/api/v1/carousels`     | List homepage carousels |
| GET    | `/api/v1/products`      | List products           |
| GET    | `/api/v1/product/:id`   | View product details    |
| GET    | `/api/v1/imgs/:id`      | View product images     |
| GET    | `/api/v1/categories`    | List product categories |

### Protected Routes (Require JWT Authentication)

| Method | Endpoint                     | Description              |
| ------ | ---------------------------- | ------------------------ |
| PUT    | `/api/v1/user`               | Update user profile      |
| POST   | `/api/v1/avater`             | Upload avatar            |
| POST   | `/api/v1/user/sending_email` | Send verification email  |
| POST   | `/api/v1/user/valid_email`   | Validate email           |
| POST   | `/api/v1/money`              | Show user balance        |
| POST   | `/api/v1/product`            | Create a product         |
| POST   | `/api/v1/products`           | Search products          |
| GET    | `/api/v1/collection`         | List favorites           |
| POST   | `/api/v1/collection`         | Add favorite             |
| DELETE | `/api/v1/collection`         | Remove favorite          |
| POST   | `/api/v1/address`            | Create address           |
| GET    | `/api/v1/address/:id`        | View address             |
| GET    | `/api/v1/address`            | List addresses           |
| PUT    | `/api/v1/address/:id`        | Update address           |
| DELETE | `/api/v1/address/:id`        | Delete address           |
| POST   | `/api/v1/order`              | Create order             |
| GET    | `/api/v1/orders`             | List orders              |
| GET    | `/api/v1/order/:id`          | View order details       |
| DELETE | `/api/v1/order/:id`          | Cancel order             |
| POST   | `/api/v1/paydown`            | Process payment          |
| GET    | `/api/v1/seckill`            | List Seckill products    |
| POST   | `/api/v1/seckill/:id`        | Purchase Seckill product |

## Running with Docker

```
docker build -t go-ec-site .
docker run -p 8080:8080 go-ec-site
```

## Contribution

Feel free to fork this repository and contribute via pull requests.

## License

This project is licensed under the MIT License.
