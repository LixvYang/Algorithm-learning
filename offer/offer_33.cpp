class Solution {
public:
    bool verifyPostorder(vector<int>& postorder) {
        return dfs(postorder, 0, postorder.size()-1);
    }

    bool dfs(vector<int>& postorder, int i, int j) 
    {
        if (i >= j) return true;

        int p = i;
        while(postorder[p] < postorder[j]) p++;
        int t = p;
        while(postorder[p] > postorder[j]) p++;
        return p == j && dfs(postorder, i, t-1) && dfs(postorder, t, j-1);
    }
};