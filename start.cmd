@echo off

rem 在新窗口中启动 Consul 开发模式代理
start "Consul Agent" consul agent -dev


rem 等待一段时间，确保 Consul 启动完成（可根据实际情况调整等待时间）
ping -n 5 127.0.0.1 >nul

rem 进入 product 目录并运行 Go 应用
cd D:\work\master\go_study\workplace\go-mall\app\product
start "Product App" go run .
cd ..

rem 进入 cart 目录并运行 Go 应用
cd D:\work\master\go_study\workplace\go-mall\app\cart
start "Cart App" go run .
cd ..

rem 进入 user 目录并运行 Go 应用
cd D:\work\master\go_study\workplace\go-mall\app\user
start "User App" go run .
cd ..

rem 进入 api 目录并运行 Go 应用
cd D:\work\master\go_study\workplace\go-mall\api
start "API App" go run .
cd ..