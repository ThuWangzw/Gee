package gee

import "strings"

type Trie struct {
	pattern  string
	part     string
	children []*Trie
	isWild   bool
}

func parsePattern(pattern string) []string {
	parts := make([]string, 0)
	splits := strings.Split(pattern, "/")
	for _, part := range splits {
		if part != "" {
			parts = append(parts, part)
			if part[0] == '*' {
				break
			}
		}
	}
	return parts
}

func parsePath(path string) []string {
	parts := make([]string, 0)
	splits := strings.Split(path, "/")
	for _, part := range splits {
		if part != "" {
			parts = append(parts, part)
		}
	}
	return parts
}

func newTrieNode(part string) *Trie {
	newNode := &Trie{
		part:     part,
		children: make([]*Trie, 0),
		isWild:   false,
	}
	if len(part) > 0 && (part[0] == ':' || part[0] == '*') {
		newNode.isWild = true
	}
	return newNode
}

func (root *Trie) insert(pattern string) {
	parts := parsePattern(pattern)
	root.insertIter(pattern, parts, 0)
}

func (root *Trie) insertIter(pattern string, parts []string, depth int) {
	if depth == len(parts) {
		root.pattern = pattern
		return
	}
	var matchChild *Trie
	for _, child := range root.children {
		if child.isWild || child.part == parts[depth] {
			matchChild = child
			break
		}
	}
	if matchChild == nil {
		matchChild = newTrieNode(parts[depth])
		root.children = append(root.children, matchChild)
	}
	matchChild.insertIter(pattern, parts, depth+1)
}

func (root *Trie) search(path string) (*Trie, map[string]string) {
	parts := parsePattern(path)
	return root.searchIter(parts, 0)
}

func (root *Trie) searchIter(parts []string, depth int) (*Trie, map[string]string) {
	if depth == len(parts) {
		if root.pattern == "" {
			return nil, make(map[string]string)
		}
		return root, make(map[string]string)
	}
	for _, child := range root.children {
		if child.isWild && child.part[0] == '*' {
			childparams := map[string]string{
				child.part[1:]: strings.Join(parts[depth:], "/"),
			}
			return child, childparams
		}
		if child.isWild || child.part == parts[depth] {
			node, childparams := child.searchIter(parts, depth+1)
			if node != nil {
				if child.isWild {
					childparams[child.part[1:]] = parts[depth]
				}
				return node, childparams
			}
		}
	}
	return nil, map[string]string{}
}
