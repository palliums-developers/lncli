package lnd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lightningnetwork/lnd/channeldb"
)

var (
	channelDB *channeldb.DB
	fout      *os.File = nil
	ferr      *os.File = nil
)

type Shutdown struct{}

//InitLnd initializes lnd, lndHomeDir is coming from host app.
// lndHomeDir could be for example in android /data/user/0/com.rtxwallet/files.

func InitLnd(lndHomeDir string) error {
	setDefaultVars(lndHomeDir)

	lndCfg, err := loadConfig()
	if err != nil {
		fmt.Println(err)
		return err
	}
	cfg = lndCfg
	return nil
}

func SetStdout(lndHomeDir string) {
	fileout := filepath.Join(lndHomeDir, "stdout")
	fout, _ = os.Create(fileout)
	os.Stdout = fout

	fileerr := filepath.Join(lndHomeDir, "stdout")
	ferr, _ = os.Create(fileerr)
	os.Stderr = ferr
}

func StartLnd() error {
	defer func() {
		if x := recover(); x != nil {
			ltndLog.Errorf("run time panic: %v", x)
		}
	}()
	err := Main()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func ShutdownStdout() {
	if fout != nil {
		fout.Close()
	}
	if ferr != nil {
		ferr.Close()
	}
}

func setDefaultVars(lndHomeDir string) {
	if lndHomeDir == "" {
		// If lndHomeDir is null, just leave the defaults as is.
		return
	}
	defaultLndDir = lndHomeDir
	defaultConfigFile = filepath.Join(defaultLndDir, defaultConfigFilename)
	defaultDataDir = filepath.Join(defaultLndDir, defaultDataDirname)
	defaultLogDir = filepath.Join(defaultLndDir, defaultLogDirname)

	defaultTLSCertPath = filepath.Join(defaultLndDir, defaultTLSCertFilename)
	defaultTLSKeyPath = filepath.Join(defaultLndDir, defaultTLSKeyFilename)

	defaultBtcdDir = filepath.Join(lndHomeDir, "btcd", "default")
	defaultBtcdRPCCertFile = filepath.Join(defaultBtcdDir, "rpc.cert")

	defaultLtcdDir = filepath.Join(lndHomeDir, "ltcd", "default")
	defaultLtcdRPCCertFile = filepath.Join(defaultLtcdDir, "rpc.cert")

	defaultBitcoindDir = filepath.Join(lndHomeDir, "bitcoin", "default")
	defaultLitecoindDir = filepath.Join(lndHomeDir, "litecoin", "default")
}
