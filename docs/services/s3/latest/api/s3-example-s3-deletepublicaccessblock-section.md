# Use `DeletePublicAccessBlock` with a CLI

The following code examples show how to use `DeletePublicAccessBlock`.

CLI

**AWS CLI**

**To delete the block public access configuration for a bucket**

The following `delete-public-access-block` example removes the block public access configuration on the specified bucket.

```nohighlight

aws s3api delete-public-access-block \
    --bucket amzn-s3-demo-bucket

```

This command produces no output.

- For API details, see
[DeletePublicAccessBlock](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/s3api/delete-public-access-block.html)
in _AWS CLI Command Reference_.

PowerShell

**Tools for PowerShell V4**

**Example 1: This command turns off the block public access setting for the given bucket.**

```powershell

Remove-S3PublicAccessBlock -BucketName 'amzn-s3-demo-bucket' -Force -Select '^BucketName'

```

**Output:**

```nohighlight

amzn-s3-demo-bucket
```

- For API details, see
[DeletePublicAccessBlock](https://docs.aws.amazon.com/powershell/v4/reference)
in _AWS Tools for PowerShell Cmdlet Reference (V4)_.

**Tools for PowerShell V5**

**Example 1: This command turns off the block public access setting for the given bucket.**

```powershell

Remove-S3PublicAccessBlock -BucketName 'amzn-s3-demo-bucket' -Force -Select '^BucketName'

```

**Output:**

```nohighlight

amzn-s3-demo-bucket
```

- For API details, see
[DeletePublicAccessBlock](https://docs.aws.amazon.com/powershell/v5/reference)
in _AWS Tools for PowerShell Cmdlet Reference (V5)_.

For a complete list of AWS SDK developer guides and code examples, see
[Developing with Amazon S3 using the AWS SDKs](../../../../reference/amazons3/latest/api/sdk-general-information-section.md).
This topic also includes information about getting started and details about previous SDK versions.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DeleteObjects

GetBucketAccelerateConfiguration
