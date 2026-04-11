This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::MessageTemplate SystemAttributes

The system attributes that are used with the message template.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomerEndpoint" : SystemEndpointAttributes,
  "Name" : String,
  "SystemEndpoint" : SystemEndpointAttributes
}

```

### YAML

```yaml

  CustomerEndpoint:
    SystemEndpointAttributes
  Name: String
  SystemEndpoint:
    SystemEndpointAttributes

```

## Properties

`CustomerEndpoint`

The CustomerEndpoint attribute.

_Required_: No

_Type_: [SystemEndpointAttributes](aws-properties-wisdom-messagetemplate-systemendpointattributes.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the task.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SystemEndpoint`

The SystemEndpoint attribute.

_Required_: No

_Type_: [SystemEndpointAttributes](aws-properties-wisdom-messagetemplate-systemendpointattributes.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SmsMessageTemplateContentBody

SystemEndpointAttributes

All content copied from https://docs.aws.amazon.com/.
