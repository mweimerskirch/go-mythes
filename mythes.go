package go_mythes

// #cgo linux LDFLAGS: -lmythes-1.2
//
// #include <stdlib.h>
// #include <stdio.h>
// #include "mythes.h"
import "C"
import (
	"reflect"
	"runtime"
	"strings"
	"sync"
	"unsafe"
)

// Struct that holds a single synonym set (made up of a definition (or label) and one or more synonyms.
type Synset struct {
	Label    string   `json:"label"`
	Synonyms []string `json:"synonyms"`
}

// Struct that holds a reference to the MyThes library
type MyThesHandle struct {
	handle *C.MyThes
	lock   *sync.Mutex
}

// Initialize MyThes
func MyThes(idxpath string, datpath string) *MyThesHandle {
	idxpathcs := C.CString(idxpath)
	defer C.free(unsafe.Pointer(idxpathcs))

	datpathcs := C.CString(datpath)
	defer C.free(unsafe.Pointer(datpathcs))

	h := &MyThesHandle{lock: new(sync.Mutex)}
	h.handle = C.MyThes_create(idxpathcs, datpathcs)

	runtime.SetFinalizer(h, func(handle *MyThesHandle) {
		C.MyThes_destroy(handle.handle)
		h.handle = nil
	})

	return h
}

// Lookup synonyms for a word in a MyThes file
func (handle *MyThesHandle) Lookup(word string) []Synset {
	word_c := C.CString(word)
	defer C.free(unsafe.Pointer(word_c))

	var pMeaning *C.struct_mentry

	pOriginal := pMeaning

	handle.lock.Lock()

	// Lookup a single word in the synonym list
	count_c := C.MyThes_Lookup(handle.handle, word_c, &pMeaning)
	count := int(count_c)

	var synsets []Synset

	if count > 0 {
		for i := 0; i < count; i++ {
			// Get the "meaning definition" and the synonym list for this item
			defn := C.GoString(pMeaning.defn)
			synonyms := convertArrayFromC(pMeaning.psyns, int(pMeaning.count))

			// Remove leading dash from the definition
			defn = removeLeadingDash(defn)

			// Create a "Synset" item ...
			synset := Synset{
				Label:    defn,
				Synonyms: synonyms,
			}

			// ... and append it to the list of synonyms
			synsets = append(synsets, synset)

			// If there is another result, increment the pointer (in C) to get the next item
			if i < count-1 {
				pMeaning = C.MyThes_Next(handle.handle, pMeaning)
			}
		}
	}

	// Once done, free all the memory in C
	C.MyThes_CleanUpAfterLookup(handle.handle, &pOriginal, C.int(count))
	handle.lock.Unlock()

	return synsets
}

// Convert an "array" from C (actually a list of pointers to chars) to a string array in Go
func convertArrayFromC(array **C.char, length int) []string {
	var list []string

	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(array)),
		Len:  length,
		Cap:  length,
	}

	for _, item := range *(*[]*C.char)(unsafe.Pointer(&header)) {
		list = append(list, C.GoString(item))
	}

	return list
}

// Remove the leading "- " from a string
func removeLeadingDash(string string) string {
	if strings.Index(string, "- ") == 0 {
		return string[2:]
	}

	return string
}
