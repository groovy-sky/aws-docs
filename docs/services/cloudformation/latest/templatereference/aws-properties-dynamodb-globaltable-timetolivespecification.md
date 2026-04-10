This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::GlobalTable TimeToLiveSpecification

Represents the settings used to enable or disable Time to Live (TTL) for the specified
table. All replicas will have the same time to live configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AttributeName" : String,
  "Enabled" : Boolean
}

```

### YAML

```yaml

  AttributeName: String
  Enabled: Boolean

```

## Properties

`AttributeName`

The name of the attribute used to store the expiration time for items in the
table.

Currently, you cannot directly change the attribute name used to evaluate time to live.
In order to do so, you must first disable time to live, and then re-enable it with the new
attribute name. It can take up to one hour for changes to time to live to take effect. If
you attempt to modify time to live within that time window, your stack operation might be
delayed.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Indicates whether TTL is to be enabled (true) or disabled (false) on the table.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TargetTrackingScalingPolicyConfiguration

WarmThroughput

All content copied from https://docs.aws.amazon.com/.
