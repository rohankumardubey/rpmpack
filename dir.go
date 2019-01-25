// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rpmpack

// DirIndex holds the index from files to directory names.
type DirIndex struct {
	m map[string]int32
	l []string
}

func NewDirIndex() *DirIndex {
	return &DirIndex{m: make(map[string]int32)}
}

func (d *DirIndex) Get(value string) int32 {

	if idx, ok := d.m[value]; ok {
		return idx
	}
	newIdx := int32(len(d.l))
	d.l = append(d.l, value)

	d.m[value] = newIdx
	return newIdx
}

func (d *DirIndex) AllDirs() []string {
	return d.l
}
