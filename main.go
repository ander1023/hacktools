package main

import (
    "fmt"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/dialog"
    "fyne.io/fyne/v2/layout"
    "fyne.io/fyne/v2/theme"
    "fyne.io/fyne/v2/widget"
    "os/exec"
    "runtime"
)

// Tool 定义工具箱中的工具
type Tool struct {
    Name        string
    Description string
    Icon        fyne.Resource
    Action      func()
}

func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("黑客工具箱")
    myWindow.Resize(fyne.NewSize(800, 600))

    // 定义工具列表
    tools := []Tool{
        {
            Name:        "端口扫描",
            Description: "扫描目标主机的开放端口",
            Icon:        theme.ViewRefreshIcon(),
            Action:      portScan,
        },
        {
            Name:        "子域名发现",
            Description: "查找目标域名的所有子域名",
            Icon:        theme.SearchIcon(),
            Action:      subdomainFinder,
        },
        {
            Name:        "WHOIS 查询",
            Description: "查询域名的注册信息",
            Icon:        theme.InfoIcon(),
            Action:      whoisLookup,
        },
        {
            Name:        "Ping 检测",
            Description: "检测主机是否在线",
            Icon:        theme.ConfirmIcon(),
            Action:      pingCheck,
        },
        {
            Name:        "哈希计算",
            Description: "计算文本的多种哈希值",
            Icon:        theme.DocumentIcon(),
            Action:      hashCalculator,
        },
        {
            Name:        "Base64 编码/解码",
            Description: "Base64 编码和解码工具",
            Icon:        theme.ContentCopyIcon(),
            Action:      base64Tool,
        },
        {
            Name:        "HTTP Header 分析",
            Description: "分析网站的HTTP头部信息",
            Icon:        theme.ListIcon(),
            Action:      httpHeaderAnalyzer,
        },
        {
            Name:        "密码生成器",
            Description: "生成强密码",
            Icon:        theme.ViewRefreshIcon(),
            Action:      passwordGenerator,
        },
    }

    // 创建网格布局
    grid := container.New(layout.NewGridLayout(3))
    
    // 为每个工具创建卡片
    for _, tool := range tools {
        card := createToolCard(tool)
        grid.Add(card)
    }

    // 创建顶部标题
    title := widget.NewLabel("黑客工具箱")
    title.TextStyle = fyne.TextStyle{Bold: true}
    title.Alignment = fyne.TextAlignCenter
    
    // 创建主容器
    content := container.NewBorder(
        container.NewVBox(title, widget.NewSeparator()), 
        nil, nil, nil, 
        grid,
    )

    myWindow.SetContent(content)
    myWindow.ShowAndRun()
}

// createToolCard 创建工具卡片
func createToolCard(tool Tool) fyne.CanvasObject {
    card := widget.NewCard(
        tool.Name,
        tool.Description,
        nil,
    )
    
    button := widget.NewButtonWithIcon("启动", tool.Icon, tool.Action)
    
    return container.NewVBox(
        card,
        layout.NewSpacer(),
        button,
    )
}

// 以下是工具函数的实现（示例实现）

func portScan() {
    showDialog("端口扫描", "端口扫描功能启动\n请输入要扫描的目标IP或域名")
    // 实际实现会包含更多逻辑，如输入框、扫描按钮等
}

func subdomainFinder() {
    showDialog("子域名发现", "子域名发现功能启动")
}

func whoisLookup() {
    showDialog("WHOIS 查询", "WHOIS 查询功能启动")
}

func pingCheck() {
    showDialog("Ping 检测", "Ping 检测功能启动")
}

func hashCalculator() {
    showDialog("哈希计算", "哈希计算功能启动")
}

func base64Tool() {
    showDialog("Base64 工具", "Base64 编码/解码功能启动")
}

func httpHeaderAnalyzer() {
    showDialog("HTTP Header 分析", "HTTP Header 分析功能启动")
}

func passwordGenerator() {
    showDialog("密码生成器", "密码生成器功能启动")
}

// showDialog 显示信息对话框
func showDialog(title, message string) {
    a := app.New()
    w := a.NewWindow(title)
    w.Resize(fyne.NewSize(300, 200))
    
    dialog.ShowInformation(title, message, w)
    
    // 在对话框显示后运行应用
    go func() {
        w.Show()
        a.Run()
    }()
}

// openBrowser 打开系统浏览器（辅助函数）
func openBrowser(url string) {
    var err error
    
    switch runtime.GOOS {
    case "linux":
        err = exec.Command("xdg-open", url).Start()
    case "windows":
        err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
    case "darwin":
        err = exec.Command("open", url).Start()
    default:
        err = fmt.Errorf("unsupported platform")
    }
    
    if err != nil {
        fmt.Println("无法打开浏览器:", err)
    }
}
