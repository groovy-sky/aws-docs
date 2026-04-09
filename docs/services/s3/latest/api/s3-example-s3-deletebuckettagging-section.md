# Use `DeleteBucketTagging` with a CLI

The following code examples show how to use `DeleteBucketTagging`.

CLI

**AWS CLI**

The following command deletes a tagging configuration from a bucket named `amzn-s3-demo-bucket`:

```nohighlight

aws s3api delete-bucket-tagging --bucket amzn-s3-demo-bucket

```

- For API details, see
[DeleteBucketTagging](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/s3api/delete-bucket-tagging.html)
in _AWS CLI Command Reference_.

PowerShell

**Tools for PowerShell V4**

**Example 1: This command removes all the tags associated with the given S3 bucket.**

```powershell

Remove-S3BucketTagging -BucketName 'amzn-s3-demo-bucket'

```

**Output:**

```nohighlight

Confirm
Are you sure you want to perform this action?
Performing the operation "Remove-S3BucketTagging (DeleteBucketTagging)" on target "amzn-s3-demo-bucket".
[Y] Yes  [A] Yes to All  [N] No  [L] No to All  [S] Suspend  [?] Help (default is "Y"): Y
```

- For API details, see
[DeleteBucketTagging](../../../powershell/v4/reference.md)
in _AWS Tools for PowerShell Cmdlet Reference (V4)_.

**Tools for PowerShell V5**

**Example 1: This command removes all the tags associated with the given S3 bucket.**

```powershell

Remove-S3BucketTagging -BucketName 'amzn-s3-demo-bucket'

```

**Output:**

```nohighlight

Confirm
Are you sure you want to perform this action?
Performing the operation "Remove-S3BucketTagging (DeleteBucketTagging)" on target "amzn-s3-demo-bucket".
[Y] Yes  [A] Yes to All  [N] No  [L] No to All  [S] Suspend  [?] Help (default is "Y"): Y
```

- For API details, see
[DeleteBucketTagging](../../../powershell/v5/reference.md)
in _AWS Tools for PowerShell Cmdlet Reference (V5)_.

For a complete list of AWS SDK developer guides and code examples, see
[Developing with Amazon S3 using the AWS SDKs](../../../../reference/amazons3/latest/api/sdk-general-information-section.md).
This topic also includes information about getting started and details about previous SDK versions.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteBucketReplication

DeleteBucketWebsite

All content copied from https://docs.aws.amazon.com/.
