---
title: "Layer"
---

# Layer

An object representing an Amazon ECR image layer.

## Contents

**layerAvailability**

The availability status of the image layer.

Type: String

Valid Values: `AVAILABLE | UNAVAILABLE | ARCHIVED`

Required: No

**layerDigest**

The `sha256` digest of the image layer.

Type: String

Pattern: `[a-zA-Z0-9-_+.]+:[a-fA-F0-9]+`

Required: No

**layerSize**

The size, in bytes, of the image layer.

Type: Long

Required: No

**mediaType**

The media type of the layer, such as
`application/vnd.docker.image.rootfs.diff.tar.gzip` or
`application/vnd.oci.image.layer.v1.tar+gzip`.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ecr-2015-09-21/Layer)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ecr-2015-09-21/Layer)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ecr-2015-09-21/Layer)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImageTagMutabilityExclusionFilter

LayerFailure

All content copied from https://docs.aws.amazon.com/.
