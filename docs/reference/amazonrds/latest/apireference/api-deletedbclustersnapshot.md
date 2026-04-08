# DeleteDBClusterSnapshot

Deletes a DB cluster snapshot. If the snapshot is being copied, the copy operation is terminated.

###### Note

The DB cluster snapshot must be in the `available` state to be
deleted.

For more information on Amazon Aurora, see
[What is Amazon Aurora?](../../../../services/amazonrds/latest/aurorauserguide/chap-auroraoverview.md) in the _Amazon Aurora User Guide_.

For more information on Multi-AZ DB clusters, see [Multi-AZ DB\
cluster deployments](../../../../services/amazonrds/latest/userguide/multi-az-db-clusters-concepts.md) in the _Amazon RDS User_
_Guide_.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBClusterSnapshotIdentifier**

The identifier of the DB cluster snapshot to delete.

Constraints: Must be the name of an existing DB cluster snapshot in the `available` state.

Type: String

Required: Yes

## Response Elements

The following element is returned by the service.

**DBClusterSnapshot**

Contains the details for an Amazon RDS DB cluster snapshot

This data type is used as a response element
in the `DescribeDBClusterSnapshots` action.

Type: [DBClusterSnapshot](api-dbclustersnapshot.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBClusterSnapshotNotFoundFault**

`DBClusterSnapshotIdentifier` doesn't refer to an existing DB cluster snapshot.

HTTP Status Code: 404

**InvalidDBClusterSnapshotStateFault**

The supplied value isn't a valid DB cluster snapshot state.

HTTP Status Code: 400

## Examples

### Example

This example illustrates one usage of DeleteDBClusterSnapshot.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
    ?Action=DeleteDBClusterSnapshot
    &DBClusterSnapshotIdentifier=sample-cluster-snapshot
    &SignatureMethod=HmacSHA256
    &SignatureVersion=4
    &Version=2014-10-31
    &X-Amz-Algorithm=AWS4-HMAC-SHA256
    &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20150318/us-east-1/rds/aws4_request
    &X-Amz-Date=20150318T215614Z
    &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
    &X-Amz-Signature=7aaab0a295151051bc4723f5b1f7b6b535615b8db9256bd56993c4dc6df4c2c4

```

#### Sample Response

```

<DeleteDBClusterSnapshotResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <DeleteDBClusterSnapshotResult>
    <DBClusterSnapshot>
      <Port>0</Port>
      <Status>available</Status>
      <Engine>aurora</Engine>
      <SnapshotType>manual</SnapshotType>
      <LicenseModel>aurora</LicenseModel>
      <DBClusterSnapshotIdentifier>sample-cluster-snapshot</DBClusterSnapshotIdentifier>
      <SnapshotCreateTime>2015-03-18T20:53:22.523Z</SnapshotCreateTime>
      <DBClusterIdentifier>sample-cluster</DBClusterIdentifier>
      <VpcId>vpc-3fabee54</VpcId>
      <ClusterCreateTime>2015-03-06T22:11:13.826Z</ClusterCreateTime>
      <PercentProgress>100</PercentProgress>
      <AllocatedStorage>1</AllocatedStorage>
      <MasterUsername>awsuser</MasterUsername>
    </DBClusterSnapshot>
  </DeleteDBClusterSnapshotResult>
  <ResponseMetadata>
    <RequestId>994ab08d-cdb9-2ce4-abf9-7528e6348483</RequestId>
  </ResponseMetadata>
</DeleteDBClusterSnapshotResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/deletedbclustersnapshot.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/deletedbclustersnapshot.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/deletedbclustersnapshot.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/deletedbclustersnapshot.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/deletedbclustersnapshot.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/deletedbclustersnapshot.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/deletedbclustersnapshot.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/deletedbclustersnapshot.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/deletedbclustersnapshot.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/deletedbclustersnapshot.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteDBClusterParameterGroup

DeleteDBInstance

All content copied from https://docs.aws.amazon.com/.
