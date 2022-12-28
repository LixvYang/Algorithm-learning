class Solution {
public:
    int majorityElement(vector<int>& nums) {
        int x, vote;
        for (auto num : nums) {
            if (vote == 0) x = num;
            vote += num == x ? 1 : -1;
        }
        return x;
    }
};
