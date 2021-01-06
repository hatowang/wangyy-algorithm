package insert_delete_getrandom_o1

import "math/rand"

type RandomizedSet struct {
	dic map[int]int
}


/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{dic: make(map[int]int)}
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	var ok bool
	if _, ok = this.dic[val]; !ok {
		this.dic[val] = 0
	}
	return !ok
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	var ok bool
	if _, ok = this.dic[val]; ok {
		delete(this.dic, val)
	}
	return ok
}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.dic))
	var i = 0
	for k, _ := range this.dic {
		if i == index {
			return k
		}
		i++
	}
	return 0
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
