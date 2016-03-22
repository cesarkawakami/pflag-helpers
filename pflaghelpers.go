package pflaghelpers

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func EnsureRequired(cmd *cobra.Command) {
	problematicFlags := []string{}

	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		if IsFlagRequired(flag) && !flag.Changed {
			problematicFlags = append(problematicFlags, flag.Name)
		}
	})

	if len(problematicFlags) > 0 {
		for i, flagName := range problematicFlags {
			problematicFlags[i] = fmt.Sprintf("`%s`", flagName)
		}

		var err error
		if len(problematicFlags) == 1 {
			_, err = fmt.Fprintf(cmd.Out(), "Usage error: flag %s is required", problematicFlags[0])
		} else {
			_, err = fmt.Fprintf(cmd.Out(),
				"Usage error: flags %s are required", strings.Join(problematicFlags, ", "))
		}
		if err != nil {
			panic(err)
		}

		os.Exit(1)
	}
}

func Bind(cmd *cobra.Command) {
	cobra.OnInitialize(func() {
		EnsureRequired(cmd)
	})

	cmd.SetGlobalNormalizationFunc(func(fs *pflag.FlagSet, name string) pflag.NormalizedName {
		flag := fs.Lookup(name)
		if flag == nil {
			panic("`flag` should never be nil here")
		}
		prefix := "1"
		if IsFlagRequired(flag) {
			prefix = "0"
		}
		return pflag.NormalizedName(prefix + name)
	})
}

func IsFlagRequired(flag *pflag.Flag) bool {
	return strings.HasSuffix(flag.Usage, "(required)")
}

// A series of MustGet* functions, which now make some sense.

func MustGetString(fs *pflag.FlagSet, name string) string {
	if rv, err := fs.GetString(name); err != nil {
		panic(err)
	} else {
		return rv
	}
}

func MustGetBool(fs *pflag.FlagSet, name string) bool {
	if rv, err := fs.GetBool(name); err != nil {
		panic(err)
	} else {
		return rv
	}
}

func MustGetDuration(fs *pflag.FlagSet, name string) time.Duration {
	if rv, err := fs.GetDuration(name); err != nil {
		panic(err)
	} else {
		return rv
	}
}

func MustGetFloat32(fs *pflag.FlagSet, name string) float32 {
	if rv, err := fs.GetFloat32(name); err != nil {
		panic(err)
	} else {
		return rv
	}
}

func MustGetFloat64(fs *pflag.FlagSet, name string) float64 {
	if rv, err := fs.GetFloat64(name); err != nil {
		panic(err)
	} else {
		return rv
	}
}

func MustGetIP(fs *pflag.FlagSet, name string) net.IP {
	if rv, err := fs.GetIP(name); err != nil {
		panic(err)
	} else {
		return rv
	}
}

func MustGetIPv4Mask(fs *pflag.FlagSet, name string) net.IPMask {
	if rv, err := fs.GetIPv4Mask(name); err != nil {
		panic(err)
	} else {
		return rv
	}
}

func MustGetInt(fs *pflag.FlagSet, name string) int {
	if rv, err := fs.GetInt(name); err != nil {
		panic(err)
	} else {
		return rv
	}
}

func MustGetInt32(fs *pflag.FlagSet, name string) int32 {
	if rv, err := fs.GetInt32(name); err != nil {
		panic(err)
	} else {
		return rv
	}
}

func MustGetInt64(fs *pflag.FlagSet, name string) int64 {
	if rv, err := fs.GetInt64(name); err != nil {
		panic(err)
	} else {
		return rv
	}
}

func MustGetInt8(fs *pflag.FlagSet, name string) int8 {
	if rv, err := fs.GetInt8(name); err != nil {
		panic(err)
	} else {
		return rv
	}
}
