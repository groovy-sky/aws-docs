This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMIncidents::ResponsePlan PagerDutyIncidentConfiguration

Details about the PagerDuty service where the response plan creates an
incident.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ServiceId" : String
}

```

### YAML

```yaml

  ServiceId: String

```

## Properties

`ServiceId`

The ID of the PagerDuty service that the response plan associates with an incident
when it launches.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PagerDutyConfiguration

SsmAutomation

All content copied from https://docs.aws.amazon.com/.
