# ProcessorFeature

Contains the processor features of a DB instance class.

To specify the number of CPU cores, use the `coreCount` feature name
for the `Name` parameter. To specify the number of threads per core, use the
`threadsPerCore` feature name for the `Name` parameter.

You can set the processor features of the DB instance class for a DB instance when you
call one of the following actions:

- `CreateDBInstance`

- `ModifyDBInstance`

- `RestoreDBInstanceFromDBSnapshot`

- `RestoreDBInstanceFromS3`

- `RestoreDBInstanceToPointInTime`

You can view the valid processor values for a particular instance class by calling the
`DescribeOrderableDBInstanceOptions` action and specifying the
instance class for the `DBInstanceClass` parameter.

In addition, you can use the following actions for DB instance class processor information:

- `DescribeDBInstances`

- `DescribeDBSnapshots`

- `DescribeValidDBInstanceModifications`

If you call `DescribeDBInstances`, `ProcessorFeature` returns
non-null values only if the following conditions are met:

- You are accessing an Oracle or SQL Server DB instance.

- Your Oracle or SQL Server DB instance class supports configuring the number of CPU cores and threads per core.

- The current number CPU cores and threads is set to a non-default value.

For more information, see [Configuring the processor for a DB instance class in RDS for Oracle](../../../../services/amazonrds/latest/userguide/concepts-dbinstanceclass.md#USER_ConfigureProcessor), [Optimizing your RDS for SQL Server CPU](../../../../services/amazonrds/latest/userguide/sqlserver-concepts-general-optimizecpu.md), and [DB instance classes](../../../../services/amazonrds/latest/userguide/concepts-dbinstanceclass.md)
in the _Amazon RDS User Guide._

## Contents

###### Note

In the following list, the required parameters are described first.

**Name**

The name of the processor feature. Valid names are `coreCount` and `threadsPerCore`.

Type: String

Required: No

**Value**

The value of a processor feature.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/processorfeature.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/processorfeature.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/processorfeature.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

PerformanceIssueDetails

Range
