// Package leetcode provides 
package leetcode

func combinationSum3(k int, n int) [][]int {
	var track []int
	var result [][]int
	backTree(n,k,1,&track,&result)
	return result
}

func backTree(n,k,startIndex int, track *[]int, result *[][]int) {
	if len(*track) == k {
			var sum int
			temp := make([]int, k)
			for k,v:=range *track{
					sum+=v
					temp[k]=v
			}
			if sum > n {
					return 
			}
			if sum == n {
					*result = append(*result, temp)
			}
			return
	}

	for i := startIndex;i <=9; i++ {
			*track=append(*track, i)
			backTree(n,k,i+1,track,result)
			*track=(*track)[:len(*track)-1]
	}
}

var res [][]int
func combinationSum3(k int, n int) [][]int {
    res := [][]int{}
    backtrack(1, n, k, 0, []int{})
    return res
}

func backtrack(startIndex, targetSum, k, sum int, track []int) {
    if sum > targetSum {
        return
    }


    if len(track) == k {
        if sum == targetSum {
            temp := make([]int, len(track))
            copy(temp, track)
            res = append(res, temp)
        }
        return
    }

    for i := startIndex; i < 9;i++ {
        sum += i
        track = append(track, i)
        backtrack(startIndex, targetSum, k, sum, track)
        sum-=i
        track = track[:len(track)-1]
    }
}