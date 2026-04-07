# CreateApplication

Creates an application. In AWS AppConfig, an application is simply an
organizational construct like a folder. This organizational construct has a relationship
with some unit of executable code. For example, you could create an application called
MyMobileApp to organize and manage configuration data for a mobile application installed by
your users.

## Request Syntax

```nohighlight

POST /applications HTTP/1.1
Content-type: application/json

{
   "Description": "string",
   "Name": "string",
   "Tags": {
      "string" : "string"
   }
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[Description](#API_CreateApplication_RequestSyntax)**

A description of the application.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Required: No

**[Name](#API_CreateApplication_RequestSyntax)**

A name for the application.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

**[Tags](#API_CreateApplication_RequestSyntax)**

Metadata to assign to the application. Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which
you define.

Type: String to string map

Map Entries: Minimum number of 0 items. Maximum number of 50 items.

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Value Length Constraints: Maximum length of 256.

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 201
Content-type: application/json

{
   "Description": "string",
   "Id": "string",
   "Name": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 201 response.

The following data is returned in JSON format by the service.

**[Description](#API_CreateApplication_ResponseSyntax)**

The description of the application.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

**[Id](#API_CreateApplication_ResponseSyntax)**

The application ID.

Type: String

Pattern: `[a-z0-9]{4,7}`

**[Name](#API_CreateApplication_ResponseSyntax)**

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

**ServiceQuotaExceededException**

The number of one more AWS AppConfig resources exceeds the maximum allowed. Verify that your
environment doesn't exceed the following service quotas:

Applications: 100 max

Deployment strategies: 20 max

Configuration profiles: 100 max per application

Environments: 20 max per application

To resolve this issue, you can delete one or more resources and try again. Or, you can
request a quota increase. For more information about quotas and to request an increase, see
[Service quotas for AWS AppConfig](https://docs.aws.amazon.com/general/latest/gr/appconfig.html#limits_appconfig) in the Amazon Web Services General Reference.

HTTP Status Code: 402

## Examples

### Example

This example illustrates one usage of CreateApplication.

#### Sample Request

```

POST /applications HTTP/1.1
Host: appconfig.us-east-1.amazonaws.com
Accept-Encoding: identity
User-Agent: aws-cli/2.2.4 Python/3.8.8 Linux/5.4.134-73.228.amzn2int.x86_64 exe/x86_64.amzn.2 prompt/off command/appconfig.create-application
X-Amz-Date: 20210916T175455Z
Authorization: AWS4-HMAC-SHA256 Credential=AKIAIOSFODNN7EXAMPLE/20210916/us-east-1/appconfig/aws4_request, SignedHeaders=host;x-amz-date, Signature=39c3b3042cd2aEXAMPLE
Content-Length: 91

{
	"Name": "test-application",
	"Description": "An application used for creating an example."
}
```

#### Sample Response

```

{
	"Description": "An application used for creating an example.",
	"Id": "abc1234",
	"Name": "test-application"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/appconfig-2019-10-09/CreateApplication)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/appconfig-2019-10-09/CreateApplication)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/appconfig-2019-10-09/CreateApplication)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/appconfig-2019-10-09/CreateApplication)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appconfig-2019-10-09/CreateApplication)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/appconfig-2019-10-09/CreateApplication)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/appconfig-2019-10-09/CreateApplication)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/appconfig-2019-10-09/CreateApplication)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/appconfig-2019-10-09/CreateApplication)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appconfig-2019-10-09/CreateApplication)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS AppConfig

CreateConfigurationProfile
