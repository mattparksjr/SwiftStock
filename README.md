# SwiftStock by LexiumTech

Swift Stock is an open source Warehouse Management System (WMS) developed by LexiumTech. This system is designed to streamline inventory management, order processing, and overall warehouse operations.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Inventory Management:**

  - Track and manage your warehouse inventory efficiently.
  - Add, edit, and view product details.

- **Order Processing:**

  - Process and fulfill customer orders.
  - Track order status and generate shipping labels.

- **User Management:**

  - Manage user roles, permissions, and access.

- **Reports and Analytics:**
  - Visualize key performance indicators (KPIs) for better decision-making.

## Getting Started

### Prerequisites

- Go (at least version 1.20)
- PostgresSQL(or equivalent)
- NPM

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/mattparksjr/SwiftStock
   cd SwiftStock
   ```
2. Setup the database
   ```
   todo
   ```
3. Build and run the application
   ```
    ./build
   ```

### Usage

1. Access SwiftStock through your webbrowser
2. Login with credentials
3. Nagivate through the different sections to manage warehouse operations

### Project Structure

SwiftStock holds both its server and client facing code in this repository.

- /ss-server holds the go based api server
- /ss-client holds the react based front end

### Testing

For ss-sever

```bash
cd ss-server
go test ./
```

For ss-client

```bash
cd ss-client
npm test
```

### Contributing

Contributions are welcome! If you have any suggestions or need to report an issue please open an issue or submit a pull request. For mission critical security issues, contact me directly!

### License

See LICENSE
