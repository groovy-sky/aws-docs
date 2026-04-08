# DeleteConnectorProfile

Enables you to delete an existing connector profile.

## Request Syntax

```nohighlight

POST /delete-connector-profile HTTP/1.1
Content-type: application/json

{
   "connectorProfileName": "string",
   "forceDelete": boolean
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[connectorProfileName](#API_DeleteConnectorProfile_RequestSyntax)**

The name of the connector profile. The name is unique for each
`ConnectorProfile` in your account.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `[\w/!@#+=.-]+`

Required: Yes

**[forceDelete](#API_DeleteConnectorProfile_RequestSyntax)**

Indicates whether Amazon AppFlow should delete the profile, even if it is currently
in use in one or more flows.

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

### Marketo

This example shows a sample request and response for the
`CreateConnectorProfile` API using Marketo. If `forceDelete` is
false, and there are flows associated, you will see a message similar to that shown in the
third sample.

#### Sample Request

```json

{
  "connectorProfileName": "testMarketoProfile",
  "forceDelete": "true"
}
```

#### Sample Response

```json

{
  "deleted": true,
  "flowNames": null
}
```

```json

{
  "message": "Conflict executing request: Connector profile: testMarketoProfile is associated with one or more flows. If you still want to *delete* it, *then* make *delete* request *with* forceDelete flag *as* true. *Some* of the associated flows are: [myTestFlow1]"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appflow-2020-08-23/deleteconnectorprofile.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appflow-2020-08-23/deleteconnectorprofile.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/deleteconnectorprofile.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appflow-2020-08-23/deleteconnectorprofile.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/deleteconnectorprofile.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appflow-2020-08-23/deleteconnectorprofile.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appflow-2020-08-23/deleteconnectorprofile.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appflow-2020-08-23/deleteconnectorprofile.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appflow-2020-08-23/deleteconnectorprofile.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/deleteconnectorprofile.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateFlow

DeleteFlow
