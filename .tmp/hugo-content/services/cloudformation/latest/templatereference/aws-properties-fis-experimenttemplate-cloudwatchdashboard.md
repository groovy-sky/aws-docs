This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FIS::ExperimentTemplate CloudWatchDashboard

The CloudWatch dashboards to include as data sources in the experiment report.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DashboardIdentifier" : String
}

```

### YAML

```yaml

  DashboardIdentifier: String

```

## Properties

`DashboardIdentifier`

The Amazon Resource Name (ARN) of the CloudWatch dashboard to include in the experiment report.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::FIS::ExperimentTemplate

CloudWatchLogsConfiguration

All content copied from https://docs.aws.amazon.com/.
