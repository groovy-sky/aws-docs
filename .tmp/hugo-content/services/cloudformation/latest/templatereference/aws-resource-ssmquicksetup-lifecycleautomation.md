This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMQuickSetup::LifecycleAutomation

Creates a lifecycle automation resource that executes SSM Automation documents during CloudFormation stack operations. This resource replaces inline AWS Lambda custom resources and provides a managed way to handle lifecycle events in Quick Setup configurations.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SSMQuickSetup::LifecycleAutomation",
  "Properties" : {
      "AutomationDocument" : String,
      "AutomationParameters" : {Key: Value, ...},
      "ResourceKey" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::SSMQuickSetup::LifecycleAutomation
Properties:
  AutomationDocument: String
  AutomationParameters:
    Key: Value
  ResourceKey: String
  Tags:
    Key: Value

```

## Properties

`AutomationDocument`

The name of the SSM Automation document to execute in response to CloudFormation lifecycle events (CREATE, UPDATE, DELETE).

_Required_: Yes

_Type_: String

_Pattern_: `^\S+$`

_Minimum_: `1`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutomationParameters`

A map of key-value parameters passed to the Automation document during execution. Each parameter name maps to a list of values, even for single values. Parameters can include configuration-specific values for your automation workflow.

_Required_: Yes

_Type_: Object of Array

_Pattern_: `^[a-zA-Z0-9_]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceKey`

A unique identifier used for generating the SSM Association name. This ensures uniqueness when multiple lifecycle automation resources exist in the same stack.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Tags applied to the underlying SSM Association created by this resource. Tags help identify and organize automation executions.

_Required_: No

_Type_: Object of String

_Pattern_: `^[A-Za-z0-9 +=@_\/:.-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

Returns the AssociationId of the lifecycle automation resource, which is the same as the association ID of the underlying Systems Manager association.

### Fn::GetAtt

Returns the value of an attribute from the `AWS::SSMQuickSetup::LifecycleAutomation` resource. This resource executes SSM Automation documents in response to CloudFormation lifecycle events (CREATE, UPDATE, DELETE) and replaces inline Lambda custom resources in Quick Setup templates.

`AssociationId`

Returns the ID of the SSM Association created to manage the automation document execution lifecycle.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StatusSummary

Next

All content copied from https://docs.aws.amazon.com/.
