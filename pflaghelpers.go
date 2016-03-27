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

func EnsureRequired(rootCmd *cobra.Command) {
	cmd, _, err := rootCmd.Find(os.Args[1:])
	if err != nil {
		panic(err)
	}

	problematicFlags := []string{}

	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		if IsFlagRequired(flag) && !flag.Changed {
			problematicFlags = append(problematicFlags, flag.Name)
		}
	})

	if len(problematicFlags) > 0 {
		cmd.Usage()

		for i, flagName := range problematicFlags {
			problematicFlags[i] = fmt.Sprintf("`%s`", flagName)
		}

		var err error
		if len(problematicFlags) == 1 {
			_, err = fmt.Fprintf(cmd.Out(),
				"\nUsage error: flag %s is required\n", problematicFlags[0])
		} else {
			_, err = fmt.Fprintf(cmd.Out(),
				"\nUsage error: flags %s are required\n", strings.Join(problematicFlags, ", "))
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
}

func IsFlagRequired(flag *pflag.Flag) bool {
	return strings.HasSuffix(flag.Usage, "(required)")
}

// A series of MustGet* functions, which now make some sense.

func MustGetString(fs *pflag.FlagSet, name string, allowEmpty bool) string {
	if rv, err := fs.GetString(name); err != nil {
		panic(err)
	} else {
		if !allowEmpty && rv == "" {
			panic(fmt.Errorf("Flag '%s' not allowed to be empty", name))
		}
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

func MustGetFloat32(fs *pflag.FlagSet, name string, allowEmpty bool) float32 {
	if rv, err := fs.GetFloat32(name); err != nil {
		panic(err)
	} else {
		if !allowEmpty && rv == 0 {
			panic(fmt.Errorf("Flag '%s' not allowed to be empty", name))
		}
		return rv
	}
}

func MustGetFloat64(fs *pflag.FlagSet, name string, allowEmpty bool) float64 {
	if rv, err := fs.GetFloat64(name); err != nil {
		panic(err)
	} else {
		if !allowEmpty && rv == 0 {
			panic(fmt.Errorf("Flag '%s' not allowed to be empty", name))
		}
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

func MustGetInt(fs *pflag.FlagSet, name string, allowEmpty bool) int {
	if rv, err := fs.GetInt(name); err != nil {
		panic(err)
	} else {
		if !allowEmpty && rv == 0 {
			panic(fmt.Errorf("Flag '%s' not allowed to be empty", name))
		}
		return rv
	}
}

func MustGetInt32(fs *pflag.FlagSet, name string, allowEmpty bool) int32 {
	if rv, err := fs.GetInt32(name); err != nil {
		panic(err)
	} else {
		if !allowEmpty && rv == 0 {
			panic(fmt.Errorf("Flag '%s' not allowed to be empty", name))
		}
		return rv
	}
}

func MustGetInt64(fs *pflag.FlagSet, name string, allowEmpty bool) int64 {
	if rv, err := fs.GetInt64(name); err != nil {
		panic(err)
	} else {
		if !allowEmpty && rv == 0 {
			panic(fmt.Errorf("Flag '%s' not allowed to be empty", name))
		}
		return rv
	}
}

func MustGetInt8(fs *pflag.FlagSet, name string, allowEmpty bool) int8 {
	if rv, err := fs.GetInt8(name); err != nil {
		panic(err)
	} else {
		if !allowEmpty && rv == 0 {
			panic(fmt.Errorf("Flag '%s' not allowed to be empty", name))
		}
		return rv
	}
}
