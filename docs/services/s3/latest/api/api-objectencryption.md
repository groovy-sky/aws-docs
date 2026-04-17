---
title: "ObjectEncryption"
---

# ObjectEncryption

The updated server-side encryption type for this object. The `UpdateObjectEncryption`
operation supports the SSE-S3 and SSE-KMS encryption types.

Valid Values: `SSES3` \| `SSEKMS`

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**SSEKMS**

Specifies to update the object encryption type to server-side encryption with AWS Key Management Service (AWS KMS) keys
(SSE-KMS).

Type: [SSEKMSEncryption](api-ssekmsencryption.md) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3-2006-03-01/ObjectEncryption)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3-2006-03-01/ObjectEncryption)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3-2006-03-01/ObjectEncryption)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Object

ObjectIdentifier

All content copied from https://docs.aws.amazon.com/.
