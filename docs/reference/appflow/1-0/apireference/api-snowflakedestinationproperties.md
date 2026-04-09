# SnowflakeDestinationProperties

The properties that are applied when Snowflake is being used as a destination.

## Contents

**intermediateBucketName**

The intermediate bucket that Amazon AppFlow uses when moving data into Snowflake.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 63.

Pattern: `\S+`

Required: Yes

**object**

The object specified in the Snowflake flow destination.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `\S+`

Required: Yes

**bucketPrefix**

The object key for the destination bucket in which Amazon AppFlow places the files.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `.*`

Required: No

**errorHandlingConfig**

The settings that determine how Amazon AppFlow handles an error when placing data in
the Snowflake destination. For example, this setting would determine if the flow should fail
after one insertion error, or continue and attempt to insert every record regardless of the
initial failure. `ErrorHandlingConfig` is a part of the destination connector
details.

Type: [ErrorHandlingConfig](api-errorhandlingconfig.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/snowflakedestinationproperties.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/snowflakedestinationproperties.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/snowflakedestinationproperties.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SnowflakeConnectorProfileProperties

SnowflakeMetadata

All content copied from https://docs.aws.amazon.com/.
