# GetApplication

Retrieves information about an application.

## Request Syntax

```nohighlight

GET /applications/ApplicationId HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[ApplicationId](#API_GetApplication_RequestSyntax)**

The ID of the application you want to get.

Pattern: `[a-z0-9]{4,7}`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "Description": "string",
   "Id": "string",
   "Name": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Description](#API_GetApplication_ResponseSyntax)**

The description of the application.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

**[Id](#API_GetApplication_ResponseSyntax)**

The application ID.

Type: String

Pattern: `[a-z0-9]{4,7}`

**[Name](#API_GetApplication_ResponseSyntax)**

The application name.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**BadRequestException**

The input fails to satisfy the constraints specified by an AWS service.

**Details**

Detailed information about the input that failed to satisfy the constraints specified by
a call.

HTTP Status Code: 400

**InternalServerException**

There was an internal failure in the AWS AppConfig service.

HTTP Status Code: 500

**ResourceNotFoundException**

The requested resource could not be found.

HTTP Status Code: 404

## Examples

### Example

This example illustrates one usage of GetApplication.

#### Sample Request

```

GET /applications/abc1234 HTTP/1.1
Host: appconfig.us-east-1.amazonaws.com
Accept-Encoding: identity
User-Agent: aws-cli/2.2.4 Python/3.8.8 Linux/5.4.134-73.228.amzn2int.x86_64 exe/x86_64.amzn.2 prompt/off command/appconfig.get-application
X-Amz-Date: 20210917T180608Z
Authorization: AWS4-HMAC-SHA256 Credential=AKIAIOSFODNN7EXAMPLE/20210917/us-east-1/appconfig/aws4_request, SignedHeaders=host;x-amz-date, Signature=39c3b3042cd2aEXAMPLE
```

#### Sample Response

```

{
	"Description": "An application used for creating an example.",
	"Id": "abc1234",
	"Name": "example-application"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appconfig-2019-10-09/getapplication.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appconfig-2019-10-09/getapplication.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/getapplication.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appconfig-2019-10-09/getapplication.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/getapplication.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appconfig-2019-10-09/getapplication.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appconfig-2019-10-09/getapplication.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appconfig-2019-10-09/getapplication.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appconfig-2019-10-09/getapplication.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/getapplication.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetAccountSettings

GetConfiguration
