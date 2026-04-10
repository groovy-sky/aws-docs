This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMIncidents::ResponsePlan Integration

Information about third-party services integrated into a response plan.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PagerDutyConfiguration" : PagerDutyConfiguration
}

```

### YAML

```yaml

  PagerDutyConfiguration:
    PagerDutyConfiguration

```

## Properties

`PagerDutyConfiguration`

Information about the PagerDuty service where the response plan creates an
incident.

_Required_: No

_Type_: [PagerDutyConfiguration](aws-properties-ssmincidents-responseplan-pagerdutyconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IncidentTemplate

NotificationTargetItem

All content copied from https://docs.aws.amazon.com/.
