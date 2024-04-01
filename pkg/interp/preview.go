package interp

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/wader/fq/internal/mathx"
	"github.com/wader/fq/internal/stringsx"
	"github.com/wader/fq/pkg/bitio"
	"github.com/wader/fq/pkg/scalar"
)

func previewValue(v any, df scalar.DisplayFormat) string {
	switch vv := v.(type) {
	case bool:
		if vv {
			return "true"
		}
		return "false"
	case int:
		// TODO: DisplayFormat is weird
		return mathx.PadFormatInt(int64(vv), df.FormatBase(), true, 0)
	case int64:
		// TODO: DisplayFormat is weird
		return mathx.PadFormatInt(vv, df.FormatBase(), true, 0)
	case uint64:
		return mathx.PadFormatUint(vv, df.FormatBase(), true, 0)
	case float64:
		// TODO: float32? better truncated to significant digits?
		return strconv.FormatFloat(vv, 'g', -1, 64)
	case string:
		s := strconv.Quote(stringsx.TrimN(vv, 50, "..."))
		return strings.ReplaceAll(s, `\u007f`, `\x7f`)
	case nil:
		return "null"
	case bitio.Reader,
		Binary:
		return "raw bits"
	case *big.Int:
		return mathx.PadFormatBigInt(vv, df.FormatBase(), true, 0)
	case map[string]any:
		return "{}"
	case []any:
		return "[]"

	default:
		panic(fmt.Sprintf("unreachable %v (%T)", v, v))
	}
}
