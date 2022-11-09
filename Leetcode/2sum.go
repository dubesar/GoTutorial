func twoSum(nums []int, target int) []int {
    mp := map[int]int{} //initialize a map
    ans := []int{} //initialize a arr
    
    for ind, ele := range nums { //looping
        mp[ele] = ind
    }
    for i, e := range nums {
        _, ok := mp[target - e] // checking if the val is present
        if target == 2 * e && mp[target - e] == i {
            continue
        } else if ok {
            ans = append(ans, i)
            ans = append(ans, mp[target - e]) //array append
            break
        }
    }
    return ans
}
