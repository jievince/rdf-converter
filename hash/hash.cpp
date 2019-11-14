#include <functional>
#include <string>
#include <iostream>
#include "hash.h"

std::string getHash(std::string key) {
    return std::to_string(std::hash<std::string>{}(key));
}
