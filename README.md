# go-mobile-collection-wrapper
# Based on [go-mobile-collection](https://github.com/scisci/go-mobile-collection)

Note: As a go-mobile-collection, this utility has been created to solve specific tasks and does not pretend to be the best or the most elegant solution. 

## Usage

###Preparation

What this repo does is build a command line utility that can then be automatically called using go generate semantics.

1. Find the file you that contains the struct definition you want to have a collection wrapper for.
2. At the top of the file add the following: `//go:generate go-mobile-collection-wrapper $GOFILE`

### Embedding inside slice

3. Before the struct, add a comment to flag it:
```
	// @slice-wrapper
	type Example struct {
	    ExampleField string
	}
```
4. When you build your project a new file should now be generated called ($GOFILE)_slice.go that contains the automatically generated definitions.

### Embedding inside map

1. Before the struct, add a comment to flag it. Specify a key types for map after @map-wrapper: separated by coma. Currently only primitives (string, int etc.) are supported as key types.
```
    // @map-wrapper:string,int
    type Example struct {
        ExampleField string
    }
```
2. When you build your project a new file should now be generated called ($GOFILE)_map.go that contains the automatically generated definitions.
