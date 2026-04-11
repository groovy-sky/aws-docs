This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationInsights::Application Alarm

The `AWS::ApplicationInsights::Application Alarm` property type defines a CloudWatch alarm to be monitored for the component.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AlarmName" : String,
  "Severity" : String
}

```

### YAML

```yaml

  AlarmName: String
  Severity: String

```

## Properties

`AlarmName`

The name of the CloudWatch alarm to be monitored for the component.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Severity`

Indicates the degree of outage when the alarm goes off.

_Required_: No

_Type_: String

_Allowed values_: `HIGH | MEDIUM | LOW`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ApplicationInsights::Application

AlarmMetric

All content copied from https://docs.aws.amazon.com/.
