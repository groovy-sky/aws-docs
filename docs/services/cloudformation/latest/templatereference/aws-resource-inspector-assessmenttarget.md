This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Inspector::AssessmentTarget

The `AWS::Inspector::AssessmentTarget` resource is used to create Amazon
Inspector assessment targets, which specify the Amazon EC2 instances that will be analyzed
during an assessment run.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Inspector::AssessmentTarget",
  "Properties" : {
      "AssessmentTargetName" : String,
      "ResourceGroupArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::Inspector::AssessmentTarget
Properties:
  AssessmentTargetName: String
  ResourceGroupArn: String

```

## Properties

`AssessmentTargetName`

The name of the Amazon Inspector assessment target. The name must be unique within
the AWS account.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `140`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceGroupArn`

The ARN that specifies the resource group that is used to create the assessment
target. If `resourceGroupArn` is not specified, all EC2 instances in the current AWS account
and Region are included in the assessment target.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `ResourceGroupArn` of the new assessment target.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) that specifies the assessment target that is
created.

## Examples

### Declaring an Amazon Inspector Assessment Target Resource

The following examples show how to declare an
`AWS::Inspector::AssessmentTarget` resource to create an Amazon
Inspector assessment target.

#### JSON

```json

{
    "Type": "AWS::Inspector::AssessmentTarget",
    "Properties": {
        "AssessmentTargetName": "MyAssessmentTarget",
        "ResourceGroupArn": "arn:aws:inspector:us-west-2:123456789012:resourcegroup/0-AB6DMKnv"
    }
}
```

#### YAML

```yaml

Type: AWS::Inspector::AssessmentTarget
Properties:
  AssessmentTargetName: MyAssessmentTarget
  ResourceGroupArn: arn:aws:inspector:us-west-2:123456789012:resourcegroup/0-AB6DMKnv
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Inspector classic

AWS::Inspector::AssessmentTemplate

All content copied from https://docs.aws.amazon.com/.
