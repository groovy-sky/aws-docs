# FailoverDBCluster

Forces a failover for a DB cluster.

For an Aurora DB cluster, failover for a DB cluster promotes one of the Aurora Replicas (read-only instances)
in the DB cluster to be the primary DB instance (the cluster writer).

For a Multi-AZ DB cluster, after RDS terminates the primary DB instance, the
internal monitoring system detects that the primary DB instance is unhealthy and promotes a readable standby (read-only instances)
in the DB cluster to be the primary DB instance (the cluster writer).
Failover times are typically less than 35 seconds.

An Amazon Aurora DB cluster automatically fails over to an Aurora Replica, if one exists,
when the primary DB instance fails. A Multi-AZ DB cluster automatically fails over to a readable standby
DB instance when the primary DB instance fails.

To simulate a failure of a primary instance for testing, you can force a failover.
Because each instance in a DB cluster has its own endpoint address, make sure to clean up and re-establish any existing
connections that use those endpoint addresses when the failover is complete.

For more information on Amazon Aurora DB clusters, see
[What is Amazon Aurora?](../../../../services/amazonrds/latest/aurorauserguide/chap-auroraoverview.md) in the _Amazon Aurora User Guide_.

For more information on Multi-AZ DB clusters, see [Multi-AZ DB\
cluster deployments](../../../../services/amazonrds/latest/userguide/multi-az-db-clusters-concepts.md) in the _Amazon RDS User_
_Guide_.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBClusterIdentifier**

The identifier of the DB cluster to force a failover for. This parameter isn't case-sensitive.

Constraints:

- Must match the identifier of an existing DB cluster.

Type: String

Required: Yes

**TargetDBInstanceIdentifier**

The name of the DB instance to promote to the primary DB instance.

Specify the DB instance identifier for an Aurora Replica or a Multi-AZ readable standby in the DB cluster,
for example `mydbcluster-replica1`.

This setting isn't supported for RDS for MySQL Multi-AZ DB clusters.

Type: String

Required: No

## Response Elements

The following element is returned by the service.

**DBCluster**

Contains the details of an Amazon Aurora DB cluster or Multi-AZ DB cluster.

For an Amazon Aurora DB cluster, this data type is used as a response element in the operations
`CreateDBCluster`, `DeleteDBCluster`, `DescribeDBClusters`,
`FailoverDBCluster`, `ModifyDBCluster`, `PromoteReadReplicaDBCluster`,
`RestoreDBClusterFromS3`, `RestoreDBClusterFromSnapshot`,
`RestoreDBClusterToPointInTime`, `StartDBCluster`, and `StopDBCluster`.

For a Multi-AZ DB cluster, this data type is used as a response element in the operations
`CreateDBCluster`, `DeleteDBCluster`, `DescribeDBClusters`,
`FailoverDBCluster`, `ModifyDBCluster`, `RebootDBCluster`,
`RestoreDBClusterFromSnapshot`, and `RestoreDBClusterToPointInTime`.

For more information on Amazon Aurora DB clusters, see
[What is Amazon Aurora?](../../../../services/amazonrds/latest/aurorauserguide/chap-auroraoverview.md) in the _Amazon Aurora User Guide._

For more information on Multi-AZ DB clusters, see
[Multi-AZ deployments with two readable standby DB instances](../../../../services/amazonrds/latest/userguide/multi-az-db-clusters-concepts.md) in the _Amazon RDS User Guide._

Type: [DBCluster](api-dbcluster.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBClusterNotFoundFault**

`DBClusterIdentifier` doesn't refer to an existing DB cluster.

HTTP Status Code: 404

**InvalidDBClusterStateFault**

The requested operation can't be performed while the cluster is in this state.

HTTP Status Code: 400

**InvalidDBInstanceState**

The DB instance isn't in a valid state.

HTTP Status Code: 400

## Examples

### Example

This example illustrates one usage of FailoverDBCluster.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
    ?Action=FailoverDBCluster
    &DBClusterIdentifier=sample-cluster
    &SignatureMethod=HmacSHA256
    &SignatureVersion=4
    &Version=2014-10-31
    &X-Amz-Algorithm=AWS4-HMAC-SHA256
    &X-Amz-Credential=AKIADQKE4SARGYLE/20150323/us-east-1/rds/aws4_request
    &X-Amz-Date=20150323T170232Z
    &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
    &X-Amz-Signature=9be705fa28a68244d5072722463a29a322f9ef8eb58a63c40a6f6547174dec44

```

#### Sample Response

```

<FailoverDBClusterResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <FailoverDBClusterResult>
    <DBCluster>
      <Port>3306</Port>
      <LatestRestorableTime>2015-03-23T17:00:54.893Z</LatestRestorableTime>
      <Engine>aurora</Engine>
      <Status>available</Status>
      <BackupRetentionPeriod>7</BackupRetentionPeriod>
      <VpcSecurityGroups>
        <VpcSecurityGroupMembership>
          <Status>active</Status>
          <VpcSecurityGroupId>sg-922dc2fd</VpcSecurityGroupId>
        </VpcSecurityGroupMembership>
      </VpcSecurityGroups>
      <DBSubnetGroup>sample-group</DBSubnetGroup>
      <EngineVersion>5.6.10a</EngineVersion>
      <Endpoint>sample-cluster.cluster-c1axbpgwvdfo.us-east-1.rds.amazonaws.com</Endpoint>
      <DBClusterParameterGroup>default.aurora5.6</DBClusterParameterGroup>
      <DBClusterIdentifier>sample-cluster</DBClusterIdentifier>
      <PreferredBackupWindow>05:47-06:17</PreferredBackupWindow>
      <PreferredMaintenanceWindow>mon:10:16-mon:10:46</PreferredMaintenanceWindow>
      <EarliestRestorableTime>2015-03-04T23:08:59.159Z</EarliestRestorableTime>
      <DBClusterMembers>
        <DBClusterMember>
          <IsClusterWriter>false</IsClusterWriter>
          <DBInstanceIdentifier>sample-replica1</DBInstanceIdentifier>
          <DBClusterParameterGroupStatus>in-sync</DBClusterParameterGroupStatus>
        </DBClusterMember>
        <DBClusterMember>
          <IsClusterWriter>true</IsClusterWriter>
          <DBInstanceIdentifier>sample-primary</DBInstanceIdentifier>
          <DBClusterParameterGroupStatus>in-sync</DBClusterParameterGroupStatus>
        </DBClusterMember>
      </DBClusterMembers>
      <AllocatedStorage>1</AllocatedStorage>
      <MasterUsername>awsuser</MasterUsername>
    </DBCluster>
  </FailoverDBClusterResult>
  <ResponseMetadata>
    <RequestId>659c3dba-d17e-11e4-9fd0-35e9d88e2515</RequestId>
  </ResponseMetadata>
</FailoverDBClusterResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/failoverdbcluster.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/failoverdbcluster.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/failoverdbcluster.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/failoverdbcluster.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/failoverdbcluster.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/failoverdbcluster.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/failoverdbcluster.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/failoverdbcluster.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/failoverdbcluster.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/failoverdbcluster.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

EnableHttpEndpoint

FailoverGlobalCluster
