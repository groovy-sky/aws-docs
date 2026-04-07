# Expression

See [Expression](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_billing_Expression.html). Billing view only supports `LINKED_ACCOUNT`, `Tags`, and `CostCategories`.

## Contents

**costCategories**

The filter that's based on `CostCategory` values.

Type: [CostCategoryValues](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_billing_CostCategoryValues.html) object

Required: No

**dimensions**

The specific `Dimension` to use for `Expression`.

Type: [DimensionValues](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_billing_DimensionValues.html) object

Required: No

**tags**

The specific `Tag` to use for `Expression`.

Type: [TagValues](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_billing_TagValues.html) object

Required: No

**timeRange**

Specifies a time range filter for the billing view data.

Type: [TimeRange](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_billing_TimeRange.html) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/billing-2023-09-07/Expression)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/billing-2023-09-07/Expression)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/billing-2023-09-07/Expression)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DimensionValues

ResourceTag
