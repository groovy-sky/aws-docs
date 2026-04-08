# CreateBlueGreenDeployment

Creates a blue/green deployment.

A blue/green deployment creates a staging environment that copies the production environment.
In a blue/green deployment, the blue environment is the current production environment.
The green environment is the staging environment, and it stays in sync
with the current production environment.

You can make changes to the databases in the green environment without affecting
production workloads. For example, you can upgrade the major or minor DB engine version, change
database parameters, or make schema changes in the staging environment. You can thoroughly test
changes in the green environment. When ready, you can switch over the environments to promote the
green environment to be the new production environment. The switchover typically takes under a minute.

For more information, see [Using Amazon RDS Blue/Green Deployments \
for database updates](../../../../services/amazonrds/latest/userguide/blue-green-deployments.md) in the _Amazon RDS User Guide_ and
[Using Amazon RDS Blue/Green Deployments for database updates](../../../../services/amazonrds/latest/aurorauserguide/blue-green-deployments.md) in the _Amazon Aurora_
_User Guide_.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**BlueGreenDeploymentName**

The name of the blue/green deployment.

Constraints:

- Can't be the same as an existing blue/green deployment name in the same account and AWS Region.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 60.

Pattern: `[a-zA-Z](?:-?[a-zA-Z0-9]+)*`

Required: Yes

**Source**

The Amazon Resource Name (ARN) of the source production database.

Specify the database that you want to clone. The blue/green deployment creates this database in
the green environment. You can make updates to the database in the green environment, such as an engine
version upgrade. When you are ready, you can switch the database in the green environment to be the
production database.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `arn:[A-Za-z][0-9A-Za-z-:._]*`

Required: Yes

**Tags.Tag.N**

Tags to assign to the blue/green deployment.

Type: Array of [Tag](api-tag.md) objects

Required: No

**TargetAllocatedStorage**

The amount of storage in gibibytes (GiB) to allocate for the green DB instance. You can choose to
increase or decrease the allocated storage on the green DB instance.

This setting doesn't apply to Amazon Aurora blue/green deployments.

Type: Integer

Required: No

**TargetDBClusterParameterGroupName**

The DB cluster parameter group associated with the Aurora DB cluster in the green environment.

To test parameter changes, specify a DB cluster parameter group that is different from the one associated
with the source DB cluster.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[A-Za-z](?!.*--)[0-9A-Za-z-]*[^-]|^default(?!.*--)(?!.*\.\.)[0-9A-Za-z-.]*[^-]`

Required: No

**TargetDBInstanceClass**

Specify the DB instance class for the databases in the green environment.

This parameter only applies to RDS DB instances, because DB instances within an Aurora DB cluster can
have multiple different instance classes. If you're creating a blue/green deployment from an Aurora DB cluster,
don't specify this parameter. After the green environment is created, you can individually modify the instance classes
of the DB instances within the green DB cluster.

Type: String

Length Constraints: Minimum length of 5. Maximum length of 20.

Pattern: `db\.[0-9a-z]{2,6}\.[0-9a-z]{4,9}`

Required: No

**TargetDBParameterGroupName**

The DB parameter group associated with the DB instance in the green environment.

To test parameter changes, specify a DB parameter group that is different from the one associated
with the source DB instance.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[A-Za-z](?!.*--)[0-9A-Za-z-]*[^-]|^default(?!.*--)(?!.*\.\.)[0-9A-Za-z-.]*[^-]`

Required: No

**TargetEngineVersion**

The engine version of the database in the green environment.

Specify the engine version to upgrade to in the green environment.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Pattern: `[0-9A-Za-z-_.]+`

Required: No

**TargetIops**

The amount of Provisioned IOPS (input/output operations per second) to allocate for the green DB instance.
For information about valid IOPS values, see
[Amazon RDS DB instance storage](../../../../services/amazonrds/latest/userguide/chap-storage.md)
in the _Amazon RDS User Guide_.

This setting doesn't apply to Amazon Aurora blue/green deployments.

Type: Integer

Required: No

**TargetStorageThroughput**

The storage throughput value for the green DB instance.

This setting applies only to the `gp3` storage type.

This setting doesn't apply to Amazon Aurora blue/green deployments.

Type: Integer

Required: No

**TargetStorageType**

The storage type to associate with the green DB instance.

Valid Values: `gp2 | gp3 | io1 | io2`

This setting doesn't apply to Amazon Aurora blue/green deployments.

Type: String

Required: No

**UpgradeTargetStorageConfig**

