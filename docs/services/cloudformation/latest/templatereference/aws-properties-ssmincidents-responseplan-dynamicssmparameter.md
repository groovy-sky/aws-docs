This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMIncidents::ResponsePlan DynamicSsmParameter

When you add a runbook to a response plan, you can specify the parameters for the runbook
to use at runtime. Response plans support parameters with both static and dynamic
values. For static values, you enter the value when you define the parameter in the
response plan. For dynamic values, the system determines the correct parameter value by
collecting information from the incident. Incident Manager supports the
following dynamic parameters:

**Incident ARN**

When Incident Manager creates an incident, the system captures the Amazon
Resource Name (ARN) of the corresponding incident record and enters it for this
parameter in the runbook.

###### Note

This value can only be assigned to parameters of type `String`. If
assigned to a parameter of any other type, the runbook fails to run.

**Involved resources**

When Incident Manager creates an incident, the system captures the ARNs of
the resources involved in the incident. These resource ARNs are then assigned to this
parameter in the runbook.

###### Note

This value can only be assigned to parameters of type `StringList`. If
assigned to a parameter of any other type, the runbook fails to run.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : DynamicSsmParameterValue
}

```

### YAML

```yaml

  Key: String
  Value:
    DynamicSsmParameterValue

```

## Properties

`Key`

The key parameter to use when running the Systems Manager Automation
runbook.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The dynamic parameter value.

_Required_: Yes

_Type_: [DynamicSsmParameterValue](aws-properties-ssmincidents-responseplan-dynamicssmparametervalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ChatChannel

DynamicSsmParameterValue

All content copied from https://docs.aws.amazon.com/.
