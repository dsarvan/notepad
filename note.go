package main

type note struct {
	content string
}

// derive the title from the content
func (n *note) title() string {
	return n.content
}
