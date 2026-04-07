# SnapshotDiskContainer

The disk container object for the import snapshot request.

## Contents

**Description**

The description of the disk image being imported.

Type: String

Required: No

**Format**

The format of the disk image being imported.

Valid values: `VHD` \| `VMDK` \| `RAW`

Type: String

Required: No

**Url**

The URL to the Amazon S3-based disk image being imported. It can either be a https URL (https://..) or an Amazon
S3 URL (s3://..).

Type: String

Required: No

**UserBucket**

The Amazon S3 bucket for the disk image.

Type: [UserBucket](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_UserBucket.html) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/SnapshotDiskContainer)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/SnapshotDiskContainer)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/SnapshotDiskContainer)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SnapshotDetail

SnapshotInfo
