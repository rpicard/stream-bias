# unfair

This tool generates a bunch of keystream data from stream ciphers and tracks the distribution of bytes at each position. The idea is to catch biases like those shown in RC4 here: http://www.isg.rhul.ac.uk/tls/biases.pdf

The interface and such is still in flux, but to test a new cipher you will basically add a *ciphername*\_streamer.go file which will have a type for CipherName and implement the interface in main.go.

```
Usage: unfair -c [-f] [-s] [-l]

chart potential biases in stream cipher keystreams

Options:
  -f, --format="html"    output format (html or json)
  -c, --cipher=""        which cipher are we testing?
  -s, --samples="1000"   how many samples should we take?
  -l, --length="256"     how many positions in the keystream do we care about?
```

## TODO

* [ ] Write script to spin up lots of ec2 spot instances, run this tool and post the resulting .json file to s3
* [ ] Make the sample generation parallel, and have it take advantage of all cores on the machine.

    I can't even tell if it's doing this at this point. It runs at the same speed on my 4 core machine and an ec2 machine with 36 cores. Slower on a micro instance though.

* [ ] Polish up the information on the generated chart page.
* [ ] Add the option to import the generated data from one or more json files
* [X] Give the tool a nice little CLI to set options like random key length, number of samples, etc.
* [X] Chart the probability of a byte showing up at a position instead of the number of times that byte showed up.
