# UnregisterConnector

Unregisters the custom connector registered in your account that matches the connector
label provided in the request.

## Request Syntax

```nohighlight

POST /unregister-connector HTTP/1.1
Content-type: application/json

{
   "connectorLabel": "string",
   "forceDelete": boolean
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[connectorLabel](#API_UnregisterConnector_RequestSyntax)**

The label of the connector. The label is unique for each
`ConnectorRegistration` in your AWS account.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `[a-zA-Z0-9][\w!@#.-]+`

Required: Yes

**[forceDelete](#API_UnregisterConnector_RequestSyntax)**

Indicates whether Amazon AppFlow should unregister the connector, even if it is
currently in use in one or more connector profiles. The default value is false.

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

### Unregister connectors

This example shows a sample request for the `UnregisterConnector` API and a
sample response.

#### Sample Request

```json

{
  "connectorLabel": "MySampleCustomConnector",
  "forceDelete": true
}
```

```json

{
  "message": "Conflict executing request: Connector: MySampleCustomConnector is associated with one or more connector profiles. If you still want to *delete* it, *then* make *delete* request *with* forceDelete flag *as* true. *Some* of the associated connector profiles are: [myTestProfile1, myTestProfile2]"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appflow-2020-08-23/unregisterconnector.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appflow-2020-08-23/unregisterconnector.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/unregisterconnector.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appflow-2020-08-23/unregisterconnector.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/unregisterconnector.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appflow-2020-08-23/unregisterconnector.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appflow-2020-08-23/unregisterconnector.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appflow-2020-08-23/unregisterconnector.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appflow-2020-08-23/unregisterconnector.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/unregisterconnector.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TagResource

UntagResource

All content copied from https://docs.aws.amazon.com/.
