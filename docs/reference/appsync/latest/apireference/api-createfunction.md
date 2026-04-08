# CreateFunction

Creates a `Function` object.

A function is a reusable entity. You can use multiple functions to compose the resolver
logic.

## Request Syntax

```nohighlight

POST /v1/apis/apiId/functions HTTP/1.1
Content-type: application/json

{
   "code": "string",
   "dataSourceName": "string",
   "description": "string",
   "functionVersion": "string",
   "maxBatchSize": number,
   "name": "string",
   "requestMappingTemplate": "string",
   "responseMappingTemplate": "string",
   "runtime": {
      "name": "string",
      "runtimeVersion": "string"
   },
   "syncConfig": {
      "conflictDetection": "string",
      "conflictHandler": "string",
      "lambdaConflictHandlerConfig": {
         "lambdaConflictHandlerArn": "string"
      }
   }
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[apiId](#API_CreateFunction_RequestSyntax)**

The GraphQL API ID.

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[code](#API_CreateFunction_RequestSyntax)**

The `function` code that contains the request and response functions. When
code is used, the `runtime` is required. The `runtime` value must be
`APPSYNC_JS`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 32768.

Required: No

**[dataSourceName](#API_CreateFunction_RequestSyntax)**

The `Function` `DataSource` name.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 65536.

Pattern: `[_A-Za-z][_0-9A-Za-z]*`

Required: Yes

**[description](#API_CreateFunction_RequestSyntax)**

The `Function` description.

Type: String

Required: No

**[functionVersion](#API_CreateFunction_RequestSyntax)**

The `version` of the request mapping template. Currently, the supported value
is 2018-05-29. Note that when using VTL and mapping templates, the
`functionVersion` is required.

Type: String

Required: No

**[maxBatchSize](#API_CreateFunction_RequestSyntax)**

The maximum batching size for a resolver.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 2000.

Required: No

**[name](#API_CreateFunction_RequestSyntax)**

The `Function` name. The function name does not have to be unique.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 65536.

Pattern: `[_A-Za-z][_0-9A-Za-z]*`

Required: Yes

**[requestMappingTemplate](#API_CreateFunction_RequestSyntax)**

The `Function` request mapping template. Functions support only the
2018-05-29 version of the request mapping template.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 65536.

Pattern: `^.*$`

Required: No

**[responseMappingTemplate](#API_CreateFunction_RequestSyntax)**

The `Function` response mapping template.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 65536.

Pattern: `^.*$`

Required: No

**[runtime](#API_CreateFunction_RequestSyntax)**

Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note
that if a runtime is specified, code must also be specified.

Type: [AppSyncRuntime](api-appsyncruntime.md) object

Required: No

**[syncConfig](#API_CreateFunction_RequestSyntax)**

Describes a Sync configuration for a resolver.

Specifies which Conflict Detection strategy and Resolution strategy to use when the
resolver is invoked.

Type: [SyncConfig](api-syncconfig.md) object

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "functionConfiguration": {
      "code": "string",
      "dataSourceName": "string",
      "description": "string",
      "functionArn": "string",
      "functionId": "string",
      "functionVersion": "string",
      "maxBatchSize": number,
      "name": "string",
      "requestMappingTemplate": "string",
      "responseMappingTemplate": "string",
      "runtime": {
         "name": "string",
         "runtimeVersion": "string"
      },
      "syncConfig": {
         "conflictDetection": "string",
         "conflictHandler": "string",
         "lambdaConflictHandlerConfig": {
            "lambdaConflictHandlerArn": "string"
         }
      }
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[functionConfiguration](#API_CreateFunction_ResponseSyntax)**

The `Function` object.

Type: [FunctionConfiguration](api-functionconfiguration.md) object

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

**ConcurrentModificationException**

Another modification is in progress at this time and it must complete before you can
make your change.

HTTP Status Code: 409

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appsync-2017-07-25/createfunction.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appsync-2017-07-25/createfunction.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appsync-2017-07-25/createfunction.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appsync-2017-07-25/createfunction.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appsync-2017-07-25/createfunction.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appsync-2017-07-25/createfunction.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appsync-2017-07-25/createfunction.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appsync-2017-07-25/createfunction.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appsync-2017-07-25/createfunction.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appsync-2017-07-25/createfunction.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateDomainName

CreateGraphqlApi
