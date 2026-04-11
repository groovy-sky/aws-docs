This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RTBFabric::Link Action

Describes a bid action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HeaderTag" : HeaderTagAction,
  "NoBid" : NoBidAction
}

```

### YAML

```yaml

  HeaderTag:
    HeaderTagAction
  NoBid:
    NoBidAction

```

## Properties

`HeaderTag`

Describes the header tag for a bid action.

_Required_: No

_Type_: [HeaderTagAction](aws-properties-rtbfabric-link-headertagaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NoBid`

Describes the parameters of a no bid module.

_Required_: No

_Type_: [NoBidAction](aws-properties-rtbfabric-link-nobidaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::RTBFabric::Link

ApplicationLogs

All content copied from https://docs.aws.amazon.com/.
