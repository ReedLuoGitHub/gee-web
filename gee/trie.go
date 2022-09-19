package gee

import "strings"

type node struct {
	pattern  string  // 待匹配路由，例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true
}

func (n *node) isWildFunc(part0 byte) bool {
	return part0 == ':' || part0 == '*'
}

// matchChild 找到第一个成功匹配的节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// matchChildren 找到所有匹配的节点
func (n *node) matchChildren(part string) (nodeSlice []*node) {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodeSlice = append(nodeSlice, child)
		}
	}
	return nodeSlice
}

// insert 将指定节点插入到Trie树中
// pattern /hello/:name
// parts [home :name]
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]

	// 查找子节点
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: n.isWildFunc(part[0])}
		n.children = append(n.children, child)
	}

	child.insert(pattern, parts, height+1)
}

// search 找到指定节点
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
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