Whether to upgrade the storage file system configuration on the green database. This
option migrates the green DB instance from the older 32-bit file system to the preferred
configuration. For more information, see [Upgrading the storage file system for a DB instance](../../../../services/amazonrds/latest/userguide/user-piops-storagetypes.md#USER_PIOPS.UpgradeFileSystem).

Type: Boolean

Required: No

## Response Elements

The following element is returned by the service.

**BlueGreenDeployment**

Details about a blue/green deployment.

For more information, see [Using Amazon RDS\
Blue/Green Deployments for database updates](../../../../services/amazonrds/latest/userguide/blue-green-deployments.md) in the _Amazon RDS User_
_Guide_ and [Using Amazon RDS\
Blue/Green Deployments for database updates](../../../../services/amazonrds/latest/aurorauserguide/blue-green-deployments.md) in the _Amazon Aurora_
_User Guide_.

Type: [BlueGreenDeployment](api-bluegreendeployment.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**BlueGreenDeploymentAlreadyExistsFault**

A blue/green deployment with the specified name already exists.

HTTP Status Code: 400

**DBClusterNotFoundFault**

`DBClusterIdentifier` doesn't refer to an existing DB cluster.

HTTP Status Code: 404

**DBClusterParameterGroupNotFound**

`DBClusterParameterGroupName` doesn't refer to an existing DB
cluster parameter group.

HTTP Status Code: 404

**DBClusterQuotaExceededFault**

The user attempted to create a new DB cluster and the user has already reached the
maximum allowed DB cluster quota.

HTTP Status Code: 403

**DBInstanceNotFound**

`DBInstanceIdentifier` doesn't refer to an existing DB instance.

HTTP Status Code: 404

**DBParameterGroupNotFound**

`DBParameterGroupName` doesn't refer to an
existing DB parameter group.

HTTP Status Code: 404

**InstanceQuotaExceeded**

The request would result in the user exceeding the allowed number of DB
instances.

HTTP Status Code: 400

**InvalidDBClusterStateFault**

The requested operation can't be performed while the cluster is in this state.

HTTP Status Code: 400

**InvalidDBInstanceState**

The DB instance isn't in a valid state.

HTTP Status Code: 400

**SourceClusterNotSupportedFault**

The source DB cluster isn't supported for a blue/green deployment.

HTTP Status Code: 400

**SourceDatabaseNotSupportedFault**

The source DB instance isn't supported for a blue/green deployment.

HTTP Status Code: 400

**StorageQuotaExceeded**

The request would result in the user exceeding the allowed amount of storage
available across all DB instances.

HTTP Status Code: 400

## Examples

### Example

This example illustrates one usage of CreateBlueGreenDeployment.

#### Sample Request

```

https://rds.us-west-2.amazonaws.com/
   ?Action=CreateBlueGreenDeployment
   &BlueGreenDeploymentName=my-blue-green-deployment
   &Source=arn%3Aaws%3Ards%3Aus-west-2%3A123456789012%3Adb%3Adatabase-1
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20141031/us-west-2/rds/aws4_request
   &X-Amz-Date=20230110T005253Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=8a684aebe6d5219bb3572316a341963324d6ef339bd0dcfa5854f1a01d401214
```

#### Sample Response

```

<CreateBlueGreenDeploymentResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <CreateBlueGreenDeploymentResult>
    <BlueGreenDeployment>
      <TagList/>
      <BlueGreenDeploymentName>my-blue-green-deployment</BlueGreenDeploymentName>
      <CreateTime>2023-01-10T18:42:09.330Z</CreateTime>
      <SwitchoverDetails>
        <member>
          <SourceMember>arn:aws:rds:us-west-2:123456789012:db:database-1</SourceMember>
        </member>
      </SwitchoverDetails>
      <Source>arn:aws:rds:us-west-2:123456789012:db:database-1</Source>
      <BlueGreenDeploymentIdentifier>bgd-mdoyy2mn7vbkhhgg</BlueGreenDeploymentIdentifier>
      <Tasks>
        <member>
          <Name>CREATING_READ_REPLICA_OF_SOURCE</Name>
          <Status>PENDING</Status>
        </member>
        <member>
          <Name>CONFIGURE_BACKUPS</Name>
          <Status>PENDING</Status>
        </member>
      </Tasks>
      <Status>PROVISIONING</Status>
    </BlueGreenDeployment>
  </CreateBlueGreenDeploymentResult>
  <ResponseMetadata>
    <RequestId>03b87d54-b780-4055-b44d-4a2a129bc8c2</RequestId>
  </ResponseMetadata>
</CreateBlueGreenDeploymentResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/createbluegreendeployment.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/createbluegreendeployment.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/createbluegreendeployment.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/createbluegreendeployment.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/createbluegreendeployment.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/createbluegreendeployment.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/createbluegreendeployment.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/createbluegreendeployment.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/createbluegreendeployment.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/createbluegreendeployment.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CopyOptionGroup

CreateCustomDBEngineVersion

All content copied from https://docs.aws.amazon.com/.
