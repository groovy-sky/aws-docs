# S3Storage

Describes the storage parameters for Amazon S3 and Amazon S3 buckets for an instance store-backed
AMI.

## Contents

**AWSAccessKeyId** (request), **AWSAccessKeyId** (response)

The access key ID of the owner of the bucket. Before you specify a value for your access
key ID, review and follow the guidance in [Best Practices for AWS\
accounts](../../../../services/accounts/latest/reference/best-practices.md) in the _AWS Account ManagementReference Guide_.

Type: String

Required: No

**Bucket** (request), **bucket** (response)

The bucket in which to store the AMI. You can specify a bucket that you already own or a
new bucket that Amazon EC2 creates on your behalf. If you specify a bucket that belongs to someone
else, Amazon EC2 returns an error.

Type: String

Required: No

**Prefix** (request), **prefix** (response)

The beginning of the file name of the AMI.

Type: String

Required: No

**UploadPolicy** (request), **uploadPolicy** (response)

An Amazon S3 upload policy that gives Amazon EC2 permission to upload items into Amazon S3 on your
behalf.

Type: Base64-encoded binary data object

Required: No

**UploadPolicySignature** (request), **uploadPolicySignature** (response)

The signature of the JSON document.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/s3storage.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/s3storage.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/s3storage.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

S3ObjectTag

ScheduledInstance
