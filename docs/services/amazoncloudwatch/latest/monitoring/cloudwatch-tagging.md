# Tagging your Amazon CloudWatch resources

A _tag_ is a custom attribute label that you or AWS assigns to an AWS resource.

Tags help you do the following:

- Identify and organize your AWS resources.
Many AWS services support tagging.
You can assign the same tag to resources from different services to indicate the resources are related.
For example, you can assign the tag that you assigned to a CloudWatch rule to an EC2 instance.

Tags have two parts:

- A _tag key_ (for example, `CostCenter`, `Environment`, or `Project`).
Tag keys are case sensitive.

- An optional field known as a _tag value_ (for example, `111122223333` or `Production`).
Omitting the tag value is the same as using an empty string.
Like tag keys, tag values are case sensitive.

The following sections provide more information about tags for CloudWatch.

## Supported resources in CloudWatch

The following resources in CloudWatch support tagging:

- Alarms – You can tag alarms using the [tag-resource](https://docs.aws.amazon.com/cli/latest/reference/cloudwatch/tag-resource.html) AWS CLI command and the
[TagResource](../../../../reference/amazoncloudwatch/latest/apireference/api-tagresource.md) API. You can also view and manage your
alarm tags using the _Alarms_ details page in the CloudWatch console.

- Canaries – You can tag canaries using the CloudWatch console. For more
information, see [Creating a canary](cloudwatch-synthetics-canaries-create.md).

- Contributor Insights rules – You can tag Contributor Insights rules when
you create them by using the [put-insight-rule](https://docs.aws.amazon.com/cli/latest/reference/cloudwatch/put-insight-rule.html) AWS CLI command and the
[PutInsightRule](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutInsightRule.html) API.
You can add tags to existing rules by using the [tag-resource](https://docs.aws.amazon.com/cli/latest/reference/cloudwatch/tag-resource.html) AWS CLI command and the
[TagResource](../../../../reference/amazoncloudwatch/latest/apireference/api-tagresource.md) API.

- Metric streams – You can tag metric streams when
you create them by using the [put-metric-stream](https://docs.aws.amazon.com/cli/latest/reference/cloudwatch/put-metric-stream.html) AWS CLI command and the
[PutMetricStream](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutMetricStream.html) API.
You can add tags to existing metric streams by using the [tag-resource](https://docs.aws.amazon.com/cli/latest/reference/cloudwatch/tag-resource.html) AWS CLI command and the
[TagResource](../../../../reference/amazoncloudwatch/latest/apireference/api-tagresource.md) API.

For information about adding and managing tags, see [Managing tags](#CloudWatch-tagging-add-edit-delete).

## Managing tags

Tags consist of the `Key` and `Value` properties on a resource. You
can use the CloudWatch console, the AWS CLI, or the CloudWatch API to add, edit, or delete the values
for these properties. For information about working with tags, see the following:

- [TagResource](../../../../reference/amazoncloudwatch/latest/apireference/api-tagresource.md), [UntagResource](../../../../reference/amazoncloudwatch/latest/apireference/api-untagresource.md), and [ListTagsForResource](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_ListTagsForResource.html)
in the _Amazon CloudWatch API Reference_

- [tag-resource](https://docs.aws.amazon.com/cli/latest/reference/cloudwatch/tag-resource.html), [untag-resource](https://docs.aws.amazon.com/cli/latest/reference/cloudwatch/untag-resource.html), and [list-tags-for-resource](https://docs.aws.amazon.com/cli/latest/reference/cloudwatch/list-tags-for-resource.html) in the
_Amazon CloudWatch CLI Reference_

- [Working with Tag Editor](https://docs.aws.amazon.com/ARG/latest/userguide/tag-editor.html) in the
_Resource Groups User Guide_

## Tag naming and usage conventions

The following basic naming and usage conventions apply to using tags with
CloudWatch resources:

- Each resource can have a maximum of 50 tags.

- For each resource, each tag key must be unique, and each tag key can have only one value.

- The maximum tag key length is 128 Unicode characters in UTF-8.

- The maximum tag value length is 256 Unicode characters in UTF-8.

- Allowed characters are letters, numbers, spaces representable in UTF-8, and the following
characters: **_. : \+ = @ \_ /_**
**_-_** (hyphen).

- Tag keys and values are case sensitive. As a best practice, decide on a strategy for
capitalizing tags and consistently implement that strategy across all resource
types. For example, decide whether to use `Costcenter`,
`costcenter`, or `CostCenter` and use the same
convention for all tags. Avoid using similar tags with inconsistent case
treatment.

- The `aws:` prefix is prohibited for tags because it's reserved for AWS use.
You can't edit or delete tag keys or values with this prefix. Tags with this
prefix don't count against your tags per resource limit.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Logging API and console operations with AWS CloudTrail

Grafana integration
