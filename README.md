rusticsearch
============

**rusticsearch** is a lightweight search server that includes all that is needed in a typical small to medium business website but is free of redundant features and all kinds of cruft. 

It supports fuzzy search and autocomplete.

###Get started
Download and install the source code with [go get](http://golang.org/cmd/go/) command:

    go get github.com/dborzov/rusticsearch
    go install rusticsearch
   
Now we can check if the binary is indeed available:
     
     ./rusticsearch -h
We need to add the config file: