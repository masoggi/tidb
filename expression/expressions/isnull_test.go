// Copyright 2015 PingCAP, Inc.
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

package expressions

import (
	"errors"

	. "github.com/pingcap/check"
)

var _ = Suite(&testIsNullSuite{})

type testIsNullSuite struct {
}

func (t *testIsNullSuite) TestIsNull(c *C) {
	e := &IsNull{
		Expr: Value{1},
	}

	v, err := e.Eval(nil, nil)
	c.Assert(err, IsNil)
	c.Assert(v, IsFalse)

	c.Assert(e.IsStatic(), IsTrue)

	str := e.String()
	c.Assert(len(str), Greater, 0)

	ec, err := e.Clone()
	c.Assert(err, IsNil)

	e2, ok := ec.(*IsNull)
	c.Assert(ok, IsTrue)

	e2.Not = true

	vv, err := e2.Eval(nil, nil)
	c.Assert(err, IsNil)
	c.Assert(vv, IsTrue)

	str = e2.String()
	c.Assert(len(str), Greater, 0)

	// check error
	expr := mockExpr{}
	expr.err = errors.New("must error")
	e.Expr = expr

	_, err = e.Clone()
	c.Assert(err, NotNil)

	_, err = e.Eval(nil, nil)
	c.Assert(err, NotNil)

	e.Expr = newTestRow(1, 2)
	_, err = e.Eval(nil, nil)
	c.Assert(err, NotNil)
}
