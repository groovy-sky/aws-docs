# Use `GetBucketRequestPayment` with a CLI

The following code examples show how to use `GetBucketRequestPayment`.

CLI

**AWS CLI**

**To retrieve the request payment configuration for a bucket**

The following `get-bucket-request-payment` example retrieves the requester pays configuration for the specified bucket.

```nohighlight

aws s3api get-bucket-request-payment \
    --bucket amzn-s3-demo-bucket

```

Output:

```nohighlight

{
    "Payer": "BucketOwner"
}
```

- For API details, see
[GetBucketRequestPayment](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/s3api/get-bucket-request-payment.html)
in _AWS CLI Command Reference_.

PowerShell

**Tools for PowerShell V4**

**Example 1: Returns the request payment configuration for the bucket named 'amzn-s3-demo-bucket'. By default, the bucket owner pays for downloads from the bucket.**

```powershell

Get-S3BucketRequestPayment -BucketName amzn-s3-demo-bucket

```

- For API details, see
[GetBucketRequestPayment](https://docs.aws.amazon.com/powershell/v4/reference)
in _AWS Tools for PowerShell Cmdlet Reference (V4)_.

**Tools for PowerShell V5**

**Example 1: Returns the request payment configuration for the bucket named 'amzn-s3-demo-bucket'. By default, the bucket owner pays for downloads from the bucket.**

```powershell

Get-S3BucketRequestPayment -BucketName amzn-s3-demo-bucket

```

- For API details, see
[GetBucketRequestPayment](https://docs.aws.amazon.com/powershell/v5/reference)
in _AWS Tools for PowerShell Cmdlet Reference (V5)_.

For a complete list of AWS SDK developer guides and code examples, see
[Developing with Amazon S3 using the AWS SDKs](../../../../reference/amazons3/latest/api/sdk-general-information-section.md).
This topic also includes information about getting started and details about previous SDK versions.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GetBucketReplication

GetBucketTagging
