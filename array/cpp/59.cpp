#include <vector>
using namespace std;

class Solution
{
public:
  vector<vector<int>> generateMatrix(int n)
  {
    int top = 0;
    int bottom = n - 1;
    int left = 0;
    int right = n - 1;
    int num = 1;
    vector<vector<int>> res(n, vector<int>(n, 0));
    int tar = n * n;
    while (num <= tar)
    {
      for (int i = left; i <= right; i++)
      {
        res[top][i] = num;
        num++;
      }
      top++;
      for (int i = top; i <= bottom; i++)
      {
        res[i][right] = num;
        num++;
      }
      right--;
      for (int i = right; i >= left; i--)
      {
        res[bottom][right] = num;
        num++;
      }
      bottom--;
      for (int i = bottom; i >= top; i--)
      {
        res[i][left] = num;
        num++;
      }
      left++;
    }
    return res;
  }
};