package main

/**
PROBLEM:
leetcode -> Simplify Path -> medium

TAGS:
stack, string
**/

type stack []string

func (st *stack) push(val string) {
	(*st) = append(*st, val)
}

func (st *stack) pop() (string, bool) {
	length := len(*st)
	if length == 0 {
		return "", false
	}
	out := (*st)[length-1]
	(*st) = (*st)[:length-1]
	return out, true
}

func (st *stack) isEmpty() bool {
	length := len(*st)
	if length == 0 {
		return true
	} else {
		return false
	}
}

func constructNewPath(st *stack) string {
	var newPath string = "/"
	if st.isEmpty() {
		return newPath
	}
	for dir := range *st {
		newPath = newPath + (*st)[dir] + "/"
	}
	return newPath[:len(newPath)-1]
}

func simplifyPath(path string) string {
	var st stack
	var i int = 0
	var temp string
	if path[0] == '/' {
		i++
	}
	if path[len(path)-1] != '/' {
		path = path + "/"
	}

	for i < len(path) {
		if path[i] != '/' {
			temp = temp + string(path[i])
		} else {
			switch temp {
			case "..":
				st.pop()
			case ".":
				break
			case "":
				break
			default:
				st.push(temp)
			}

			temp = ""
		}
		i++
	}
	return constructNewPath(&st)
}
