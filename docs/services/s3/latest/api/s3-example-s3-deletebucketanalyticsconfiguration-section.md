# Use `DeleteBucketAnalyticsConfiguration` with a CLI

The following code examples show how to use `DeleteBucketAnalyticsConfiguration`.

CLI

**AWS CLI**

**To delete an analytics configuration for a bucket**

The following `delete-bucket-analytics-configuration` example removes the analytics configuration for the specified bucket and ID.

```nohighlight

aws s3api delete-bucket-analytics-configuration \
    --bucket amzn-s3-demo-bucket \
    --id 1

```

This command produces no output.

- For API details, see
[DeleteBucketAnalyticsConfiguration](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/s3api/delete-bucket-analytics-configuration.html)
in _AWS CLI Command Reference_.

PowerShell

**Tools for PowerShell V4**

**Example 1: The command removes the analytics filter with name 'testfilter' in the given S3 bucket.**

```powershell

Remove-S3BucketAnalyticsConfiguration -BucketName 'amzn-s3-demo-bucket' -AnalyticsId 'testfilter'

```

- For API details, see
[DeleteBucketAnalyticsConfiguration](../../../powershell/v4/reference.md)
in _AWS Tools for PowerShell Cmdlet Reference (V4)_.

**Tools for PowerShell V5**

**Example 1: The command removes the analytics filter with name 'testfilter' in the given S3 bucket.**

```powershell

Remove-S3BucketAnalyticsConfiguration -BucketName 'amzn-s3-demo-bucket' -AnalyticsId 'testfilter'

```

- For API details, see
[DeleteBucketAnalyticsConfiguration](../../../powershell/v5/reference.md)
in _AWS Tools for PowerShell Cmdlet Reference (V5)_.

For a complete list of AWS SDK developer guides and code examples, see
[Developing with Amazon S3 using the AWS SDKs](../../../../reference/amazons3/latest/api/sdk-general-information-section.md).
This topic also includes information about getting started and details about previous SDK versions.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteBucket

DeleteBucketCors

All content copied from https://docs.aws.amazon.com/.
