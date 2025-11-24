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