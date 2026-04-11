# NoncurrentVersionTransition

The container for the noncurrent version transition.

## Contents

**NoncurrentDays**

Specifies the number of days an object is noncurrent before Amazon S3 can perform the
associated action. For information about the noncurrent days calculations, see [How\
Amazon S3 Calculates How Long an Object Has Been Noncurrent](../dev/intro-lifecycle-rules.md#non-current-days-calculations) in the
_Amazon S3 User Guide_.

Type: Integer

Required: No

**StorageClass**

The class of storage used to store the object.

Type: String

Valid Values: `GLACIER | STANDARD_IA | ONEZONE_IA | INTELLIGENT_TIERING | DEEP_ARCHIVE`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/noncurrentversiontransition.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/noncurrentversiontransition.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/noncurrentversiontransition.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NoncurrentVersionExpiration

NotSSEFilter

All content copied from https://docs.aws.amazon.com/.
