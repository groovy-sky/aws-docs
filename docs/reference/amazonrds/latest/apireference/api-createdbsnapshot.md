# CreateDBSnapshot

Creates a snapshot of a DB instance. The source DB instance must be in the `available` or
`storage-optimization` state.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBInstanceIdentifier**

The identifier of the DB instance that you want to create the snapshot of.

Constraints:

- Must match the identifier of an existing DBInstance.

Type: String

Required: Yes

**DBSnapshotIdentifier**

The identifier for the DB snapshot.

Constraints:

- Can't be null, empty, or blank

- Must contain from 1 to 255 letters, numbers, or hyphens

- First character must be a letter

- Can't end with a hyphen or contain two consecutive hyphens

Example: `my-snapshot-id`

Type: String

Required: Yes

**Tags.Tag.N**

A list of tags.

For more information, see
[Tagging Amazon RDS resources](../../../../services/amazonrds/latest/userguide/user-tagging.md) in the _Amazon RDS User Guide_ or
[Tagging Amazon Aurora and Amazon RDS resources](../../../../services/amazonrds/latest/aurorauserguide/user-tagging.md) in the _Amazon Aurora User Guide_.

Type: Array of [Tag](api-tag.md) objects

Required: No

## Response Elements

The following element is returned by the service.

**DBSnapshot**

Contains the details of an Amazon RDS DB snapshot.

This data type is used as a response element
in the `DescribeDBSnapshots` action.

Type: [DBSnapshot](api-dbsnapshot.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBInstanceNotFound**

`DBInstanceIdentifier` doesn't refer to an existing DB instance.

HTTP Status Code: 404

**DBSnapshotAlreadyExists**

`DBSnapshotIdentifier` is already used by an existing snapshot.

HTTP Status Code: 400

**InvalidDBInstanceState**

The DB instance isn't in a valid state.

HTTP Status Code: 400

**SnapshotQuotaExceeded**

The request would result in the user exceeding the allowed number of DB
snapshots.

HTTP Status Code: 400

## Examples

### Example

This example illustrates one usage of CreateDBSnapshot.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
   ?Action=CreateDBSnapshot
   &DBInstanceIdentifier=mysqldb-02
   &DBSnapshotIdentifier=mySQLdb-snap-1
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AKIADQKE4SARGYLE/20140423/us-east-1/rds/aws4_request
   &X-Amz-Date=20140423T161105Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=e9649af6edcfbab4016f04d72e1b7fc16d8734c37477afcf25b3def625484ed2

```

#### Sample Response

```

<CreateDBSnapshotResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <CreateDBSnapshotResult>
    <DBSnapshot>
      <Port>3306</Port>
      <OptionGroupName>default:mysql-5-6</OptionGroupName>
      <Engine>mysql</Engine>
      <Status>creating</Status>
      <SnapshotType>manual</SnapshotType>
      <LicenseModel>general-public-license</LicenseModel>
      <EngineVersion>5.6.13</EngineVersion>
      <DBInstanceIdentifier>mysqldb-02</DBInstanceIdentifier>
      <DBSnapshotIdentifier>mysqldb-snap-1</DBSnapshotIdentifier>
      <AvailabilityZone>us-east-1a</AvailabilityZone>
      <InstanceCreateTime>2014-04-21T22:24:26.573Z</InstanceCreateTime>
      <PercentProgress>0</PercentProgress>
      <AllocatedStorage>100</AllocatedStorage>
      <MasterUsername>myawsuser</MasterUsername>
    </DBSnapshot>
  </CreateDBSnapshotResult>
  <ResponseMetadata>
    <RequestId>bd80a25a-af09-11c3-ed11-23c32f9aa7d3</RequestId>
  </ResponseMetadata>
</CreateDBSnapshotResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/createdbsnapshot.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/createdbsnapshot.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/createdbsnapshot.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/createdbsnapshot.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/createdbsnapshot.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/createdbsnapshot.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/createdbsnapshot.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/createdbsnapshot.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/createdbsnapshot.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/createdbsnapshot.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateDBShardGroup

CreateDBSubnetGroup
