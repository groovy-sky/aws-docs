# DeleteExtension

Deletes an AWS AppConfig extension. You must delete all associations to an
extension before you delete the extension.

## Request Syntax

```nohighlight

DELETE /extensions/ExtensionIdentifier?version=VersionNumber HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[ExtensionIdentifier](#API_DeleteExtension_RequestSyntax)**

The name, ID, or Amazon Resource Name (ARN) of the extension you want to delete.

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: Yes

**[VersionNumber](#API_DeleteExtension_RequestSyntax)**

A specific version of an extension to delete. If omitted, the highest version is
deleted.

## Request Body

The request does not have a request body.

## Response Syntax

```

HTTP/1.1 204

```

## Response Elements

If the action is successful, the service sends back an HTTP 204 response with an empty HTTP body.

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

This example illustrates one usage of DeleteExtension.

#### Sample Request

```

DELETE /extensions/my-test-extension HTTP/1.1
Host: appconfig.us-west-2.amazonaws.com
Accept-Encoding: identity
User-Agent: aws-cli/2.7.19 Python/3.9.11 Windows/10 exe/AMD64 prompt/off command/appconfig.delete-extension
X-Amz-Date: 20220803T223549Z
Authorization: AWS4-HMAC-SHA256 Credential=AKIAIOSFODNN7EXAMPLE/20220803/us-west-2/appconfig/aws4_request, SignedHeaders=host;x-amz-date, Signature=39c3b3042cd2aEXAMPLE
Content-Length: 0
```

#### Sample Response

```

{}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appconfig-2019-10-09/deleteextension.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appconfig-2019-10-09/deleteextension.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/deleteextension.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appconfig-2019-10-09/deleteextension.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/deleteextension.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appconfig-2019-10-09/deleteextension.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appconfig-2019-10-09/deleteextension.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appconfig-2019-10-09/deleteextension.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appconfig-2019-10-09/deleteextension.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/deleteextension.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteEnvironment

DeleteExtensionAssociation
