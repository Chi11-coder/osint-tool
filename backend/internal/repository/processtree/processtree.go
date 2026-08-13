package processtree

import (
	"slices"
	"strings"

	"example.com/security/internal/models"
)

// processes_createdに存在するプロセスのみ抽出
func Prune(processCreated []string, nodes []models.ProcessTree) []models.ProcessTree {
	result := make([]models.ProcessTree, 0, len(nodes))

	for _, node := range nodes {
		prunedChildren := Prune(processCreated, node.Children)
		if isMatchProcessName(processCreated, node.Name) || len(prunedChildren) > 0 {
			node.Children = prunedChildren
			result = append(result, node)
		}
	}

	return result
}

// processes_createdにprocessNameが存在するか
func isMatchProcessName(processesCreated []string, processName string) bool {
	return slices.Contains(processesCreated, processName)
}

func Build(nodes []models.ProcessTree) []models.ProcessTree {
	seen := make(map[string]struct{}, len(nodes))
	result := make([]models.ProcessTree, 0, len(nodes))

	for _, node := range nodes {
		key := nodeKey(node)
		if _, ok := seen[key]; ok {
			continue
		}

		seen[key] = struct{}{}
		result = append(result, toModel(node))
	}

	return result
}

// プロセス名を結合
func nodeKey(node models.ProcessTree) string {
	sb := new(strings.Builder)
	sb.WriteString("(")
	sb.WriteString(node.Name)
	for _, c := range node.Children {
		sb.WriteString(nodeKey(c))
	}
	sb.WriteString(")")
	return sb.String()
}

func toModel(node models.ProcessTree) models.ProcessTree {
	return models.ProcessTree{
		Name:      node.Name,
		ProcessID: node.ProcessID,
		Children:  Build(node.Children),
	}
}
