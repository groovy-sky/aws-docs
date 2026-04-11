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

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/successresponsehandlingconfig.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/successresponsehandlingconfig.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/successresponsehandlingconfig.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SourceFlowConfig

SupportedFieldTypeDetails

All content copied from https://docs.aws.amazon.com/.
