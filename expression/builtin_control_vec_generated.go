// Copyright 2019 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go generate in expression/generator; DO NOT EDIT.

package expression

import (
	"time"

	"github.com/pingcap/tidb/types"
	"github.com/pingcap/tidb/util/chunk"
)

func (b *builtinCaseWhenIntSig) vecEvalInt(input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()
	args, l := b.getArgs(), len(b.getArgs())
	whens := make([]*chunk.Column, l/2)
	whensSlice := make([][]int64, l/2)
	thens := make([]*chunk.Column, l/2)
	var eLse *chunk.Column

	thensSlice := make([][]int64, l/2)
	var eLseSlice []int64

	for j := 0; j < l-1; j += 2 {
		bufWhen, err := b.bufAllocator.get(types.ETInt, n)
		if err != nil {
			return err
		}
		defer b.bufAllocator.put(bufWhen)
		if err := args[j].VecEvalInt(b.ctx, input, bufWhen); err != nil {
			return err
		}
		whens[j/2] = bufWhen
		whensSlice[j/2] = bufWhen.Int64s()

		bufThen, err := b.bufAllocator.get(types.ETInt, n)
		if err != nil {
			return err
		}
		defer b.bufAllocator.put(bufThen)
		if err := args[j+1].VecEvalInt(b.ctx, input, bufThen); err != nil {
			return err
		}
		thens[j/2] = bufThen

		thensSlice[j/2] = bufThen.Int64s()

	}
	// when clause(condition, result) -> args[i], args[i+1]; (i >= 0 && i+1 < l-1)
	// else clause -> args[l-1]
	// If case clause has else clause, l%2 == 1.
	if l%2 == 1 {
		bufElse, err := b.bufAllocator.get(types.ETInt, n)
		if err != nil {
			return err
		}
		defer b.bufAllocator.put(bufElse)
		if err := args[l-1].VecEvalInt(b.ctx, input, bufElse); err != nil {
			return err
		}
		eLse = bufElse

		eLseSlice = bufElse.Int64s()

	}

	result.ResizeInt64(n, false)
	resultSlice := result.Int64s()

ROW:
	for i := 0; i < n; i++ {
		for j := 0; j < l/2; j++ {
			if whens[j].IsNull(i) || whensSlice[j][i] == 0 {
				continue
			}

			resultSlice[i] = thensSlice[j][i]
			result.SetNull(i, thens[j].IsNull(i))

			continue ROW
		}
		if eLse != nil {

			resultSlice[i] = eLseSlice[i]
			result.SetNull(i, eLse.IsNull(i))

		} else {

			result.SetNull(i, true)

		}
	}
	return nil
}

func (b *builtinCaseWhenIntSig) vectorized() bool {
	return true
}

func (b *builtinCaseWhenRealSig) vecEvalReal(input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()
	args, l := b.getArgs(), len(b.getArgs())
	whens := make([]*chunk.Column, l/2)
	whensSlice := make([][]int64, l/2)
	thens := make([]*chunk.Column, l/2)
	var eLse *chunk.Column

	thensSlice := make([][]float64, l/2)
	var eLseSlice []float64

	for j := 0; j < l-1; j += 2 {
		bufWhen, err := b.bufAllocator.get(types.ETInt, n)
		if err != nil {
			return err
		}
		defer b.bufAllocator.put(bufWhen)
		if err := args[j].VecEvalInt(b.ctx, input, bufWhen); err != nil {
			return err
		}
		whens[j/2] = bufWhen
		whensSlice[j/2] = bufWhen.Int64s()

		bufThen, err := b.bufAllocator.get(types.ETReal, n)
		if err != nil {
			return err
		}
		defer b.bufAllocator.put(bufThen)
		if err := args[j+1].VecEvalReal(b.ctx, input, bufThen); err != nil {
			return err
		}
		thens[j/2] = bufThen

		thensSlice[j/2] = bufThen.Float64s()

	}
	// when clause(condition, result) -> args[i], args[i+1]; (i >= 0 && i+1 < l-1)
	// else clause -> args[l-1]
	// If case clause has else clause, l%2 == 1.
	if l%2 == 1 {
		bufElse, err := b.bufAllocator.get(types.ETReal, n)
		if err != nil {
			return err
		}
		defer b.bufAllocator.put(bufElse)
		if err := args[l-1].VecEvalReal(b.ctx, input, bufElse); err != nil {
			return err
		}
		eLse = bufElse

		eLseSlice = bufElse.Float64s()

	}

	result.ResizeFloat64(n, false)
	resultSlice := result.Float64s()

ROW:
	for i := 0; i < n; i++ {
		for j := 0; j < l/2; j++ {
			if whens[j].IsNull(i) || whensSlice[j][i] == 0 {
				continue
			}

			resultSlice[i] = thensSlice[j][i]
			result.SetNull(i, thens[j].IsNull(i))

			continue ROW
		}
		if eLse != nil {

			resultSlice[i] = eLseSlice[i]
			result.SetNull(i, eLse.IsNull(i))

		} else {

			result.SetNull(i, true)

		}
	}
	return nil
}

