# StopDBInstance

Stops an Amazon RDS DB instance temporarily. When you stop a DB instance, Amazon RDS retains the DB instance's metadata,
including its endpoint, DB parameter group, and option group membership. Amazon RDS also retains
the transaction logs so you can do a point-in-time restore if necessary. The instance restarts automatically
after 7 days.

For more information, see
[Stopping an Amazon RDS DB Instance Temporarily](../../../../services/amazonrds/latest/userguide/user-stopinstance.md) in the
_Amazon RDS User Guide._

###### Note

This command doesn't apply to RDS Custom, Aurora MySQL, and Aurora PostgreSQL.
For Aurora clusters, use `StopDBCluster` instead.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBInstanceIdentifier**

The user-supplied instance identifier.

Type: String

Required: Yes

**DBSnapshotIdentifier**

The user-supplied instance identifier of the DB Snapshot created immediately before the DB instance is stopped.

Type: String

Required: No

## Response Elements

The following element is returned by the service.

**DBInstance**

Contains the details of an Amazon RDS DB instance.

This data type is used as a response element in the operations `CreateDBInstance`,
`CreateDBInstanceReadReplica`, `DeleteDBInstance`, `DescribeDBInstances`,
`ModifyDBInstance`, `PromoteReadReplica`, `RebootDBInstance`,
`RestoreDBInstanceFromDBSnapshot`, `RestoreDBInstanceFromS3`, `RestoreDBInstanceToPointInTime`,
`StartDBInstance`, and `StopDBInstance`.

Type: [DBInstance](api-dbinstance.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBInstanceNotFound**

`DBInstanceIdentifier` doesn't refer to an existing DB instance.

HTTP Status Code: 404

**DBSnapshotAlreadyExists**

`DBSnapshotIdentifier` is already used by an existing snapshot.

HTTP Status Code: 400

**InvalidDBClusterStateFault**

The requested operation can't be performed while the cluster is in this state.

HTTP Status Code: 400

**InvalidDBInstanceState**

The DB instance isn't in a valid state.

HTTP Status Code: 400

**SnapshotQuotaExceeded**

The request would result in the user exceeding the allowed number of DB
snapshots.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/stopdbinstance.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/stopdbinstance.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/stopdbinstance.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/stopdbinstance.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/stopdbinstance.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/stopdbinstance.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/stopdbinstance.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/stopdbinstance.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/stopdbinstance.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/stopdbinstance.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StopDBCluster

StopDBInstanceAutomatedBackupsReplication

All content copied from https://docs.aws.amazon.com/.
