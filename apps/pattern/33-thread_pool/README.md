
You'll need the Go command-line tools. Ginkgo is tested with Go 1.6+, but preferably you should get the latest. Follow the installation instructions if you don't have it installed.
install the tool

```
# installs the ginkgo CLI
go get -u github.com/onsi/ginkgo/ginkgo  
# fetches the matcher library
go get -u github.com/onsi/gomega/...     

```

Enter into the test directory:
```

cd path/to/package/you/want/to/test
```

Basic *ginkgo* commands:

```
ginkgo bootstrap # set up a new ginkgo suite
ginkgo generate  # will create a sample test file.  edit this file and add your tests then...

go test # to run your tests

ginkgo  # also runs your tests

```
