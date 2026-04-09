# CustomConnectorSourceProperties

The properties that are applied when the custom connector is being used as a
source.

## Contents

**entityName**

The entity specified in the custom connector as a source in the flow.

Type: String

Length Constraints: Maximum length of 1024.

Pattern: `\S+`

Required: Yes

**customProperties**

Custom properties that are required to use the custom connector as a source.

Type: String to string map

Map Entries: Minimum number of 0 items. Maximum number of 50 items.

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Key Pattern: `[\w]+`

Value Length Constraints: Maximum length of 2048.

Value Pattern: `\S+`

Required: No

**dataTransferApi**

The API of the connector application that Amazon AppFlow uses to transfer your
data.

Type: [DataTransferApi](api-datatransferapi.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/customconnectorsourceproperties.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/customconnectorsourceproperties.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/customconnectorsourceproperties.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomConnectorProfileProperties

CustomerProfilesDestinationProperties

All content copied from https://docs.aws.amazon.com/.
