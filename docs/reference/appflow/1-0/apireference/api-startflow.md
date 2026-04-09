# StartFlow

Activates an existing flow. For on-demand flows, this operation runs the flow
immediately. For schedule and event-triggered flows, this operation activates the flow.

## Request Syntax

```nohighlight

POST /start-flow HTTP/1.1
Content-type: application/json

{
   "clientToken": "string",
   "flowName": "string"
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[clientToken](#API_StartFlow_RequestSyntax)**

The `clientToken` parameter is an idempotency token. It ensures that your
`StartFlow` request completes only once. You choose the value to pass. For
example, if you don't receive a response from your request, you can safely retry the request
with the same `clientToken` parameter value.

If you omit a `clientToken` value, the AWS SDK that you are
using inserts a value for you. This way, the SDK can safely retry requests multiple times
after a network error. You must provide your own value for other use cases.

If you specify input parameters that differ from your first request, an error occurs for
flows that run on a schedule or based on an event. However, the error doesn't occur for flows
that run on demand. You set the conditions that initiate your flow for the
`triggerConfig` parameter.

If you use a different value for `clientToken`, Amazon AppFlow considers
it a new call to `StartFlow`. The token is active for 8 hours.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `[ -~]+`

Required: No

**[flowName](#API_StartFlow_RequestSyntax)**

The specified name of the flow. Spaces are not allowed. Use underscores (\_) or hyphens
(-) only.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `[a-zA-Z0-9][\w!@#.-]+`

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "executionId": "string",
   "flowArn": "string",
   "flowStatus": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[executionId](#API_StartFlow_ResponseSyntax)**

Returns the internal execution ID of an on-demand flow when the flow is started. For
scheduled or event-triggered flows, this value is null.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `\S+`

**[flowArn](#API_StartFlow_ResponseSyntax)**

The flow's Amazon Resource Name (ARN).

Type: String

Length Constraints: Maximum length of 512.

Pattern: `arn:aws:appflow:.*:[0-9]+:.*`

**[flowStatus](#API_StartFlow_ResponseSyntax)**

Indicates the current status of the flow.

Type: String

Valid Values: `Active | Deprecated | Deleted | Draft | Errored | Suspended`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ConflictException**

There was a conflict when processing the request (for example, a flow with the given name
already exists within the account. Check for conflicting resource names and try again.

HTTP Status Code: 409

**InternalServerException**

An internal service error occurred during the processing of your request. Try again
later.

HTTP Status Code: 500

**ResourceNotFoundException**

The resource specified in the request (such as the source or destination connector
profile) is not found.

HTTP Status Code: 404

**ServiceQuotaExceededException**

The request would cause a service quota (such as the number of flows) to be exceeded.

HTTP Status Code: 402

## Examples

### StartFlow example

This example shows a sample request and response for the `StartFlow` API.

#### Sample Request

```json

{
  "flowName": "name"
}
```

#### Sample Response

```json

{
  "flowArn": "arn:aws:appflow:region:<AccountId>:flow/test_flow_ondemand_1",
  "flowStatus": "Active"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appflow-2020-08-23/startflow.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appflow-2020-08-23/startflow.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/startflow.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appflow-2020-08-23/startflow.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/startflow.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appflow-2020-08-23/startflow.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appflow-2020-08-23/startflow.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appflow-2020-08-23/startflow.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appflow-2020-08-23/startflow.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/startflow.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResetConnectorMetadataCache

StopFlow

All content copied from https://docs.aws.amazon.com/.
