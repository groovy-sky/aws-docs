# MarketoDestinationProperties

The properties that Amazon AppFlow applies when you use Marketo as a flow
destination.

## Contents

**object**

The object specified in the Marketo flow destination.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `\S+`

Required: Yes

**errorHandlingConfig**

The settings that determine how Amazon AppFlow handles an error when placing data in
the destination. For example, this setting would determine if the flow should fail after one
insertion error, or continue and attempt to insert every record regardless of the initial
failure. `ErrorHandlingConfig` is a part of the destination connector details.

Type: [ErrorHandlingConfig](api-errorhandlingconfig.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/marketodestinationproperties.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/marketodestinationproperties.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/marketodestinationproperties.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MarketoConnectorProfileProperties

MarketoMetadata

All content copied from https://docs.aws.amazon.com/.
