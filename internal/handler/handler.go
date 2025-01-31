// SPDX-License-Identifier: Apache-2.0

package handler

import ()

// Handler ...
type Handler interface {
	Run() error
	Complete() error
}
