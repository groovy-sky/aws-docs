# ListPluginTypeMetadata

Lists metadata for all Amazon Q Business plugin types.

## Request Syntax

```nohighlight

GET /pluginTypeMetadata?maxResults=maxResults&nextToken=nextToken HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[maxResults](#API_ListPluginTypeMetadata_RequestSyntax)**

The maximum number of plugin metadata items to return.

Valid Range: Minimum value of 1. Maximum value of 50.

**[nextToken](#API_ListPluginTypeMetadata_RequestSyntax)**

If the metadata returned exceeds `maxResults`, Amazon Q Business
returns a next token as a pagination token to retrieve the next set of metadata.

Length Constraints: Minimum length of 1. Maximum length of 800.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "items": [
      {
         "category": "string",
         "description": "string",
         "type": "string"
      }
   ],
   "nextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[items](#API_ListPluginTypeMetadata_ResponseSyntax)**

An array of information on plugin metadata.

Type: Array of [PluginTypeMetadataSummary](api-plugintypemetadatasummary.md) objects

**[nextToken](#API_ListPluginTypeMetadata_ResponseSyntax)**

If the response is truncated, Amazon Q Business returns this token, which you
can use in a later request to list the next set of plugin metadata.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 800.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You don't have access to perform this action. Make sure you have the required
permission policies and user accounts and try again.

HTTP Status Code: 403

**InternalServerException**

An issue occurred with the internal server used for your Amazon Q Business service. Wait
some minutes and try again, or contact [Support](http://aws.amazon.com/contact-us) for help.

HTTP Status Code: 500

**ThrottlingException**

The request was denied due to throttling. Reduce the number of requests and try
again.

HTTP Status Code: 429

**ValidationException**

The input doesn't meet the constraints set by the Amazon Q Business service. Provide the
correct input and try again.

**fields**

The input field(s) that failed validation.

**message**

The message describing the `ValidationException`.

**reason**

The reason for the `ValidationException`.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/qbusiness-2023-11-27/listplugintypemetadata.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/qbusiness-2023-11-27/listplugintypemetadata.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/listplugintypemetadata.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/qbusiness-2023-11-27/listplugintypemetadata.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/listplugintypemetadata.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/qbusiness-2023-11-27/listplugintypemetadata.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/qbusiness-2023-11-27/listplugintypemetadata.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/qbusiness-2023-11-27/listplugintypemetadata.md)

- [AWS SDK for Python](../../../goto/boto3/qbusiness-2023-11-27/listplugintypemetadata.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/listplugintypemetadata.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListPluginTypeActions

ListRetrievers

All content copied from https://docs.aws.amazon.com/.
