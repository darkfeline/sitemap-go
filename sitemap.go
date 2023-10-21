// Copyright (C) 2019  Allen Li
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package sitemap implements sitemap generation according to
// https://www.sitemaps.org/protocol.html.
package sitemap

import (
	"encoding/xml"
	"fmt"
	"io"
)

// XMLNamespace is the XML namespace for sitemap elements.
const XMLNamespace = "http://www.sitemaps.org/schemas/sitemap/0.9"

// URLSet is the Go representation of a urlset element.
type URLSet struct {
	XMLName   xml.Name `xml:"urlset"`
	Namespace string   `xml:"namespace,attr"`
	URLs      []URL
}

// URL is the Go representation of a url element.
type URL struct {
	XMLName    xml.Name `xml:"url"`
	Loc        string   `xml:"loc"`
	LastMod    string   `xml:"lastmod,omitempty"`
	ChangeFreq Freq     `xml:"changefreq,omitempty"`
	Priority   string   `xml:"priority,omitempty"`
}

// Freq are the valid values for the ChangeFreq field in URL.
type Freq string

const (
	Always  Freq = "always"
	Hourly  Freq = "hourly"
	Daily   Freq = "daily"
	Weekly  Freq = "weekly"
	Monthly Freq = "monthly"
	Yearly  Freq = "yearly"
	Never   Freq = "never"
)

// Write writes a sitemap.
func Write(w io.Writer, u *URLSet) error {
	setDefaultNamespace(u)
	if _, err := io.WriteString(w, xml.Header); err != nil {
		return fmt.Errorf("write sitemap: %w", err)
	}
	e := xml.NewEncoder(w)
	if err := e.Encode(u); err != nil {
		return fmt.Errorf("write sitemap: %w", err)
	}
	return nil
}

// WritePretty writes a sitemap in a human-friendly format.
func WritePretty(w io.Writer, u *URLSet) error {
	setDefaultNamespace(u)
	if _, err := io.WriteString(w, xml.Header); err != nil {
		return fmt.Errorf("write sitemap: %w", err)
	}
	e := xml.NewEncoder(w)
	e.Indent("", "  ")
	if err := e.Encode(u); err != nil {
		return fmt.Errorf("write sitemap: %w", err)
	}
	return nil
}

func setDefaultNamespace(u *URLSet) {
	if u.Namespace == "" {
		u.Namespace = XMLNamespace
	}
}
