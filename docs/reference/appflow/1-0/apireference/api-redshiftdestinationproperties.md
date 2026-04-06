# RedshiftDestinationProperties

The properties that are applied when Amazon Redshift is being used as a destination.

## Contents

**intermediateBucketName**

The intermediate bucket that Amazon AppFlow uses when moving data into Amazon Redshift.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 63.

Pattern: `\S+`

Required: Yes

**object**

The object specified in the Amazon Redshift flow destination.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `\S+`

Required: Yes

**bucketPrefix**

The object key for the bucket in which Amazon AppFlow places the destination files.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `.*`

Required: No

**errorHandlingConfig**

The settings that determine how Amazon AppFlow handles an error when placing data in
the Amazon Redshift destination. For example, this setting would determine if the flow
should fail after one insertion error, or continue and attempt to insert every record
regardless of the initial failure. `ErrorHandlingConfig` is a part of the
destination connector details.

Type: [ErrorHandlingConfig](api-errorhandlingconfig.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/appflow-2020-08-23/RedshiftDestinationProperties)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appflow-2020-08-23/RedshiftDestinationProperties)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appflow-2020-08-23/RedshiftDestinationProperties)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RedshiftConnectorProfileProperties

RedshiftMetadata
