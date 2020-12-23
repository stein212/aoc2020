#include <vector>
#include <stdio.h>
using namespace std;

int main()
{
    char line[28];
    int qs[26];
    int count = 0;
    int groupSize = 0;
    int lineLength;

    for (int i = 0; i < 26; ++i)
        qs[i] = 0;

    while (fgets(line, sizeof(line), stdin) != NULL)
    {
        if (line[0] == '\n')
        {
            for (int i = 0; i < 26; ++i)
            {
                if (qs[i] == groupSize)
                    ++count;
                qs[i] = 0;
            }
            groupSize = 0;
        }
        else
        {
            ++groupSize;
            sscanf(line, "%s%n", line, &lineLength);
            // printf("%s %d\n", line, lineLength);
            for (int i = 0; i < lineLength; ++i)
            {
                ++qs[line[i] - 'a'];
            }
        }
    }

    // printf("counting %d ", count);
    for (int i = 0; i < 26; ++i)
    {
        if (qs[i] == groupSize)
            ++count;
    }
    printf("\n%d\n", count);

    return 0;
}