func (b *builtinCaseWhenRealSig) vectorized() bool {
	return true
}

func (b *builtinCaseWhenDecimalSig) vecEvalDecimal(input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()
	args, l := b.getArgs(), len(b.getArgs())
	whens := make([]*chunk.Column, l/2)
	whensSlice := make([][]int64, l/2)
	thens := make([]*chunk.Column, l/2)
	var eLse *chunk.Column

	thensSlice := make([][]types.MyDecimal, l/2)
	var eLseSlice []types.MyDecimal

	for j := 0; j < l-1; j += 2 {
		bufWhen, err := b.bufAllocator.get(types.ETInt, n)
		if err != nil {
			return err
		}
		defer b.bufAllocator.put(bufWhen)
		if err := args[j].VecEvalInt(b.ctx, input, bufWhen); err != nil {
			return err
		}
		whens[j/2] = bufWhen
		whensSlice[j/2] = bufWhen.Int64s()

		bufThen, err := b.bufAllocator.get(types.ETDecimal, n)
		if err != nil {
			return err
		}
		defer b.bufAllocator.put(bufThen)
		if err := args[j+1].VecEvalDecimal(b.ctx, input, bufThen); err != nil {
			return err
		}
		thens[j/2] = bufThen

		thensSlice[j/2] = bufThen.Decimals()

	}
	// when clause(condition, result) -> args[i], args[i+1]; (i >= 0 && i+1 < l-1)
	// else clause -> args[l-1]
	// If case clause has else clause, l%2 == 1.
	if l%2 == 1 {
		bufElse, err := b.bufAllocator.get(types.ETDecimal, n)
		if err != nil {
			return err
		}
		defer b.bufAllocator.put(bufElse)
		if err := args[l-1].VecEvalDecimal(b.ctx, input, bufElse); err != nil {
			return err
		}
		eLse = bufElse

		eLseSlice = bufElse.Decimals()

	}

	result.ResizeDecimal(n, false)
	resultSlice := result.Decimals()

ROW:
	for i := 0; i < n; i++ {
		for j := 0; j < l/2; j++ {
			if whens[j].IsNull(i) || whensSlice[j][i] == 0 {
				continue
			}

			resultSlice[i] = thensSlice[j][i]
			result.SetNull(i, thens[j].IsNull(i))

			continue ROW
		}
		if eLse != nil {

			resultSlice[i] = eLseSlice[i]
			result.SetNull(i, eLse.IsNull(i))

		} else {

			result.SetNull(i, true)

		}
	}
	return nil
}

func (b *builtinCaseWhenDecimalSig) vectorized() bool {
	return true
}

