#include "hash.h"
#include "bridge.h"

const char* getHash(char* key) {
    return getHash(std::string(key)).c_str();
}
    
