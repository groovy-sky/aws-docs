# LifecycleRuleAndOperator

This is used in a Lifecycle Rule Filter to apply a logical AND to two or more predicates. The
Lifecycle Rule will apply to any object matching all of the predicates configured inside the And
operator.

## Contents

**ObjectSizeGreaterThan**

Minimum object size to which the rule applies.

Type: Long

Required: No

**ObjectSizeLessThan**

Maximum object size to which the rule applies.

Type: Long

Required: No

**Prefix**

Prefix identifying one or more objects to which the rule applies.

Type: String

Required: No

**Tags**

All of these tags must exist in the object's tag set in order for the rule to apply.

Type: Array of [Tag](api-tag.md) data types

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/lifecycleruleandoperator.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/lifecycleruleandoperator.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/lifecycleruleandoperator.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LifecycleRule

LifecycleRuleFilter

All content copied from https://docs.aws.amazon.com/.
