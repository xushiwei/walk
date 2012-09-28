// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package declarative

import (
	"github.com/lxn/walk"
)

type LineEdit struct {
	Widget     **walk.LineEdit
	Name       string
	HStretch   int
	VStretch   int
	Row        int
	RowSpan    int
	Column     int
	ColumnSpan int
	Font       Font
	Text       string
	ReadOnly   bool
	MaxLength  int
}

func (le LineEdit) Create(parent walk.Container) (walk.Widget, error) {
	w, err := walk.NewLineEdit(parent)
	if err != nil {
		return nil, err
	}

	var succeeded bool
	defer func() {
		if !succeeded {
			w.Dispose()
		}
	}()

	w.SetName(le.Name)

	f, err := le.Font.Create()
	if err != nil {
		return nil, err
	}

	if f != nil {
		w.SetFont(f)
	}

	if err := w.SetText(le.Text); err != nil {
		return nil, err
	}

	w.SetReadOnly(le.ReadOnly)

	w.SetMaxLength(le.MaxLength)

	if le.Widget != nil {
		*le.Widget = w
	}

	succeeded = true

	return w, nil
}
