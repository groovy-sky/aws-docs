This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMIncidents::ResponsePlan DynamicSsmParameterValue

The dynamic parameter value.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Variable" : String
}

```

### YAML

```yaml

  Variable: String

```

## Properties

`Variable`

Variable dynamic parameters. A parameter value is determined when an incident is
created.

_Required_: No

_Type_: String

_Allowed values_: `INCIDENT_RECORD_ARN | INVOLVED_RESOURCES`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DynamicSsmParameter

IncidentTemplate

All content copied from https://docs.aws.amazon.com/.
