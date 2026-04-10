This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIGuardrail GuardrailTopicConfig

Topic configuration in topic policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Definition" : String,
  "Examples" : [ String, ... ],
  "Name" : String,
  "Type" : String
}

```

### YAML

```yaml

  Definition: String
  Examples:
    - String
  Name: String
  Type: String

```

## Properties

`Definition`

Definition of topic in topic policy.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Examples`

Text example in topic policy.

_Required_: No

_Type_: Array of String

_Maximum_: `100`

_Minimum_: `1 | 0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Name of topic in topic policy.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-zA-Z-_ !?.]+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Type of topic in a policy.

_Required_: Yes

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GuardrailRegexConfig

GuardrailWordConfig

All content copied from https://docs.aws.amazon.com/.