func (b *builtinCaseWhenStringSig) vecEvalString(input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()
	args, l := b.getArgs(), len(b.getArgs())
	whens := make([]*chunk.Column, l/2)
	whensSlice := make([][]int64, l/2)
	thens := make([]*chunk.Column, l/2)
	var eLse *chunk.Column

	for j := 0; j < l-1; j += 2 {
		bufWhen, err := b.bufAllocator.get(types.ETInt, n)
		if err != nil {
			return err
		}
		defer b.bufAllocator.put(bufWhen)
		if err := args[j].VecEvalInt(b.ctx, input, bufWhen); err != nil {
			return err
		}
		whens[j/2] = bufWhen
		whensSlice[j/2] = bufWhen.Int64s()

		bufThen, err := b.bufAllocator.get(types.ETString, n)
		if err != nil {
			return err
		}
		defer b.bufAllocator.put(bufThen)
		if err := args[j+1].VecEvalString(b.ctx, input, bufThen); err != nil {
			return err
		}
		thens[j/2] = bufThen

	}
	// when clause(condition, result) -> args[i], args[i+1]; (i >= 0 && i+1 < l-1)
	// else clause -> args[l-1]
	// If case clause has else clause, l%2 == 1.
	if l%2 == 1 {
		bufElse, err := b.bufAllocator.get(types.ETString, n)
		if err != nil {
			return err
		}
		defer b.bufAllocator.put(bufElse)
		if err := args[l-1].VecEvalString(b.ctx, input, bufElse); err != nil {
			return err
		}
		eLse = bufElse

	}

	result.ReserveString(n)

ROW:
	for i := 0; i < n; i++ {
		for j := 0; j < l/2; j++ {
			if whens[j].IsNull(i) || whensSlice[j][i] == 0 {
				continue
			}

			if thens[j].IsNull(i) {
				result.AppendNull()
			} else {
				result.AppendString(thens[j].GetString(i))
			}

			continue ROW
		}
		if eLse != nil {

			if eLse.IsNull(i) {
				result.AppendNull()
			} else {
				result.AppendString(eLse.GetString(i))
			}

		} else {

			result.AppendNull()

		}
	}
	return nil
}

func (b *builtinCaseWhenStringSig) vectorized() bool {
	return true
}

func (b *builtinCaseWhenTimeSig) vecEvalTime(input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()
	args, l := b.getArgs(), len(b.getArgs())
	whens := make([]*chunk.Column, l/2)
	whensSlice := make([][]int64, l/2)
	thens := make([]*chunk.Column, l/2)
	var eLse *chunk.Column

	thensSlice := make([][]types.Time, l/2)
	var eLseSlice []types.Time

	for j := 0; j < l-1; j += 2 {
		bufWhen, err := b.bufAllocator.get(types.ETInt, n)
		if err != nil {
			return err
		}
		defer b.bufAllocator.put(bufWhen)
		if err := args[j].VecEvalInt(b.ctx, input, bufWhen); err != nil {
			return err
		}
		whens[j/2] = bufWhen
		whensSlice[j/2] = bufWhen.Int64s()

		bufThen, err := b.bufAllocator.get(types.ETDatetime, n)
		if err != nil {
			return err
		}
		defer b.bufAllocator.put(bufThen)
		if err := args[j+1].VecEvalTime(b.ctx, input, bufThen); err != nil {
			return err
		}
		thens[j/2] = bufThen

		thensSlice[j/2] = bufThen.Times()

	}
	// when clause(condition, result) -> args[i], args[i+1]; (i >= 0 && i+1 < l-1)
	// else clause -> args[l-1]
	// If case clause has else clause, l%2 == 1.
	if l%2 == 1 {
		bufElse, err := b.bufAllocator.get(types.ETDatetime, n)
		if err != nil {
			return err
		}
		defer b.bufAllocator.put(bufElse)
		if err := args[l-1].VecEvalTime(b.ctx, input, bufElse); err != nil {
			return err
		}
		eLse = bufElse

		eLseSlice = bufElse.Times()

	}

	result.ResizeTime(n, false)
	resultSlice := result.Times()

ROW:
	for i := 0; i < n; i++ {
		for j := 0; j < l/2; j++ {
			if whens[j].IsNull(i) || whensSlice[j][i] == 0 {
				continue
			}

			resultSlice[i] = thensSlice[j][i]
			result.SetNull(i, thens[j].IsNull(i))

			continue ROW
		}
		if eLse != nil {

			resultSlice[i] = eLseSlice[i]
			result.SetNull(i, eLse.IsNull(i))

		} else {

			result.SetNull(i, true)

		}
	}
	return nil
}

func (b *builtinCaseWhenTimeSig) vectorized() bool {
	return true
}

