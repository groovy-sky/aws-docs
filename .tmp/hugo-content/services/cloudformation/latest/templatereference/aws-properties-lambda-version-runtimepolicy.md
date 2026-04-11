This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::Version RuntimePolicy

Sets the runtime management configuration for a function's version. For more information,
see [Runtime updates](../../../lambda/latest/dg/runtimes-update.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RuntimeVersionArn" : String,
  "UpdateRuntimeOn" : String
}

```

### YAML

```yaml

  RuntimeVersionArn: String
  UpdateRuntimeOn: String

```

## Properties

`RuntimeVersionArn`

The ARN of the runtime version you want the function to use.

###### Note

This is only required if you're using the **Manual** runtime update mode.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*):lambda:(eusc-)?[a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1}::runtime:.+$`

_Minimum_: `26`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UpdateRuntimeOn`

Specify the runtime update mode.

- **Auto (default)** \- Automatically update to the most recent and secure runtime version using a [Two-phase runtime version rollout](../../../lambda/latest/dg/runtimes-update.md#runtime-management-two-phase). This is the best
choice for most customers to ensure they always benefit from runtime updates.

- **FunctionUpdate** \- Lambda updates the runtime of you function to the most recent and secure runtime version when you update your
function. This approach synchronizes runtime updates with function deployments, giving you control over when runtime updates are applied and allowing you to detect and
mitigate rare runtime update incompatibilities early. When using this setting, you need to regularly update your functions to keep their runtime up-to-date.

- **Manual** \- You specify a runtime version in your function configuration. The function will use this runtime version indefinitely.
In the rare case where a new runtime version is incompatible with an existing function, this allows you to roll back your function to an earlier runtime version. For more information,
see [Roll back a runtime version](../../../lambda/latest/dg/runtimes-update.md#runtime-management-rollback).

_Valid Values_: `Auto` \| `FunctionUpdate` \| `Manual`

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProvisionedConcurrencyConfiguration

Next

All content copied from https://docs.aws.amazon.com/.
