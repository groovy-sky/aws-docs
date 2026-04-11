# CustomConnectorDestinationProperties

The properties that are applied when the custom connector is being used as a
destination.

## Contents

**entityName**

The entity specified in the custom connector as a destination in the flow.

Type: String

Length Constraints: Maximum length of 1024.

Pattern: `\S+`

Required: Yes

**customProperties**

The custom properties that are specific to the connector when it's used as a destination
in the flow.

Type: String to string map

Map Entries: Minimum number of 0 items. Maximum number of 50 items.

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Key Pattern: `[\w]+`

Value Length Constraints: Maximum length of 2048.

Value Pattern: `\S+`

Required: No

**errorHandlingConfig**

The settings that determine how Amazon AppFlow handles an error when placing data in
the custom connector as destination.

Type: [ErrorHandlingConfig](api-errorhandlingconfig.md) object

Required: No

**idFieldNames**

The name of the field that Amazon AppFlow uses as an ID when performing a write
operation such as update, delete, or upsert.

Type: Array of strings

Array Members: Minimum number of 0 items.

Length Constraints: Maximum length of 128.

Pattern: `\S+`

Required: No

**writeOperationType**

Specifies the type of write operation to be performed in the custom connector when it's
used as destination.

Type: String

Valid Values: `INSERT | UPSERT | UPDATE | DELETE`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/customconnectordestinationproperties.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/customconnectordestinationproperties.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/customconnectordestinationproperties.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomAuthCredentials

CustomConnectorProfileCredentials

All content copied from https://docs.aws.amazon.com/.
