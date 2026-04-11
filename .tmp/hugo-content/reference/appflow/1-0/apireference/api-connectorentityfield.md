# ConnectorEntityField

Describes the data model of a connector field. For example, for an
_account_ entity, the fields would be _account name_,
_account ID_, and so on.

## Contents

**identifier**

The unique identifier of the connector field.

Type: String

Length Constraints: Maximum length of 128.

Pattern: `\S+`

Required: Yes

**customProperties**

A map that has specific properties related to the ConnectorEntityField.

Type: String to string map

Map Entries: Minimum number of 0 items. Maximum number of 50 items.

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Key Pattern: `[\w]+`

Value Length Constraints: Maximum length of 2048.

Value Pattern: `\S+`

Required: No

**defaultValue**

Default value that can be assigned to this field.

Type: String

Length Constraints: Maximum length of 2048.

Pattern: `.*`

Required: No

**description**

A description of the connector entity field.

Type: String

Length Constraints: Maximum length of 1024.

Pattern: `[\s\w/!@#+=.-]*`

Required: No

**destinationProperties**

The properties applied to a field when the connector is being used as a destination.

Type: [DestinationFieldProperties](api-destinationfieldproperties.md) object

Required: No

**isDeprecated**

Booelan value that indicates whether this field is deprecated or not.

Type: Boolean

Required: No

**isPrimaryKey**

Booelan value that indicates whether this field can be used as a primary key.

Type: Boolean

Required: No

**label**

The label applied to a connector entity field.

Type: String

Length Constraints: Maximum length of 128.

Pattern: `.*`

Required: No

**parentIdentifier**

The parent identifier of the connector field.

Type: String

Length Constraints: Maximum length of 128.

Pattern: `\S+`

Required: No

**sourceProperties**

The properties that can be applied to a field when the connector is being used as a
source.

Type: [SourceFieldProperties](api-sourcefieldproperties.md) object

Required: No

**supportedFieldTypeDetails**

Contains details regarding the supported `FieldType`, including the
corresponding `filterOperators` and `supportedValues`.

Type: [SupportedFieldTypeDetails](api-supportedfieldtypedetails.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/connectorentityfield.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/connectorentityfield.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/connectorentityfield.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectorEntity

ConnectorMetadata

All content copied from https://docs.aws.amazon.com/.
