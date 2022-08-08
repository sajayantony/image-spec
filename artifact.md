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

  This REQUIRED property specifies the mediaType of the referenced content and MAY be registered with [IANA][iana].

- **`blobs`** *array of objects* 

  This OPTIONAL property contains a list of blob [descriptors](descriptor.md).

- **`refers`** *[descriptor](descriptor.md)*

  This OPTIONAL property specifies a [descriptor](descriptor.md) of a container image or another artifact. The purpose of this property is to provide a reference to the container image or artifact this artifact is related to. The "Referrers" API in the distribution specification looks for this property to list all artifacts that refer to a given artifact or container image.

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

*Example showing an artifact manifest for an image signature:*

```jsonc,title=Manifest&mediatype=application/vnd.oci.artifact.manifest.v1%2Bjson
{
  "schemaVersion": 2,
  "mediaType": "application/vnd.oci.artifact.manifest.v1+json",
  "artifactType" "application/example"
  "blobs": [
    {
      "mediaType": "application/vnd.dev.cosign.simplesigning.v1+json",
      "size": 210,
      "digest": "sha256:1119abab63e605dcc281019bad0424744178b6f61ba57378701fe7391994c999"
    }
  ],
  "refers": {
    "mediaType": "application/vnd.oci.image.manifest.v1+json",
    "size": 7682,
    "digest": "sha256:5b0bcabd1ed22e9fb1310cf6c2dec7cdef19f0ad69efa1f392e94a4333501270"
  },
  "annotations": [
    "org.opencontainers.artifact.description": "cosign signature for image:tag",
    "org.opencontainers.artifact.created": "2022-04-05T14:30Z"
  ]
}
```
[iana]:         https://www.iana.org/assignments/media-types/media-types.xhtml
[rfc-3339]:     https://tools.ietf.org/html/rfc3339#section-5.6