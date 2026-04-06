# SuccessResponseHandlingConfig

Determines how Amazon AppFlow handles the success response that it gets from the
connector after placing data.

For example, this setting would determine where to write the response from the destination
connector upon a successful insert operation.

## Contents

**bucketName**

The name of the Amazon S3 bucket.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 63.

Pattern: `\S+`

Required: No

**bucketPrefix**

The Amazon S3 bucket prefix.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `.*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/appflow-2020-08-23/SuccessResponseHandlingConfig)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appflow-2020-08-23/SuccessResponseHandlingConfig)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appflow-2020-08-23/SuccessResponseHandlingConfig)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SourceFlowConfig

SupportedFieldTypeDetails
