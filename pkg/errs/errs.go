// Copyright 2020 TiKV Project Authors.
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

package errs

import (
	"github.com/pingcap/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapError is used to make the log output eaiser.
func ZapError(err *errors.Error, causeError error) zap.Field {
	e := err.Wrap(causeError).FastGenWithCause()
	return zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: e}
}
