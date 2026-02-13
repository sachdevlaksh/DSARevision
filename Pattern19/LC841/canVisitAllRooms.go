/*
LeetCode Problem #841: Keys and Rooms
Difficulty: Medium

There are n rooms labeled from 0 to n - 1, and all the rooms are locked except for room 0. Your goal is to visit all the rooms. However, you cannot enter a locked room without having its key.
*/

package LC841

func CanVisitAllRooms(rooms [][]int) bool {

	out := make([]int, len(rooms))
	q := make([]int, 0)
	q = append(q, 0)
	out[0] = 1

	for len(q) > 0 {
		size := len(q)
		for size > 0 {
			qpop := q[0]
			q = q[1:]
			for _, v := range rooms[qpop] {
				if out[v] == 0 {
					out[v] = 1
					q = append(q, v)
				}
			}
			size--
		}
	}

	for _, v := range out {
		if v == 0 {
			return false
		}
	}
	return true
}
