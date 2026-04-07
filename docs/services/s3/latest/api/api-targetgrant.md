# TargetGrant

Container for granting information.

Buckets that use the bucket owner enforced setting for Object Ownership don't support target grants.
For more information, see [Permissions server access log delivery](../userguide/enable-server-access-logging.md#grant-log-delivery-permissions-general) in the _Amazon S3 User Guide_.

## Contents

**Grantee**

Container for the person being granted permissions.

Type: [Grantee](api-grantee.md) data type

Required: No

**Permission**

Logging permissions assigned to the grantee for the bucket.

Type: String

Valid Values: `FULL_CONTROL | READ | WRITE`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3-2006-03-01/TargetGrant)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3-2006-03-01/TargetGrant)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3-2006-03-01/TargetGrant)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tagging

TargetObjectKeyFormat
