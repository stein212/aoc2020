#include <vector>
#include <stdio.h>
using namespace std;

int boardingPassRowToSeatId(char *row)
{
    int x = 0, i = 0;

    for (; i < 7; ++i)
    {
        x <<= 1;
        if (row[i] == 'B')
        {
            x += 1;
        }
    }

    for (; i < 10; ++i)
    {
        x <<= 1;
        if (row[i] == 'R')
        {
            x += 1;
        }
    }

    return x;
}

int main()
{
    // char line[11];
    // int max = 0;
    // while (scanf("%s", line) != EOF)
    // {
    //     int row = boardingPassRowToSeatId(line);
    //     if (row > max)
    //     {
    //         max = row;
    //     }
    // }
    // printf("%d\n", max);

    char line[11];
    vector<int> seatIds;

    while (scanf("%s", line) != EOF)
    {
        int row = boardingPassRowToSeatId(line);
        // printf("%d, ", row);
        seatIds.push_back(row);
    }

    sort(seatIds.begin(), seatIds.end());

    int i = seatIds[0];
    for (vector<int>::iterator it = seatIds.begin(); it != seatIds.end(); ++it)
    {
        if (*it != i)
        {
            printf("%d\n", i);
            break;
        }
        ++i;
    }

    return 0;
}