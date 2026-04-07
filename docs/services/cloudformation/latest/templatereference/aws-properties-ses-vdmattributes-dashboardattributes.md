This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::VdmAttributes DashboardAttributes

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

- `ENABLED` – Amazon SES enables engagement metrics for your
account.

- `DISABLED` – Amazon SES disables engagement metrics for your
account.

_Required_: No

_Type_: String

_Pattern_: `ENABLED|DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::SES::VdmAttributes

GuardianAttributes
