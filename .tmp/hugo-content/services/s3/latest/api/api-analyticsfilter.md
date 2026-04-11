# AnalyticsFilter

The filter used to describe a set of objects for analyses. A filter must have exactly one prefix,
one tag, or one conjunction (AnalyticsAndOperator). If no filter is provided, all objects will be
considered in any analysis.

## Contents

**And**

A conjunction (logical AND) of predicates, which is used in evaluating an analytics filter. The
operator must have at least two predicates.

Type: [AnalyticsAndOperator](api-analyticsandoperator.md) data type

Required: No

**Prefix**

The prefix to use when evaluating an analytics filter.

Type: String

Required: No

**Tag**

The tag to use when evaluating an analytics filter.

Type: [Tag](api-tag.md) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/analyticsfilter.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/analyticsfilter.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/analyticsfilter.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AnalyticsExportDestination

AnalyticsS3BucketDestination

All content copied from https://docs.aws.amazon.com/.
