# unfair

This tool generates a bunch of keystream data from stream ciphers and tracks the distribution of bytes at each position. The idea is to catch biases like those shown in RC4 here: http://www.isg.rhul.ac.uk/tls/biases.pdf

The interface and such is still in flux, but to test a new cipher you will basically add a *ciphername*\_streamer.go file which will have a type for CipherName and implement the interface in main.go.

## TODO

* [ ] Make the sample generation parallel, and have it take advantage of all cores on the machine.
* [ ] Polish up the information on the generated chart page.
* [ ] Add the option to import the generated data from one or more json files
* [X] Give the tool a nice little CLI to set options like random key length, number of samples, etc.
* [X] Chart the probability of a byte showing up at a position instead of the number of times that byte showed up.
