# ObjectLockRule

The container element for an Object Lock rule.

## Contents

**DefaultRetention**

The default Object Lock retention mode and period that you want to apply to new objects placed in
the specified bucket. Bucket settings require both a mode and a period. The period can be either
`Days` or `Years` but you must select one. You cannot specify `Days`
and `Years` at the same time.

Type: [DefaultRetention](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DefaultRetention.html) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3-2006-03-01/ObjectLockRule)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3-2006-03-01/ObjectLockRule)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3-2006-03-01/ObjectLockRule)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ObjectLockRetention

ObjectPart
