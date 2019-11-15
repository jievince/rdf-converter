#include <functional>
#include <string>
#include <iostream>
#include "hash.h"

std::string getHash(std::string key) {
    int64_t val = std::hash<std::string>{}(key);
    return std::to_string(val);
}
