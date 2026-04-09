# StartDataSourceIntrospection

Creates a new introspection. Returns the `introspectionId` of the new
introspection after its creation.

## Request Syntax

```nohighlight

POST /v1/datasources/introspections HTTP/1.1
Content-type: application/json

{
   "rdsDataApiConfig": {
      "databaseName": "string",
      "resourceArn": "string",
      "secretArn": "string"
   }
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[rdsDataApiConfig](#API_StartDataSourceIntrospection_RequestSyntax)**

The `rdsDataApiConfig` object data.

Type: [RdsDataApiConfig](api-rdsdataapiconfig.md) object

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "introspectionId": "string",
   "introspectionStatus": "string",
   "introspectionStatusDetail": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[introspectionId](#API_StartDataSourceIntrospection_ResponseSyntax)**

The introspection ID. Each introspection contains a unique ID that can be used to
reference the instrospection record.

Type: String

**[introspectionStatus](#API_StartDataSourceIntrospection_ResponseSyntax)**

The status of the introspection during creation. By default, when a new instrospection
has been created, the status will be set to `PROCESSING`. Once the operation has
been completed, the status will change to `SUCCESS` or `FAILED`
depending on how the data was parsed. A `FAILED` operation will return an error
and its details as an `introspectionStatusDetail`.

Type: String

Valid Values: `PROCESSING | FAILED | SUCCESS`

**[introspectionStatusDetail](#API_StartDataSourceIntrospection_ResponseSyntax)**

The error detail field. When a `FAILED` `introspectionStatus` is returned, the `introspectionStatusDetail`
will also return the exact error that was generated during the operation.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**BadRequestException**

The request is not well formed. For example, a value is invalid or a required field is
missing. Check the field values, and then try again.

**detail**

Provides further details for the reason behind the bad request. For reason type
`CODE_ERROR`, the detail will contain a list of code errors.

**reason**

Provides context for the cause of the bad request. The only supported value is
`CODE_ERROR`.

HTTP Status Code: 400

**InternalFailureException**

An internal AWS AppSync error occurred. Try your request again.

HTTP Status Code: 500

**NotFoundException**

The resource specified in the request was not found. Check the resource, and then try
again.

HTTP Status Code: 404

**UnauthorizedException**

You aren't authorized to perform this operation.

HTTP Status Code: 401

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appsync-2017-07-25/startdatasourceintrospection.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appsync-2017-07-25/startdatasourceintrospection.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appsync-2017-07-25/startdatasourceintrospection.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appsync-2017-07-25/startdatasourceintrospection.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appsync-2017-07-25/startdatasourceintrospection.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appsync-2017-07-25/startdatasourceintrospection.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appsync-2017-07-25/startdatasourceintrospection.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appsync-2017-07-25/startdatasourceintrospection.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appsync-2017-07-25/startdatasourceintrospection.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appsync-2017-07-25/startdatasourceintrospection.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PutGraphqlApiEnvironmentVariables

StartSchemaCreation

All content copied from https://docs.aws.amazon.com/.
