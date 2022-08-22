#include <string>
using namespace std;

class Solution
{
public:
  string replaceSpace(string s)
  {
    int count = 0;
    int oldSize = s.size();
    for (int i = 0; i < s.size(); i++)
    {
      if (s[i] == ' ')
      {
        count++;
      }
    }
    int newSize = oldSize + 2 * count;
    s.resize(newSize);
    for (int i = newSize - 1, j = oldSize - 1; j < i; i--, j--)
    {
      if (s[j] != ' ')
      {
        s[i] = s[j];
      }
      else
      {
        s[i] = '0';
        s[i - 1] = '2';
        s[i - 2] = '%';
        i -= 2;
      }
    }
    return s;
  }
};