This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::AnalysisTemplate AnalysisParameter

Optional. The member who can query can provide this placeholder for a literal data value
in an analysis template.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultValue" : String,
  "Name" : String,
  "Type" : String
}

```

### YAML

```yaml

  DefaultValue: String
  Name: String
  Type: String

```

## Properties

`DefaultValue`

Optional. The default value that is applied in the analysis template. The member who can
query can override this value in the query editor.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the parameter. The name must use only alphanumeric or underscore (\_) characters.

_Required_: Yes

_Type_: String

_Pattern_: `[0-9a-zA-Z_]+`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The type of parameter.

_Required_: Yes

_Type_: String

_Allowed values_: `SMALLINT | INTEGER | BIGINT | DECIMAL | REAL | DOUBLE_PRECISION | BOOLEAN | CHAR | VARCHAR | DATE | TIMESTAMP | TIMESTAMPTZ | TIME | TIMETZ | VARBYTE | BINARY | BYTE | CHARACTER | DOUBLE | FLOAT | INT | LONG | NUMERIC | SHORT | STRING | TIMESTAMP_LTZ | TIMESTAMP_NTZ | TINYINT`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::CleanRooms::AnalysisTemplate

AnalysisSchema
