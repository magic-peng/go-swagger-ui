package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

//go:embed third_party/swagger-ui/dist/*
var swaggerUI embed.FS

var (
	swaggerPath string
	port        int
)

var rootCmd = &cobra.Command{
	Use:   "swagger-server",
	Short: "启动一个 Swagger UI Web 服务器",
	Long: `使用 Go 的 embed 包嵌入 Swagger UI 静态文件，并通过命令行参数指定 swagger.json 文件的路径。
    
示例:
  swagger-server --swagger ./swagger/swagger.json --port 8080
`,
	Run: func(cmd *cobra.Command, args []string) {
		if swaggerPath == "" {
			log.Fatal("请通过 --swagger 参数指定 swagger.json 文件的路径")
		}

		// 检查 swagger.json 文件是否存在
		if _, err := os.Stat(swaggerPath); os.IsNotExist(err) {
			log.Fatalf("指定的 swagger.json 文件不存在: %s", swaggerPath)
		}

		// 提取子文件系统
		swaggerUIFS, err := fs.Sub(swaggerUI, "third_party/swagger-ui/dist")
		if err != nil {
			log.Fatalf("无法提取 swagger-ui 子文件系统: %v", err)
		}

		// 服务 Swagger UI 静态文件
		fsHandler := http.FileServer(http.FS(swaggerUIFS))
		http.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", fsHandler))

		// 服务 Swagger JSON
		http.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
			absPath, err := filepath.Abs(swaggerPath)
			if err != nil {
				http.Error(w, "无法解析 swagger.json 文件路径", http.StatusInternalServerError)
				return
			}

			// 确保请求方法为 GET
			if r.Method != http.MethodGet {
				http.Error(w, "方法不被允许", http.StatusMethodNotAllowed)
				return
			}

			// 确保文件存在且可读
			file, err := os.Open(absPath)
			if err != nil {
				http.Error(w, "无法打开 swagger.json 文件", http.StatusInternalServerError)
				return
			}
			defer file.Close()

			// 设置正确的 Content-Type
			w.Header().Set("Content-Type", "application/json")

			// 将文件内容写入响应
			http.ServeFile(w, r, absPath)
		})

		// 重定向根路径到 Swagger UI
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/swagger-ui/index.html", http.StatusFound)
		})

		addr := fmt.Sprintf(":%d", port)
		log.Printf("Swagger UI 正在运行在 http://localhost%s", addr)
		log.Fatal(http.ListenAndServe(addr, nil))
	},
}

func main() {
	rootCmd.Flags().StringVarP(&swaggerPath, "swagger", "s", "", "Path to the swagger.json file (required)")
	rootCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to run the server on")
	err := rootCmd.MarkFlagRequired("swagger")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
