# ZendeskDestinationProperties

The properties that are applied when Zendesk is used as a destination.

## Contents

**object**

The object specified in the Zendesk flow destination.

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

**writeOperationType**

The possible write operations in the destination connector. When this value is not
provided, this defaults to the `INSERT` operation.

Type: String

Valid Values: `INSERT | UPSERT | UPDATE | DELETE`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/appflow-2020-08-23/ZendeskDestinationProperties)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appflow-2020-08-23/ZendeskDestinationProperties)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appflow-2020-08-23/ZendeskDestinationProperties)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ZendeskConnectorProfileProperties

ZendeskMetadata
