package main

import "strings"

type note struct {
	content string
}

// derive the title from the content
func (n *note) title() string {
	/* split on new line and only have
	   two lines and return the first line */
	return strings.SplitN(n.content, "\n", 2)[0]
}
