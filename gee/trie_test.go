package gee

import (
	"reflect"
	"testing"
)

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/"), []string{})
	ok = ok && reflect.DeepEqual(parsePattern("/info/zhaowei"), []string{"info", "zhaowei"})
	ok = ok && reflect.DeepEqual(parsePattern("/info/*name"), []string{"info", "*name"})
	ok = ok && reflect.DeepEqual(parsePattern("/info/*name/balabala"), []string{"info", "*name"})
	if !ok {
		t.Error("parse pattern wrong")
	}
}

func TestParsePath(t *testing.T) {
	ok := reflect.DeepEqual(parsePath("/"), []string{})
	ok = ok && reflect.DeepEqual(parsePath("/info/zhaowei"), []string{"info", "zhaowei"})
	ok = ok && reflect.DeepEqual(parsePath("/info/*name"), []string{"info", "*name"})
	ok = ok && reflect.DeepEqual(parsePath("/info/*name/balabala"), []string{"info", "*name", "balabala"})
	if !ok {
		t.Error("parse path wrong")
	}
}

func newTestTrie() *Trie {
	trie := newTrieNode("")

	trie.insert("/")
	trie.insert("/person/zhaowei")
	trie.insert("/book/:id/info")
	trie.insert("/static/*filepath")
	return trie
}

func TestTrieInsert(t *testing.T) {
	newTestTrie()
}

func TestTrieSearch(t *testing.T) {
	trie := newTestTrie()
	node, _ := trie.search("/")
	if node == nil || node.pattern != "/" {
		t.Error("failed to search /")
	}

	node, _ = trie.search("/person/zhaowei")
	if node == nil || node.pattern != "/person/zhaowei" {
		t.Error("failed to search /person/zhaowei")
	}

	node, _ = trie.search("/person/another")
	if node != nil {
		t.Error("failed to search /person/another")
	}

	node, _ = trie.search("/person")
	if node != nil {
		t.Error("failed to search /person")
	}

	node, params := trie.search("/book/golangtutorial/info")
	if node == nil || node.pattern != "/book/:id/info" || !reflect.DeepEqual(params, map[string]string{"id": "golangtutorial"}) {
		t.Error("failed to search /book/golangtutorial/info")
	}

	node, params = trie.search("/static/main.js")
	if node == nil || node.pattern != "/static/*filepath" || !reflect.DeepEqual(params, map[string]string{"filepath": "main.js"}) {
		t.Error("failed to search /static/main.js")
	}

	node, params = trie.search("/static/css/style.css")
	if node == nil || node.pattern != "/static/*filepath" || !reflect.DeepEqual(params, map[string]string{"filepath": "css/style.css"}) {
		t.Error("failed to search /static/css/style.css")
	}
}
