// Copyright 2023 The CortexTheseus Authors
// This file is part of the CortexTheseus library.
//
// The CortexTheseus library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The CortexTheseus library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the CortexTheseus library. If not, see <http://www.gnu.org/licenses/>.

// Package core implements the Cortex consensus protocol.

package metrics

import (
	"testing"
)

func TestGaugeInfoJsonString(t *testing.T) {
	g := NewGaugeInfo()
	g.Update(GaugeInfoValue{
		"chain_id":   "5",
		"anotherKey": "any_string_value",
		"third_key":  "anything",
	},
	)
	want := `{"anotherKey":"any_string_value","chain_id":"5","third_key":"anything"}`

	original := g.Snapshot()
	g.Update(GaugeInfoValue{"value": "updated"})

	if have := original.Value().String(); have != want {
		t.Errorf("\nhave: %v\nwant: %v\n", have, want)
	}
	if have, want := g.Snapshot().Value().String(), `{"value":"updated"}`; have != want {
		t.Errorf("\nhave: %v\nwant: %v\n", have, want)
	}
}

func TestGetOrRegisterGaugeInfo(t *testing.T) {
	r := NewRegistry()
	NewRegisteredGaugeInfo("foo", r).Update(
		GaugeInfoValue{"chain_id": "5"})
	g := GetOrRegisterGaugeInfo("foo", r).Snapshot()
	if have, want := g.Value().String(), `{"chain_id":"5"}`; have != want {
		t.Errorf("have\n%v\nwant\n%v\n", have, want)
	}
}
