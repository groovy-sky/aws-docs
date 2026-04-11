This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BCMDataExports::Export DataQuery

The SQL query of column selections and row filters from the data table you want.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "QueryStatement" : String,
  "TableConfigurations" : {Key: Value, ...}
}

```

### YAML

```yaml

  QueryStatement: String
  TableConfigurations:
    Key: Value

```

## Properties

`QueryStatement`

The query statement.

_Required_: Yes

_Type_: String

_Pattern_: `^[\S\s]*$`

_Minimum_: `1`

_Maximum_: `36000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableConfigurations`

The table configuration.

_Required_: No

_Type_: Object of String

_Pattern_: `^[\S\s]*$`

_Minimum_: `0`

_Maximum_: `16384`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::BCMDataExports::Export

DestinationConfigurations

All content copied from https://docs.aws.amazon.com/.
