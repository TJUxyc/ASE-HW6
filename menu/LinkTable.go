package menu

// LinkTable节点类定义
type LinkTableNode struct {
	next *LinkTableNode
	data int
}

// LinkTable链表类定义
type LinkTable struct {
	pHead     *LinkTableNode
	pTail     *LinkTableNode
	SumOfNode int
	// mutex_t   Mutex
}

func CreateLinkTable() *LinkTable {
	var pLinkTable LinkTable
	pLinkTable.pHead = nil
	pLinkTable.pTail = nil
	pLinkTable.SumOfNode = 0
	return &pLinkTable
}

// DeleteLinkTable方法在go语言中无需实现，内存将自动回收

// 添加链表节点，创建成功返回1，否则返回0
func AddLinkTableNode(pLinkTable *LinkTable, pNode *LinkTableNode) int {
	if pLinkTable == nil || pNode == nil {
		return 0
	}
	pNode.next = nil
	if pLinkTable.pHead == nil {
		pLinkTable.pHead = pNode
	}
	if pLinkTable.pTail == nil {
		pLinkTable.pTail = pNode
	} else {
		pLinkTable.pTail.next = pNode
		pLinkTable.pTail = pNode
	}
	pLinkTable.SumOfNode++
	return 1
}

// 删除链表节点，删除成功返回1，否则返回0
func DelLinkTableNode(pLinkTable *LinkTable, pNode *LinkTableNode) int {
	if pLinkTable == nil || pNode == nil {
		return 0
	}
	if pLinkTable.pHead == pNode {
		pLinkTable.pHead = pLinkTable.pHead.next
		pLinkTable.SumOfNode--
		if pLinkTable.SumOfNode == 0 {
			pLinkTable.pTail = nil
		}
		return 1
	}
	for pTempNode := pLinkTable.pHead; pTempNode != nil; pTempNode = pTempNode.next {
		if pTempNode.next == pNode {
			pTempNode.next = pTempNode.next.next
			pLinkTable.SumOfNode--
			if pLinkTable.SumOfNode == 0 {
				pLinkTable.pTail = nil
			}
			return 1
		}
	}
	return 0
}

// 获取链表首节点
// ps.此方法似乎有些多余，直接调用 pLinkTable.head 就可得到结果
func GetLinkTableHead(pLinkTable *LinkTable) *LinkTableNode {
	// TODO: 这里如果 pLinkTable指针为空或者pLinkTable节点为空都是返回nil，应当抛出一个异常
	if pLinkTable == nil {
		return nil
	}
	return pLinkTable.pHead
}

// 获取下一个链表节点
// ps.此方法似乎有些多余，直接调用 pNode.next 就可完成函数操作
func GetNextLinkTableNode(pLinkTable *LinkTable, pNode *LinkTableNode) *LinkTableNode {
	if pLinkTable == nil || pNode == nil {
		return nil
	}
	for pTempNode := pLinkTable.pHead; pTempNode != nil; pTempNode = pTempNode.next {
		if pTempNode == pNode {
			return pTempNode.next
		}
	}
	return nil
}
