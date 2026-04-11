# DeleteExtensionAssociation

Deletes an extension association. This action doesn't delete extensions defined in the
association.

## Request Syntax

```nohighlight

DELETE /extensionassociations/ExtensionAssociationId HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[ExtensionAssociationId](#API_DeleteExtensionAssociation_RequestSyntax)**

The ID of the extension association to delete.

Pattern: `[a-z0-9]{4,7}`

Required: Yes

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

This example illustrates one usage of DeleteExtensionAssociation.

#### Sample Request

```

DELETE /extensionassociations/rnekru4 HTTP/1.1
Host: appconfig.us-west-2.amazonaws.com
Accept-Encoding: identity
User-Agent: aws-cli/2.7.19 Python/3.9.11 Windows/10 exe/AMD64 prompt/off command/appconfig.delete-extension-association
X-Amz-Date: 20220803T223105Z
Authorization: AWS4-HMAC-SHA256 Credential=AWS_ACCESS_KEY_ID_REDACTED/20220803/us-west-2/appconfig/aws4_request, SignedHeaders=host;x-amz-date, Signature=39c3b3042cd2aEXAMPLE
Content-Length: 0
```

#### Sample Response

```

{}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appconfig-2019-10-09/deleteextensionassociation.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appconfig-2019-10-09/deleteextensionassociation.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/deleteextensionassociation.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appconfig-2019-10-09/deleteextensionassociation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/deleteextensionassociation.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appconfig-2019-10-09/deleteextensionassociation.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appconfig-2019-10-09/deleteextensionassociation.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appconfig-2019-10-09/deleteextensionassociation.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appconfig-2019-10-09/deleteextensionassociation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/deleteextensionassociation.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteExtension

DeleteHostedConfigurationVersion

All content copied from https://docs.aws.amazon.com/.
