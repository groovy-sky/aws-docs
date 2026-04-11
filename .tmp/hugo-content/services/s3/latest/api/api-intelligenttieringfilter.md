# IntelligentTieringFilter

The `Filter` is used to identify objects that the S3 Intelligent-Tiering configuration
applies to.

## Contents

**And**

A conjunction (logical AND) of predicates, which is used in evaluating a metrics filter. The
operator must have at least two predicates, and an object must match all of the predicates in order for
the filter to apply.

Type: [IntelligentTieringAndOperator](api-intelligenttieringandoperator.md) data type

Required: No

**Prefix**

An object key name prefix that identifies the subset of objects to which the rule applies.

###### Important

Replacement must be made for object keys containing special characters (such as carriage returns) when using
XML requests. For more information, see [XML related object key constraints](../userguide/object-keys.md#object-key-xml-related-constraints).

Type: String

Required: No

**Tag**

A container of a key value name pair.

Type: [Tag](api-tag.md) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/intelligenttieringfilter.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/intelligenttieringfilter.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/intelligenttieringfilter.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IntelligentTieringConfiguration

InventoryConfiguration

All content copied from https://docs.aws.amazon.com/.
