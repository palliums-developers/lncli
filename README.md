## lncli is library for lnd to build liblnd

### Preliminaries
  In order to build liblnd, the
  following build dependencies are required:

  * **Go:** To compile, run one of the following commands:


    **Note**: The minimum version of Go supported is Go 1.12. We recommend that
    users use the latest version of Go, which at the time of writing is
    [`1.12`](https://blog.golang.org/go1.12).


    On Linux:

    (x86-64)
    ```
    wget https://dl.google.com/go/go1.12.3.linux-amd64.tar.gz
    sha256sum go1.12.3.linux-amd64.tar.gz | awk -F " " '{ print $1 }'
    ```

    The final output of the command above should be
    `3924819eed16e55114f02d25d03e77c916ec40b7fd15c8acb5838b63135b03df`. If it
    isn't, then the target REPO HAS BEEN MODIFIED, and you shouldn't install
    this version of Go. If it matches, then proceed to install Go:
    ```
    tar -C /usr/local -xzf go1.12.3.linux-amd64.tar.gz
    export PATH=$PATH:/usr/local/go/bin
    ```

    ```

    On Mac OS X:
    ```
    brew install go@1.12
    ```


    Alternatively, one can download the pre-compiled binaries hosted on the
    [Golang download page](https://golang.org/dl/). If one seeks to install
    from source, then more detailed installation instructions can be found
    [here](https://golang.org/doc/install).

    At this point, you should set your `$GOPATH` environment variable, which
    represents the path to your workspace. By default, `$GOPATH` is set to
    `~/go`. You will also need to add `$GOPATH/bin` to your `PATH`. This ensures
    that your shell will be able to detect the binaries you install.

    ```bash
    export GOPATH=~/gocode
    export PATH=$PATH:$GOPATH/bin
    ```

    We recommend placing the above in your .bashrc or in a setup script so that
    you can avoid typing this every time you open a new terminal window.

  * **Go modules:** This project uses [Go modules](https://github.com/golang/go/wiki/Modules) 
    to manage dependencies as well as to provide *reproducible builds*.
	
	Additionaly, in order to compile the mobile library, you need to install [xgo](https://github.com/karalabe/xgo).

### Compile liblnd

    With the preliminary steps completed, to compile `liblnd`,all
    related dependencies run the following commands:
    ```
    go get -d github.com/lightningnetwork/lnd
    cd $GOPATH/src/github.com/lightningnetwork/lnd
    make
	
	git clone https://github.com/palliums-developers/lncli.git
	mv lncli/api.go cmd/lnd/
	mv lncli/api_config.go ./
	sed -i '/host, err := os.Hostname/{N;N;N;N;N;s/.*/\thost := "localhost"/}' lnd.go
	

### Cross compiling liblnd
  
    Depending on which platform you're targeting, compile using xgo:
    * **Mac**, run:
    cd $GOPATH/src/github.com/lightningnetwork/liblnd/cmd/lnd
    xgo -out liblnd -buildmode=c-archive -tags="experimental autopilotrpc" --targets=darwin-10.11/*  ./
    
    * **Windows**, run:
    cd $GOPATH/src/github.com/lightningnetwork/liblnd/cmd/lnd
    xgo -out liblnd -buildmode=c-archive -tags="experimental autopilotrpc" --targets=windows/*  ./
    
    * **Linux**, run:
    cd $GOPATH/src/github.com/lightningnetwork/liblnd/cmd/lnd
    xgo -out liblnd -buildmode=c-archive -tags="experimental autopilotrpc" --targets=linux/*  ./
    
    
    * **Android**, run:
    ```
	cd $GOPATH/src/github.com/lightningnetwork/liblnd/cmd/lnd
    xgo -tags="experimental autopilotrpc" -out liblnd_export --targets=android/aar  ./
    ```
    This will produce an **liblnd_export-android-16.aar** file and you need to copy it
    to the app's directory.

    * **iOS**, run:
    ```
	cd $GOPATH/src/github.com/lightningnetwork/liblnd/cmd/lnd
    xgo -tags="experimental autopilotrpc" -out liblnd_export --targets=ios/framework ./
    ```
    This will produce a **liblnd_export-ios-5.0-framework/Liblnd_export.framework/** folder
    and you need to copy it to the app's directory.

```
