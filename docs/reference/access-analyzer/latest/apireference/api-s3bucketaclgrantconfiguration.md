# S3BucketAclGrantConfiguration

A proposed access control list grant configuration for an Amazon S3 bucket. For more
information, see [How to Specify an\
ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#setting-acls).

## Contents

**grantee**

The grantee to whom you’re assigning access rights.

Type: [AclGrantee](api-aclgrantee.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: Yes

**permission**

The permissions being granted.

Type: String

Valid Values: `READ | WRITE | READ_ACP | WRITE_ACP | FULL_CONTROL`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/accessanalyzer-2019-11-01/S3BucketAclGrantConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/accessanalyzer-2019-11-01/S3BucketAclGrantConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/accessanalyzer-2019-11-01/S3BucketAclGrantConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

S3AccessPointConfiguration

S3BucketConfiguration