func (b *builtinCaseWhenDurationSig) vecEvalDuration(input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()
	args, l := b.getArgs(), len(b.getArgs())
	whens := make([]*chunk.Column, l/2)
	whensSlice := make([][]int64, l/2)
	thens := make([]*chunk.Column, l/2)
	var eLse *chunk.Column

	thensSlice := make([][]time.Duration, l/2)
	var eLseSlice []time.Duration

	for j := 0; j < l-1; j += 2 {
		bufWhen, err := b.bufAllocator.get(types.ETInt, n)
		if err != nil {
			return err
		}
		defer b.bufAllocator.put(bufWhen)
		if err := args[j].VecEvalInt(b.ctx, input, bufWhen); err != nil {
			return err
		}
		whens[j/2] = bufWhen
		whensSlice[j/2] = bufWhen.Int64s()

		bufThen, err := b.bufAllocator.get(types.ETDuration, n)
		if err != nil {
			return err
		}
		defer b.bufAllocator.put(bufThen)
		if err := args[j+1].VecEvalDuration(b.ctx, input, bufThen); err != nil {
			return err
		}
		thens[j/2] = bufThen

		thensSlice[j/2] = bufThen.GoDurations()

	}
	// when clause(condition, result) -> args[i], args[i+1]; (i >= 0 && i+1 < l-1)
	// else clause -> args[l-1]
	// If case clause has else clause, l%2 == 1.
	if l%2 == 1 {
		bufElse, err := b.bufAllocator.get(types.ETDuration, n)
		if err != nil {
			return err
		}
		defer b.bufAllocator.put(bufElse)
		if err := args[l-1].VecEvalDuration(b.ctx, input, bufElse); err != nil {
			return err
		}
		eLse = bufElse

		eLseSlice = bufElse.GoDurations()

	}

	result.ResizeGoDuration(n, false)
	resultSlice := result.GoDurations()

ROW:
	for i := 0; i < n; i++ {
		for j := 0; j < l/2; j++ {
			if whens[j].IsNull(i) || whensSlice[j][i] == 0 {
				continue
			}

			resultSlice[i] = thensSlice[j][i]
			result.SetNull(i, thens[j].IsNull(i))

			continue ROW
		}
		if eLse != nil {

			resultSlice[i] = eLseSlice[i]
			result.SetNull(i, eLse.IsNull(i))

		} else {

			result.SetNull(i, true)

		}
	}
	return nil
}

func (b *builtinCaseWhenDurationSig) vectorized() bool {
	return true
}

func (b *builtinCaseWhenJSONSig) vecEvalJSON(input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()
	args, l := b.getArgs(), len(b.getArgs())
	whens := make([]*chunk.Column, l/2)
	whensSlice := make([][]int64, l/2)
	thens := make([]*chunk.Column, l/2)
	var eLse *chunk.Column

	for j := 0; j < l-1; j += 2 {
		bufWhen, err := b.bufAllocator.get(types.ETInt, n)
		if err != nil {
			return err
		}
		defer b.bufAllocator.put(bufWhen)
		if err := args[j].VecEvalInt(b.ctx, input, bufWhen); err != nil {
			return err
		}
		whens[j/2] = bufWhen
		whensSlice[j/2] = bufWhen.Int64s()

		bufThen, err := b.bufAllocator.get(types.ETJson, n)
		if err != nil {
			return err
		}
		defer b.bufAllocator.put(bufThen)
		if err := args[j+1].VecEvalJSON(b.ctx, input, bufThen); err != nil {
			return err
		}
		thens[j/2] = bufThen

	}
	// when clause(condition, result) -> args[i], args[i+1]; (i >= 0 && i+1 < l-1)
	// else clause -> args[l-1]
	// If case clause has else clause, l%2 == 1.
	if l%2 == 1 {
		bufElse, err := b.bufAllocator.get(types.ETJson, n)
		if err != nil {
			return err
		}
		defer b.bufAllocator.put(bufElse)
		if err := args[l-1].VecEvalJSON(b.ctx, input, bufElse); err != nil {
			return err
		}
		eLse = bufElse

	}

	result.ReserveJSON(n)

ROW:
	for i := 0; i < n; i++ {
		for j := 0; j < l/2; j++ {
			if whens[j].IsNull(i) || whensSlice[j][i] == 0 {
				continue
			}

			if thens[j].IsNull(i) {
				result.AppendNull()
			} else {
				result.AppendJSON(thens[j].GetJSON(i))
			}

			continue ROW
		}
		if eLse != nil {

			if eLse.IsNull(i) {
				result.AppendNull()
			} else {
				result.AppendJSON(eLse.GetJSON(i))
			}

		} else {

			result.AppendNull()

		}
	}
	return nil
}

func (b *builtinCaseWhenJSONSig) vectorized() bool {
	return true
}

