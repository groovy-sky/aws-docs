# AddTagsToResource

Adds metadata tags to an Amazon RDS resource. These tags can also be used with cost allocation reporting to track cost associated with Amazon RDS resources, or used in a Condition statement in an IAM policy for Amazon RDS.

For an overview on tagging your relational database resources,
see [Tagging Amazon RDS Resources](../../../../services/amazonrds/latest/userguide/user-tagging.md)
or [Tagging Amazon Aurora and Amazon RDS Resources](../../../../services/amazonrds/latest/aurorauserguide/user-tagging.md).

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**ResourceName**

The Amazon RDS resource that the tags are added to. This value is an Amazon Resource Name (ARN). For information about
creating an ARN,
see [Constructing an RDS Amazon Resource Name (ARN)](../../../../services/amazonrds/latest/userguide/user-tagging-arn.md#USER_Tagging.ARN.Constructing).

Type: String

Required: Yes

**Tags.Tag.N**

The tags to be assigned to the Amazon RDS resource.

Type: Array of [Tag](api-tag.md) objects

Required: Yes

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**BlueGreenDeploymentNotFoundFault**

`BlueGreenDeploymentIdentifier` doesn't refer to an existing blue/green deployment.

HTTP Status Code: 404

**DBClusterNotFoundFault**

`DBClusterIdentifier` doesn't refer to an existing DB cluster.

HTTP Status Code: 404

**DBInstanceNotFound**

`DBInstanceIdentifier` doesn't refer to an existing DB instance.

HTTP Status Code: 404

**DBProxyEndpointNotFoundFault**

The DB proxy endpoint doesn't exist.

HTTP Status Code: 404

**DBProxyNotFoundFault**

The specified proxy name doesn't correspond to a proxy owned by your AWS account in the specified AWS Region.

HTTP Status Code: 404

**DBProxyTargetGroupNotFoundFault**

The specified target group isn't available for a proxy owned by your AWS account in the specified AWS Region.

HTTP Status Code: 404

**DBShardGroupNotFound**

The specified DB shard group name wasn't found.

HTTP Status Code: 404

**DBSnapshotNotFound**

`DBSnapshotIdentifier` doesn't refer to an existing DB snapshot.

HTTP Status Code: 404

**DBSnapshotTenantDatabaseNotFoundFault**

The specified snapshot tenant database wasn't found.

HTTP Status Code: 404

**IntegrationNotFoundFault**

The specified integration could not be found.

HTTP Status Code: 404

**InvalidDBClusterEndpointStateFault**

The requested operation can't be performed on the endpoint while the endpoint is in this state.

HTTP Status Code: 400

**InvalidDBClusterStateFault**

The requested operation can't be performed while the cluster is in this state.

HTTP Status Code: 400

**InvalidDBInstanceState**

The DB instance isn't in a valid state.

HTTP Status Code: 400

**TenantDatabaseNotFound**

The specified tenant database wasn't found in the DB instance.

HTTP Status Code: 404

## Examples

### Example

This example illustrates one usage of AddTagsToResource.

#### Sample Request

```

https://rds.us-west-2.amazonaws.com/
    ?Action=AddTagsToResource
    &ResourceName=arn%3Aaws%3Ards%3Aus-west-2%3A123456789012%3Adb%3Asample
    &SignatureMethod=HmacSHA256
    &SignatureVersion=4
    &Tags.member.1.Key=InstanceType
    &Tags.member.1.Value=Development
    &Tags.member.2.Key=Owner
    &Tags.member.2.Value=Admin123
    &Version=2014-10-31
    &X-Amz-Algorithm=AWS4-HMAC-SHA256
    &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20160913/us-west-2/rds/aws4_request
    &X-Amz-Date=20160913T173915Z
    &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
    &X-Amz-Signature=90a257aa949fab364b7db0964a255986922f933f2e55e7b582ce6f9ccca2a4e0
```

#### Sample Response

```

<AddTagsToResourceResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <ResponseMetadata>
    <RequestId>fd9cd844-79d8-11e6-956c-915ad715fa2f</RequestId>
  </ResponseMetadata>
</AddTagsToResourceResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/addtagstoresource.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/addtagstoresource.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/addtagstoresource.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/addtagstoresource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/addtagstoresource.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/addtagstoresource.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/addtagstoresource.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/addtagstoresource.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/addtagstoresource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/addtagstoresource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AddSourceIdentifierToSubscription

ApplyPendingMaintenanceAction

All content copied from https://docs.aws.amazon.com/.
