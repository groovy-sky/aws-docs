# DimensionCondition

Specifies a condition for filtering capacity data based on dimension values. Used to create precise filters for metric queries and dimension lookups.

## Contents

**Comparison**

The comparison operator to use for the filter.

Type: String

Valid Values: `equals | in`

Required: No

**Dimension**

The name of the dimension to filter by.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 300.

Valid Values: `resource-region | availability-zone-id | account-id | instance-family | instance-type | instance-platform | reservation-arn | reservation-id | reservation-type | reservation-create-timestamp | reservation-start-timestamp | reservation-end-timestamp | reservation-end-date-type | tenancy | reservation-state | reservation-instance-match-criteria | reservation-unused-financial-owner`

Required: No

**Value.N**

The list of values to match against the specified dimension. For 'equals' comparison, only the first value is used. For 'in' comparison, any matching value will satisfy the condition.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 10 items.

Length Constraints: Minimum length of 0. Maximum length of 300.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/dimensioncondition.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/dimensioncondition.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/dimensioncondition.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DhcpOptions

DirectoryServiceAuthentication
