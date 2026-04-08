# StopDBCluster

Stops an Amazon Aurora DB cluster. When you stop a DB cluster, Aurora retains the DB cluster's
metadata, including its endpoints and DB parameter groups. Aurora also
retains the transaction logs so you can do a point-in-time restore if necessary.

For more information, see
[Stopping and Starting an Aurora Cluster](../../../../services/amazonrds/latest/aurorauserguide/aurora-cluster-stop-start.md) in the _Amazon Aurora User Guide_.

###### Note

This operation only applies to Aurora DB clusters.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBClusterIdentifier**

The DB cluster identifier of the Amazon Aurora DB cluster to be stopped. This parameter is stored as
a lowercase string.

Type: String

Required: Yes

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

**InvalidDBShardGroupState**

The DB shard group must be in the available state.

HTTP Status Code: 400

## Examples

### Example

This example illustrates one usage of StopDBCluster.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
    ?Action=StopDBCluster
    &DBClusterIdentifier=mydbcluster
    &SignatureMethod=HmacSHA256
    &SignatureVersion=4
    &Version=2014-10-31
    &X-Amz-Algorithm=AWS4-HMAC-SHA256
    &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20131016/us-west-1/rds/aws4_request
    &X-Amz-Date=20131016T233051Z
    &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
    &X-Amz-Signature=087a8eb41cb1ab5f99e81575f23e73757ffc6a1e42d7d2b30b9cc0be988cff97
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/stopdbcluster.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/stopdbcluster.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/stopdbcluster.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/stopdbcluster.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/stopdbcluster.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/stopdbcluster.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/stopdbcluster.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/stopdbcluster.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/stopdbcluster.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/stopdbcluster.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StopActivityStream

StopDBInstance

All content copied from https://docs.aws.amazon.com/.
