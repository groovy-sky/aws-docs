# Transition

Specifies when an object transitions to a specified storage class. For more information about Amazon S3
lifecycle configuration rules, see [Transitioning Objects Using\
Amazon S3 Lifecycle](../dev/lifecycle-transition-general-considerations.md) in the _Amazon S3 User Guide_.

## Contents

**Date**

Indicates when objects are transitioned to the specified storage class. The date value must be in
ISO 8601 format. The time is always midnight UTC.

Type: Timestamp

Required: No

**Days**

Indicates the number of days after creation when objects are transitioned to the specified storage
class. If the specified storage class is `INTELLIGENT_TIERING`, `GLACIER_IR`,
`GLACIER`, or `DEEP_ARCHIVE`, valid values are `0` or positive
integers. If the specified storage class is `STANDARD_IA` or `ONEZONE_IA`, valid
values are positive integers greater than `30`. Be aware that some storage classes have a
minimum storage duration and that you're charged for transitioning objects before their minimum storage
duration. For more information, see [Constraints and considerations for transitions](../userguide/lifecycle-transition-general-considerations.md#lifecycle-configuration-constraints) in the _Amazon S3 User_
_Guide_.

Type: Integer

Required: No

**StorageClass**

The storage class to which you want the object to transition.

Type: String

Valid Values: `GLACIER | STANDARD_IA | ONEZONE_IA | INTELLIGENT_TIERING | DEEP_ARCHIVE | GLACIER_IR`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/transition.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/transition.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/transition.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TopicConfigurationDeprecated

VersioningConfiguration

All content copied from https://docs.aws.amazon.com/.
