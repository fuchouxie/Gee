package gee

import "strings"

type node struct {
	pattern  string  //待匹配的路由
	part     string  //路由中的一部分
	children []*node //子结点
	isWild   bool    //是否精确匹配
}

// 取出当前结点孩子列表中与part相同或可模糊匹配的那个，不存在返回空
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) insert(pattern string, parts []string, height int) {
	//1.判断是否已到parts数组的最后一个
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	//2.取出分解后url的当前部分
	part := parts[height]
	//3.判断当前部分是否已在该层级存在，存在则取出
	child := n.matchChild(part)
	//4.不存在则进行构建
	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		res := child.search(parts, height+1)
		if res != nil {
			return res
		}
	}
	return nil
}
