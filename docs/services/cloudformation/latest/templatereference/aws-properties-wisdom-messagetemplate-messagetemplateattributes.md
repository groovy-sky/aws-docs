This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::MessageTemplate MessageTemplateAttributes

The attributes that are used with the message template.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AgentAttributes" : AgentAttributes,
  "CustomAttributes" : {Key: Value, ...},
  "CustomerProfileAttributes" : CustomerProfileAttributes,
  "SystemAttributes" : SystemAttributes
}

```

### YAML

```yaml

  AgentAttributes:
    AgentAttributes
  CustomAttributes:
    Key: Value
  CustomerProfileAttributes:
    CustomerProfileAttributes
  SystemAttributes:
    SystemAttributes

```

## Properties

`AgentAttributes`

The agent attributes that are used with the message template.

_Required_: No

_Type_: [AgentAttributes](aws-properties-wisdom-messagetemplate-agentattributes.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomAttributes`

The custom attributes that are used with the message template.

_Required_: No

_Type_: Object of String

_Pattern_: `^[a-zA-Z0-9\s._:/=+@-]*$`

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomerProfileAttributes`

The customer profile attributes that are used with the message template.

_Required_: No

_Type_: [CustomerProfileAttributes](aws-properties-wisdom-messagetemplate-customerprofileattributes.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SystemAttributes`

The system attributes that are used with the message template.

_Required_: No

_Type_: [SystemAttributes](aws-properties-wisdom-messagetemplate-systemattributes.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MessageTemplateAttachment

MessageTemplateBodyContentProvider

All content copied from https://docs.aws.amazon.com/.
