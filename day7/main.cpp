#include <iostream>
#include <fstream>
#include <vector>
#include <sstream>
#include <regex>
#include <map>

using namespace std;

vector<string> split(const string &s, string delimiter)
{
    vector<string> tokens;
    std::size_t current, previous = 0;
    current = s.find(delimiter);
    while (current != string::npos)
    {
        tokens.push_back(s.substr(previous, current - previous));
        previous = current + 1;
        current = s.find(delimiter, previous);
    }

    if (tokens.size() < 1)
        tokens.push_back(s.substr(0, string::npos));

    return tokens;
}

void dfs(const map<string, vector<string> > & bagNeighbours, map<string, int> & visited, const string bagColor, int *sum) {
    // cout << bagColor << "\n";
    visited.insert(pair<string, int>(bagColor, 1));

    if (bagNeighbours.find(bagColor) == bagNeighbours.end())
        return;

    auto neighbours = bagNeighbours.find(bagColor)->second;

    for (auto const & v : neighbours)
    {
        if (visited.find(v) == visited.end()) {
            ++*sum;
            dfs(bagNeighbours, visited, v, sum);
        }
    }
}

int count(const map<string, vector<pair<string, int> > > & bagNeighbours, const string bagColor) {
    int sum = 0;

    auto neighbours = bagNeighbours.find(bagColor)->second;
    if (neighbours.size() == 0)
        return 0;

    for (auto const & v : neighbours) {
        cout << v.second << " " << v.first << "\n";
        sum += v.second * count(bagNeighbours, v.first) + v.second;
    }

    return sum;
}

int main(void)
{
    ifstream inFile;
    inFile.open("1.in");

    string line;
    size_t linecap = 0;
    ssize_t linelen;

    regex bagRegex = regex("(\\d+) (.*?) bags?,?");

    map<string, vector<string> > bagNeighbours1;
    map<string, vector<pair<string, int> > > bagNeighbours2;

    string d = "contain";
    while (getline(inFile, line))
    {
        // split by 'contain'
        size_t containIndex = line.find(d);
        string parentBag = line.substr(0, containIndex - 6);
        string childStr = line.substr(containIndex + 8, line.size() - containIndex - 9);

        vector<pair<string, int> > neighbours2;

        if (bagNeighbours2.find(parentBag) == bagNeighbours2.end()) {
            bagNeighbours2.insert(pair<string, vector<pair<string, int> > >(parentBag, neighbours2));
        }


        if (childStr != "no other bags")
        {
            smatch matches;
            string::const_iterator searchStart(childStr.cbegin());
            while (regex_search(searchStart, childStr.cend(), matches, bagRegex))
            {
                string bagNumberStr = matches[1].str();
                string bagColor = matches[2].str();
                int bagNumber = stoi(bagNumberStr);
                if (bagNeighbours1.count(bagColor) == 0)
                {
                    vector<string> neighbours;
                    bagNeighbours1.insert(pair<string, vector<string> >(bagColor, neighbours));
                }
                bagNeighbours1.find(bagColor)->second.push_back(parentBag);

                bagNeighbours2.find(parentBag)->second.push_back(pair<string, int>(bagColor, bagNumber));

                searchStart = matches.suffix().first;
            }
        }
    }

    map<string, int> visited;
    string start = "shiny gold";
    int sum = 0;
    dfs(bagNeighbours1, visited, start, &sum);

    cout << "prob 1: " << sum << "\n";

    cout << "prob 2: " << count(bagNeighbours2, start) << "\n";
}