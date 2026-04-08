# Expression

See [Expression](api-billing-expression.md). Billing view only supports `LINKED_ACCOUNT`, `Tags`, and `CostCategories`.

## Contents

**costCategories**

The filter that's based on `CostCategory` values.

Type: [CostCategoryValues](api-billing-costcategoryvalues.md) object

Required: No

**dimensions**

The specific `Dimension` to use for `Expression`.

Type: [DimensionValues](api-billing-dimensionvalues.md) object

Required: No

**tags**

The specific `Tag` to use for `Expression`.

Type: [TagValues](api-billing-tagvalues.md) object

Required: No

**timeRange**

Specifies a time range filter for the billing view data.

Type: [TimeRange](api-billing-timerange.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/billing-2023-09-07/expression.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/billing-2023-09-07/expression.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/billing-2023-09-07/expression.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DimensionValues

ResourceTag
