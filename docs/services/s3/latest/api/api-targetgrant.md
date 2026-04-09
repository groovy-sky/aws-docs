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

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/targetgrant.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/targetgrant.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/targetgrant.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tagging

TargetObjectKeyFormat

All content copied from https://docs.aws.amazon.com/.
