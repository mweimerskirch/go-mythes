# go-mythes
Go wrapper for the MyThes library

# Usage (on Ubuntu/Debian)

Install the library:

<pre>sudo apt install libmythes-1.2-0</pre>

Import the wrapper in your go application:

<pre>
import (
    "github.com/spellchecker-lu/go-mythes"
)
</pre>

Example usage:

<pre>
var idxFile = "dic/th_xyz.idx"
var datFile = "dic/th_xyz.dat"

var mythesHandle = mythes.MyThes(idxFile, datFile)

synsets := mythesHandle.Lookup(strings.ToLower(word))
for _, synset := range synsets {
    log.Println(synset)
}
</pre>