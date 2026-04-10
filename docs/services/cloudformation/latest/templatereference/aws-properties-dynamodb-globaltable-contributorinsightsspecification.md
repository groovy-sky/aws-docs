This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::GlobalTable ContributorInsightsSpecification

Configures contributor insights settings for a replica or one of its indexes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "Mode" : String
}

```

### YAML

```yaml

  Enabled: Boolean
  Mode: String

```

## Properties

`Enabled`

Indicates whether CloudWatch Contributor Insights are to be enabled (true) or disabled
(false).

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Mode`

Specifies the CloudWatch Contributor Insights mode for a global table. Valid values are
`ACCESSED_AND_THROTTLED_KEYS` (tracks all access and throttled events) or
`THROTTLED_KEYS` (tracks only throttled events). This setting determines what
type of contributor insights data is collected for the global table.

_Required_: No

_Type_: String

_Allowed values_: `ACCESSED_AND_THROTTLED_KEYS | THROTTLED_KEYS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CapacityAutoScalingSettings

GlobalReadProvisionedThroughputSettings

All content copied from https://docs.aws.amazon.com/.
