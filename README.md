# EC Site Backend

This is a backend system for an e-commerce (EC) site, built with Go using the Gin framework. It provides APIs for user management, product management, orders, and payments.

It is mainly referenced from this https://www.bilibili.com/video/BV1Zd4y1U7D8?spm_id_from=333.788.videopod.episodes&vd_source=461c37833af8427a514d9f0b39901286.

In addition to the original implementation, we have added a "Second Kill" feature.

## Second Kill feature
Process for second kill
<img width="760" alt="Screenshot 2025-02-17 at 23 48 15" src="https://github.com/user-attachments/assets/3107290f-abba-4ce4-ad1c-064fe950382a" />

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

## Structure
.
├── README.md
├── api
│   └── v1
│       ├── address.go
│       ├── carousel.go
│       ├── catrgory.go
│       ├── common.go
│       ├── favorites.go
│       ├── order.go
│       ├── pay.go
│       ├── product.go
│       ├── product_img.go
│       ├── seckill.go
│       └── user.go
├── cache
│   ├── common.go
│   ├── key.go
│   ├── product.go
│   └── seckill.go
├── cmd
│   └── main.go
├── conf
│   ├── config.go
│   └── config.ini
├── dao
│   ├── address.go
│   ├── carousel.go
│   ├── category.go
│   ├── favorite.go
│   ├── init.go
│   ├── migration.go
│   ├── notice.go
│   ├── order.go
│   ├── product.go
│   ├── productimg.go
│   ├── seckill.go
│   └── user.go
├── go.mod
├── go.sum
├── logs
├── middleware
│   ├── cors.go
│   └── jwt.go
├── model
│   ├── address.go
│   ├── admin.go
│   ├── base.go
│   ├── carousel.go
│   ├── cart.go
│   ├── category.go
│   ├── favorite.go
│   ├── notice.go
│   ├── order.go
│   ├── product.go
│   ├── product_img.go
│   ├── seckill.go
│   └── user.go
├── output.md
├── pkg
│   ├── e
│   │   ├── code.go
│   │   └── msg.go
│   └── util
│       ├── encrypt.go
│       ├── jwt.go
│       ├── logger.go
│       └── preheat.go
├── routes
│   └── router.go
├── serializer
│   ├── address.go
│   ├── carousel.go
│   ├── category.go
│   ├── common.go
│   ├── favorite.go
│   ├── money.go
│   ├── order.go
│   ├── product.go
│   ├── product_img.go
│   ├── seckill.go
│   └── user.go
├── service
│   ├── address.go
│   ├── carousel.go
│   ├── category.go
│   ├── favorite.go
│   ├── order.go
│   ├── paydown.go
│   ├── product.go
│   ├── product_img.go
│   ├── seckill.go
│   ├── upload.go
│   └── user.go
└── static
    └── imgs

## Contribution

Feel free to fork this repository and contribute via pull requests.



