# MetricsAndOperator

A conjunction (logical AND) of predicates, which is used in evaluating a metrics filter. The
operator must have at least two predicates, and an object must match all of the predicates in order for
the filter to apply.

## Contents

**AccessPointArn**

The access point ARN used when evaluating an `AND` predicate.

Type: String

Required: No

**Prefix**

The prefix used when evaluating an AND predicate.

Type: String

Required: No

**Tags**

The list of tags used when evaluating an AND predicate.

###### Note

`Tag` filters are not supported for directory buckets.

Type: Array of [Tag](api-tag.md) data types

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3-2006-03-01/MetricsAndOperator)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3-2006-03-01/MetricsAndOperator)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3-2006-03-01/MetricsAndOperator)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Metrics

MetricsConfiguration
