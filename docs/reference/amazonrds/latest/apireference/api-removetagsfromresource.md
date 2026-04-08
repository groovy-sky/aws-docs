# RemoveTagsFromResource

Removes metadata tags from an Amazon RDS resource.

For an overview on tagging an Amazon RDS resource,
see [Tagging Amazon RDS Resources](../../../../services/amazonrds/latest/userguide/user-tagging.md) in the _Amazon RDS User Guide_
or [Tagging Amazon Aurora and Amazon RDS Resources](../../../../services/amazonrds/latest/aurorauserguide/user-tagging.md) in the _Amazon Aurora User Guide_.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**ResourceName**

The Amazon RDS resource that the tags are removed from. This value is an Amazon Resource Name (ARN). For information about
creating an ARN,
see [Constructing an ARN for Amazon RDS](../../../../services/amazonrds/latest/userguide/user-tagging-arn-user-tagging-arn-constructing.md) in the _Amazon RDS User Guide._

Type: String

Required: Yes

**TagKeys.member.N**

The tag key (name) of the tag to be removed.

Type: Array of strings

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

This example illustrates one usage of RemoveTagsFromResource.

#### Sample Request

```

https://rds.us-west-2.amazonaws.com/
    ?Action=RemoveTagsFromResource
    &ResourceName=arn%3Aaws%3Ards%3Aus-west-2%3A123456789012%3Adb%3Asample
    &SignatureMethod=HmacSHA256
    &SignatureVersion=4
    &TagKeys.member.1=InstanceType
    &TagKeys.member.2=Owner
    &Version=2014-10-31
    &X-Amz-Algorithm=AWS4-HMAC-SHA256
    &X-Amz-Credential=AKIADQKE4SARGYLE/20160913/us-west-2/rds/aws4_request
    &X-Amz-Date=20160913T174918Z
    &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
    &X-Amz-Signature=4c72f307a75444461bd9b9ccb7de361fec75b8adad66a52824226320d0a33ca8
```

#### Sample Response

```

<RemoveTagsFromResourceResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <ResponseMetadata>
    <RequestId>126d40cc-79da-11e6-b8e4-29f0c684be5d</RequestId>
  </ResponseMetadata>
</RemoveTagsFromResourceResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/removetagsfromresource.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/removetagsfromresource.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/removetagsfromresource.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/removetagsfromresource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/removetagsfromresource.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/removetagsfromresource.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/removetagsfromresource.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/removetagsfromresource.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/removetagsfromresource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/removetagsfromresource.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

RemoveSourceIdentifierFromSubscription

ResetDBClusterParameterGroup
