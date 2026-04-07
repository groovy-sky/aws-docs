This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::InspectorV2::CisScanConfiguration

The CIS scan configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::InspectorV2::CisScanConfiguration",
  "Properties" : {
      "ScanName" : String,
      "Schedule" : Schedule,
      "SecurityLevel" : String,
      "Tags" : {Key: Value, ...},
      "Targets" : CisTargets
    }
}

```

### YAML

```yaml

Type: AWS::InspectorV2::CisScanConfiguration
Properties:
  ScanName: String
  Schedule:
    Schedule
  SecurityLevel: String
  Tags:
    Key: Value
  Targets:
    CisTargets

```

## Properties

`ScanName`

The name of the CIS scan configuration.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Schedule`

The CIS scan configuration's schedule.

_Required_: Yes

_Type_: [Schedule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-cisscanconfiguration-schedule.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityLevel`

The CIS scan configuration's CIS Benchmark level.

_Required_: Yes

_Type_: String

_Allowed values_: `LEVEL_1 | LEVEL_2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The CIS scan configuration's tags.

_Required_: No

_Type_: Object of String

_Pattern_: `^.{2,127}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Targets`

The CIS scan configuration's targets.

_Required_: Yes

_Type_: [CisTargets](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-cisscanconfiguration-cistargets.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the CIS scan configuration. For
example:

`arn:aws:inspector2:us-east-1:012345678901:owner/012345678901/cis-configuration/c1c0fe5d28e39baa`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The CIS scan configuration's scan configuration ARN.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Inspector

CisTargets
