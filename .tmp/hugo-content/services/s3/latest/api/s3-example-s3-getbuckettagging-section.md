# Use `GetBucketTagging` with a CLI

The following code examples show how to use `GetBucketTagging`.

CLI

**AWS CLI**

The following command retrieves the tagging configuration for a bucket named `amzn-s3-demo-bucket`:

```nohighlight

aws s3api get-bucket-tagging --bucket amzn-s3-demo-bucket

```

Output:

```nohighlight

{
    "TagSet": [
        {
            "Value": "marketing",
            "Key": "organization"
        }
    ]
}
```

- For API details, see
[GetBucketTagging](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/s3api/get-bucket-tagging.html)
in _AWS CLI Command Reference_.

PowerShell

**Tools for PowerShell V4**

**Example 1: This command returns all the tags associated with the given bucket.**

```powershell

Get-S3BucketTagging -BucketName 'amzn-s3-demo-bucket'

```

- For API details, see
[GetBucketTagging](../../../powershell/v4/reference.md)
in _AWS Tools for PowerShell Cmdlet Reference (V4)_.

**Tools for PowerShell V5**

**Example 1: This command returns all the tags associated with the given bucket.**

```powershell

Get-S3BucketTagging -BucketName 'amzn-s3-demo-bucket'

```

- For API details, see
[GetBucketTagging](../../../powershell/v5/reference.md)
in _AWS Tools for PowerShell Cmdlet Reference (V5)_.

For a complete list of AWS SDK developer guides and code examples, see
[Developing with Amazon S3 using the AWS SDKs](../../../../reference/amazons3/latest/api/sdk-general-information-section.md).
This topic also includes information about getting started and details about previous SDK versions.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetBucketRequestPayment

GetBucketVersioning

All content copied from https://docs.aws.amazon.com/.
