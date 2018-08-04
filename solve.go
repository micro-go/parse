package parse

import (
	"errors"
	"go/constant"
	"go/token"
	"go/types"
)

func SolveInt(s string) (int, error) {
	// Evaluate
	fs := token.NewFileSet()
	tv, err := types.Eval(fs, nil, token.NoPos, s)
	if err != nil {
		return 0, err
	}

	// Extract. Try both float and int
	vi := constant.ToInt(tv.Value)
	if vi.Kind() == constant.Int {
		i, ok := constant.Int64Val(vi)
		if !ok {
			return 0, errors.New("Unsolvable: " + s)
		}
		return int(i), nil
	}

	vf := constant.ToFloat(tv.Value)
	if vf.Kind() == constant.Float {
		f, ok := constant.Float64Val(vf)
		if !ok {
			return 0, errors.New("Unsolvable: " + s)
		}
		return int(f), nil
	}

	return 0, errors.New("Unsolvable: " + s)
}
