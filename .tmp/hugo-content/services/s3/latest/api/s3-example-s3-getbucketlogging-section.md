# Use `GetBucketLogging` with a CLI

The following code examples show how to use `GetBucketLogging`.

CLI

**AWS CLI**

**To retrieve the logging status for a bucket**

The following `get-bucket-logging` example retrieves the logging status for the specified bucket.

```nohighlight

aws s3api get-bucket-logging \
    --bucket amzn-s3-demo-bucket

```

Output:

```nohighlight

{
    "LoggingEnabled": {
        "TargetPrefix": "",
        "TargetBucket": "amzn-s3-demo-bucket-logs"
          }
}
```

- For API details, see
[GetBucketLogging](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/s3api/get-bucket-logging.html)
in _AWS CLI Command Reference_.

PowerShell

**Tools for PowerShell V4**

**Example 1: This command returns the logging status for the specified bucket.**

```powershell

Get-S3BucketLogging -BucketName 'amzn-s3-demo-bucket'

```

**Output:**

```nohighlight

TargetBucketName   Grants TargetPrefix
----------------   ------ ------------
testbucket1        {}     testprefix
```

- For API details, see
[GetBucketLogging](../../../powershell/v4/reference.md)
in _AWS Tools for PowerShell Cmdlet Reference (V4)_.

**Tools for PowerShell V5**

**Example 1: This command returns the logging status for the specified bucket.**

```powershell

Get-S3BucketLogging -BucketName 'amzn-s3-demo-bucket'

```

**Output:**

```nohighlight

TargetBucketName   Grants TargetPrefix
----------------   ------ ------------
testbucket1        {}     testprefix
```

- For API details, see
[GetBucketLogging](../../../powershell/v5/reference.md)
in _AWS Tools for PowerShell Cmdlet Reference (V5)_.

For a complete list of AWS SDK developer guides and code examples, see
[Developing with Amazon S3 using the AWS SDKs](../../../../reference/amazons3/latest/api/sdk-general-information-section.md).
This topic also includes information about getting started and details about previous SDK versions.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetBucketLocation

GetBucketMetricsConfiguration

All content copied from https://docs.aws.amazon.com/.
