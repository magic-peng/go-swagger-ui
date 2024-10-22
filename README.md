# Swagger UI Project

This is a simple project that uses Golang to start Swagger UI.

## Multilingual Versions

- 🇨🇳 [简体中文](./README-cn.md)

## Open Source Projects Used

This project uses the following open source projects:

- [Swagger UI](https://github.com/swagger-api/swagger-ui): A powerful tool for generating and displaying API documentation. Swagger UI makes it easy for users to interact with APIs, view available endpoints, and see request and response examples.

Thanks to the Swagger UI team for their excellent work and contributions!

## Usage Steps

1. **Install**:

    ```bash
    go install github.com/magic-peng/go-swagger-ui
    ```

2. **Start the Server**:

    ```bash
    go-swagger-ui --swagger=./swagger.json --port=8080
    ```

3. **Access Swagger UI**:

   Open [http://localhost:8080](http://localhost:8080) in your browser to view the API documentation.

## License

This project is licensed under the MIT License.
