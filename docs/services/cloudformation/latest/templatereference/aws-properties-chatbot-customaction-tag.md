This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Chatbot::CustomAction Tag

###### Note

AWS Chatbot is now Amazon Q Developer. [Learn more](../../../chatbot/latest/adminguide/service-rename.md)

`Type` attribute values remain unchanged.

A map of tags assigned to a resource. A tag is a string-to-string map of key-value pairs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

A string used to identify this tag. You can specify a maximum of 128 characters for a tag key. Tags owned by Amazon Web Services (AWS) have the reserved prefix: `aws:`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

A string containing the value for this tag. You can specify a maximum of 256 characters for a tag value.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomActionDefinition

AWS::Chatbot::MicrosoftTeamsChannelConfiguration

All content copied from https://docs.aws.amazon.com/.
