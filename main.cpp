/*
// Sample code to perform I/O:

#include <iostream>

using namespace std;

int main() {
	int num;
	cin >> num;										// Reading input from STDIN
	cout << "Input number is " << num << endl;		// Writing output to STDOUT
}

// Warning: Printing unwanted or ill-formatted data to output will cause the test cases to fail
*/
#include <iostream>
#include <vector>

using namespace std;

void generate(int x, int n, int sum, int index, char* buffer) {
	if (index > n || sum > x)
        return;
    
    if (index == n) {
        if (sum == 0) {
            buffer[index] = '\0';
            cout << buffer << " ";
        }
    }

    for (int i = 0; i < 10; ++i) {
        buffer[index] = i + '0';
        generate(x, n, sum - i, index + 1, buffer);
    }
}

// Write your code here
int main() {
	int x, n;
	cin >> x >> n;
    char buffer[n + 1];

    for (int i = 1; i < 10; ++i) {
        buffer[0] = i + '0';
        generate(x, n, x - i, 1, buffer);
    }
    cout << endl;
}