This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMIncidents::ResponsePlan PagerDutyConfiguration

Details about the PagerDuty configuration for a response plan.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "PagerDutyIncidentConfiguration" : PagerDutyIncidentConfiguration,
  "SecretId" : String
}

```

### YAML

```yaml

  Name: String
  PagerDutyIncidentConfiguration:
    PagerDutyIncidentConfiguration
  SecretId: String

```

## Properties

`Name`

The name of the PagerDuty configuration.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PagerDutyIncidentConfiguration`

Details about the PagerDuty service associated with the configuration.

_Required_: Yes

_Type_: [PagerDutyIncidentConfiguration](aws-properties-ssmincidents-responseplan-pagerdutyincidentconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretId`

The ID of the AWS Secrets Manager secret that stores your PagerDuty key, either a
General Access REST API Key or User Token REST API Key, and other user
credentials.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NotificationTargetItem

PagerDutyIncidentConfiguration

All content copied from https://docs.aws.amazon.com/.
