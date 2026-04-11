# LifecycleRuleFilter

The container for the filter of the lifecycle rule.

## Contents

**And**

The container for the `AND` condition for the lifecycle rule.

Type: [LifecycleRuleAndOperator](api-control-lifecycleruleandoperator.md) data type

Required: No

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

###### Important

When you're using XML requests, you must
replace special characters (such as carriage returns) in object keys with their equivalent XML entity codes.
For more information, see [XML-related object key constraints](../userguide/object-keys.md#object-key-xml-related-constraints) in the _Amazon S3 User Guide_.

Type: String

Required: No

**Tag**

A container for a key-value name pair.

Type: [S3Tag](api-control-s3tag.md) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/lifecyclerulefilter.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/lifecyclerulefilter.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/lifecyclerulefilter.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LifecycleRuleAndOperator

ListAccessGrantEntry

All content copied from https://docs.aws.amazon.com/.
