# CostCategory

The structure of Cost Categories. This includes detailed metadata and the set of rules
for the `CostCategory` object.

## Contents

**CostCategoryArn**

The unique identifier for your cost category.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:[-a-zA-Z0-9/:_]+`

Required: Yes

**EffectiveStart**

The effective start date of your cost category.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 25.

Pattern: `^\d{4}-\d\d-\d\dT\d\d:\d\d:\d\d(([+-]\d\d:\d\d)|Z)$`

Required: Yes

**Name**

The unique name of the cost category.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 50.

Pattern: `^(?! )[\p{L}\p{N}\p{Z}-_]*(?<! )$`

Required: Yes

**Rules**

The rules are processed in order. If there are multiple rules that match the line
item, then the first rule to match is used to determine that cost category value.

Type: Array of [CostCategoryRule](api-costcategoryrule.md) objects

Array Members: Minimum number of 1 item. Maximum number of 500 items.

Required: Yes

**RuleVersion**

The rule schema version in this particular cost category.

Type: String

Valid Values: `CostCategoryExpression.v1`

Required: Yes

**DefaultValue**

The
default value for the cost category.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 50.

Pattern: `^(?! )[\p{L}\p{N}\p{Z}-_]*(?<! )$`

Required: No

**EffectiveEnd**

The effective end date of your cost category.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 25.

Pattern: `^\d{4}-\d\d-\d\dT\d\d:\d\d:\d\d(([+-]\d\d:\d\d)|Z)$`

Required: No

**ProcessingStatus**

The list of processing statuses for Cost Management products for a specific cost
category.

Type: Array of [CostCategoryProcessingStatus](api-costcategoryprocessingstatus.md) objects

Required: No

**SplitChargeRules**

The split charge rules that are used to allocate your charges between your cost category values.

Type: Array of [CostCategorySplitChargeRule](api-costcategorysplitchargerule.md) objects

Array Members: Minimum number of 1 item. Maximum number of 10 items.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ce-2017-10-25/costcategory.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ce-2017-10-25/costcategory.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ce-2017-10-25/costcategory.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CostAndUsageComparison

CostCategoryInheritedValueDimension
