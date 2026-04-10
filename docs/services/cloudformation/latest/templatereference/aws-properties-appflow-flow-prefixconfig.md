This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow PrefixConfig

Specifies elements that Amazon AppFlow includes in the file and folder names in the flow
destination.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PathPrefixHierarchy" : [ String, ... ],
  "PrefixFormat" : String,
  "PrefixType" : String
}

```

### YAML

```yaml

  PathPrefixHierarchy:
    - String
  PrefixFormat: String
  PrefixType: String

```

## Properties

`PathPrefixHierarchy`

Specifies whether the destination file path includes either or both of the following
elements:

EXECUTION\_ID

The ID that Amazon AppFlow assigns to the flow run.

SCHEMA\_VERSION

The version number of your data schema. Amazon AppFlow assigns this version
number. The version number increases by one when you change any of the following
settings in your flow configuration:

- Source-to-destination field mappings

- Field data types

- Partition keys

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrefixFormat`

Determines the level of granularity for the date and time that's included in the prefix.

_Required_: No

_Type_: String

_Allowed values_: `YEAR | MONTH | DAY | HOUR | MINUTE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrefixType`

Determines the format of the prefix, and whether it applies to the file name, file path,
or both.

_Required_: No

_Type_: String

_Allowed values_: `FILENAME | PATH | PATH_AND_FILENAME`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [PrefixConfig](../../../../reference/appflow/1-0/apireference/api-prefixconfig.md) in the _Amazon AppFlow API Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PardotSourceProperties

RedshiftDestinationProperties

All content copied from https://docs.aws.amazon.com/.
