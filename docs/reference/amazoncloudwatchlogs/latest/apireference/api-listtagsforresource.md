# ListTagsForResource

Displays the tags associated with a CloudWatch Logs resource. Currently, log groups and
destinations support tagging.

## Request Syntax

```nohighlight

{
   "resourceArn": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[resourceArn](#API_ListTagsForResource_RequestSyntax)**

The ARN of the resource that you want to view tags for.

The ARN format of a log group is
`arn:aws:logs:Region:account-id:log-group:log-group-name
                  `

The ARN format of a destination is
`arn:aws:logs:Region:account-id:destination:destination-name
                  `

For more information about ARN format, see [CloudWatch Logs\
resources and operations](../../../../services/amazoncloudwatch/latest/logs/iam-access-control-overview-cwl.md).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `[\w+=/:,.@-]*`

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

**[tags](#API_ListTagsForResource_ResponseSyntax)**

The list of tags associated with the requested resource.>

Type: String to string map

Map Entries: Maximum number of 50 items.

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Key Pattern: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]+)$`

Value Length Constraints: Maximum length of 256.

Value Pattern: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

A parameter is specified incorrectly.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource does not exist.

HTTP Status Code: 400

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/listtagsforresource.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/listtagsforresource.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/listtagsforresource.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/listtagsforresource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/listtagsforresource.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/listtagsforresource.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/listtagsforresource.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/listtagsforresource.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/listtagsforresource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/listtagsforresource.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListSourcesForS3TableIntegration

ListTagsLogGroup
