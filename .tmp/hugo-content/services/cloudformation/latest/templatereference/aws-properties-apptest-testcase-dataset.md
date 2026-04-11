This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppTest::TestCase DataSet

Defines a data set.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Ccsid" : String,
  "Format" : String,
  "Length" : Number,
  "Name" : String,
  "Type" : String
}

```

### YAML

```yaml

  Ccsid: String
  Format: String
  Length: Number
  Name: String
  Type: String

```

## Properties

`Ccsid`

The CCSID of the data set.

_Required_: Yes

_Type_: String

_Pattern_: `^\S{1,50}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Format`

The format of the data set.

_Required_: Yes

_Type_: String

_Allowed values_: `FIXED | VARIABLE | LINE_SEQUENTIAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Length`

The length of the data set.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the data set.

_Required_: Yes

_Type_: String

_Pattern_: `^\S{1,100}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the data set.

_Required_: Yes

_Type_: String

_Allowed values_: `PS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DatabaseCDC

FileMetadata

All content copied from https://docs.aws.amazon.com/.
