// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = BetweenFunction{}
)

func NewBetweenFunction() function.Function {
	return BetweenFunction{}
}

type BetweenFunction struct{}

func (r BetweenFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "between"
}

func (r BetweenFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a number is within a given range",
		Parameters: []function.Parameter{
			function.NumberParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The beginning of the range",
				Name:               "begin",
			},
			function.NumberParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The end of the range",
				Name:               "end",
			},
			function.NumberParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The number to check",
				Name:               "number",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r BetweenFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var begin, end, number *big.Float

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &begin, &end, &number))
	if resp.Error != nil {
		return
	}
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, isInRange(number, begin, end)))
}

func isInRange(number, start, end *big.Float) bool {
	return number.Cmp(start) != -1 && number.Cmp(end) != 1
}
