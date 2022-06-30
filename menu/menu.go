package menu

import "fmt"

// 定义操作符节点类（继承自LinkTable）
type CmdNode struct {
	*LinkTableNode
}

// 定义数字节点类（继承自Linktable）
type NumNode struct {
	*LinkTableNode
}

// 操作符对照表
var cmds = [3]string{"help", "version", "quit"}

// 定义接口
type FindCmder interface {
	FindCmd()
}

// 操作符节点类的接口实现
func (a *CmdNode) FindCmd() string {
	return cmds[a.data]
}

// 数字类的接口实现
func (a *NumNode) FindCmd() int {
	return a.data
}

// 主函数
func test() {
	pLinkTable := CreateLinkTable()
	var dataNums = [3]int{2, 0, 1}
	// 添加节点
	for i := 0; i < len(dataNums); i++ {
		pNode := &LinkTableNode{nil, dataNums[i]}
		if AddLinkTableNode(pLinkTable, pNode) == 0 {
			fmt.Println("Add Node Failed!")
			break
		}
	}
	// 测试接口的实现
	fmt.Println("==========测试CmdNode类的接口实现==========")
	for pTempNode := pLinkTable.pHead; pTempNode != nil; pTempNode = pTempNode.next {
		d := CmdNode{pTempNode}
		fmt.Print(d.FindCmd() + " ")
	}
	fmt.Println()
	fmt.Println("==========测试NumNode类的接口实现==========")
	for pTempNode := pLinkTable.pHead; pTempNode != nil; pTempNode = pTempNode.next {
		d := NumNode{pTempNode}
		fmt.Print(d.FindCmd())
		fmt.Print(" ")
	}
	fmt.Println()
	// 测试删除节点
	fmt.Println("==========测试删除节点==========")
	pHeadNode := GetLinkTableHead(pLinkTable)
	DelLinkTableNode(pLinkTable, pHeadNode)
	for pTempNode := pLinkTable.pHead; pTempNode != nil; pTempNode = pTempNode.next {
		d := NumNode{pTempNode}
		fmt.Print(d.FindCmd())
		fmt.Print(" ")
	}
	fmt.Println()
}
