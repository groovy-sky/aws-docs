# ErrorHandlingConfig

The settings that determine how Amazon AppFlow handles an error when placing data in
the destination. For example, this setting would determine if the flow should fail after one
insertion error, or continue and attempt to insert every record regardless of the initial
failure. `ErrorHandlingConfig` is a part of the destination connector details.

## Contents

**bucketName**

Specifies the name of the Amazon S3 bucket.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 63.

Pattern: `\S+`

Required: No

**bucketPrefix**

Specifies the Amazon S3 bucket prefix.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `.*`

Required: No

**failOnFirstDestinationError**

Specifies if the flow should fail after the first instance of a failure when attempting
to place data in the destination.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/errorhandlingconfig.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/errorhandlingconfig.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/errorhandlingconfig.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DynatraceSourceProperties

ErrorInfo

All content copied from https://docs.aws.amazon.com/.
