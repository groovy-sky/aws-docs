# Image

An object representing an Amazon ECR image.

## Contents

**imageId**

An object containing the image tag and image digest associated with an image.

Type: [ImageIdentifier](api-imageidentifier.md) object

Required: No

**imageManifest**

The image manifest associated with the image.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 4194304.

Required: No

**imageManifestMediaType**

The manifest media type of the image.

Type: String

Required: No

**registryId**

The AWS account ID associated with the registry containing the image.

Type: String

Pattern: `[0-9]{12}`

Required: No

**repositoryName**

The name of the repository associated with the image.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 256.

Pattern: `[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/image.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/image.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/image.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EnhancedImageScanFinding

ImageDetail

All content copied from https://docs.aws.amazon.com/.
