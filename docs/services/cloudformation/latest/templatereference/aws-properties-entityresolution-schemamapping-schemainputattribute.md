This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EntityResolution::SchemaMapping SchemaInputAttribute

A configuration object for defining input data fields in AWS Entity Resolution. The `SchemaInputAttribute` specifies how individual fields in your input data
should be processed and matched.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldName" : String,
  "GroupName" : String,
  "Hashed" : Boolean,
  "MatchKey" : String,
  "SubType" : String,
  "Type" : String
}

```

### YAML

```yaml

  FieldName: String
  GroupName: String
  Hashed: Boolean
  MatchKey: String
  SubType: String
  Type: String

```

## Properties

`FieldName`

A string containing the field name.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z_0-9- \t]*$`

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GroupName`

A string that instructs AWS Entity Resolution to combine several columns into a unified
column with the identical attribute type.

For example, when working with columns such as `NAME_FIRST`,
`NAME_MIDDLE`, and `NAME_LAST`, assigning them a common
`groupName` will prompt AWS Entity Resolution to concatenate them into a single
value.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z_0-9- \t]*$`

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Hashed`

Indicates if the column values are hashed in the schema input.

If the value is set to `TRUE`, the column values are hashed.

If the value is set to `FALSE`, the column values are cleartext.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MatchKey`

A key that allows grouping of multiple input attributes into a unified matching group.

For example, consider a scenario where the source table contains various addresses, such
as `business_address` and `shipping_address`. By assigning a
`matchKey` called `address` to both attributes, AWS Entity Resolution
will match records across these fields to create a consolidated matching group.

If no `matchKey` is specified for a column, it won't be utilized for matching
purposes but will still be included in the output table.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z_0-9- \t]*$`

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubType`

The subtype of the attribute, selected from a list of values.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the attribute, selected from a list of values.

LiveRamp supports: `NAME` \| `NAME_FIRST` \|
`NAME_MIDDLE` \| `NAME_LAST` \| `ADDRESS` \|
`ADDRESS_STREET1` \| `ADDRESS_STREET2` \|
`ADDRESS_STREET3` \| `ADDRESS_CITY` \| `ADDRESS_STATE` \|
`ADDRESS_COUNTRY` \| `ADDRESS_POSTALCODE` \| `PHONE` \|
`PHONE_NUMBER` \| `EMAIL_ADDRESS` \| `UNIQUE_ID` \|
`PROVIDER_ID`

TransUnion supports: `NAME` \| `NAME_FIRST` \|
`NAME_LAST` \| `ADDRESS` \| `ADDRESS_CITY` \|
`ADDRESS_STATE` \| `ADDRESS_COUNTRY` \|
`ADDRESS_POSTALCODE` \| `PHONE_NUMBER` \| `EMAIL_ADDRESS`
\| `UNIQUE_ID` \| `IPV4` \| `IPV6` \| `MAID`

Unified ID 2.0 supports: `PHONE_NUMBER` \| `EMAIL_ADDRESS` \|
`UNIQUE_ID`

###### Note

Normalization is only supported for `NAME`, `ADDRESS`,
`PHONE`, and `EMAIL_ADDRESS`.

If you want to normalize `NAME_FIRST`, `NAME_MIDDLE`, and
`NAME_LAST`, you must group them by assigning them to the
`NAME` `groupName`.

If you want to normalize `ADDRESS_STREET1`, `ADDRESS_STREET2`,
`ADDRESS_STREET3`, `ADDRESS_CITY`, `ADDRESS_STATE`,
`ADDRESS_COUNTRY`, and `ADDRESS_POSTALCODE`, you must group
them by assigning them to the `ADDRESS` `groupName`.

If you want to normalize `PHONE_NUMBER` and
`PHONE_COUNTRYCODE`, you must group them by assigning them to the
`PHONE` `groupName`.

_Required_: Yes

_Type_: String

_Allowed values_: `NAME | NAME_FIRST | NAME_MIDDLE | NAME_LAST | ADDRESS | ADDRESS_STREET1 | ADDRESS_STREET2 | ADDRESS_STREET3 | ADDRESS_CITY | ADDRESS_STATE | ADDRESS_COUNTRY | ADDRESS_POSTALCODE | PHONE | PHONE_NUMBER | PHONE_COUNTRYCODE | EMAIL_ADDRESS | UNIQUE_ID | DATE | STRING | PROVIDER_ID`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EntityResolution::SchemaMapping

Tag
