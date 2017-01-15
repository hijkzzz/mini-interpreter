package ast

type ASTList interface {
	ASTree
}

type astList struct {
	children []ASTree
}

func (self *astList) Child(i int) ASTree {
	return self.children[i]
}

func (self *astList) NumChildren() int{
	return len(self.children)
}

func (self *astList) Children() []ASTree {
	return self.children
}

func (self *astList) Location() string {
	for _, t := range self.children {
		if t.Location() != "" {
			return t.Location()
		}
	}
	return ""
}
