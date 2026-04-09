# ListTagsForResource

Retrieves the tags that are associated with a specified flow.

## Request Syntax

```nohighlight

GET /tags/resourceArn HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[resourceArn](#API_ListTagsForResource_RequestSyntax)**

The Amazon Resource Name (ARN) of the specified flow.

Length Constraints: Maximum length of 512.

Pattern: `arn:aws:.*:.*:[0-9]+:.*`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "tags": {
      "string" : "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[tags](#API_ListTagsForResource_ResponseSyntax)**

The tags used to organize, track, or control access for your flow.

Type: String to string map

Map Entries: Minimum number of 0 items. Maximum number of 50 items.

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Key Pattern: `^(?!aws:)[a-zA-Z+-=._:/]+$`

Value Length Constraints: Maximum length of 256.

Value Pattern: `[\s\w+-=\.:/@]*`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServerException**

An internal service error occurred during the processing of your request. Try again
later.

HTTP Status Code: 500

**ResourceNotFoundException**

The resource specified in the request (such as the source or destination connector
profile) is not found.

HTTP Status Code: 404

**ValidationException**

The request has invalid or missing parameters.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appflow-2020-08-23/listtagsforresource.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appflow-2020-08-23/listtagsforresource.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/listtagsforresource.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appflow-2020-08-23/listtagsforresource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/listtagsforresource.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appflow-2020-08-23/listtagsforresource.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appflow-2020-08-23/listtagsforresource.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appflow-2020-08-23/listtagsforresource.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appflow-2020-08-23/listtagsforresource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/listtagsforresource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListFlows

RegisterConnector

All content copied from https://docs.aws.amazon.com/.
