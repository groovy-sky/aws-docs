# DeleteClusterPolicy

Deletes the resource-based policy attached to a cluster. This removes all access permissions defined by the policy, reverting to default access controls.

## Request Syntax

```nohighlight

DELETE /cluster/identifier/policy?client-token=clientToken&expected-policy-version=expectedPolicyVersion HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[clientToken](#API_DeleteClusterPolicy_RequestSyntax)**

Idempotency token so a request is only processed once.

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `[!-~]+`

**[expectedPolicyVersion](#API_DeleteClusterPolicy_RequestSyntax)**

The expected version of the policy to delete. This parameter ensures that you're deleting the correct version of the policy and helps prevent accidental deletions.

**[identifier](#API_DeleteClusterPolicy_RequestSyntax)**

The ID of the cluster.

Pattern: `[a-z0-9]{26}`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "policyVersion": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[policyVersion](#API_DeleteClusterPolicy_ResponseSyntax)**

The version of the policy that was deleted.

Type: String

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dsql-2018-05-10/deleteclusterpolicy.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dsql-2018-05-10/deleteclusterpolicy.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dsql-2018-05-10/deleteclusterpolicy.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dsql-2018-05-10/deleteclusterpolicy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dsql-2018-05-10/deleteclusterpolicy.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dsql-2018-05-10/deleteclusterpolicy.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dsql-2018-05-10/deleteclusterpolicy.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dsql-2018-05-10/deleteclusterpolicy.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dsql-2018-05-10/deleteclusterpolicy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dsql-2018-05-10/deleteclusterpolicy.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteCluster

GetCluster

All content copied from https://docs.aws.amazon.com/.
