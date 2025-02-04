package medium

// Time: O(v + e)
// Space: O(v)
func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	inDegree := make(map[string]int)
	graph := make(map[string][]string)

	for i := range recipes {
		inDegree[recipes[i]] = len(ingredients[i])
		for _, ingredient := range ingredients[i] {
			graph[ingredient] = append(graph[ingredient], recipes[i])
		}
	}

	var result []string
	for len(supplies) > 0 {
		supply := supplies[0]
		supplies = supplies[1:]

		for i := range graph[supply] {
			ngh := graph[supply][i]
			inDegree[ngh]--
			if inDegree[ngh] == 0 {
				result = append(result, ngh)
				supplies = append(supplies, ngh)
			}
		}
	}

	return result
}
