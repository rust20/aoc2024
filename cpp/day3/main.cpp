#include <algorithm>
#include <boost/tokenizer.hpp>
#include <cstddef>
#include <cstdlib>
#include <cstring>
#include <execution>
#include <fstream>
#include <iostream>
#include <ostream>
#include <sstream>
#include <string>
#include <unordered_map>
#include <utility>
#include <vector>

#include <chrono>

std::pair<int, bool> getnum(std::string str, size_t &i) {
  int total = 0;
  bool err = 1;
  while ('0' <= str.at(i) && str.at(i) <= '9') {
    err = 0;

    total *= 10;
    total += int(str.at(i) - '0');
    i++;
  }

  return std::make_pair(total, err);
}

int part1(std::string item) {
  auto start = std::chrono::high_resolution_clock::now();

  int sum = 0;

  size_t cursor = 0;

  while (cursor < item.length()) {
    switch (item.at(cursor)) {
    case 'm':
      if (item.substr(cursor, 4) == "mul(") {
        cursor += 4;
        auto [val1, err1] = getnum(item, cursor);
        if (item.at(cursor) != ',' || err1) {
          continue;
        }
        cursor++;
        auto [val2, err2] = getnum(item, cursor);
        if (item.at(cursor) != ')' || err2) {
          continue;
        }
        sum += val1 * val2;
      }
    }
    cursor++;
  }

  std::cout << "star1: " << sum << std::endl;

  auto stop = std::chrono::high_resolution_clock::now();
  auto duration =
      std::chrono::duration_cast<std::chrono::microseconds>(stop - start);
  std::cout << "Time taken by function: " << duration.count() << " microseconds"
            << std::endl;

  return 0;
}

int part2(std::string item) {
  auto start = std::chrono::high_resolution_clock::now();

  long sum2 = 0;

  size_t cursor = 0;

  bool is_add = true;

  while (cursor < item.length()) {
    switch (item.at(cursor)) {
    case 'm':
      if (item.substr(cursor, 4) == "mul(") {
        cursor += 4;
        auto [val1, err1] = getnum(item, cursor);
        if (item.at(cursor) != ',' || err1) {
          continue;
        }
        cursor++;
        auto [val2, err2] = getnum(item, cursor);
        if (item.at(cursor) != ')' || err2) {
          continue;
        }
        if (is_add)
          sum2 += val1 * val2;
      }
    case 'd':
      if (item.substr(cursor, 7) == "don't()") {
        cursor += 6;
        is_add = false;
      } else if (item.substr(cursor, 4) == "do()") {
        cursor += 3;
        is_add = true;
      }
    }

    cursor++;
  }

  std::cout << "star2: " << sum2 << std::endl;

  auto stop = std::chrono::high_resolution_clock::now();
  auto duration =
      std::chrono::duration_cast<std::chrono::microseconds>(stop - start);
  std::cout << "Time taken by function: " << duration.count() << " microseconds"
            << std::endl;

  return 0;
}

int main(int argc, char *argv[]) {

  std::string s;
  std::unordered_map<long, long> count(64);

  std::vector<int> levels;
  levels.reserve(100);

  std::ifstream inp_file(argv[1]);
  std::stringstream ss;
  ss << inp_file.rdbuf();
  std::string item = ss.str();

  part1(item);
  part2(item);
  return 0;
}

// vim: set ts=2 sts=2 sw=2 et:

// vim: set ts=2 sts=2 sw=2 et:
