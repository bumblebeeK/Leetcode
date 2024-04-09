package solutiongo

/*
 * @lc app=leetcode.cn id=1600 lang=golang
 *
 * [1600] 王位继承顺序
 */

// @lc code=start
type ThroneInheritance struct {
	root  *Successor
	index map[string]*Successor
}

type Successor struct {
	Children []*Successor
	IsDeath  bool
	Name     string
}

func Constructor(kingName string) ThroneInheritance {
	k := &Successor{Name: kingName}
	return ThroneInheritance{root: k, index: map[string]*Successor{
		kingName: k,
	}}
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	s := &Successor{
		Name: childName,
	}
	this.index[parentName].Children = append(this.index[parentName].Children, s)
	this.index[childName] = s
}

func (this *ThroneInheritance) Death(name string) {
	this.index[name].IsDeath = true
}

func (this *ThroneInheritance) GetInheritanceOrder() []string {
	ans := []string{}
	var preOrder func(node *Successor)
	preOrder = func(node *Successor) {
		if !node.IsDeath {
			ans = append(ans, node.Name)
		}
		for _, v := range node.Children {
			preOrder(v)
		}
	}
	preOrder(this.root)

	return ans
}

/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * obj := Constructor(kingName);
 * obj.Birth(parentName,childName);
 * obj.Death(name);
 * param_3 := obj.GetInheritanceOrder();
 */
// @lc code=end
