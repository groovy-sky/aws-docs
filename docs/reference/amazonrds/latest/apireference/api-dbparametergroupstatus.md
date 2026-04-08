# DBParameterGroupStatus

The status of the DB parameter group.

This data type is used as a response element in the following actions:

- `CreateDBInstance`

- `CreateDBInstanceReadReplica`

- `DeleteDBInstance`

- `ModifyDBInstance`

- `RebootDBInstance`

- `RestoreDBInstanceFromDBSnapshot`

## Contents

###### Note

In the following list, the required parameters are described first.

**DBParameterGroupName**

The name of the DB parameter group.

Type: String

Required: No

**ParameterApplyStatus**

The status of parameter updates. Valid values are:

- `applying`: The parameter group change is being applied to the
database.

- `failed-to-apply`: The parameter group is in an invalid
state.

- `in-sync`: The parameter group change is synchronized with the
database.

- `pending-database-upgrade`: The parameter group change will be
applied after the DB instance is upgraded.

- `pending-reboot`: The parameter group change will be applied after
the DB instance reboots.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/dbparametergroupstatus.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/dbparametergroupstatus.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/dbparametergroupstatus.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DBParameterGroup

DBProxy

All content copied from https://docs.aws.amazon.com/.
