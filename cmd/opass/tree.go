package main

import (
	"fmt"
	"sort"

	"github.com/disiqueira/gotree"
)

func PrintSimpleTree(nodes []string) {
	root := gotree.New("1Password")

	sort.Strings(nodes)
	for _, node := range nodes {
		root.Add(node)
	}

	fmt.Println(root.Print())
}

func PrintNestedTree(branchName string, nodes []string) {
	root := gotree.New("1Password")
	branch := gotree.New(branchName)

	sort.Strings(nodes)
	for _, node := range nodes {
		branch.Add(node)
	}

	root.AddTree(branch)
	fmt.Println(root.Print())
}

func PrintMapTree(tree map[string][]string) {
	root := gotree.New("1Password")
	branches := make([]string, 0)

	for branchName := range tree {
		branches = append(branches, branchName)
	}

	sort.Strings(branches)

	for _, branchName := range branches {
		nodes := tree[branchName]
		branch := gotree.New(branchName)

		sort.Strings(nodes)
		for _, node := range nodes {
			branch.Add(node)
		}
		root.AddTree(branch)
	}

	fmt.Println(root.Print())

}
