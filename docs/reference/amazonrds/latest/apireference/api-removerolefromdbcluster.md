# RemoveRoleFromDBCluster

Removes the asssociation of an AWS Identity and Access Management (IAM) role from a
DB cluster.

For more information on Amazon Aurora DB clusters, see
[What is Amazon Aurora?](../../../../services/amazonrds/latest/aurorauserguide/chap-auroraoverview.md) in the _Amazon Aurora User Guide_.

For more information on Multi-AZ DB clusters, see [Multi-AZ DB\
cluster deployments](../../../../services/amazonrds/latest/userguide/multi-az-db-clusters-concepts.md) in the _Amazon RDS User_
_Guide._

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBClusterIdentifier**

The name of the DB cluster to disassociate the IAM role from.

Type: String

Required: Yes

**RoleArn**

The Amazon Resource Name (ARN) of the IAM role to disassociate from the Aurora DB cluster, for example
`arn:aws:iam::123456789012:role/AuroraAccessRole`.

Type: String

Required: Yes

**FeatureName**

The name of the feature for the DB cluster that the IAM role is to be disassociated from.
For information about supported feature names, see [DBEngineVersion](api-dbengineversion.md).

Type: String

Required: No

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBClusterNotFoundFault**

`DBClusterIdentifier` doesn't refer to an existing DB cluster.

HTTP Status Code: 404

**DBClusterRoleNotFound**

The specified IAM role Amazon Resource Name (ARN) isn't associated with the specified DB cluster.

HTTP Status Code: 404

**InvalidDBClusterStateFault**

The requested operation can't be performed while the cluster is in this state.

HTTP Status Code: 400

## Examples

### Example

This example illustrates one usage of RemoveRoleFromDBCluster.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
    ?Action=RemoveRoleFromDBCluster
    &DBClusterIdentifier=sample-cluster
    &RoleArn=arn%3Aaws%3Aiam%3A%3A123456789012%3Arole%2Fsample-role
    &SignatureMethod=HmacSHA256
    &SignatureVersion=4
    &Version=2014-10-31
    &X-Amz-Algorithm=AWS4-HMAC-SHA256
    &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20161012/us-east-1/rds/aws4_request
    &X-Amz-Date=20161012T204525Z
    &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
    &X-Amz-Signature=cd7d5005d56a505b4e2a878c297e6f8a3cc26b19a335ede018ba41f3185c92a2
```

#### Sample Response

```

<RemoveRoleFromDBClusterResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <ResponseMetadata>
    <RequestId>ccfca75a-90bc-11e6-8533-cd6377e421f8</RequestId>
  </ResponseMetadata>
</RemoveRoleFromDBClusterResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/removerolefromdbcluster.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/removerolefromdbcluster.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/removerolefromdbcluster.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/removerolefromdbcluster.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/removerolefromdbcluster.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/removerolefromdbcluster.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/removerolefromdbcluster.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/removerolefromdbcluster.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/removerolefromdbcluster.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/removerolefromdbcluster.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RemoveFromGlobalCluster

RemoveRoleFromDBInstance

All content copied from https://docs.aws.amazon.com/.
