# Enabling log file integrity validation for CloudTrail

You can enable log file integrity validation by using the AWS Management Console, AWS Command Line
Interface (AWS CLI), or CloudTrail API. CloudTrail starts delivering digest files in about an hour.

## AWS Management Console

To enable log file integrity validation with the CloudTrail console, choose
**Yes** for the **Enable log file validation**
option when you create or update a trail. By default, this feature is enabled for new
trails. For more information, see [Creating and updating a trail with the console](cloudtrail-create-and-update-a-trail-by-using-the-console.md).

## AWS CLI

To enable log file integrity validation with the AWS CLI, use the
`--enable-log-file-validation` option with the [create-trail](../../../cli/latest/reference/cloudtrail/create-trail.md) or [update-trail](../../../cli/latest/reference/cloudtrail/update-trail.md) commands. To
disable log file integrity validation, use the
`--no-enable-log-file-validation` option.

**Example**

The following `update-trail` command enables log file validation and starts
delivering digest files to the Amazon S3 bucket for the specified trail.

```nohighlight

aws cloudtrail update-trail --name your-trail-name --enable-log-file-validation
```

## CloudTrail API

To enable log file integrity validation with the CloudTrail API, set the
`EnableLogFileValidation` request parameter to `true` when calling
`CreateTrail` or `UpdateTrail`.

For more information, see [CreateTrail](../../../../reference/awscloudtrail/latest/apireference/api-createtrail.md) and [UpdateTrail](../../../../reference/awscloudtrail/latest/apireference/api-updatetrail.md) in the [AWS CloudTrail API Reference](../../../../reference/awscloudtrail/latest/apireference.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Validating CloudTrail log file integrity

Validating CloudTrail log file integrity with the AWS CLI

All content copied from https://docs.aws.amazon.com/.
