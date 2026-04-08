# Criterion

The criteria to use in the filter that defines the archive rule. For more information on
available filter keys, see [IAM Access Analyzer filter\
keys](../../../../services/iam/latest/userguide/access-analyzer-reference-filter-keys.md).

## Contents

**contains**

A "contains" operator to match for the filter used to create the rule.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 20 items.

Required: No

**eq**

An "equals" operator to match for the filter used to create the rule.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 20 items.

Required: No

**exists**

An "exists" operator to match for the filter used to create the rule.

Type: Boolean

Required: No

**neq**

A "not equals" operator to match for the filter used to create the rule.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 20 items.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/criterion.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/criterion.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/criterion.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Configuration

DynamodbStreamConfiguration
