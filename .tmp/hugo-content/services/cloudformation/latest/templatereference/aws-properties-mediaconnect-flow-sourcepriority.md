This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::Flow SourcePriority

The priority you want to assign to a source. You can have a primary stream and a backup stream or two equally prioritized streams.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PrimarySource" : String
}

```

### YAML

```yaml

  PrimarySource: String

```

## Properties

`PrimarySource`

The name of the source you choose as the primary source for this flow.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SourceMonitoringConfig

Tag

All content copied from https://docs.aws.amazon.com/.
