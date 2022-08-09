# OCI Artifact Manifest Specification

The goal of the Artifact Manifest Specification is to define content addressable artifacts in order to store them along side container images in a registry. Like [OCI Images](manifest.md), OCI Artifacts may be referenced by the hash of their manifest. Unlike OCI Images, OCI Artifacts are not meant to be used by any container runtime.

Examples of artifacts that may be stored along with container images are Software Bill of Materials (SBOM), Digital Signatures, Provenance data, Supply Chain Attestations, scan results, and Helm charts.

This section defines the `application/vnd.oci.artifact.manifest.v1+json` [media type](media-types.md).
For the media type(s) that this is compatible with see the [matrix](media-types.md#compatibility-matrix).

# Artifact Manifest

## *Artifact Manifest* Property Descriptions

- **`mediaType`** *string*

  This property MUST be used and contain the media type `application/vnd.oci.artifact.manifest.v1+json`.

- **`artifactType`** *string*

  This property SHOULD specify the mediaType of the referenced content and MAY be registered with [IANA][iana].

- **`blobs`** *string*

  This OPTIONAL property contains a list of [descriptors](descriptor.md).
  Each descriptor represents an artifact of any IANA mediaType.
  The list MAY be ordered for certain artifact types like scan results.

- **`refers`** *string*

  This OPTIONAL property specifies a [descriptor](descriptor.md) of a container image or another artifact.
  The purpose of this property is to provide a reference to the container image or artifact this artifact is related to.
  The "Referrers" API in the distribution specification looks for this property to list all artifacts that refer to a given artifact or container image.

- **`annotations`** *string-string map*

  This OPTIONAL property contains additional metadata for the artifact manifest.
  This OPTIONAL property MUST use the [annotation rules](annotations.md#rules).

  The following annotations MAY be used:

  - `org.opencontainers.artifact.description`: human readable description for the artifact
  - `org.opencontainers.artifact.created`: creation time of the artifact expressed as string defined by [RFC 3339][rfc-3339]

  Additionally, the following annotations SHOULD be used when deploying multi-arch container images:

  - `org.opencontainers.platform.architecture`: CPU architecture for binaries
  - `org.opencontainers.platform.os`: operating system for binaries
  - `org.opencontainers.platform.os.version`: operating system version for binaries
  - `org.opencontainers.platform.variant`: variant of the CPU architecture for binaries

    Also, see [Pre-Defined Annotation Keys](annotations.md#pre-defined-annotation-keys).

  User defined annotations MAY be used to filter various artifact types, e.g. signature public key hash, attestation type, and SBOM schema.

## Examples

*Example showing an artifact manifest for an ice cream flavor referencing an ice cream:*

```jsonc,title=Manifest&mediatype=application/vnd.oci.artifact.manifest.v1%2Bjson
{
  "schemaVersion": 2,
  "mediaType": "application/vnd.oci.artifact.manifest.v1+json",
  "artifactType" "application/example"
  "blobs": [
    {
      "mediaType": "application/vnd.icecream.flavor",
      "size": 123,
      "digest": "sha256:87923725d74f4bfb94c9e86d64170f7521aad8221a5de834851470ca142da630"
    }
  ],
  "refers": {
    "mediaType": "application/vnd.icecream",
    "size": 1234,
    "digest": "sha256:cc06a2839488b8bd2a2b99dcdc03d5cfd818eed72ad08ef3cc197aac64c0d0a0"
  },
  "annotations": [
    "org.opencontainers.artifact.description": "vanilla surprise",
    "org.opencontainers.artifact.created": "2022-04-05T14:30Z"
  ]
}
```

[iana]:         https://www.iana.org/assignments/media-types/media-types.xhtml
[rfc-3339]:     https://tools.ietf.org/html/rfc3339#section-5.6