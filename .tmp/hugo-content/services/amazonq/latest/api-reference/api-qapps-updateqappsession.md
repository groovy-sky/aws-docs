# UpdateQAppSession

Updates the session for a given Q App `sessionId`. This is only valid when at
least one card of the session is in the `WAITING` state. Data for each
`WAITING` card can be provided as input. If inputs are not provided, the call
will be accepted but session will not move forward. Inputs for cards that are not in the
`WAITING` status will be ignored.

## Request Syntax

```nohighlight

POST /runtime.updateQAppSession HTTP/1.1
instance-id: instanceId
Content-type: application/json

{
   "sessionId": "string",
   "values": [
      {
         "cardId": "string",
         "submissionMutation": {
            "mutationType": "string",
            "submissionId": "string"
         },
         "value": "string"
      }
   ]
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[instanceId](#API_qapps_UpdateQAppSession_RequestSyntax)**

The unique identifier of the Amazon Q Business application environment instance.

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[sessionId](#API_qapps_UpdateQAppSession_RequestSyntax)**

The unique identifier of the Q App session to provide input for.

Type: String

Pattern: `[\da-f]{8}-[\da-f]{4}-[45][\da-f]{3}-[89ABab][\da-f]{3}-[\da-f]{12}`

Required: Yes

**[values](#API_qapps_UpdateQAppSession_RequestSyntax)**

The input values to provide for the current state of the Q App session.

Type: Array of [CardValue](api-qapps-cardvalue.md) objects

Array Members: Minimum number of 0 items. Maximum number of 20 items.

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "sessionArn": "string",
   "sessionId": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[sessionArn](#API_qapps_UpdateQAppSession_ResponseSyntax)**

The Amazon Resource Name (ARN) of the updated Q App session.

Type: String

**[sessionId](#API_qapps_UpdateQAppSession_ResponseSyntax)**

The unique identifier of the updated Q App session.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

The client is not authorized to perform the requested operation.

HTTP Status Code: 403

**InternalServerException**

An internal service error occurred while processing the request.

**retryAfterSeconds**

The number of seconds to wait before retrying the operation

HTTP Status Code: 500

**ResourceNotFoundException**

The requested resource could not be found.

**resourceId**

The unique identifier of the resource

**resourceType**

The type of the resource

HTTP Status Code: 404

**ServiceQuotaExceededException**

The requested operation could not be completed because it would exceed the service's quota
or limit.

**quotaCode**

The code of the quota that was exceeded

**resourceId**

The unique identifier of the resource

**resourceType**

The type of the resource

**serviceCode**

The code for the service where the quota was exceeded

HTTP Status Code: 402

**ThrottlingException**

The requested operation could not be completed because too many requests were sent at
once. Wait a bit and try again later.

**quotaCode**

The code of the quota that was exceeded

**retryAfterSeconds**

The number of seconds to wait before retrying the operation

**serviceCode**

The code for the service where the quota was exceeded

HTTP Status Code: 429

**UnauthorizedException**

The client is not authenticated or authorized to perform the requested operation.

HTTP Status Code: 401

**ValidationException**

The input failed to satisfy the constraints specified by the service.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/qapps-2023-11-27/updateqappsession.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/qapps-2023-11-27/updateqappsession.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qapps-2023-11-27/updateqappsession.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/qapps-2023-11-27/updateqappsession.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qapps-2023-11-27/updateqappsession.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/qapps-2023-11-27/updateqappsession.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/qapps-2023-11-27/updateqappsession.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/qapps-2023-11-27/updateqappsession.md)

- [AWS SDK for Python](../../../goto/boto3/qapps-2023-11-27/updateqappsession.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qapps-2023-11-27/updateqappsession.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateQAppPermissions

UpdateQAppSessionMetadata

All content copied from https://docs.aws.amazon.com/.
