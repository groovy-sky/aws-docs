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

Type: Array of [CostCategoryRule](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_CostCategoryRule.html) objects

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

Type: Array of [CostCategoryProcessingStatus](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_CostCategoryProcessingStatus.html) objects

Required: No

**SplitChargeRules**

The split charge rules that are used to allocate your charges between your cost category values.

Type: Array of [CostCategorySplitChargeRule](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_CostCategorySplitChargeRule.html) objects

Array Members: Minimum number of 1 item. Maximum number of 10 items.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ce-2017-10-25/CostCategory)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ce-2017-10-25/CostCategory)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ce-2017-10-25/CostCategory)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CostAndUsageComparison

CostCategoryInheritedValueDimension