func (b *builtinIfNullIntSig) vecEvalInt(input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()

	if err := b.args[0].VecEvalInt(b.ctx, input, result); err != nil {
		return err
	}
	buf1, err := b.bufAllocator.get(types.ETInt, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)
	if err := b.args[1].VecEvalInt(b.ctx, input, buf1); err != nil {
		return err
	}

	arg0 := result.Int64s()
	arg1 := buf1.Int64s()
	for i := 0; i < n; i++ {
		if result.IsNull(i) && !buf1.IsNull(i) {
			result.SetNull(i, false)
			arg0[i] = arg1[i]
		}
	}

	return nil
}

func (b *builtinIfNullIntSig) vectorized() bool {
	return true
}

func (b *builtinIfNullRealSig) vecEvalReal(input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()

	if err := b.args[0].VecEvalReal(b.ctx, input, result); err != nil {
		return err
	}
	buf1, err := b.bufAllocator.get(types.ETReal, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)
	if err := b.args[1].VecEvalReal(b.ctx, input, buf1); err != nil {
		return err
	}

	arg0 := result.Float64s()
	arg1 := buf1.Float64s()
	for i := 0; i < n; i++ {
		if result.IsNull(i) && !buf1.IsNull(i) {
			result.SetNull(i, false)
			arg0[i] = arg1[i]
		}
	}

	return nil
}

func (b *builtinIfNullRealSig) vectorized() bool {
	return true
}

func (b *builtinIfNullDecimalSig) vecEvalDecimal(input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()

	if err := b.args[0].VecEvalDecimal(b.ctx, input, result); err != nil {
		return err
	}
	buf1, err := b.bufAllocator.get(types.ETDecimal, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)
	if err := b.args[1].VecEvalDecimal(b.ctx, input, buf1); err != nil {
		return err
	}

	arg0 := result.Decimals()
	arg1 := buf1.Decimals()
	for i := 0; i < n; i++ {
		if result.IsNull(i) && !buf1.IsNull(i) {
			result.SetNull(i, false)
			arg0[i] = arg1[i]
		}
	}

	return nil
}

func (b *builtinIfNullDecimalSig) vectorized() bool {
	return true
}

func (b *builtinIfNullStringSig) vecEvalString(input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()

	buf0, err := b.bufAllocator.get(types.ETString, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf0)
	if err := b.args[0].VecEvalString(b.ctx, input, buf0); err != nil {
		return err
	}
	buf1, err := b.bufAllocator.get(types.ETString, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)
	if err := b.args[1].VecEvalString(b.ctx, input, buf1); err != nil {
		return err
	}
	result.ReserveString(n)

	for i := 0; i < n; i++ {
		if !buf0.IsNull(i) {
			result.AppendString(buf0.GetString(i))
		} else if !buf1.IsNull(i) {
			result.AppendString(buf1.GetString(i))
		} else {
			result.AppendNull()
		}
	}

	return nil
}

func (b *builtinIfNullStringSig) vectorized() bool {
	return true
}

func (b *builtinIfNullTimeSig) vecEvalTime(input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()

	if err := b.args[0].VecEvalTime(b.ctx, input, result); err != nil {
		return err
	}
	buf1, err := b.bufAllocator.get(types.ETDatetime, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)
	if err := b.args[1].VecEvalTime(b.ctx, input, buf1); err != nil {
		return err
	}

	arg0 := result.Times()
	arg1 := buf1.Times()
	for i := 0; i < n; i++ {
		if result.IsNull(i) && !buf1.IsNull(i) {
			result.SetNull(i, false)
			arg0[i] = arg1[i]
		}
	}

	return nil
}

func (b *builtinIfNullTimeSig) vectorized() bool {
	return true
}

func (b *builtinIfNullDurationSig) vecEvalDuration(input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()

	if err := b.args[0].VecEvalDuration(b.ctx, input, result); err != nil {
		return err
	}
	buf1, err := b.bufAllocator.get(types.ETDuration, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)
	if err := b.args[1].VecEvalDuration(b.ctx, input, buf1); err != nil {
		return err
	}

	arg0 := result.GoDurations()
	arg1 := buf1.GoDurations()
	for i := 0; i < n; i++ {
		if result.IsNull(i) && !buf1.IsNull(i) {
			result.SetNull(i, false)
			arg0[i] = arg1[i]
		}
	}

	return nil
}

func (b *builtinIfNullDurationSig) vectorized() bool {
	return true
}

