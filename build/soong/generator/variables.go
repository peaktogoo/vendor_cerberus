package generator

import (
	"fmt"

	"android/soong/android"
)

func cerberusExpandVariables(ctx android.ModuleContext, in string) string {
	cerberusVars := ctx.Config().VendorConfig("cerberusVarsPlugin")

	out, err := android.Expand(in, func(name string) (string, error) {
		if cerberusVars.IsSet(name) {
			return cerberusVars.String(name), nil
		}
		// This variable is not for us, restore what the original
		// variable string will have looked like for an Expand
		// that comes later.
		return fmt.Sprintf("$(%s)", name), nil
	})

	if err != nil {
		ctx.PropertyErrorf("%s: %s", in, err.Error())
		return ""
	}

	return out
}

