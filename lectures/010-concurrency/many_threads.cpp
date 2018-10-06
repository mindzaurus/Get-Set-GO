#include <iostream>
#include <utility>
#include <thread>
#include <chrono>
#include <climits>
#include <vector>

using namespace std;

unsigned long long int x = 50000, y = 10000, NT = 5000;

void cu() {
  volatile unsigned long long int v = 0;

  for(unsigned long long int i=0; i<x; i++) {
    for(unsigned long long int j=0; j<y; j++) {
      v = j * j * j;
    }
  }

  std::cout << v << std::endl;
}

int main()
{
  vector <std::thread> t;

  for(int i=0; i<NT; i++) {
    t.push_back(std::thread(&cu));
  }

  for(int i=0; i<t.size(); i++) {
    t[i].join();
  }

}
