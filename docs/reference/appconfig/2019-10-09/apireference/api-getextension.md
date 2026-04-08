# GetExtension

Returns information about an AWS AppConfig extension.

## Request Syntax

```nohighlight

GET /extensions/ExtensionIdentifier?version_number=VersionNumber HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[ExtensionIdentifier](#API_GetExtension_RequestSyntax)**

The name, the ID, or the Amazon Resource Name (ARN) of the extension.

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: Yes

**[VersionNumber](#API_GetExtension_RequestSyntax)**

The extension version number. If no version number was defined, AWS AppConfig uses
the highest version.

## Request Body

The request does not have a request body.

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

**[Actions](#API_GetExtension_ResponseSyntax)**

The actions defined in the extension.

Type: String to array of [Action](api-action.md) objects map

Map Entries: Maximum number of 5 items.

Valid Keys: `PRE_CREATE_HOSTED_CONFIGURATION_VERSION | PRE_START_DEPLOYMENT | AT_DEPLOYMENT_TICK | ON_DEPLOYMENT_START | ON_DEPLOYMENT_STEP | ON_DEPLOYMENT_BAKING | ON_DEPLOYMENT_COMPLETE | ON_DEPLOYMENT_ROLLED_BACK`

Array Members: Fixed number of 1 item.

**[Arn](#API_GetExtension_ResponseSyntax)**

The system-generated Amazon Resource Name (ARN) for the extension.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:(aws[a-zA-Z-]*)?:[a-z]+:((eusc-)?[a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1})?:(\d{12})?:[a-zA-Z0-9-_/:.]+`

**[Description](#API_GetExtension_ResponseSyntax)**

Information about the extension.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

**[Id](#API_GetExtension_ResponseSyntax)**

The system-generated ID of the extension.

Type: String

Pattern: `[a-z0-9]{4,7}`

**[Name](#API_GetExtension_ResponseSyntax)**

The extension name.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

**[Parameters](#API_GetExtension_ResponseSyntax)**

The parameters accepted by the extension. You specify parameter values when you
associate the extension to an AWS AppConfig resource by using the
`CreateExtensionAssociation` API action. For AWS Lambda extension
actions, these parameters are included in the Lambda request object.

Type: String to [Parameter](api-parameter.md) object map

Map Entries: Maximum number of 10 items.

Key Pattern: `^[^\/#:\n]{1,64}$`

**[VersionNumber](#API_GetExtension_ResponseSyntax)**

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

**InternalServerException**

There was an internal failure in the AWS AppConfig service.

HTTP Status Code: 500

**ResourceNotFoundException**

The requested resource could not be found.

HTTP Status Code: 404

## Examples

### Example

This example illustrates one usage of GetExtension.

#### Sample Request

```

GET /extensions/my-test-extension HTTP/1.1
Host: appconfig.us-west-2.amazonaws.com
Accept-Encoding: identity
User-Agent: aws-cli/2.7.19 Python/3.9.11 Windows/10 exe/AMD64 prompt/off command/appconfig.get-extension
X-Amz-Date: 20220803T220925Z
Authorization: AWS4-HMAC-SHA256 Credential=AKIAIOSFODNN7EXAMPLE/20220803/us-west-2/appconfig/aws4_request, SignedHeaders=host;x-amz-date, Signature=39c3b3042cd2aEXAMPLE
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
	}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appconfig-2019-10-09/getextension.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appconfig-2019-10-09/getextension.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/getextension.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appconfig-2019-10-09/getextension.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/getextension.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appconfig-2019-10-09/getextension.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appconfig-2019-10-09/getextension.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appconfig-2019-10-09/getextension.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appconfig-2019-10-09/getextension.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/getextension.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetEnvironment

GetExtensionAssociation
