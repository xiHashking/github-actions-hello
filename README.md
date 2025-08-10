# GitHub Actions Hello World 项目

这是一个使用 GitHub Actions 每5分钟自动运行的 Go 脚本示例。

## 项目结构
```
github-actions-hello/
├── .github/
│   └── workflows/
│       └── hello-world.yml    # GitHub Actions 配置
├── main.go                    # Go 脚本
├── go.mod                     # Go 模块文件
└── README.md                  # 项目说明
```

## 功能
- 每5分钟自动运行一次
- 打印当前时间和 Hello World 消息
- 支持手动触发
- 代码推送时自动运行测试

## 使用方法
1. 推送到 GitHub 仓库
2. 在 Actions 页面查看运行结果
3. 可以手动点击 "Run workflow" 立即测试