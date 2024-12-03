#include <algorithm>
#include <boost/tokenizer.hpp>
#include <cstdlib>
#include <cstring>
#include <fstream>
#include <iostream>
#include <ostream>
#include <unordered_map>
#include <vector>

#include <chrono>

bool issafe(std::vector<int> levels) {
  bool isincr = levels[0] < levels[1];

  bool is_safe = true;
  for (int i = 0; i < levels.size() - 1; i++) {
    auto a = levels[i];
    auto b = levels[i + 1];
    auto diff = std::abs(a - b);
    if (1 > diff || diff > 3 || diff == 0 || (isincr ? a > b : a < b)) {

      is_safe = false;
      break;
    }
  }

  return is_safe;
}

bool issafe_dampen(std::vector<int> level) {
  bool is_safe = false;
  for (int i = 0; i < level.size(); i++) {
    std::vector<int> levels(level);
    levels.erase(levels.begin() + i);
    if (issafe(levels)) {
      return true;
    }
  }
  return false;
}

int main(int argc, char *argv[]) {

  auto start = std::chrono::high_resolution_clock::now();

  std::string s;
  std::unordered_map<long, long> count(64);

  std::vector<int> levels;
  levels.reserve(100);

  std::ifstream inp_file(argv[1]);
  std::string item;

  int sum = 0;
  long sum2 = 0;

  while (getline(inp_file, item)) {

    if (item == "") {
      continue;
    }
    boost::tokenizer<boost::char_separator<char>> tok(item);

    auto k = tok.begin();
    levels.clear();

    for (; k != tok.end(); k++) {
      levels.push_back(std::stoi(*k));
    }

    bool isincr = levels[0] < levels[1];
    int dampen = 1;

    bool is_safe = issafe(levels);

    if (is_safe) {
      sum2++;
      sum++;
    } else {
      if (issafe_dampen(levels)) {
        sum2++;
      }
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
