# S3BucketAclGrantConfiguration

A proposed access control list grant configuration for an Amazon S3 bucket. For more
information, see [How to Specify an\
ACL](../../../../services/s3/latest/dev/acl-overview-setting-acls.md).

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

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/s3bucketaclgrantconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/s3bucketaclgrantconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/s3bucketaclgrantconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

S3AccessPointConfiguration

S3BucketConfiguration
