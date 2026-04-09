# Transition

Specifies when an object transitions to a specified storage class. For more information
about Amazon S3 Lifecycle configuration rules, see [Transitioning objects using Amazon S3 Lifecycle](../dev/lifecycle-transition-general-considerations.md) in the
_Amazon S3 User Guide_.

## Contents

**Date**

Indicates when objects are transitioned to the specified storage class. The date value
must be in ISO 8601 format. The time is always midnight UTC.

Type: Timestamp

Required: No

**Days**

Indicates the number of days after creation when objects are transitioned to the
specified storage class. The value must be a positive integer.

Type: Integer

Required: No

**StorageClass**

The storage class to which you want the object to transition.

Type: String

Valid Values: `GLACIER | STANDARD_IA | ONEZONE_IA | INTELLIGENT_TIERING | DEEP_ARCHIVE`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/transition.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/transition.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/transition.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tagging

VersioningConfiguration

All content copied from https://docs.aws.amazon.com/.
