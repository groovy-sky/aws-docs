This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMIncidents::ResponsePlan SsmAutomation

The `SsmAutomation` property type specifies details about the
Systems Manager Automation runbook that will be used as the runbook during an incident.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DocumentName" : String,
  "DocumentVersion" : String,
  "DynamicParameters" : [ DynamicSsmParameter, ... ],
  "Parameters" : [ SsmParameter, ... ],
  "RoleArn" : String,
  "TargetAccount" : String
}

```

### YAML

```yaml

  DocumentName: String
  DocumentVersion: String
  DynamicParameters:
    - DynamicSsmParameter
  Parameters:
    - SsmParameter
  RoleArn: String
  TargetAccount: String

```

## Properties

`DocumentName`

The automation document's name.

_Required_: Yes

_Type_: String

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocumentVersion`

The version of the runbook to use when running.

_Required_: No

_Type_: String

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DynamicParameters`

The key-value pairs to resolve dynamic parameter values when processing a Systems Manager Automation runbook.

_Required_: No

_Type_: Array of [DynamicSsmParameter](aws-properties-ssmincidents-responseplan-dynamicssmparameter.md)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

The key-value pair parameters to use when running the runbook.

_Required_: No

_Type_: Array of [SsmParameter](aws-properties-ssmincidents-responseplan-ssmparameter.md)

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the role that the automation document will assume
when running commands.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(-(cn|us-gov))?:[a-z-]+:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetAccount`

The account that the automation document will be run in. This can be in either the
management account or an application account.

_Required_: No

_Type_: String

_Allowed values_: `IMPACTED_ACCOUNT | RESPONSE_PLAN_OWNER_ACCOUNT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PagerDutyIncidentConfiguration

SsmParameter

All content copied from https://docs.aws.amazon.com/.
