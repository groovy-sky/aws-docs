# StartResourceScan

Immediately starts a scan of the policies applied to the specified resource.

###### Note

This action is supported only for external access analyzers.

## Request Syntax

```nohighlight

POST /resource/scan HTTP/1.1
Content-type: application/json

{
   "analyzerArn": "string",
   "resourceArn": "string",
   "resourceOwnerAccount": "string"
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[analyzerArn](#API_StartResourceScan_RequestSyntax)**

The [ARN of\
the analyzer](../../../../services/iam/latest/userguide/access-analyzer-getting-started.md#permission-resources) to use to scan the policies applied to the specified
resource.

Type: String

Pattern: `[^:]*:[^:]*:[^:]*:[^:]*:[^:]*:analyzer/.{1,255}`

Required: Yes

**[resourceArn](#API_StartResourceScan_RequestSyntax)**

The ARN of the resource to scan.

Type: String

Pattern: `arn:[^:]*:[^:]*:[^:]*:[^:]*:.*`

Required: Yes

**[resourceOwnerAccount](#API_StartResourceScan_RequestSyntax)**

The AWS account ID that owns the resource. For most AWS resources, the owning
account is the account in which the resource was created.

Type: String

Required: No

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/accessanalyzer-2019-11-01/startresourcescan.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/accessanalyzer-2019-11-01/startresourcescan.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/startresourcescan.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/accessanalyzer-2019-11-01/startresourcescan.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/startresourcescan.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/accessanalyzer-2019-11-01/startresourcescan.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/accessanalyzer-2019-11-01/startresourcescan.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/accessanalyzer-2019-11-01/startresourcescan.md)

- [AWS SDK for Python](../../../../services/goto/boto3/accessanalyzer-2019-11-01/startresourcescan.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/startresourcescan.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

StartPolicyGeneration

TagResource
