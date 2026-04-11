# UpdateExtensionAssociation

Updates an association. For more information about extensions and associations, see
[Extending\
workflows](../../../../services/appconfig/latest/userguide/working-with-appconfig-extensions.md) in the _AWS AppConfig User Guide_.

## Request Syntax

```nohighlight

PATCH /extensionassociations/ExtensionAssociationId HTTP/1.1
Content-type: application/json

{
   "Parameters": {
      "string" : "string"
   }
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[ExtensionAssociationId](#API_UpdateExtensionAssociation_RequestSyntax)**

The system-generated ID for the association.

Pattern: `[a-z0-9]{4,7}`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[Parameters](#API_UpdateExtensionAssociation_RequestSyntax)**

The parameter names and values defined in the extension.

Type: String to string map

Map Entries: Minimum number of 0 items. Maximum number of 10 items.

Key Pattern: `^[^\/#:\n]{1,64}$`

Value Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
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

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Arn](#API_UpdateExtensionAssociation_ResponseSyntax)**

The system-generated Amazon Resource Name (ARN) for the extension.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:(aws[a-zA-Z-]*)?:[a-z]+:((eusc-)?[a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1})?:(\d{12})?:[a-zA-Z0-9-_/:.]+`

**[ExtensionArn](#API_UpdateExtensionAssociation_ResponseSyntax)**

The ARN of the extension defined in the association.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:(aws[a-zA-Z-]*)?:[a-z]+:((eusc-)?[a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1})?:(\d{12})?:[a-zA-Z0-9-_/:.]+`

**[ExtensionVersionNumber](#API_UpdateExtensionAssociation_ResponseSyntax)**

The version number for the extension defined in the association.

Type: Integer

**[Id](#API_UpdateExtensionAssociation_ResponseSyntax)**

The system-generated ID for the association.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

**[Parameters](#API_UpdateExtensionAssociation_ResponseSyntax)**

The parameter names and values defined in the association.

Type: String to string map

Map Entries: Minimum number of 0 items. Maximum number of 10 items.

Key Pattern: `^[^\/#:\n]{1,64}$`

Value Length Constraints: Minimum length of 1. Maximum length of 2048.

**[ResourceArn](#API_UpdateExtensionAssociation_ResponseSyntax)**

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

## Examples

### Example

This example illustrates one usage of UpdateExtensionAssociation.

#### Sample Request

```

PATCH /extensionassociations/rnekru4 HTTP/1.1
Host: appconfig.us-west-2.amazonaws.com
Accept-Encoding: identity
Content-Type: application/json
User-Agent: aws-cli/2.7.19 Python/3.9.11 Windows/10 exe/AMD64 prompt/off command/appconfig.update-extension-association
X-Amz-Date: 20220803T222640Z
Authorization: AWS4-HMAC-SHA256 Credential=AWS_ACCESS_KEY_ID_REDACTED/20220803/us-west-2/appconfig/aws4_request, SignedHeaders=content-type;host;x-amz-date, Signature=39c3b3042cd2aEXAMPLE
Content-Length: 53

{"Parameters": {"MyParamKey": "DifferentParamValue"}}
```

#### Sample Response

```

{
	"Arn": null,
	"ExtensionArn": "arn:aws:appconfig:us-west-2:111122223333:extension/6czExample/1",
	"ExtensionVersionNumber": 1,
	"Id": "rnekru4",
	"Parameters": {
		"MyParamKey": "DifferentParamValue"
	},
	"ResourceArn": "arn:aws:appconfig:us-west-2:111122223333:application/xlmExample"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appconfig-2019-10-09/updateextensionassociation.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appconfig-2019-10-09/updateextensionassociation.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/updateextensionassociation.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appconfig-2019-10-09/updateextensionassociation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/updateextensionassociation.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appconfig-2019-10-09/updateextensionassociation.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appconfig-2019-10-09/updateextensionassociation.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appconfig-2019-10-09/updateextensionassociation.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appconfig-2019-10-09/updateextensionassociation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/updateextensionassociation.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateExtension

ValidateConfiguration

All content copied from https://docs.aws.amazon.com/.