func (b *builtinIfNullJSONSig) vecEvalJSON(input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()

	buf0, err := b.bufAllocator.get(types.ETJson, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf0)
	if err := b.args[0].VecEvalJSON(b.ctx, input, buf0); err != nil {
		return err
	}
	buf1, err := b.bufAllocator.get(types.ETJson, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)
	if err := b.args[1].VecEvalJSON(b.ctx, input, buf1); err != nil {
		return err
	}
	result.ReserveJSON(n)

	for i := 0; i < n; i++ {
		if !buf0.IsNull(i) {
			result.AppendJSON(buf0.GetJSON(i))
		} else if !buf1.IsNull(i) {
			result.AppendJSON(buf1.GetJSON(i))
		} else {
			result.AppendNull()
		}
	}

	return nil
}

func (b *builtinIfNullJSONSig) vectorized() bool {
	return true
}

func (b *builtinIfIntSig) vecEvalInt(input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()
	buf0, err := b.bufAllocator.get(types.ETInt, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf0)
	if err := b.args[0].VecEvalInt(b.ctx, input, buf0); err != nil {
		return err
	}

	if err := b.args[1].VecEvalInt(b.ctx, input, result); err != nil {
		return err
	}

	buf2, err := b.bufAllocator.get(types.ETInt, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf2)
	if err := b.args[2].VecEvalInt(b.ctx, input, buf2); err != nil {
		return err
	}

	arg0 := buf0.Int64s()

	arg2 := buf2.Int64s()
	rs := result.Int64s()

	for i := 0; i < n; i++ {
		arg := arg0[i]
		isNull0 := buf0.IsNull(i)
		switch {
		case isNull0 || arg == 0:

			if buf2.IsNull(i) {
				result.SetNull(i, true)
			} else {
				result.SetNull(i, false)
				rs[i] = arg2[i]
			}

		case arg != 0:

		}
	}
	return nil
}

func (b *builtinIfIntSig) vectorized() bool {
	return true
}

func (b *builtinIfRealSig) vecEvalReal(input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()
	buf0, err := b.bufAllocator.get(types.ETInt, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf0)
	if err := b.args[0].VecEvalInt(b.ctx, input, buf0); err != nil {
		return err
	}

	if err := b.args[1].VecEvalReal(b.ctx, input, result); err != nil {
		return err
	}

	buf2, err := b.bufAllocator.get(types.ETReal, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf2)
	if err := b.args[2].VecEvalReal(b.ctx, input, buf2); err != nil {
		return err
	}

	arg0 := buf0.Int64s()

	arg2 := buf2.Float64s()
	rs := result.Float64s()

	for i := 0; i < n; i++ {
		arg := arg0[i]
		isNull0 := buf0.IsNull(i)
		switch {
		case isNull0 || arg == 0:

			if buf2.IsNull(i) {
				result.SetNull(i, true)
			} else {
				result.SetNull(i, false)
				rs[i] = arg2[i]
			}

		case arg != 0:

		}
	}
	return nil
}

func (b *builtinIfRealSig) vectorized() bool {
	return true
}

func (b *builtinIfDecimalSig) vecEvalDecimal(input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()
	buf0, err := b.bufAllocator.get(types.ETInt, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf0)
	if err := b.args[0].VecEvalInt(b.ctx, input, buf0); err != nil {
		return err
	}

	if err := b.args[1].VecEvalDecimal(b.ctx, input, result); err != nil {
		return err
	}

	buf2, err := b.bufAllocator.get(types.ETDecimal, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf2)
	if err := b.args[2].VecEvalDecimal(b.ctx, input, buf2); err != nil {
		return err
	}

	arg0 := buf0.Int64s()

	arg2 := buf2.Decimals()
	rs := result.Decimals()

	for i := 0; i < n; i++ {
		arg := arg0[i]
		isNull0 := buf0.IsNull(i)
		switch {
		case isNull0 || arg == 0:

			if buf2.IsNull(i) {
				result.SetNull(i, true)
			} else {
				result.SetNull(i, false)
				rs[i] = arg2[i]
			}

		case arg != 0:

		}
	}
	return nil
}

func (b *builtinIfDecimalSig) vectorized() bool {
	return true
}

