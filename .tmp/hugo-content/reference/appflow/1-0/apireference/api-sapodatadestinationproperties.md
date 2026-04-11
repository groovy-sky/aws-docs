# SAPODataDestinationProperties

The properties that are applied when using SAPOData as a flow destination

## Contents

**objectPath**

The object path specified in the SAPOData flow destination.

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

**idFieldNames**

A list of field names that can be used as an ID field when performing a write operation.

Type: Array of strings

Array Members: Minimum number of 0 items.

Length Constraints: Maximum length of 128.

Pattern: `\S+`

Required: No

**successResponseHandlingConfig**

Determines how Amazon AppFlow handles the success response that it gets from the
connector after placing data.

For example, this setting would determine where to write the response from a destination
connector upon a successful insert operation.

Type: [SuccessResponseHandlingConfig](api-successresponsehandlingconfig.md) object

Required: No

**writeOperationType**

The possible write operations in the destination connector. When this value is not
provided, this defaults to the `INSERT` operation.

Type: String

Valid Values: `INSERT | UPSERT | UPDATE | DELETE`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/sapodatadestinationproperties.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/sapodatadestinationproperties.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/sapodatadestinationproperties.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SAPODataConnectorProfileProperties

SAPODataMetadata

All content copied from https://docs.aws.amazon.com/.
