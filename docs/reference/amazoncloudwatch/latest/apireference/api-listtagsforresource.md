---
title: "ListTagsForResource"
---

# ListTagsForResource

Displays the tags associated with a CloudWatch resource. Currently, alarms,
dashboards, metric streams and Contributor Insights rules support tagging.

## Request Parameters

**ResourceARN**

The ARN of the CloudWatch resource that you want to view tags for.

The ARN format of an alarm is
`arn:aws:cloudwatch:Region:account-id:alarm:alarm-name
                  `

The ARN format of a Contributor Insights rule is
`arn:aws:cloudwatch:Region:account-id:insight-rule/insight-rule-name
                  `

The ARN format of a dashboard is
`arn:aws:cloudwatch::account-id:dashboard/dashboard-name
                  `

The ARN format of a metric stream is
`arn:aws:cloudwatch:Region:account-id:metric-stream/metric-stream-name
                  `

For more information about ARN format, see [Resource Types Defined by Amazon CloudWatch](../../../../services/iam/latest/userguide/list-amazoncloudwatch.md#amazoncloudwatch-resources-for-iam-policies) in the _Amazon Web_
_Services General Reference_.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: Yes

## Response Elements

The following element is returned by the service.

**Tags**

The list of tag keys and values associated with the resource you specified.

Type: Array of [Tag](api-tag.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/monitoring-2010-08-01/ListTagsForResource)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/monitoring-2010-08-01/ListTagsForResource)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/monitoring-2010-08-01/ListTagsForResource)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/monitoring-2010-08-01/ListTagsForResource)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/monitoring-2010-08-01/ListTagsForResource)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/monitoring-2010-08-01/ListTagsForResource)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/monitoring-2010-08-01/ListTagsForResource)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/monitoring-2010-08-01/ListTagsForResource)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/monitoring-2010-08-01/ListTagsForResource)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/monitoring-2010-08-01/ListTagsForResource)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListMetricStreams

PutAlarmMuteRule

All content copied from https://docs.aws.amazon.com/.
