package main

/*
#ifdef SWIG
%newobject InitLnd;
#endif
*/

import (
	"C"
	"fmt"

	"github.com/lightningnetwork/lnd"
	"github.com/lightningnetwork/lnd/lncli"
)

//InitLnd initializes lnd, lndHomeDir is coming from host app.
// lndHomeDir could be for example in android /data/user/0/com.rtxwallet/files.
//export InitLnd
func InitLnd(lndHomeDir *C.char) *C.char {
	lndHomeDirString := C.GoString(lndHomeDir)
	err := lnd.InitLnd(lndHomeDirString)
	if err != nil {
		lnd.ShutdownStdout()
		return C.CString(err.Error())
	}
	return C.CString("")
}

//export StartLnd
func StartLnd() *C.char {
	err := lnd.StartLnd()
	if err != nil {
		fmt.Println(err)
		return C.CString(err.Error())
	}
	fmt.Println("lnd start succeed")
	return C.CString("")
}

//export SetStdout
func SetStdout(lndHomeDir *C.char) {
	lnd.SetStdout(C.GoString(lndHomeDir))
}

//export CommandExecute
func CommandExecute(args *C.char) *C.char {
	ret := lncli.CommandExecute(C.GoString(args))
	return C.CString(ret)
}
