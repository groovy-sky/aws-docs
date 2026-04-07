This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::TaskTemplate DefaultFieldValue

Describes a default field and its corresponding value.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultValue" : String,
  "Id" : FieldIdentifier
}

```

### YAML

```yaml

  DefaultValue: String
  Id:
    FieldIdentifier

```

## Properties

`DefaultValue`

Default value for the field.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

Identifier of a field.

_Required_: Yes

_Type_: [FieldIdentifier](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-tasktemplate-fieldidentifier.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Constraints

Field
