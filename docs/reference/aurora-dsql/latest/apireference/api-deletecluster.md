# DeleteCluster

Deletes a cluster in Amazon Aurora DSQL.

## Request Syntax

```nohighlight

DELETE /cluster/identifier?client-token=clientToken HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[clientToken](#API_DeleteCluster_RequestSyntax)**

A unique, case-sensitive identifier that you provide to ensure the idempotency of the
request. Idempotency ensures that an API request completes only once. With an idempotent
request, if the original request completes successfully. The subsequent retries with the
same client token return the result from the original successful request and they have no
additional effect.

If you don't specify a client token, the AWS SDK automatically generates
one.

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `[!-~]+`

**[identifier](#API_DeleteCluster_RequestSyntax)**

The ID of the cluster to delete.

Pattern: `[a-z0-9]{26}`

Required: Yes

## Request Body

The request does not have a request body.

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

**[arn](#API_DeleteCluster_ResponseSyntax)**

The ARN of the deleted cluster.

Type: String

Pattern: `arn:aws(-[^:]+)?:dsql:[a-z0-9-]{1,20}:[0-9]{12}:cluster/[a-z0-9]{26}`

**[creationTime](#API_DeleteCluster_ResponseSyntax)**

The time of when the cluster was created.

Type: Timestamp

**[identifier](#API_DeleteCluster_ResponseSyntax)**

The ID of the deleted cluster.

Type: String

Pattern: `[a-z0-9]{26}`

**[status](#API_DeleteCluster_ResponseSyntax)**

The status of the cluster.

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dsql-2018-05-10/deletecluster.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dsql-2018-05-10/deletecluster.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dsql-2018-05-10/deletecluster.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dsql-2018-05-10/deletecluster.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dsql-2018-05-10/deletecluster.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dsql-2018-05-10/deletecluster.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dsql-2018-05-10/deletecluster.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dsql-2018-05-10/deletecluster.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dsql-2018-05-10/deletecluster.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dsql-2018-05-10/deletecluster.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateCluster

DeleteClusterPolicy

All content copied from https://docs.aws.amazon.com/.
