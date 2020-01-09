package prototype

import (
	"fmt"
	"testing"
)

func TestPrototype(t *testing.T) {
	directorNode := NewElement("工程师主任")

	engManagerNode := NewElement("工程师经理")
	engManagerNode.AddChild(NewElement("软件工程师技术负责人"))
	engManagerNode.AddChild(NewElement("项目负责人"))

	directorNode.AddChild(engManagerNode)

	officeManagerNode := NewElement("办公室经理")
	directorNode.AddChild(officeManagerNode)

	fmt.Println("")
	fmt.Println("# 公司职能层次结构：")
	fmt.Print(directorNode)
	fmt.Println("")
	fmt.Println("# 团队成员层次结构")
	fmt.Print(engManagerNode.Clone())
	fmt.Println("")
}
