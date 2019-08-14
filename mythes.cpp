#include <string.h>
#include <stdlib.h>
#include <stdio.h>
#include "mythes.hxx"
#include "mythes.h"

#ifdef __cplusplus
extern "C"{
#endif

// Wrapper for the constructor
MyThes* MyThes_create(const char* idx_path, const char* dat_path){
    return reinterpret_cast<MyThes*>( new MyThes(idx_path, dat_path) );
}

// Wrapper for the destructor
void MyThes_destroy(MyThes* self){
    delete reinterpret_cast<MyThes*>(self);
}

// Wrapper for method Lookup
int MyThes_Lookup(MyThes* self, const char * pText, mentry** pMeaning){
    int len = strlen(pText);
    int count = reinterpret_cast<MyThes*>(self)->Lookup(pText, len, pMeaning);

    mentry* pm = *pMeaning;

    return count;
}

// Wrapper for method CleanUpAfterLookup
void MyThes_CleanUpAfterLookup(MyThes* self, mentry** pme, int len){
    reinterpret_cast<MyThes*>(self)->CleanUpAfterLookup(pme, len);
}

// Method that increments the pointer (easier to to in C than in Go)
mentry * MyThes_Next(MyThes* self, mentry* pMeaning){
    pMeaning++;
    return pMeaning;
}

#ifdef __cplusplus
}
#endif