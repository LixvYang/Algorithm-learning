#include <vector>
using namespace std;

class Solution
{
public:
  int removeElement(vector<int> &nums, int val)
  {
    int res = 0;
    for (int i = 0; i < nums.size() - 1; i++)
    {
      if (nums[i] != val)
      {
        nums[res++] = nums[i];
      }
    }
    return res;
  }
};