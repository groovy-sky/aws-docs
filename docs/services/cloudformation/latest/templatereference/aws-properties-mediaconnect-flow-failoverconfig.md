This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::Flow FailoverConfig

The settings for source failover.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FailoverMode" : String,
  "RecoveryWindow" : Integer,
  "SourcePriority" : SourcePriority,
  "State" : String
}

```

### YAML

```yaml

  FailoverMode: String
  RecoveryWindow: Integer
  SourcePriority:
    SourcePriority
  State: String

```

## Properties

`FailoverMode`

The type of failover you choose for this flow. MERGE combines the source streams
into a single stream, allowing graceful recovery from any single-source loss.
FAILOVER allows switching between different streams. The string for this property must be entered as MERGE or FAILOVER. No other string entry is valid.

_Required_: No

_Type_: String

_Allowed values_: `MERGE | FAILOVER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecoveryWindow`

Search window time to look for dash-7 packets.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourcePriority`

The priority you want to assign to a source. You can have a primary stream and a backup stream or two equally prioritized streams.

_Required_: No

_Type_: [SourcePriority](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediaconnect-flow-sourcepriority.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`State`

The state of source failover on the flow. If the state is inactive, the flow can have only one source. If the state is active, the flow can have one or two sources.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Encryption

FlowTransitEncryption
