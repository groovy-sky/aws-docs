# ObjectLockConfiguration

The container element for Object Lock configuration parameters.

## Contents

**ObjectLockEnabled**

Indicates whether this bucket has an Object Lock configuration enabled. Enable
`ObjectLockEnabled` when you apply `ObjectLockConfiguration` to a bucket.

Type: String

Valid Values: `Enabled`

Required: No

**Rule**

Specifies the Object Lock rule for the specified object. Enable the this rule when you apply
`ObjectLockConfiguration` to a bucket. Bucket settings require both a mode and a period.
The period can be either `Days` or `Years` but you must select one. You cannot
specify `Days` and `Years` at the same time.

Type: [ObjectLockRule](api-objectlockrule.md) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/objectlockconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/objectlockconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/objectlockconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ObjectIdentifier

ObjectLockLegalHold

All content copied from https://docs.aws.amazon.com/.
