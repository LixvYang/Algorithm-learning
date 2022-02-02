// Package leetcode provides 
package leetcode

import "strings"
var res [][]string

func isValid(board [][]string, row, col int) (res bool){
    n := len(board)
    for i:=0; i < row; i++ {
        if board[i][col] == "Q" {
            return false
        }
    }
    for i := 0; i < n; i++{
        if board[row][i] == "Q" {
            return false
        }
    }

    for i ,j := row, col; i >= 0 && j >=0 ; i, j = i - 1, j- 1{
        if board[i][j] == "Q"{
            return false
        }
    }
    for i, j := row, col; i >=0 && j < n; i,j = i-1, j+1 {
        if board[i][j] == "Q" {
            return false
        }
    }
    return true
}

func backtrack(board [][]string, row int) {
    size := len(board)
    if row == size{
        temp := make([]string, size)
        for i := 0; i<size;i++{
            temp[i] = strings.Join(board[i],"")
        }
        res =append(res,temp)
        return 
    }
    for col := 0; col < size; col++ {
        if !isValid(board, row, col){
            continue
        }
        board[row][col] = "Q"
        backtrack(board, row+1)
        board[row][col] = "."
    }
}

func solveNQueens(n int) [][]string {
    res = [][]string{}
    board := make([][]string, n)
    for i := 0; i < n; i++{
        board[i] = make([]string, n)
    }
    for i := 0; i < n; i++{
        for j := 0; j<n;j++{
            board[i][j] = "."
        }
    }
    backtrack(board, 0)

    return res
}