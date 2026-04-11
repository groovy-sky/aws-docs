This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PinpointEmail::ConfigurationSet SendingOptions

Used to enable or disable email sending for messages that use this configuration set in
the current AWS Region.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SendingEnabled" : Boolean
}

```

### YAML

```yaml

  SendingEnabled: Boolean

```

## Properties

`SendingEnabled`

If `true`, email sending is enabled for the configuration set. If
`false`, email sending is disabled for the configuration set.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReputationOptions

Tags

All content copied from https://docs.aws.amazon.com/.
