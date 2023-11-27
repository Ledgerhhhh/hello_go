package try_designPatterns

import "fmt"

// Color 实现部分接口
type Color interface {
	Fill() string
}

// RedColor 红色实现
type RedColor struct{}

func (c *RedColor) Fill() string {
	return "红色"
}

// GreenColor 绿色实现
type GreenColor struct{}

func (c *GreenColor) Fill() string {
	return "绿色"
}

// Shape 抽象部分接口
type Shape interface {
	Draw() string
	SetColor(color Color)
}

// Circle 圆形抽象类
type Circle struct {
	color Color
}

func (c *Circle) SetColor(color Color) {
	c.color = color
}

func (c *Circle) Draw() string {
	return fmt.Sprintf("画一个%s的圆形", c.color.Fill())
}

func main() {
	// 创建圆形对象
	circle := &Circle{}

	// 设置颜色
	redColor := &RedColor{}
	circle.SetColor(redColor)

	// 绘制圆形
	fmt.Println(circle.Draw())

	// 切换颜色
	greenColor := &GreenColor{}
	circle.SetColor(greenColor)

	// 绘制圆形
	fmt.Println(circle.Draw())
}
