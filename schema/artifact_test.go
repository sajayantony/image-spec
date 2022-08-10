// Copyright 2016 The Linux Foundation
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

package schema_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/opencontainers/image-spec/schema"
)

func TestArtifact(t *testing.T) {
	for i, tt := range []struct {
		manifest string
		fail     bool
	}{
		// expected failure: mediaType does not match pattern
		{
			manifest: `
{
  "mediaType" : "invalid",
  "artifactType" : "application/example",
  "blobs": [
    {
      "mediaType": "application/vnd.oci.image.layer.v1.tar+gzip",
      "size": 148,
      "digest": "sha256:c57089565e894899735d458f0fd4bb17a0f1e0df8d72da392b85c9b35ee777cd"
    }
  ]
}
`,
			fail: true,
		},

		// expected failure: invalid artifact mediaType
		{
			manifest: `
{
  "mediaType" : "application/vnd.oci.artifact.manifest.v1+json",
  "artifactType" : "invalid",
  "blobs": [
    {
      "mediaType": "application/vnd.oci.image.layer.v1.tar+gzip",
      "size": "148",
      "digest": "sha256:c57089565e894899735d458f0fd4bb17a0f1e0df8d72da392b85c9b35ee777cd"
    }
  ]
}
`,
			fail: true,
		},

		// expected failure: blob[0].size is a string, expected integer
		{
			manifest: `
{
  "mediaType" : "application/vnd.oci.artifact.manifest.v1+json",
  "artifactType" : "application/example",
  "blobs": [
    {
      "mediaType": "application/vnd.oci.image.layer.v1.tar+gzip",
      "size": "148",
      "digest": "sha256:c57089565e894899735d458f0fd4bb17a0f1e0df8d72da392b85c9b35ee777cd"
    }
  ]
}
`,
			fail: true,
		},

		// valid manifest with optional fields
		{
			manifest: `
{
  "mediaType" : "application/vnd.oci.artifact.manifest.v1+json",
  "artifactType" : "application/example",
  "blobs": [
    {
      "mediaType": "application/vnd.oci.image.layer.v1.tar+gzip",
      "size": 675598,
      "digest": "sha256:9d3dd9504c685a304985025df4ed0283e47ac9ffa9bd0326fddf4d59513f0827"
    },
    {
      "mediaType": "application/vnd.oci.image.layer.v1.tar+gzip",
      "size": 156,
      "digest": "sha256:2b689805fbd00b2db1df73fae47562faac1a626d5f61744bfe29946ecff5d73d"
    },
    {
      "mediaType": "application/vnd.oci.image.layer.v1.tar+gzip",
      "size": 148,
      "digest": "sha256:c57089565e894899735d458f0fd4bb17a0f1e0df8d72da392b85c9b35ee777cd"
    }
  ],
  "annotations": {
    "key1": "value1",
    "key2": "value2"
  }
}
`,
			fail: false,
		},
	} {
		r := strings.NewReader(tt.manifest)
		err := schema.ValidatorMediaTypeArtifact.Validate(r)

		if got := err != nil; tt.fail != got {
			t.Errorf("test %d: expected validation failure %t but got %t, err %v", i, tt.fail, got, err)
			fmt.Println(tt.manifest)
		}
	}
}
