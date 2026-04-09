# StopFlow

Deactivates the existing flow. For on-demand flows, this operation returns an
`unsupportedOperationException` error message. For schedule and event-triggered
flows, this operation deactivates the flow.

## Request Syntax

```nohighlight

POST /stop-flow HTTP/1.1
Content-type: application/json

{
   "flowName": "string"
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[flowName](#API_StopFlow_RequestSyntax)**

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
   "flowArn": "string",
   "flowStatus": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[flowArn](#API_StopFlow_ResponseSyntax)**

The flow's Amazon Resource Name (ARN).

Type: String

Length Constraints: Maximum length of 512.

Pattern: `arn:aws:appflow:.*:[0-9]+:.*`

**[flowStatus](#API_StopFlow_ResponseSyntax)**

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

**UnsupportedOperationException**

The requested operation is not supported for the current flow.

HTTP Status Code: 400

## Examples

### StopFlow example

This example shows a sample request and response for the `StopFlow` API.

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
  "flowStatus": "Suspended"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appflow-2020-08-23/stopflow.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appflow-2020-08-23/stopflow.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/stopflow.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appflow-2020-08-23/stopflow.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/stopflow.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appflow-2020-08-23/stopflow.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appflow-2020-08-23/stopflow.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appflow-2020-08-23/stopflow.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appflow-2020-08-23/stopflow.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/stopflow.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StartFlow

TagResource

All content copied from https://docs.aws.amazon.com/.
