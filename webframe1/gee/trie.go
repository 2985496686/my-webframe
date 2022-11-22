package gee

import "strings"

type node struct {
	pattern  string
	part     string
	children []*node
	isVague  bool
}

// 找出第一个与part匹配的子节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isVague {
			return child
		}
	}
	return nil
}

// 找出和part匹配的全部子节点·
func (n *node) matchChildren(part string) []*node {
	children := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isVague {
			children = append(children, child)
		}
	}
	return children
}

// 将一个路由插入到路由树上
func (n *node) insert(pattern string, parts []string, height int) {
	//pattern 中的所有part都已经插入
	if height == len(parts)-1 {
		n.pattern = pattern
		return
	}
	//下一个要匹配的part
	part := parts[height+1]
	//找到子节点中第一个和下一个part匹配的节点
	child := n.matchChild(part)
	//下一个要匹配的part不存在，创建并插入
	if child == nil {
		child = &node{part: part, isVague: part[0] == '*' || part[0] == ':'}
		n.children = append(n.children, child)
	}
	//递归匹配下一个part
	child.insert(pattern, parts, height+1)
}

// 查找路由
func (n *node) search(parts []string, height int) *node {
	//height == len(parts) 判断路由是否全部匹配完毕 或者这是一个模糊匹配
	if height == len(parts)-1 || strings.HasPrefix(n.part, "*") {
		//n.pattern != "" 判断当前路由是否是前缀树上的一个完整路由
		if n.pattern != "" {
			return n
		}
		return nil
	}
	children := n.matchChildren(parts[height+1])
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
