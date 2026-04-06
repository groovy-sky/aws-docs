# CreateExtensionAssociation

When you create an extension or configure an AWS authored extension, you
associate the extension with an AWS AppConfig application, environment, or
configuration profile. For example, you can choose to run the `
         AWS AppConfig
            deployment events to Amazon SNS`
AWS authored extension and receive notifications on an Amazon SNS
topic anytime a configuration deployment is started for a specific application. Defining
which extension to associate with an AWS AppConfig resource is called an
_extension association_. An extension association is a specified
relationship between an extension and an AWS AppConfig resource, such as an
application or a configuration profile. For more information about extensions and
associations, see [Extending\
workflows](../../../../services/appconfig/latest/userguide/working-with-appconfig-extensions.md) in the _AWS AppConfig User Guide_.

## Request Syntax

```nohighlight

POST /extensionassociations HTTP/1.1
Content-type: application/json

{
   "ExtensionIdentifier": "string",
   "ExtensionVersionNumber": number,
   "Parameters": {
      "string" : "string"
   },
   "ResourceIdentifier": "string",
   "Tags": {
      "string" : "string"
   }
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[ExtensionIdentifier](#API_CreateExtensionAssociation_RequestSyntax)**

The name, the ID, or the Amazon Resource Name (ARN) of the extension.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: Yes

**[ExtensionVersionNumber](#API_CreateExtensionAssociation_RequestSyntax)**

The version number of the extension. If not specified, AWS AppConfig uses the
maximum version of the extension.

Type: Integer

Required: No

**[Parameters](#API_CreateExtensionAssociation_RequestSyntax)**

The parameter names and values defined in the extensions. Extension parameters marked
`Required` must be entered for this field.

Type: String to string map

Map Entries: Minimum number of 0 items. Maximum number of 10 items.

Key Pattern: `^[^\/#:\n]{1,64}$`

Value Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**[ResourceIdentifier](#API_CreateExtensionAssociation_RequestSyntax)**

The ARN of an application, configuration profile, or environment.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: Yes

**[Tags](#API_CreateExtensionAssociation_RequestSyntax)**

Adds one or more tags for the specified extension association. Tags are metadata that
help you categorize resources in different ways, for example, by purpose, owner, or
environment. Each tag consists of a key and an optional value, both of which you define.

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
   "Arn": "string",
   "ExtensionArn": "string",
   "ExtensionVersionNumber": number,
   "Id": "string",
   "Parameters": {
      "string" : "string"
   },
   "ResourceArn": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 201 response.

The following data is returned in JSON format by the service.

**[Arn](#API_CreateExtensionAssociation_ResponseSyntax)**

The system-generated Amazon Resource Name (ARN) for the extension.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:(aws[a-zA-Z-]*)?:[a-z]+:((eusc-)?[a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1})?:(\d{12})?:[a-zA-Z0-9-_/:.]+`

**[ExtensionArn](#API_CreateExtensionAssociation_ResponseSyntax)**

The ARN of the extension defined in the association.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:(aws[a-zA-Z-]*)?:[a-z]+:((eusc-)?[a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1})?:(\d{12})?:[a-zA-Z0-9-_/:.]+`

**[ExtensionVersionNumber](#API_CreateExtensionAssociation_ResponseSyntax)**

The version number for the extension defined in the association.

Type: Integer

**[Id](#API_CreateExtensionAssociation_ResponseSyntax)**

The system-generated ID for the association.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

**[Parameters](#API_CreateExtensionAssociation_ResponseSyntax)**

The parameter names and values defined in the association.

Type: String to string map

Map Entries: Minimum number of 0 items. Maximum number of 10 items.

Key Pattern: `^[^\/#:\n]{1,64}$`

Value Length Constraints: Minimum length of 1. Maximum length of 2048.

**[ResourceArn](#API_CreateExtensionAssociation_ResponseSyntax)**

The ARNs of applications, configuration profiles, or environments defined in the
association.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:(aws[a-zA-Z-]*)?:[a-z]+:((eusc-)?[a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1})?:(\d{12})?:[a-zA-Z0-9-_/:.]+`

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

This example illustrates one usage of CreateExtensionAssociation.

#### Sample Request

```

POST /extensionassociations HTTP/1.1
Host: appconfig.us-west-2.amazonaws.com
Accept-Encoding: identity
Content-Type: application/json
User-Agent: aws-cli/2.7.19 Python/3.9.11 Windows/10 exe/AMD64 prompt/off command/appconfig.create-extension-association
X-Amz-Date: 20220803T214332Z
Authorization: AWS4-HMAC-SHA256 Credential=AKIAIOSFODNN7EXAMPLE/20220803/us-west-2/appconfig/aws4_request, SignedHeaders=content-type;host;x-amz-date, Signature=39c3b3042cd2aEXAMPLE
Content-Length: 180

{"ExtensionIdentifier": "my-test-extension", "ResourceIdentifier": "arn:aws:appconfig:us-west-2:111122223333:application/xlmtExample", "Parameters": {"MyParamKey": "MyParamValue"}}
```

#### Sample Response

```

{
	"Arn": null,
	"ExtensionArn": "arn:aws:appconfig:us-west-2:111122223333:extension/6czatq1/1",
	"ExtensionVersionNumber": 1,
	"Id": "rnekru4",
	"Parameters": {
		"MyParamKey": "MyParamValue"
	},
	"ResourceArn": "arn:aws:appconfig:us-west-2:111122223333:application/xlmtExample"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/appconfig-2019-10-09/CreateExtensionAssociation)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/appconfig-2019-10-09/CreateExtensionAssociation)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/appconfig-2019-10-09/CreateExtensionAssociation)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/appconfig-2019-10-09/CreateExtensionAssociation)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appconfig-2019-10-09/CreateExtensionAssociation)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/appconfig-2019-10-09/CreateExtensionAssociation)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/appconfig-2019-10-09/CreateExtensionAssociation)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/appconfig-2019-10-09/CreateExtensionAssociation)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/appconfig-2019-10-09/CreateExtensionAssociation)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appconfig-2019-10-09/CreateExtensionAssociation)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CreateExtension

CreateHostedConfigurationVersion
