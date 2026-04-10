This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ConfigurationSet DashboardOptions

An object containing additional settings for your VDM configuration as applicable to
the Dashboard.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EngagementMetrics" : String
}

```

### YAML

```yaml

  EngagementMetrics: String

```

## Properties

`EngagementMetrics`

Specifies the status of your VDM engagement metrics collection. Can be one of the
following:

- `ENABLED` – Amazon SES enables engagement metrics for the
configuration set.

- `DISABLED` – Amazon SES disables engagement metrics for the
configuration set.

_Required_: Yes

_Type_: String

_Pattern_: `ENABLED|DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConditionThreshold

DeliveryOptions

All content copied from https://docs.aws.amazon.com/.
