# ListExtensionAssociations

Lists all AWS AppConfig extension associations in the account. For more
information about extensions and associations, see [Extending\
workflows](../../../../services/appconfig/latest/userguide/working-with-appconfig-extensions.md) in the _AWS AppConfig User Guide_.

## Request Syntax

```nohighlight

GET /extensionassociations?extension_identifier=ExtensionIdentifier&extension_version_number=ExtensionVersionNumber&max_results=MaxResults&next_token=NextToken&resource_identifier=ResourceIdentifier HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[ExtensionIdentifier](#API_ListExtensionAssociations_RequestSyntax)**

The name, the ID, or the Amazon Resource Name (ARN) of the extension.

Length Constraints: Minimum length of 1. Maximum length of 2048.

**[ExtensionVersionNumber](#API_ListExtensionAssociations_RequestSyntax)**

The version number for the extension defined in the association.

**[MaxResults](#API_ListExtensionAssociations_RequestSyntax)**

The maximum number of items to return for this call. The call also returns a token that
you can specify in a subsequent call to get the next set of results.

Valid Range: Minimum value of 1. Maximum value of 50.

**[NextToken](#API_ListExtensionAssociations_RequestSyntax)**

A token to start the list. Use this token to get the next set of results or pass null to
get the first set of results.

Length Constraints: Minimum length of 1. Maximum length of 2048.

**[ResourceIdentifier](#API_ListExtensionAssociations_RequestSyntax)**

The ARN of an application, configuration profile, or environment.

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:(aws[a-zA-Z-]*)?:[a-z]+:((eusc-)?[a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1})?:(\d{12})?:[a-zA-Z0-9-_/:.]+`

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "Items": [
      {
         "ExtensionArn": "string",
         "Id": "string",
         "ResourceArn": "string"
      }
   ],
   "NextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Items](#API_ListExtensionAssociations_ResponseSyntax)**

The list of extension associations. Each item represents an extension association to an
application, environment, or configuration profile.

Type: Array of [ExtensionAssociationSummary](api-extensionassociationsummary.md) objects

**[NextToken](#API_ListExtensionAssociations_ResponseSyntax)**

The token for the next set of items to return. Use this token to get the next set of
results.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

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

## Examples

### Example

This example illustrates one usage of ListExtensionAssociations.

#### Sample Request

```

GET /extensionassociations HTTP/1.1
Host: appconfig.us-west-2.amazonaws.com
Accept-Encoding: identity
User-Agent: aws-cli/2.7.19 Python/3.9.11 Windows/10 exe/AMD64 prompt/off command/appconfig.list-extension-associations
X-Amz-Date: 20220803T215900Z
Authorization: AWS4-HMAC-SHA256 Credential=AKIAIOSFODNN7EXAMPLE/20220803/us-west-2/appconfig/aws4_request, SignedHeaders=host;x-amz-date, Signature=39c3b3042cd2aEXAMPLE
```

#### Sample Response

```

{
	"Items": [{
		"ExtensionArn": "arn:aws:appconfig:us-west-2:111122223333:extension/6czExample/1",
		"Id": "rnekru4",
		"ResourceArn": "arn:aws:appconfig:us-west-2:111122223333:application/xlmtnms"
	}],
	"NextToken": null
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appconfig-2019-10-09/listextensionassociations.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appconfig-2019-10-09/listextensionassociations.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/listextensionassociations.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appconfig-2019-10-09/listextensionassociations.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/listextensionassociations.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appconfig-2019-10-09/listextensionassociations.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appconfig-2019-10-09/listextensionassociations.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appconfig-2019-10-09/listextensionassociations.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appconfig-2019-10-09/listextensionassociations.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/listextensionassociations.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListEnvironments

ListExtensions
