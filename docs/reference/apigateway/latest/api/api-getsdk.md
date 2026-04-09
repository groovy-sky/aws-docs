# GetSdk

Generates a client SDK for a RestApi and Stage.

## Request Syntax

```nohighlight

GET /restapis/restapi_id/stages/stage_name/sdks/sdk_type?parameters HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[parameters](#API_GetSdk_RequestSyntax)**

A string-to-string key-value map of query parameters `sdkType`-dependent properties of the SDK. For `sdkType` of `objectivec` or `swift`, a parameter named `classPrefix` is required. For `sdkType` of `android`, parameters named `groupId`, `artifactId`, `artifactVersion`, and `invokerPackage` are required. For `sdkType` of `java`, parameters named `serviceName` and `javaPackageName` are required.

**[restapi\_id](#API_GetSdk_RequestSyntax)**

The string identifier of the associated RestApi.

Required: Yes

**[sdk\_type](#API_GetSdk_RequestSyntax)**

The language for the generated SDK. Currently `java`, `javascript`, `android`, `objectivec` (for iOS), `swift` (for iOS), and `ruby` are supported.

Required: Yes

**[stage\_name](#API_GetSdk_RequestSyntax)**

The name of the Stage that the SDK will use.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-Type: contentType
Content-Disposition: contentDisposition

body
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The response returns the following HTTP headers.

**[contentDisposition](#API_GetSdk_ResponseSyntax)**

The content-disposition header value in the HTTP response.

**[contentType](#API_GetSdk_ResponseSyntax)**

The content-type header value in the HTTP response.

The response returns the following as the HTTP body.

**[body](#API_GetSdk_ResponseSyntax)**

The binary blob response to GetSdk, which contains the generated SDK.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](../../../../services/apigateway/latest/api/commonerrors.md).

**BadRequestException**

The submitted request is not valid, for example, the input is incomplete or incorrect. See the accompanying error message for details.

HTTP Status Code: 400

**ConflictException**

The request configuration has conflicts. For details, see the accompanying error message.

HTTP Status Code: 409

**LimitExceededException**

The request exceeded the rate limit. Retry after the specified time period.

HTTP Status Code: 429

**NotFoundException**

The requested resource is not found. Make sure that the request URI is correct.

HTTP Status Code: 404

**TooManyRequestsException**

The request has reached its throttling limit. Retry after the specified time period.

HTTP Status Code: 429

**UnauthorizedException**

The request is denied because the caller has insufficient permissions.

HTTP Status Code: 401

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/getsdk.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/apigateway-2015-07-09/getsdk.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/apigateway-2015-07-09/getsdk.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/apigateway-2015-07-09/getsdk.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/apigateway-2015-07-09/getsdk.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/apigateway-2015-07-09/getsdk.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/apigateway-2015-07-09/getsdk.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/apigateway-2015-07-09/getsdk.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/getsdk.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/apigateway-2015-07-09/getsdk.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetRestApis

GetSdkType

All content copied from https://docs.aws.amazon.com/.
