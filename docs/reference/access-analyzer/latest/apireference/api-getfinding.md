# GetFinding

Retrieves information about the specified finding. GetFinding and GetFindingV2 both use
`access-analyzer:GetFinding` in the `Action` element of an IAM
policy statement. You must have permission to perform the
`access-analyzer:GetFinding` action.

###### Note

GetFinding is supported only for external access analyzers. You must use GetFindingV2
for internal and unused access analyzers.

## Request Syntax

```nohighlight

GET /finding/id?analyzerArn=analyzerArn HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[analyzerArn](#API_GetFinding_RequestSyntax)**

The [ARN of\
the analyzer](../../../../services/iam/latest/userguide/access-analyzer-getting-started.md#permission-resources) that generated the finding.

Pattern: `[^:]*:[^:]*:[^:]*:[^:]*:[^:]*:analyzer/.{1,255}`

Required: Yes

**[id](#API_GetFinding_RequestSyntax)**

The ID of the finding to retrieve.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "finding": {
      "action": [ "string" ],
      "analyzedAt": "string",
      "condition": {
         "string" : "string"
      },
      "createdAt": "string",
      "error": "string",
      "id": "string",
      "isPublic": boolean,
      "principal": {
         "string" : "string"
      },
      "resource": "string",
      "resourceControlPolicyRestriction": "string",
      "resourceOwnerAccount": "string",
      "resourceType": "string",
      "sources": [
         {
            "detail": {
               "accessPointAccount": "string",
               "accessPointArn": "string"
            },
            "type": "string"
         }
      ],
      "status": "string",
      "updatedAt": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[finding](#API_GetFinding_ResponseSyntax)**

A `finding` object that contains finding details.

Type: [Finding](api-finding.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You do not have sufficient access to perform this action.

HTTP Status Code: 403

**InternalServerException**

Internal server error.

**retryAfterSeconds**

The seconds to wait to retry.

HTTP Status Code: 500

**ResourceNotFoundException**

The specified resource could not be found.

**resourceId**

The ID of the resource.

**resourceType**

The type of the resource.

HTTP Status Code: 404

**ThrottlingException**

Throttling limit exceeded error.

**retryAfterSeconds**

The seconds to wait to retry.

HTTP Status Code: 429

**ValidationException**

Validation exception error.

**fieldList**

A list of fields that didn't validate.

**reason**

The reason for the exception.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/accessanalyzer-2019-11-01/GetFinding)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/accessanalyzer-2019-11-01/GetFinding)

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/getfinding.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/accessanalyzer-2019-11-01/getfinding.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/getfinding.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/accessanalyzer-2019-11-01/getfinding.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/accessanalyzer-2019-11-01/getfinding.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/accessanalyzer-2019-11-01/getfinding.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/accessanalyzer-2019-11-01/GetFinding)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/getfinding.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetArchiveRule

GetFindingRecommendation
