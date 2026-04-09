# UpdateCluster

The _UpdateCluster_ API allows you to modify both single-Region and multi-Region cluster configurations. With the _multiRegionProperties_ parameter, you can add or modify witness Region support and manage peer relationships with clusters in other Regions.

###### Note

Note that updating multi-Region clusters requires additional IAM permissions
beyond those needed for standard cluster updates, as detailed in the Permissions
section.

**Required permissions**

dsql:UpdateCluster

Permission to update a DSQL cluster.

Resources: `arn:aws:dsql:region:account-id:cluster/cluster-id
               `

dsql:PutMultiRegionProperties

Permission to configure multi-Region properties for a cluster.

Resources: `arn:aws:dsql:region:account-id:cluster/cluster-id
               `

dsql:GetCluster

Permission to retrieve cluster information.

Resources: `arn:aws:dsql:region:account-id:cluster/cluster-id
               `

dsql:AddPeerCluster

Permission to add peer clusters.

Resources:

- Local cluster: `arn:aws:dsql:region:account-id:cluster/cluster-id
                       `

- Each peer cluster: exact ARN of each specified peer cluster

dsql:RemovePeerCluster

Permission to remove peer clusters. The _dsql:RemovePeerCluster_ permission
uses a wildcard ARN pattern to simplify permission management during updates.

Resources:
`arn:aws:dsql:*:account-id:cluster/*`

dsql:PutWitnessRegion

Permission to set a witness Region.

Resources: `arn:aws:dsql:region:account-id:cluster/cluster-id
               `

Condition Keys: dsql:WitnessRegion (matching the specified witness
Region)

**This permission is checked both in the cluster Region and in the witness**
**Region.**

###### Important

- The witness region specified in
`multiRegionProperties.witnessRegion` cannot be the same as the
cluster's Region.

- When updating clusters with peer relationships, permissions are checked for both adding and removing peers.

- The `dsql:RemovePeerCluster` permission uses a wildcard ARN pattern to simplify permission management during updates.

## Request Syntax

```nohighlight

POST /cluster/identifier HTTP/1.1
Content-type: application/json

{
   "clientToken": "string",
   "deletionProtectionEnabled": boolean,
   "kmsEncryptionKey": "string",
   "multiRegionProperties": {
      "clusters": [ "string" ],
      "witnessRegion": "string"
   }
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[identifier](#API_UpdateCluster_RequestSyntax)**

The ID of the cluster you want to update.

Pattern: `[a-z0-9]{26}`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[clientToken](#API_UpdateCluster_RequestSyntax)**

A unique, case-sensitive identifier that you provide to ensure the idempotency of the
request. Idempotency ensures that an API request completes only once. With an idempotent
request, if the original request completes successfully. The subsequent retries with the
same client token return the result from the original successful request and they have no
additional effect.

If you don't specify a client token, the AWS SDK automatically generates
one.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `[!-~]+`

Required: No

**[deletionProtectionEnabled](#API_UpdateCluster_RequestSyntax)**

Specifies whether to enable deletion protection in your cluster.

Type: Boolean

Required: No

**[kmsEncryptionKey](#API_UpdateCluster_RequestSyntax)**

The AWS KMS key that encrypts and protects the data on your cluster. You can specify the ARN, ID, or alias of an existing key or have AWS create a default key for you.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `[a-zA-Z0-9:/_-]+`

Required: No

**[multiRegionProperties](#API_UpdateCluster_RequestSyntax)**

The new multi-Region cluster configuration settings to be applied during an update operation.

Type: [MultiRegionProperties](api-multiregionproperties.md) object

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "arn": "string",
   "creationTime": number,
   "identifier": "string",
   "status": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[arn](#API_UpdateCluster_ResponseSyntax)**

The ARN of the updated cluster.

Type: String

Pattern: `arn:aws(-[^:]+)?:dsql:[a-z0-9-]{1,20}:[0-9]{12}:cluster/[a-z0-9]{26}`

**[creationTime](#API_UpdateCluster_ResponseSyntax)**

The time of when the cluster was created.

Type: Timestamp

**[identifier](#API_UpdateCluster_ResponseSyntax)**

The ID of the cluster to update.

Type: String

Pattern: `[a-z0-9]{26}`

**[status](#API_UpdateCluster_ResponseSyntax)**

The status of the updated cluster.

Type: String

Valid Values: `CREATING | ACTIVE | IDLE | INACTIVE | UPDATING | DELETING | DELETED | FAILED | PENDING_SETUP | PENDING_DELETE`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You do not have sufficient access to perform this action.

HTTP Status Code: 403

**ConflictException**

The submitted action has conflicts.

**resourceId**

Resource Id

**resourceType**

Resource Type

HTTP Status Code: 409

**InternalServerException**

The request processing has failed because of an unknown error, exception or
failure.

**retryAfterSeconds**

Retry after seconds.

HTTP Status Code: 500

**ResourceNotFoundException**

The resource could not be found.

**resourceId**

The resource ID could not be found.

**resourceType**

The resource type could not be found.

HTTP Status Code: 404

**ThrottlingException**

The request was denied due to request throttling.

**message**

The message that the request was denied due to request throttling.

**quotaCode**

The request exceeds a request rate quota.

**retryAfterSeconds**

The request exceeds a request rate quota. Retry after seconds.

**serviceCode**

The request exceeds a service quota.

HTTP Status Code: 429

**ValidationException**

The input failed to satisfy the constraints specified by an AWS service.

**fieldList**

A list of fields that didn't validate.

**reason**

The reason for the validation exception.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dsql-2018-05-10/updatecluster.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dsql-2018-05-10/updatecluster.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dsql-2018-05-10/updatecluster.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dsql-2018-05-10/updatecluster.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dsql-2018-05-10/updatecluster.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dsql-2018-05-10/updatecluster.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dsql-2018-05-10/updatecluster.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dsql-2018-05-10/updatecluster.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dsql-2018-05-10/updatecluster.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dsql-2018-05-10/updatecluster.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UntagResource

Data Types

All content copied from https://docs.aws.amazon.com/.
