package pflaghelpers

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// This should be called by the root command PersistentPreRunE and also
// by every command that overrides PreRun.
func EnsureRequired(cmd *cobra.Command, args []string) error {
	problematicFlags := []string{}

	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		if strings.HasSuffix(flag.Usage, "(required)") && !flag.Changed {
			problematicFlags = append(problematicFlags, flag.Name)
		}
	})

	if len(problematicFlags) > 0 {
		for i, flagName := range problematicFlags {
			problematicFlags[i] = fmt.Sprintf("`%s`", flagName)
		}

		if len(problematicFlags) == 1 {
			return fmt.Errorf("Usage error: flag %s is required", problematicFlags[0])
		} else {
			return fmt.Errorf(
				"Usage error: flags %s are required", strings.Join(problematicFlags, ", "))
		}
	}

	return nil
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
