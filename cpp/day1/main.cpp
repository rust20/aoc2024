#include <algorithm>
#include <boost/tokenizer.hpp>
#include <cstdlib>
#include <cstring>
#include <fstream>
#include <iostream>
#include <ostream>
#include <sstream>
#include <unordered_map>
#include <vector>

#include <chrono>

int main(int argc, char *argv[]) {

  auto start = std::chrono::high_resolution_clock::now();

  std::string s;
  std::vector<int> l;
  std::vector<int> r;
  std::unordered_map<long, long> count(1000);

  l.reserve(1000);
  r.reserve(1000);

  std::ifstream inp_file(argv[1]);
  // std::stringstream sss;
  // sss << inp_file.rdbuf();
  std::string item;

  while (getline(inp_file, item)) {
    boost::tokenizer<boost::char_separator<char>> tok(item);

    auto k = tok.begin();

    l.push_back(std::stoi(*k));
    r.push_back(std::stoi(*++k));

  }

  std::sort(l.begin(), l.end());
  std::sort(r.begin(), r.end());

  int sum = 0;
  for (int i = 0; i < l.size(); i++) {
    sum += std::abs(l[i] - r[i]);
  }



  for (auto val : r) {
    if (count.find(val) != count.end()) {
      count[val] += 1;
    } else {
      count[val] = 1;
    }
  }

  long sum2 = 0;
  for (auto val : l) {
    if (count.find(val) != count.end()) {
      sum2 += val * count[val];
    }
  }

  std::cout << "star1: " << sum << std::endl;
  std::cout << "star2: " << sum2 << std::endl;

  auto stop = std::chrono::high_resolution_clock::now();
  auto duration =
      std::chrono::duration_cast<std::chrono::microseconds>(stop - start);
  std::cout << "Time taken by function: " << duration.count() << " microseconds"
            << std::endl;

  return 0;
}

// vim: set ts=2 sts=2 sw=2 et:
