#include <cstdint>
#include <vector>
using namespace std;

class Solution
{
public:
  int minSubArrayLen(int target, vector<int> &nums)
  {
    int i = 0, j = 0;
    int subLength = 0;
    int result = INT32_MAX;
    int sum = 0;
    for (; j < nums.size(); j++)
    {
      sum += nums[j];
      if (sum >= target)
      {
        subLength = j - i + 1;
        result = result < subLength ? result : subLength;
        sum -= nums[i++];
      }
    }
    return result != INT32_MAX ? result : 0;
  }
};