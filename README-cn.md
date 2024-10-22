# Swagger UI 项目

这是一个使用 Golang 启动 Swagger UI 的简单项目。

## 使用的开源项目

本项目使用了以下开源项目：

- [Swagger UI](https://github.com/swagger-api/swagger-ui): 一个强大的工具，用于生成和展示 API 文档。Swagger UI 使得用户可以方便地与 API 进行交互，查看可用的端点及其请求和响应示例。

感谢 Swagger UI 团队的出色工作和贡献！

## 使用步骤

1. **安装**：

    ```bash
    go install github.com/magic-peng/go-swagger-ui
    ```

2. **启动服务**：

    ```bash
    go-swagger-ui --swagger=./swagger.json --port=8080
    ```

3. **访问 Swagger UI**：

   在浏览器中打开 [http://localhost:8080](http://localhost:8080)，查看 API 文档。

## 许可证

本项目采用 MIT 许可证。
