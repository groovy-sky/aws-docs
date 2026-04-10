This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::SignalMap

The `AWS::MediaLive::SignalMap` resource Property description not available. for MediaLive.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaLive::SignalMap",
  "Properties" : {
      "CloudWatchAlarmTemplateGroupIdentifiers" : [ String, ... ],
      "Description" : String,
      "DiscoveryEntryPointArn" : String,
      "EventBridgeRuleTemplateGroupIdentifiers" : [ String, ... ],
      "ForceRediscovery" : Boolean,
      "Name" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::MediaLive::SignalMap
Properties:
  CloudWatchAlarmTemplateGroupIdentifiers:
    - String
  Description: String
  DiscoveryEntryPointArn: String
  EventBridgeRuleTemplateGroupIdentifiers:
    - String
  ForceRediscovery: Boolean
  Name: String
  Tags:
    Key: Value

```

## Properties

`CloudWatchAlarmTemplateGroupIdentifiers`

A cloudwatch alarm template group's identifier. Can be either be its id or current name.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A resource's optional description.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DiscoveryEntryPointArn`

A top-level supported Amazon Web Services resource ARN to discover a signal map from.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventBridgeRuleTemplateGroupIdentifiers`

An eventbridge rule template group's identifier. Can be either be its id or current name.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ForceRediscovery`

If true, will force a rediscovery of a signal map if an unchanged discoveryEntryPointArn is provided.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A resource's name. Names must be unique within the scope of a resource type in a specific region.

_Required_: Yes

_Type_: String

_Pattern_: `^[^\s]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`Arn`

A signal map's ARN (Amazon Resource Name)

`CloudWatchAlarmTemplateGroupIds`

An alarm template group's id.

`CreatedAt`

The date and time of resource creation.

`ErrorMessage`

Error message associated with a failed creation or failed update attempt of a signal map.

`EventBridgeRuleTemplateGroupIds`

An eventbridge rule template group's id.

`Id`

A signal map's id.

`Identifier`

Property description not available.

`LastDiscoveredAt`

The date and time of latest discovery.

`ModifiedAt`

The date and time of latest resource modification.

`MonitorChangesPendingDeployment`

If true, there are pending monitor changes for this signal map that can be deployed.

`Status`

A signal map's current status, which is dependent on its lifecycle actions or associated jobs.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tags

MediaResource

All content copied from https://docs.aws.amazon.com/.
