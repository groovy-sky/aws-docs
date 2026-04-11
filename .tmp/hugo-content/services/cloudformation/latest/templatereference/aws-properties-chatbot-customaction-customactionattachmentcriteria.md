This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Chatbot::CustomAction CustomActionAttachmentCriteria

###### Note

AWS Chatbot is now Amazon Q Developer. [Learn more](../../../chatbot/latest/adminguide/service-rename.md)

`Type` attribute values remain unchanged.

A criteria for when a button should be shown based on values in the notification.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Operator" : String,
  "Value" : String,
  "VariableName" : String
}

```

### YAML

```yaml

  Operator: String
  Value: String
  VariableName: String

```

## Properties

`Operator`

The operation to perform on the named variable.

_Required_: Yes

_Type_: String

_Allowed values_: `HAS_VALUE | EQUALS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

A value that is compared with the actual value of the variable based on the behavior of the operator.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VariableName`

The name of the variable to operate on.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomActionAttachment

CustomActionDefinition

All content copied from https://docs.aws.amazon.com/.
