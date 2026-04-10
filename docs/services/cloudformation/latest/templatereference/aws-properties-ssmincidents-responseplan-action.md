This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMIncidents::ResponsePlan Action

The `Action` property type specifies the configuration to launch.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SsmAutomation" : SsmAutomation
}

```

### YAML

```yaml

  SsmAutomation:
    SsmAutomation

```

## Properties

`SsmAutomation`

Details about the Systems Manager automation document that will be used as a
runbook during an incident.

_Required_: No

_Type_: [SsmAutomation](aws-properties-ssmincidents-responseplan-ssmautomation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SSMIncidents::ResponsePlan

ChatChannel

All content copied from https://docs.aws.amazon.com/.
