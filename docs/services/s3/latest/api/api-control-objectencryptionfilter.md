---
title: "ObjectEncryptionFilter"
---

# ObjectEncryptionFilter

An optional filter for the S3JobManifestGenerator that identifies the subset of objects by encryption type.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**DSSEKMS**

Filters for objects that are encrypted by dual-layer server-side encryption with AWS Key Management Service (KMS) keys (DSSE-KMS).

Type: [DSSEKMSFilter](api-control-dssekmsfilter.md) data type

Required: No

**NOTSSE**

Filters for objects that are not encrypted by server-side encryption.

Type: [NotSSEFilter](api-control-notssefilter.md) data type

Required: No

**SSEC**

Filters for objects that are encrypted by server-side encryption with customer-provided keys (SSE-C).

Type: [SSECFilter](api-control-ssecfilter.md) data type

Required: No

**SSEKMS**

Filters for objects that are encrypted by server-side encryption with AWS Key Management Service (KMS) keys (SSE-KMS).

Type: [SSEKMSFilter](api-control-ssekmsfilter.md) data type

Required: No

**SSES3**

Filters for objects that are encrypted by server-side encryption with Amazon S3 managed keys (SSE-S3).

Type: [SSES3Filter](api-control-sses3filter.md) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3control-2018-08-20/ObjectEncryptionFilter)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3control-2018-08-20/ObjectEncryptionFilter)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3control-2018-08-20/ObjectEncryptionFilter)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ObjectEncryption

ObjectLambdaAccessPoint

All content copied from https://docs.aws.amazon.com/.
