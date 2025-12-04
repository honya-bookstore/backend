package repositorypostgres

import (
	"math"
	"math/big"

	"github.com/jackc/pgx/v5/pgtype"
)

func numericToInt64(n pgtype.Numeric) int64 {
	if !n.Valid {
		return 0
	}
	var result big.Int
	n.Int.Mul(n.Int, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(n.Exp)), nil))
	result.Set(n.Int)
	return result.Int64()
}

func int64ToNumeric(value int64) pgtype.Numeric {
	return pgtype.Numeric{
		Int:   big.NewInt(value),
		Exp:   0,
		Valid: true,
	}
}

func numericToFloat64(n pgtype.Numeric) float64 {
	if !n.Valid {
		return 0
	}
	ratVal := new(big.Rat).SetInt(n.Int)
	ratVal = ratVal.Mul(ratVal, big.NewRat(1, int64(math.Pow10(int(n.Exp*-1)))))
	floatVal, ok := ratVal.Float64()
	if !ok {
		return 0
	}
	return floatVal
}

func float64ToNumeric(value float64) pgtype.Numeric {
	var n pgtype.Numeric
	err := n.Scan(value)
	if err != nil {
		return pgtype.Numeric{}
	}
	return n
}
