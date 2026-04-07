This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelPackageGroup

A group of versioned models in the Model Registry.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SageMaker::ModelPackageGroup",
  "Properties" : {
      "ModelPackageGroupDescription" : String,
      "ModelPackageGroupName" : String,
      "ModelPackageGroupPolicy" : Json,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SageMaker::ModelPackageGroup
Properties:
  ModelPackageGroupDescription: String
  ModelPackageGroupName: String
  ModelPackageGroupPolicy: Json
  Tags:
    - Tag

```

## Properties

`ModelPackageGroupDescription`

The description for the model group.

_Required_: No

_Type_: String

_Pattern_: `[\p{L}\p{M}\p{Z}\p{S}\p{N}\p{P}]*`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModelPackageGroupName`

The name of the model group.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*$`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModelPackageGroupPolicy`

A resouce policy to control access to a model group. For information about resoure policies, see [Identity-based policies\
and resource-based policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_identity-vs-resource.html) in the _AWS Identity and Access Management User_
_Guide._.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelpackagegroup-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the model package group.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreationTime`

The time when the model group was created.

`ModelPackageGroupArn`

The Amazon Resource Name (ARN) of the model group.

`ModelPackageGroupStatus`

The status of the model group.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ValidationSpecification

Tag
