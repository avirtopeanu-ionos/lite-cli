package paths

import (
	"fmt"
	"github.com/spf13/cobra"
	"sort"
	"strings"
)

func SortByPartCount(paths []string) []string {
	sort.Slice(paths, func(i, j int) bool {
		parts := func(idx int) int {
			return strings.Count(paths[idx], "/") - strings.Count(paths[idx], "{")
		}
		if parts(i) == parts(j) {
			return strings.Count(paths[i], "{") < strings.Count(paths[j], "{")
		}
		return parts(i) < parts(j)
	})
	return paths
}

type VariablePart string // GET datacenters/{datacenterId} 	-> datacenterId is a VariablePart
type SimplePart string   // POST datacenters/servers 		-> 2 simple parts

// TODO: Solve code duplication
func ExtractVariables(path string) []VariablePart {
	var vars []VariablePart
	parts := strings.Split(path, "/")
	for _, part := range parts {
		if strings.Contains(part, "{") && strings.Contains(part, "}") {
			vars = append(vars, VariablePart(part))
		}
	}
	return vars
}

// TODO: Solve code duplication
func ExtractSimpleParts(path string) []SimplePart {
	fmt.Printf("ExtractSimpleParts(%s)\n", path)
	var simpleParts []SimplePart
	parts := strings.Split(path, "/")
	parts = parts[1:]
	for _, part := range parts {
		if !strings.Contains(part, "{") || !strings.Contains(part, "}") {
			simpleParts = append(simpleParts, SimplePart(part))
		}
	}
	return simpleParts
}

type NewCommandOpt func(c *cobra.Command)

func BuildCommands(root *cobra.Command, paths []string) {
	//fmt.Printf("Building cmd chain: %+v\n", paths)

	for _, p := range paths {
		sel := ExtractSimpleParts(p)
		BuildPath(root, sel)
	}

}

func BuildPath(root *cobra.Command, path []SimplePart, opts ...NewCommandOpt) {
	fmt.Printf("BuildCommands(%+v)\n", path)
	fmt.Printf("path[0] = (%s)\n", path[0])
	if len(path) == 0 {
		return
	}
	c := &cobra.Command{
		Use: string(path[0]),
	}

	for _, o := range opts {
		o(c)
	}
	root.AddCommand(c)
}

/// NEW

type commandNode struct {
	name     SimplePart
	children map[SimplePart]*commandNode
	cmd      *cobra.Command
}

func newCommandNode(name SimplePart) *commandNode {
	return &commandNode{
		name:     name,
		children: make(map[SimplePart]*commandNode),
	}
}

func BuildCommandsNew(root *cobra.Command, paths []string) {
	rootNode := newCommandNode("")
	for _, p := range paths {
		parts := ExtractSimpleParts(p)
		addPath(rootNode, parts)
	}

	buildCommandsHelper(root, rootNode)
}

func addPath(rootNode *commandNode, parts []SimplePart) {
	node := rootNode
	for _, part := range parts {
		child, ok := node.children[part]
		if !ok {
			child = newCommandNode(part)
			node.children[part] = child
		}
		node = child
	}
	if node.cmd == nil {
		node.cmd = &cobra.Command{
			Use:   string(node.name),
			Short: fmt.Sprintf("Handle %s operations", node.name),
		}
		if len(parts) > 0 {
			node.cmd.Args = cobra.MaximumNArgs(1)
		}
	}
}

func buildCommandsHelper(parent *cobra.Command, node *commandNode) {
	if node.cmd != nil {
		parent.AddCommand(node.cmd)
	}
	for _, child := range node.children {
		childCmd := &cobra.Command{
			Use:   string(child.name),
			Short: fmt.Sprintf("Handle %s operations", child.name),
		}
		if len(child.children) > 0 {
			buildCommandsHelper(childCmd, child)
		} else if child.cmd != nil {
			childCmd.Args = child.cmd.Args
			childCmd.RunE = child.cmd.RunE
		}
		parent.AddCommand(childCmd)
	}
}
