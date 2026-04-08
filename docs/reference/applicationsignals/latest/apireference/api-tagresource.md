# TagResource

Assigns one or more tags (key-value pairs) to the specified CloudWatch resource, such as a service level objective.

Tags can help you organize and categorize your resources. You can also use them to scope user
permissions by granting a user
permission to access or change only resources with certain tag values.

Tags don't have any semantic meaning to AWS and are interpreted strictly as strings of characters.

You can use the `TagResource` action with an alarm that already has tags. If you specify a new tag key for the alarm,
this tag is appended to the list of tags associated
with the alarm. If you specify a tag key that is already associated with the alarm, the new tag value that you specify replaces
the previous value for that tag.

You can associate as many as 50 tags with a CloudWatch resource.

## Request Syntax

```nohighlight

POST /tag-resource HTTP/1.1
Content-type: application/json

{
   "ResourceArn": "string",
   "Tags": [
      {
         "Key": "string",
         "Value": "string"
      }
   ]
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[ResourceArn](#API_TagResource_RequestSyntax)**

The Amazon Resource Name (ARN) of the CloudWatch resource that you want to set tags for.

The ARN format of an Application Signals SLO is
`arn:aws:cloudwatch:Region:account-id:slo:slo-name
                  `

For more information about ARN format, see [Resource\
Types Defined by Amazon CloudWatch](../../../../services/iam/latest/userguide/list-amazoncloudwatch-amazoncloudwatch-resources-for-iam-policies.md) in the _Amazon Web Services General_
_Reference_.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: Yes

**[Tags](#API_TagResource_RequestSyntax)**

The list of key-value pairs to associate with the alarm.

Type: Array of [Tag](api-tag.md) objects

Array Members: Minimum number of 0 items. Maximum number of 200 items.

Required: Yes

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ResourceNotFoundException**

Resource not found.

**ResourceId**

Can't find the resource id.

**ResourceType**

The resource type is not valid.

HTTP Status Code: 404

**ServiceQuotaExceededException**

This request exceeds a service quota.

HTTP Status Code: 402

**ThrottlingException**

The request was throttled because of quota limits.

HTTP Status Code: 429

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/application-signals-2024-04-15/tagresource.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/application-signals-2024-04-15/tagresource.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/application-signals-2024-04-15/tagresource.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/application-signals-2024-04-15/tagresource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/application-signals-2024-04-15/tagresource.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/application-signals-2024-04-15/tagresource.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/application-signals-2024-04-15/tagresource.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/application-signals-2024-04-15/tagresource.md)

- [AWS SDK for Python](../../../../services/goto/boto3/application-signals-2024-04-15/tagresource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/application-signals-2024-04-15/tagresource.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

StartDiscovery

UntagResource
