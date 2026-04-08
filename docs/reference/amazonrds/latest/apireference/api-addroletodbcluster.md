# AddRoleToDBCluster

Associates an AWS Identity and Access Management (IAM) role with a DB cluster.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBClusterIdentifier**

The name of the DB cluster to associate the IAM role with.

Type: String

Required: Yes

**RoleArn**

The Amazon Resource Name (ARN) of the IAM role to associate with the Aurora DB
cluster, for example `arn:aws:iam::123456789012:role/AuroraAccessRole`.

Type: String

Required: Yes

**FeatureName**

The name of the feature for the DB cluster that the IAM role is to be associated with.
For information about supported feature names, see [DBEngineVersion](api-dbengineversion.md).

Type: String

Required: No

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBClusterNotFoundFault**

`DBClusterIdentifier` doesn't refer to an existing DB cluster.

HTTP Status Code: 404

**DBClusterRoleAlreadyExists**

The specified IAM role Amazon Resource Name (ARN) is already associated with the specified DB cluster.

HTTP Status Code: 400

**DBClusterRoleQuotaExceeded**

You have exceeded the maximum number of IAM roles that can be associated with the specified DB cluster.

HTTP Status Code: 400

**InvalidDBClusterStateFault**

The requested operation can't be performed while the cluster is in this state.

HTTP Status Code: 400

## Examples

### Example

This example illustrates one usage of AddRoleToDBCluster.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
    ?Action=AddRoleToDBCluster
    &DBClusterIdentifier=sample-cluster
    &RoleArn=arn%3Aaws%3Aiam%3A%3A123456789012%3Arole%2Fsample-role
    &SignatureMethod=HmacSHA256
    &SignatureVersion=4
    &Version=2014-10-31
    &X-Amz-Algorithm=AWS4-HMAC-SHA256
    &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20161012/us-east-1/rds/aws4_request
    &X-Amz-Date=20161012T204524Z
    &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
    &X-Amz-Signature=d73c069210f98e5377851fa4c4ab2fdd53e8bd5d5f02f4f8ef15d4daa5b04567
```

#### Sample Response

```

<AddRoleToDBClusterResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <ResponseMetadata>
    <RequestId>ccccbdb6-90bc-11e6-8533-cd6447e421f8</RequestId>
  </ResponseMetadata>
</AddRoleToDBClusterResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/addroletodbcluster.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/addroletodbcluster.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/addroletodbcluster.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/addroletodbcluster.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/addroletodbcluster.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/addroletodbcluster.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/addroletodbcluster.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/addroletodbcluster.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/addroletodbcluster.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/addroletodbcluster.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Actions

AddRoleToDBInstance

All content copied from https://docs.aws.amazon.com/.
