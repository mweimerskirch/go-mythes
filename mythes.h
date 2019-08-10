#ifndef __MYTHES_H__
#define __MYTHES_H__

#ifdef __cplusplus
extern "C"{
#endif

// a meaning with definition, count of synonyms and synonym list
typedef struct mentry mentry;
struct mentry {
  char*  defn;
  int  count;
  char** psyns;
};

// Wrapper struct to hold a pointer to MyThes object in C
struct   MyThes;
typedef  struct MyThes MyThes;

// Wrapper for the constructor
MyThes*  MyThes_create(const char* idx_path, const char* dat_path);

// Wrapper for the destructor
void     MyThes_destroy(MyThes* self);

// Wrapper for method Lookup
int      MyThes_Lookup(MyThes* self, const char * pText, mentry** pMeaning);

// Wrapper for method CleanUpAfterLookup
void     MyThes_CleanUpAfterLookup(MyThes* self, mentry** pMeaning, int len);

// Method that increments the pointer (easier to to in C than in Go)
mentry*  MyThes_Next(MyThes* self, mentry* pMeaning);

#ifdef __cplusplus
}
#endif

#endif
