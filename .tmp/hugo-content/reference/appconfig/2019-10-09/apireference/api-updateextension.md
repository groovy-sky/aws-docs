# UpdateExtension

Updates an AWS AppConfig extension. For more information about extensions, see
[Extending\
workflows](../../../../services/appconfig/latest/userguide/working-with-appconfig-extensions.md) in the _AWS AppConfig User Guide_.

## Request Syntax

```nohighlight

PATCH /extensions/ExtensionIdentifier HTTP/1.1
Content-type: application/json

{
   "Actions": {
      "string" : [
         {
            "Description": "string",
            "Name": "string",
            "RoleArn": "string",
            "Uri": "string"
         }
      ]
   },
   "Description": "string",
   "Parameters": {
      "string" : {
         "Description": "string",
         "Dynamic": boolean,
         "Required": boolean
      }
   },
   "VersionNumber": number
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[ExtensionIdentifier](#API_UpdateExtension_RequestSyntax)**

The name, the ID, or the Amazon Resource Name (ARN) of the extension.

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[Actions](#API_UpdateExtension_RequestSyntax)**

The actions defined in the extension.

Type: String to array of [Action](api-action.md) objects map

Map Entries: Maximum number of 5 items.

Valid Keys: `PRE_CREATE_HOSTED_CONFIGURATION_VERSION | PRE_START_DEPLOYMENT | AT_DEPLOYMENT_TICK | ON_DEPLOYMENT_START | ON_DEPLOYMENT_STEP | ON_DEPLOYMENT_BAKING | ON_DEPLOYMENT_COMPLETE | ON_DEPLOYMENT_ROLLED_BACK`

Array Members: Fixed number of 1 item.

Required: No

**[Description](#API_UpdateExtension_RequestSyntax)**

Information about the extension.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Required: No

**[Parameters](#API_UpdateExtension_RequestSyntax)**

One or more parameters for the actions called by the extension.

Type: String to [Parameter](api-parameter.md) object map

Map Entries: Maximum number of 10 items.

Key Pattern: `^[^\/#:\n]{1,64}$`

Required: No

**[VersionNumber](#API_UpdateExtension_RequestSyntax)**

The extension version number.

Type: Integer

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "Actions": {
      "string" : [
         {
            "Description": "string",
            "Name": "string",
            "RoleArn": "string",
            "Uri": "string"
         }
      ]
   },
   "Arn": "string",
   "Description": "string",
   "Id": "string",
   "Name": "string",
   "Parameters": {
      "string" : {
         "Description": "string",
         "Dynamic": boolean,
         "Required": boolean
      }
   },
   "VersionNumber": number
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Actions](#API_UpdateExtension_ResponseSyntax)**

The actions defined in the extension.

Type: String to array of [Action](api-action.md) objects map

Map Entries: Maximum number of 5 items.

Valid Keys: `PRE_CREATE_HOSTED_CONFIGURATION_VERSION | PRE_START_DEPLOYMENT | AT_DEPLOYMENT_TICK | ON_DEPLOYMENT_START | ON_DEPLOYMENT_STEP | ON_DEPLOYMENT_BAKING | ON_DEPLOYMENT_COMPLETE | ON_DEPLOYMENT_ROLLED_BACK`

Array Members: Fixed number of 1 item.

**[Arn](#API_UpdateExtension_ResponseSyntax)**

The system-generated Amazon Resource Name (ARN) for the extension.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:(aws[a-zA-Z-]*)?:[a-z]+:((eusc-)?[a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1})?:(\d{12})?:[a-zA-Z0-9-_/:.]+`

**[Description](#API_UpdateExtension_ResponseSyntax)**

Information about the extension.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

**[Id](#API_UpdateExtension_ResponseSyntax)**

The system-generated ID of the extension.

Type: String

Pattern: `[a-z0-9]{4,7}`

**[Name](#API_UpdateExtension_ResponseSyntax)**

The extension name.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

**[Parameters](#API_UpdateExtension_ResponseSyntax)**

The parameters accepted by the extension. You specify parameter values when you
associate the extension to an AWS AppConfig resource by using the
`CreateExtensionAssociation` API action. For AWS Lambda extension
actions, these parameters are included in the Lambda request object.

Type: String to [Parameter](api-parameter.md) object map

Map Entries: Maximum number of 10 items.

Key Pattern: `^[^\/#:\n]{1,64}$`

**[VersionNumber](#API_UpdateExtension_ResponseSyntax)**

The extension version number.

Type: Integer

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**BadRequestException**

The input fails to satisfy the constraints specified by an AWS service.

**Details**

Detailed information about the input that failed to satisfy the constraints specified by
a call.

HTTP Status Code: 400

**ConflictException**

The request could not be processed because of conflict in the current state of the
resource.

HTTP Status Code: 409

**InternalServerException**

There was an internal failure in the AWS AppConfig service.

HTTP Status Code: 500

**ResourceNotFoundException**

The requested resource could not be found.

HTTP Status Code: 404

## Examples

### Example

This example illustrates one usage of UpdateExtension.

#### Sample Request

```

PATCH /extensions/my-test-extension HTTP/1.1
Host: appconfig.us-west-2.amazonaws.com
Accept-Encoding: identity
Content-Type: application/json
User-Agent: aws-cli/2.7.19 Python/3.9.11 Windows/10 exe/AMD64 prompt/off command/appconfig.update-extension
X-Amz-Date: 20220803T222153Z
Authorization: AWS4-HMAC-SHA256 Credential=AWS_ACCESS_KEY_ID_REDACTED/20220803/us-west-2/appconfig/aws4_request, SignedHeaders=content-type;host;x-amz-date, Signature=39c3b3042cd2aEXAMPLE
Content-Length: 90

{"Parameters": {"MyParamKey": {"Required": true}, "MyOtherParamKey": {"Required": false}}}
```

#### Sample Response

```

{
	"Actions": {
		"PRE_CREATE_HOSTED_CONFIGURATION_VERSION": [{
			"Description": null,
			"Name": "S3backup",
			"RoleArn": "arn:aws:iam::111122223333:role/mytestextensionrole",
			"Uri": "arn:aws:lambda:us-west-2:111122223333:function:mytestfunction"
		}]
	},
	"Arn": "arn:aws:appconfig:us-west-2:111122223333:extension/6czExample/1",
	"Description": null,
	"Id": "6czExample",
	"Name": "my-test-extension",
	"Parameters": {
		"MyOtherParamKey": {
			"Description": null,
			"Required": false
		},
		"MyParamKey": {
			"Description": null,
			"Required": true
		}
	},
	"VersionNumber": 1
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appconfig-2019-10-09/updateextension.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appconfig-2019-10-09/updateextension.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/updateextension.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appconfig-2019-10-09/updateextension.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/updateextension.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appconfig-2019-10-09/updateextension.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appconfig-2019-10-09/updateextension.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appconfig-2019-10-09/updateextension.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appconfig-2019-10-09/updateextension.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/updateextension.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateEnvironment

UpdateExtensionAssociation

All content copied from https://docs.aws.amazon.com/.
