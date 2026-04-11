This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EntityResolution::MatchingWorkflow OutputAttribute

A list of `OutputAttribute` objects, each of which have the fields
`Name` and `Hashed`. Each of these objects selects a column to be
included in the output table, and whether the values of the column should be hashed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Hashed" : Boolean,
  "Name" : String
}

```

### YAML

```yaml

  Hashed: Boolean
  Name: String

```

## Properties

`Hashed`

Enables the ability to hash the column values in the output.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A name of a column to be written to the output. This must be an `InputField`
name in the schema mapping.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z_0-9- \t]*$`

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IntermediateSourceConfiguration

OutputSource

All content copied from https://docs.aws.amazon.com/.
