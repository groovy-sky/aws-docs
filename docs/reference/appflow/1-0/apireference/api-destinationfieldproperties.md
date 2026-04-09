# DestinationFieldProperties

The properties that can be applied to a field when connector is being used as a
destination.

## Contents

**isCreatable**

Specifies if the destination field can be created by the current user.

Type: Boolean

Required: No

**isDefaultedOnCreate**

Specifies whether the field can use the default value during a Create operation.

Type: Boolean

Required: No

**isNullable**

Specifies if the destination field can have a null value.

Type: Boolean

Required: No

**isUpdatable**

Specifies whether the field can be updated during an `UPDATE` or
`UPSERT` write operation.

Type: Boolean

Required: No

**isUpsertable**

Specifies if the flow run can either insert new rows in the destination field if they do
not already exist, or update them if they do.

Type: Boolean

Required: No

**supportedWriteOperations**

A list of supported write operations. For each write operation listed, this field can be
used in `idFieldNames` when that write operation is present as a destination
option.

Type: Array of strings

Valid Values: `INSERT | UPSERT | UPDATE | DELETE`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/destinationfieldproperties.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/destinationfieldproperties.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/destinationfieldproperties.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DestinationConnectorProperties

DestinationFlowConfig

All content copied from https://docs.aws.amazon.com/.
