This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Partition SerdeInfo

Information about a serialization/deserialization program (SerDe) that serves as an
extractor and loader.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Parameters" : Json,
  "SerializationLibrary" : String
}

```

### YAML

```yaml

  Name: String
  Parameters: Json
  SerializationLibrary: String

```

## Properties

`Name`

Name of the SerDe.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

These key-value pairs define initialization parameters for the SerDe.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SerializationLibrary`

Usually the class that implements the SerDe. An example is
`org.apache.hadoop.hive.serde2.columnar.ColumnarSerDe`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SchemaReference

SkewedInfo

All content copied from https://docs.aws.amazon.com/.
