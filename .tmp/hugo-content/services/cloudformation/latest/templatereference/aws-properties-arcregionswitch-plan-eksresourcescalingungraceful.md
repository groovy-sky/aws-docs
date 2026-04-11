This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ARCRegionSwitch::Plan EksResourceScalingUngraceful

The ungraceful settings for AWS EKS resource scaling.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MinimumSuccessPercentage" : Number
}

```

### YAML

```yaml

  MinimumSuccessPercentage: Number

```

## Properties

`MinimumSuccessPercentage`

The minimum success percentage for the configuration.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `99`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EksResourceScalingConfiguration

ExecutionApprovalConfiguration

All content copied from https://docs.aws.amazon.com/.
