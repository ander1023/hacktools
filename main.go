package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Tool 定义工具箱中的工具
type Tool struct {
	Name        string
	Description string
	Action      func(w fyne.Window)
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
			Action: func(w fyne.Window) {
				dialog.ShowInformation("端口扫描", "端口扫描功能启动", w)
			},
		},
		{
			Name:        "子域名发现",
			Description: "查找目标域名的所有子域名",
			Action: func(w fyne.Window) {
				dialog.ShowInformation("子域名发现", "子域名发现功能启动", w)
			},
		},
		{
			Name:        "WHOIS 查询",
			Description: "查询域名的注册信息",
			Action: func(w fyne.Window) {
				dialog.ShowInformation("WHOIS 查询", "WHOIS 查询功能启动", w)
			},
		},
		{
			Name:        "漏洞扫描",
			Description: "扫描目标系统的安全漏洞",
			Action: func(w fyne.Window) {
				dialog.ShowInformation("漏洞扫描", "漏洞扫描功能启动", w)
			},
		},
		{
			Name:        "密码破解",
			Description: "尝试破解加密的密码",
			Action: func(w fyne.Window) {
				dialog.ShowInformation("密码破解", "密码破解功能启动", w)
			},
		},
		{
			Name:        "网络嗅探",
			Description: "捕获和分析网络流量",
			Action: func(w fyne.Window) {
				dialog.ShowInformation("网络嗅探", "网络嗅探功能启动", w)
			},
		},
		{
			Name:        "SQL注入检测",
			Description: "检测Web应用的SQL注入漏洞",
			Action: func(w fyne.Window) {
				dialog.ShowInformation("SQL注入检测", "SQL注入检测功能启动", w)
			},
		},
		{
			Name:        "XSS检测",
			Description: "检测Web应用的跨站脚本漏洞",
			Action: func(w fyne.Window) {
				dialog.ShowInformation("XSS检测", "XSS检测功能启动", w)
			},
		},
		{
			Name:        "目录遍历",
			Description: "尝试发现Web服务器的隐藏目录",
			Action: func(w fyne.Window) {
				dialog.ShowInformation("目录遍历", "目录遍历功能启动", w)
			},
		},
		{
			Name:        "目录遍历",
			Description: "尝试发现Web服务器的隐藏目录",
			Action: func(w fyne.Window) {
				dialog.ShowInformation("目录遍历", "目录遍历功能启动", w)
			},
		},
		{
			Name:        "目录遍历",
			Description: "尝试发现Web服务器的隐藏目录",
			Action: func(w fyne.Window) {
				dialog.ShowInformation("目录遍历", "目录遍历功能启动", w)
			},
		},
		{
			Name:        "目录遍历",
			Description: "尝试发现Web服务器的隐藏目录",
			Action: func(w fyne.Window) {
				dialog.ShowInformation("目录遍历", "目录遍历功能启动", w)
			},
		},
		{
			Name:        "目录遍历",
			Description: "尝试发现Web服务器的隐藏目录",
			Action: func(w fyne.Window) {
				dialog.ShowInformation("目录遍历", "目录遍历功能启动", w)
			},
		},
		{
			Name:        "目录遍历",
			Description: "尝试发现Web服务器的隐藏目录",
			Action: func(w fyne.Window) {
				dialog.ShowInformation("目录遍历", "目录遍历功能启动", w)
			},
		},
		{
			Name:        "目录遍历",
			Description: "尝试发现Web服务器的隐藏目录",
			Action: func(w fyne.Window) {
				dialog.ShowInformation("目录遍历", "目录遍历功能启动", w)
			},
		},
		{
			Name:        "目录遍历",
			Description: "尝试发现Web服务器的隐藏目录",
			Action: func(w fyne.Window) {
				dialog.ShowInformation("目录遍历", "目录遍历功能启动", w)
			},
		},
		{
			Name:        "目录遍历",
			Description: "尝试发现Web服务器的隐藏目录",
			Action: func(w fyne.Window) {
				dialog.ShowInformation("目录遍历", "目录遍历功能启动", w)
			},
		},
	}

	// 创建网格布局
	grid := container.New(layout.NewGridLayout(3))

	// 为每个工具创建卡片
	for _, tool := range tools {
		card := createToolCard(tool, myWindow)
		grid.Add(card)
	}

	// 创建顶部标题
	title := widget.NewLabel("黑客工具箱")
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	// 创建滚动容器
	scrollContainer := container.NewVScroll(grid)
	scrollContainer.SetMinSize(fyne.NewSize(780, 500))

	// 创建主容器
	content := container.NewBorder(
		container.NewVBox(title, widget.NewSeparator()),
				       nil, nil, nil,
				       scrollContainer,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

// createToolCard 创建工具卡片
func createToolCard(tool Tool, parentWindow fyne.Window) fyne.CanvasObject {
	card := widget.NewCard(
		tool.Name,
		tool.Description,
		nil,
	)

	button := widget.NewButton("启动", func() {
		tool.Action(parentWindow)
	})

	return container.NewVBox(
		card,
		layout.NewSpacer(),
				 button,
	)
}
