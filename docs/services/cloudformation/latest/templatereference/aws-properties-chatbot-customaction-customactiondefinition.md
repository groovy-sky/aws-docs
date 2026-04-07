This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Chatbot::CustomAction CustomActionDefinition

###### Note

AWS Chatbot is now Amazon Q Developer. [Learn more](https://docs.aws.amazon.com/chatbot/latest/adminguide/service-rename.html)

`Type` attribute values remain unchanged.

The definition of the command to run when invoked as an alias or as an action button.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CommandText" : String
}

```

### YAML

```yaml

  CommandText: String

```

## Properties

`CommandText`

The command string to run which may include variables by prefixing with a dollar sign ($).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `5000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CustomActionAttachmentCriteria

Tag
