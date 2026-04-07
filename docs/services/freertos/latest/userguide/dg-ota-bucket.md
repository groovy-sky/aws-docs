# Create an Amazon S3 bucket to store your update

OTA update files are stored in Amazon S3 buckets.

If you're using Code Signing for AWS IoT, the command
that you use to create a code-signing job takes a source bucket (where
the unsigned firmware image is located) and a destination bucket (where the signed
firmware image is written). You can specify the same bucket for the
source and destination. The file names are changed to GUIDs so the original files are
not overwritten.

###### To create an Amazon S3 bucket

1. Sign in to the Amazon S3 console at [https://console.aws.amazon.com/s3/](https://console.aws.amazon.com/s3).

2. Choose **Create bucket**.

3. Enter a **bucket name**.

4. Under **Bucket settings for Block Public Access** keep **Block all**
**public access** selected to accept the default permissions.

5. Under **Bucket Versioning**, select **Enable** to keep all
    versions in the same bucket.

6. Choose **Create bucket**.

For more information about Amazon S3, see [Amazon Simple Storage Service User Guide](../../../s3/latest/userguide.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

OTA update prerequisites

Create an OTA Update service role
