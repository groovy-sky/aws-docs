# ImageReferrer

An object representing an artifact associated with a subject image.

## Contents

**digest**

The digest of the artifact manifest.

Type: String

Required: Yes

**mediaType**

The media type of the artifact manifest.

Type: String

Required: Yes

**size**

The size, in bytes, of the artifact.

Type: Long

Required: Yes

**annotations**

A map of annotations associated with the artifact.

Type: String to string map

Required: No

**artifactStatus**

The status of the artifact. Valid values are `ACTIVE`, `ARCHIVED`, or `ACTIVATING`.

Type: String

Valid Values: `ACTIVE | ARCHIVED | ACTIVATING`

Required: No

**artifactType**

A string identifying the type of artifact.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/imagereferrer.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/imagereferrer.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/imagereferrer.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImageIdentifier

ImageReplicationStatus

All content copied from https://docs.aws.amazon.com/.
