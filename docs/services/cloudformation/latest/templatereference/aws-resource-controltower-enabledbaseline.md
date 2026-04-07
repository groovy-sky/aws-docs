This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ControlTower::EnabledBaseline

The `AWS::ControlTower::EnabledBaseline` resource Property description not available. for ControlTower.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ControlTower::EnabledBaseline",
  "Properties" : {
      "BaselineIdentifier" : String,
      "BaselineVersion" : String,
      "Parameters" : [ Parameter, ... ],
      "Tags" : [ Tag, ... ],
      "TargetIdentifier" : String
    }
}

```

### YAML

```yaml

Type: AWS::ControlTower::EnabledBaseline
Properties:
  BaselineIdentifier: String
  BaselineVersion: String
  Parameters:
    - Parameter
  Tags:
    - Tag
  TargetIdentifier: String

```

## Properties

`BaselineIdentifier`

The specific `Baseline` enabled as part of the `EnabledBaseline` resource.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[0-9a-zA-Z_\-:\/]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BaselineVersion`

The enabled version of the `Baseline`.

_Required_: Yes

_Type_: String

_Pattern_: `^\d+(?:\.\d+){0,2}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

Shows the parameters that are applied when enabling this `Baseline`.

_Required_: No

_Type_: Array of [Parameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-controltower-enabledbaseline-parameter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-controltower-enabledbaseline-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetIdentifier`

The target on which to enable the `Baseline`.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[0-9a-zA-Z_\-:\/]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`EnabledBaselineIdentifier`

Property description not available.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Control Tower

Parameter
