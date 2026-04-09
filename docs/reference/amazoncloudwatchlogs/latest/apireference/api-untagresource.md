# UntagResource

Removes one or more tags from the specified resource.

## Request Syntax

```nohighlight

{
   "resourceArn": "string",
   "tagKeys": [ "string" ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[resourceArn](#API_UntagResource_RequestSyntax)**

The ARN of the CloudWatch Logs resource that you're removing tags from.

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

**[tagKeys](#API_UntagResource_RequestSyntax)**

The list of tag keys to remove from the resource.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 50 items.

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]+)$`

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/untagresource.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/untagresource.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/untagresource.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/untagresource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/untagresource.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/untagresource.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/untagresource.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/untagresource.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/untagresource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/untagresource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UntagLogGroup

UpdateAnomaly

All content copied from https://docs.aws.amazon.com/.
