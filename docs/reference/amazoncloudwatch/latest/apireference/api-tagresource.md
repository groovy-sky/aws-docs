# TagResource

Assigns one or more tags (key-value pairs) to the specified CloudWatch resource.
Currently, the only CloudWatch resources that can be tagged are alarms and Contributor
Insights rules.

Tags can help you organize and categorize your resources. You can also use them to
scope user permissions by granting a user permission to access or change only resources
with certain tag values.

Tags don't have any semantic meaning to AWS and are interpreted
strictly as strings of characters.

You can use the `TagResource` action with an alarm that already has tags.
If you specify a new tag key for the alarm, this tag is appended to the list of tags
associated with the alarm. If you specify a tag key that is already associated with the
alarm, the new tag value that you specify replaces the previous value for that
tag.

You can associate as many as 50 tags with a CloudWatch resource.

## Request Parameters

**ResourceARN**

The ARN of the CloudWatch resource that you're adding tags to.

The ARN format of an alarm is
`arn:aws:cloudwatch:Region:account-id:alarm:alarm-name
                  `

The ARN format of a Contributor Insights rule is
`arn:aws:cloudwatch:Region:account-id:insight-rule/insight-rule-name
                  `

For more information about ARN format, see [Resource Types Defined by Amazon CloudWatch](../../../../services/iam/latest/userguide/list-amazoncloudwatch-amazoncloudwatch-resources-for-iam-policies.md) in the _Amazon Web_
_Services General Reference_.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: Yes

**Tags**

The list of key-value pairs to associate with the alarm.

Type: Array of [Tag](api-tag.md) objects

Required: Yes

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ConcurrentModificationException**

More than one process tried to modify a resource at the same time.

HTTP Status Code: 429

**ConflictException**

This operation attempted to create a resource that already exists.

HTTP Status Code: 409

**InternalServiceError**

Request processing has failed due to some unknown error, exception, or
failure.

**Message**

HTTP Status Code: 500

**InvalidParameterValue**

The value of an input parameter is bad or out-of-range.

**message**

HTTP Status Code: 400

**ResourceNotFoundException**

The named resource does not exist.

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/monitoring-2010-08-01/tagresource.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/monitoring-2010-08-01/tagresource.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/monitoring-2010-08-01/tagresource.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/monitoring-2010-08-01/tagresource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/monitoring-2010-08-01/tagresource.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/monitoring-2010-08-01/tagresource.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/monitoring-2010-08-01/tagresource.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/monitoring-2010-08-01/tagresource.md)

- [AWS SDK for Python](../../../../services/goto/boto3/monitoring-2010-08-01/tagresource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/monitoring-2010-08-01/tagresource.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

StopMetricStreams

UntagResource
