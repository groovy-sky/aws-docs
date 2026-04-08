# SupportedEngineLifecycle

This data type is used as a response element in the operation
`DescribeDBMajorEngineVersions`.

You can use the information that this data type returns to plan for upgrades.

This data type only returns information for the open source engines Amazon RDS for
MariaDB, Amazon RDS for MySQL, Amazon RDS for PostgreSQL, Aurora MySQL, and Aurora
PostgreSQL.

## Contents

###### Note

In the following list, the required parameters are described first.

**LifecycleSupportEndDate**

The end date for the type of support returned by `LifecycleSupportName`.

Type: Timestamp

Required: Yes

**LifecycleSupportName**

The type of lifecycle support that the engine version is in.

This parameter returns the following values:

- `open-source-rds-standard-support` \- Indicates RDS standard support or Aurora standard support.

- `open-source-rds-extended-support` \- Indicates Amazon RDS Extended Support.

For Amazon RDS for MySQL, Amazon RDS for PostgreSQL, Aurora MySQL, and Aurora
PostgreSQL, this parameter returns both `open-source-rds-standard-support`
and `open-source-rds-extended-support`.

For Amazon RDS for MariaDB, this parameter only returns the value
`open-source-rds-standard-support`.

For information about Amazon RDS Extended Support, see [Amazon RDS Extended Support\
with Amazon RDS](../../../../services/amazonrds/latest/userguide/extended-support.md) in the _Amazon RDS User Guide_ and [Amazon RDS Extended Support with Amazon Aurora](../../../../services/amazonrds/latest/aurorauserguide/extended-support.md) in the _Amazon_
_Aurora User Guide_.

Type: String

Valid Values: `open-source-rds-standard-support | open-source-rds-extended-support`

Required: Yes

**LifecycleSupportStartDate**

The start date for the type of support returned by `LifecycleSupportName`.

Type: Timestamp

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/supportedenginelifecycle.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/supportedenginelifecycle.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/supportedenginelifecycle.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Subnet

SwitchoverDetail

All content copied from https://docs.aws.amazon.com/.
