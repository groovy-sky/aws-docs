# ListTagsLogGroup

_This action has been deprecated._

###### Important

The ListTagsLogGroup operation is on the path to deprecation. We recommend that you use
[ListTagsForResource](api-listtagsforresource.md) instead.

Lists the tags for the specified log group.

## Request Syntax

```nohighlight

{
   "logGroupName": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[logGroupName](#API_ListTagsLogGroup_RequestSyntax)**

The name of the log group.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `[\.\-_/#A-Za-z0-9]+`

Required: Yes

## Response Syntax

```nohighlight

{
   "tags": {
      "string" : "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[tags](#API_ListTagsLogGroup_ResponseSyntax)**

The tags for the log group.

Type: String to string map

Map Entries: Maximum number of 50 items.

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Key Pattern: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]+)$`

Value Length Constraints: Maximum length of 256.

Value Pattern: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ResourceNotFoundException**

The specified resource does not exist.

HTTP Status Code: 400

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/listtagsloggroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/listtagsloggroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/listtagsloggroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/listtagsloggroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/listtagsloggroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/listtagsloggroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/listtagsloggroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/listtagsloggroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/listtagsloggroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/listtagsloggroup.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListTagsForResource

PutAccountPolicy
