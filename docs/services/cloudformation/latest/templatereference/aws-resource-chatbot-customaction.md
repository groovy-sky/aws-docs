This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Chatbot::CustomAction

###### Note

AWS Chatbot is now Amazon Q Developer. [Learn more](../../../chatbot/latest/adminguide/service-rename.md)

`Type` attribute values remain unchanged.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Chatbot::CustomAction",
  "Properties" : {
      "ActionName" : String,
      "AliasName" : String,
      "Attachments" : [ CustomActionAttachment, ... ],
      "Definition" : CustomActionDefinition,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Chatbot::CustomAction
Properties:
  ActionName: String
  AliasName: String
  Attachments:
    - CustomActionAttachment
  Definition:
    CustomActionDefinition
  Tags:
    - Tag

```

## Properties

`ActionName`

The name of the custom action. This name is included in the Amazon Resource Name (ARN).

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]{1,64}$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AliasName`

The name used to invoke this action in a chat channel. For example, `@Amazon Q run my-alias`.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9-_]+$`

_Minimum_: `1`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Attachments`

Defines when this custom action button should be attached to a notification.

_Required_: No

_Type_: Array of [CustomActionAttachment](aws-properties-chatbot-customaction-customactionattachment.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Definition`

The definition of the command to run when invoked as an alias or as an action button.

_Required_: Yes

_Type_: [CustomActionDefinition](aws-properties-chatbot-customaction-customactiondefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to add to the configuration.

_Required_: No

_Type_: Array of [Tag](aws-properties-chatbot-customaction-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic Ref function, Ref returns the ARN of the configuration created.

### Fn::GetAtt

`CustomActionArn`

The fully defined ARN of the custom action.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Q Developer in chat applications

CustomActionAttachment

All content copied from https://docs.aws.amazon.com/.
