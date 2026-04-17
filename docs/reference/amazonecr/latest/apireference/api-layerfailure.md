---
title: "LayerFailure"
---

# LayerFailure

An object representing an Amazon ECR image layer failure.

## Contents

**failureCode**

The failure code associated with the failure.

Type: String

Valid Values: `InvalidLayerDigest | MissingLayerDigest`

Required: No

**failureReason**

The reason for the failure.

Type: String

Required: No

**layerDigest**

The layer digest associated with the failure.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1000.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ecr-2015-09-21/LayerFailure)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ecr-2015-09-21/LayerFailure)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ecr-2015-09-21/LayerFailure)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Layer

LifecyclePolicyPreviewFilter

All content copied from https://docs.aws.amazon.com/.