func (b *builtinIfStringSig) vecEvalString(input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()
	buf0, err := b.bufAllocator.get(types.ETInt, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf0)
	if err := b.args[0].VecEvalInt(b.ctx, input, buf0); err != nil {
		return err
	}

	buf1, err := b.bufAllocator.get(types.ETString, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)
	if err := b.args[1].VecEvalString(b.ctx, input, buf1); err != nil {
		return err
	}

	buf2, err := b.bufAllocator.get(types.ETString, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf2)
	if err := b.args[2].VecEvalString(b.ctx, input, buf2); err != nil {
		return err
	}

	result.ReserveString(n)

	arg0 := buf0.Int64s()

	for i := 0; i < n; i++ {
		arg := arg0[i]
		isNull0 := buf0.IsNull(i)
		switch {
		case isNull0 || arg == 0:

			if buf2.IsNull(i) {
				result.AppendNull()
			} else {
				result.AppendString(buf2.GetString(i))
			}

		case arg != 0:

			if buf1.IsNull(i) {
				result.AppendNull()
			} else {
				result.AppendString(buf1.GetString(i))
			}

		}
	}
	return nil
}

func (b *builtinIfStringSig) vectorized() bool {
	return true
}

func (b *builtinIfTimeSig) vecEvalTime(input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()
	buf0, err := b.bufAllocator.get(types.ETInt, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf0)
	if err := b.args[0].VecEvalInt(b.ctx, input, buf0); err != nil {
		return err
	}

	if err := b.args[1].VecEvalTime(b.ctx, input, result); err != nil {
		return err
	}

	buf2, err := b.bufAllocator.get(types.ETDatetime, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf2)
	if err := b.args[2].VecEvalTime(b.ctx, input, buf2); err != nil {
		return err
	}

	arg0 := buf0.Int64s()

	arg2 := buf2.Times()
	rs := result.Times()

	for i := 0; i < n; i++ {
		arg := arg0[i]
		isNull0 := buf0.IsNull(i)
		switch {
		case isNull0 || arg == 0:

			if buf2.IsNull(i) {
				result.SetNull(i, true)
			} else {
				result.SetNull(i, false)
				rs[i] = arg2[i]
			}

		case arg != 0:

		}
	}
	return nil
}

func (b *builtinIfTimeSig) vectorized() bool {
	return true
}

func (b *builtinIfDurationSig) vecEvalDuration(input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()
	buf0, err := b.bufAllocator.get(types.ETInt, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf0)
	if err := b.args[0].VecEvalInt(b.ctx, input, buf0); err != nil {
		return err
	}

	if err := b.args[1].VecEvalDuration(b.ctx, input, result); err != nil {
		return err
	}

	buf2, err := b.bufAllocator.get(types.ETDuration, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf2)
	if err := b.args[2].VecEvalDuration(b.ctx, input, buf2); err != nil {
		return err
	}

	arg0 := buf0.Int64s()

	arg2 := buf2.GoDurations()
	rs := result.GoDurations()

	for i := 0; i < n; i++ {
		arg := arg0[i]
		isNull0 := buf0.IsNull(i)
		switch {
		case isNull0 || arg == 0:

			if buf2.IsNull(i) {
				result.SetNull(i, true)
			} else {
				result.SetNull(i, false)
				rs[i] = arg2[i]
			}

		case arg != 0:

		}
	}
	return nil
}

func (b *builtinIfDurationSig) vectorized() bool {
	return true
}

func (b *builtinIfJSONSig) vecEvalJSON(input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()
	buf0, err := b.bufAllocator.get(types.ETInt, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf0)
	if err := b.args[0].VecEvalInt(b.ctx, input, buf0); err != nil {
		return err
	}

	buf1, err := b.bufAllocator.get(types.ETJson, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)
	if err := b.args[1].VecEvalJSON(b.ctx, input, buf1); err != nil {
		return err
	}

	buf2, err := b.bufAllocator.get(types.ETJson, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf2)
	if err := b.args[2].VecEvalJSON(b.ctx, input, buf2); err != nil {
		return err
	}

	result.ReserveJSON(n)

	arg0 := buf0.Int64s()

	for i := 0; i < n; i++ {
		arg := arg0[i]
		isNull0 := buf0.IsNull(i)
		switch {
		case isNull0 || arg == 0:

			if buf2.IsNull(i) {
				result.AppendNull()
			} else {
				result.AppendJSON(buf2.GetJSON(i))
			}

		case arg != 0:

			if buf1.IsNull(i) {
				result.AppendNull()
			} else {
				result.AppendJSON(buf1.GetJSON(i))
			}

		}
	}
	return nil
}

func (b *builtinIfJSONSig) vectorized() bool {
	return true
}