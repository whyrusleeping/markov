package markov

import (
	"math/rand"
	"strings"
)

var WordSet map[string]*Node

func init() {
	WordSet = make(map[string]*Node)
}

type Node struct {
	Value string
	Next  []*Link
	Final int
}

type Link struct {
	Weight int
	Target *Node
}

func (n *Node) InsertPhrase(p string) {
	n.Insert(SplitPhrase(p))
}

func (n *Node) Insert(words []string) {
	if len(words) == 0 {
		n.Final++
		return
	}

	for _, l := range n.Next {
		if l.Target.Value == words[0] {
			l.Weight++
			l.Target.Insert(words[1:])
			return
		}
	}

	var next *Node
	if nd, ok := WordSet[words[0]]; ok {
		next = nd
	} else {
		next = &Node{
			Value: words[0],
		}
		WordSet[words[0]] = next
	}

	n.Next = append(n.Next, &Link{
		Weight: 1,
		Target: next,
	})

	next.Insert(words[1:])
}

func (n *Node) Generate() []string {
	return n.generate(nil)
}

func (n *Node) GeneratePhrase() string {
	return strings.Join(n.Generate(), " ")
}

func (n *Node) selectNext() *Node {
	total := n.Final
	for _, l := range n.Next {
		total += l.Weight
	}

	i := rand.Intn(total)
	if i < n.Final {
		return nil
	}

	i -= n.Final
	for _, l := range n.Next {
		if i < l.Weight {
			return l.Target
		}
		i -= l.Weight
	}

	panic("shouldnt get here")
}

func (n *Node) generate(cur []string) []string {
	l := n.selectNext()
	if l == nil {
		return cur
	}

	return l.generate(append(cur, l.Value))
}

func SplitPhrase(p string) []string {
	p = strings.ToLower(p)
	puncts := []string{
		"(",
		")",
		".",
		",",
		":",
		"\"",
		"'",
	}
	for _, c := range puncts {
		p = strings.Replace(p, c, " "+c+" ", -1)
	}
	return strings.Fields(p)
}
