#ifndef _MYTHES_HXX_
#define _MYTHES_HXX_

#include "mythes.h"

class MyThes
{
    int nw;
    char** list;
    unsigned int* offst;
    char* encoding;

    FILE  *pdfile;

	MyThes();
	MyThes(const MyThes &);
	MyThes & operator = (const MyThes &);

    public:
        MyThes(const char* idx_path, const char* dat_path);
        ~MyThes();

        int Lookup(const char * pText, int len, mentry** p_meaning);

        void CleanUpAfterLookup(mentry** p_meaning, int len);

        char* get_th_encoding();
};

#endif
