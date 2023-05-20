
Installation
============

gameset is a proof of concept of modeling simple game elements in Golang
as packages. It includes several command line programs demonstrating
the use of the gameset packages.  It install these programs you need
to compile them from source code with the Golang compiler version
1.20 or better, pandoc 3 or better for compilation and documentation.

Quick install with curl
-----------------------

You can a release of gameset on macOS and Linux using with curl.

~~~
curl https://rsdoiel.github.io/gameset/installer.sh | sh
~~~


Getting the source code
-----------------------

The source code for gameset is available from GitHub in the
[github.com/rsdoiel/gameset](https://github.com/rsdoiel/gameset). You
can clone with with the `git` command or your favorite "git" GUI.

Compiling the source
--------------------

After you cloned your repository change into the repository directory
and run `make`. This will invoke the go compiler as well as Pandoc to
build the project. You can then use `make test` to test the project
and `make install` to install the project. By default this programs
and man pages are installed in your "home" directory. If you wish
to install them at another location set the PREFIX variable to
the desired location.

### Installing on home directory

```
    git clone https://github.com/rsdoiel/gameset
    cd gameset
    make
    make test
    make install
```

### Installing in `/usr/local`

```
    git clone https://github.com/rsdoiel/gameset
    cd gameset
    make
    make test
    make install PREFIX=/usr/local
```


