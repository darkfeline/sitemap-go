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

package sitemap

import "os"

func ExampleWritePretty() {
	u := URLSet{
		URLs: []URL{
			{
				Loc:        "http://www.example.com/foo",
				LastMod:    "2001-02-03",
				ChangeFreq: Always,
				Priority:   "0.5",
			},
			{
				Loc: "http://www.example.com/bar",
			},
		},
	}
	if err := WritePretty(os.Stdout, &u); err != nil {
		panic(err)
	}
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <urlset Namespace="http://www.sitemaps.org/schemas/sitemap/0.9">
	//   <url>
	//     <loc>http://www.example.com/foo</loc>
	//     <lastmod>2001-02-03</lastmod>
	//     <changefreq>always</changefreq>
	//     <priority>0.5</priority>
	//   </url>
	//   <url>
	//     <loc>http://www.example.com/bar</loc>
	//   </url>
	// </urlset>
}
