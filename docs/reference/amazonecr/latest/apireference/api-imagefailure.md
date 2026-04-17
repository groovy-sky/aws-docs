---
title: "ImageFailure"
---

# ImageFailure

An object representing an Amazon ECR image failure.

## Contents

**failureCode**

The code associated with the failure.

Type: String

Valid Values: `InvalidImageDigest | InvalidImageTag | ImageTagDoesNotMatchDigest | ImageNotFound | MissingDigestAndTag | ImageReferencedByManifestList | KmsError | UpstreamAccessDenied | UpstreamTooManyRequests | UpstreamUnavailable | ImageInaccessible`

Required: No

**failureReason**

The reason for the failure.

Type: String

Required: No

**imageId**

The image ID associated with the failure.

Type: [ImageIdentifier](api-imageidentifier.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ecr-2015-09-21/ImageFailure)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ecr-2015-09-21/ImageFailure)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ecr-2015-09-21/ImageFailure)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImageDetail

ImageIdentifier

All content copied from https://docs.aws.amazon.com/.
