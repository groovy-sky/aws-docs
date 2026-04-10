This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::Campaign OverrideButtonConfiguration

Specifies the configuration of a button with settings that are specific to a certain
device type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ButtonAction" : String,
  "Link" : String
}

```

### YAML

```yaml

  ButtonAction: String
  Link: String

```

## Properties

`ButtonAction`

The action that occurs when a recipient chooses a button in an in-app message.
You can specify one of the following:

- `LINK` – A link to a web destination.

- `DEEP_LINK` – A link to a specific page in an
application.

- `CLOSE` – Dismisses the message.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Link`

The destination (such as a URL) for a button.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MessageConfiguration

QuietTime

All content copied from https://docs.aws.amazon.com/.
