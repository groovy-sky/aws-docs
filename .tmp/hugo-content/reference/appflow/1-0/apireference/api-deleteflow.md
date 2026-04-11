# DeleteFlow

Enables your application to delete an existing flow. Before deleting the flow, Amazon AppFlow validates the request by checking the flow configuration and status. You can
delete flows one at a time.

## Request Syntax

```nohighlight

POST /delete-flow HTTP/1.1
Content-type: application/json

{
   "flowName": "string",
   "forceDelete": boolean
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[flowName](#API_DeleteFlow_RequestSyntax)**

The specified name of the flow. Spaces are not allowed. Use underscores (\_) or hyphens
(-) only.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `[a-zA-Z0-9][\w!@#.-]+`

Required: Yes

**[forceDelete](#API_DeleteFlow_RequestSyntax)**

Indicates whether Amazon AppFlow should delete the flow, even if it is currently in
use.

Type: Boolean

Required: No

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

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

## Examples

### An active scheduled flow with forceDelete false

This example shows an active scheduled flow with `forceDelete` false. The
result is that a conflict exception is thrown.

#### Sample Request

```json

{
  "flowName": "flowName_value"
}
```

#### Sample Response

```json

{
  "message": "Conflict executing request: Flow is in active state, please set forceDelete to true or deactivate the flow: flowName_value"
}
```

### An active scheduled flow with forceDelete true

This example shows an active scheduled flow with `forceDelete` true. The
result is that the flow is deleted successfully.

#### Sample Request

```json

{
  "flowName": "flowName_value",
  "forceDelete": true
}
```

#### Sample Response

```json

{}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appflow-2020-08-23/deleteflow.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appflow-2020-08-23/deleteflow.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/deleteflow.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appflow-2020-08-23/deleteflow.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/deleteflow.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appflow-2020-08-23/deleteflow.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appflow-2020-08-23/deleteflow.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appflow-2020-08-23/deleteflow.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appflow-2020-08-23/deleteflow.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/deleteflow.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteConnectorProfile

DescribeConnector

All content copied from https://docs.aws.amazon.com/.
