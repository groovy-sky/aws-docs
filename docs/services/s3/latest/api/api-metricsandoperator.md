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

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/metricsandoperator.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/metricsandoperator.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/metricsandoperator.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Metrics

MetricsConfiguration

All content copied from https://docs.aws.amazon.com/.
