// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = GreaterOrEqualFunction{}
)

func NewGreaterOrEqualFunction() function.Function {
	return GreaterOrEqualFunction{}
}

type GreaterOrEqualFunction struct{}

func (r GreaterOrEqualFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "greater_or_equal"
}

func (r GreaterOrEqualFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a number is greater than or equal to a given number",
		Parameters: []function.Parameter{
			function.NumberParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The number to compare against",
				Name:               "compare_against",
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

func (r GreaterOrEqualFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var compareAgainst *big.Float
	var number *big.Float

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &compareAgainst, &number))
	if resp.Error != nil {
		return
	}
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, isGreaterOrEqual(number, compareAgainst)))
}

func isGreaterOrEqual(number, compareAgainst *big.Float) bool {
	return number.Cmp(compareAgainst) >= 0
}
