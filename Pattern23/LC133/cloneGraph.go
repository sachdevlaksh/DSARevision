/*
LeetCode Problem #133: Clone Graph
Difficulty: Medium

Given a reference of a node in a connected undirected graph. Return a deep copy (clone) of the graph. Each node in the graph contains a val (int) and a list (List[Node]) of its neighbors.
*/

package LC133

/**
* Definition for a Node.*/
 type Node struct {
     Val int
     Neighbors []*Node
 }


func cloneGraph(node *Node) *Node {

    if node == nil{
        return nil
    }

    seen := make(map[*Node] *Node)

    var dfs func(nod *Node) *Node

    dfs = func(nod *Node) *Node {
        if nod == nil{
            return nil
        }
        if see, ok := seen[nod]; ok{
            return see
        }
        clone := &Node{Val: nod.Val}
        seen[nod] = clone  

        for _, nei := range nod.Neighbors{
            clone.Neighbors = append(clone.Neighbors, dfs(nei))
        }

        return clone
    }
    return dfs(node)
}