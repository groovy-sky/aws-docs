This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ImageVersion

Creates a version of the SageMaker image specified by `ImageName`. The version represents the Amazon
Container Registry (ECR) container image specified by `BaseImage`.

###### Note

You can use the `DependsOn` attribute to specify that the creation of a specific resource follows
another. You can use it for the following use cases. For more information, see [`DependsOn`\
attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html).

1\. `DependsOn` can be used to establish a parent/child relationship between
`ImageVersion` and `Image` where the `ImageVersion` `DependsOn` the
`Image`.

2\. `DependsOn` can be used to establish order among `ImageVersion` s within the same
`Image` namespace. For example, if ImageVersionB `DependsOn` ImageVersionA and both share the
same parent `Image`, then ImageVersionA is version N and ImageVersionB is N+1.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SageMaker::ImageVersion",
  "Properties" : {
      "Alias" : String,
      "Aliases" : [ String, ... ],
      "BaseImage" : String,
      "Horovod" : Boolean,
      "ImageName" : String,
      "JobType" : String,
      "MLFramework" : String,
      "Processor" : String,
      "ProgrammingLang" : String,
      "ReleaseNotes" : String,
      "VendorGuidance" : String
    }
}

```

### YAML

```yaml

Type: AWS::SageMaker::ImageVersion
Properties:
  Alias: String
  Aliases:
    - String
  BaseImage: String
  Horovod: Boolean
  ImageName: String
  JobType: String
  MLFramework: String
  Processor: String
  ProgrammingLang: String
  ReleaseNotes: String
  VendorGuidance: String

```

## Properties

`Alias`

The alias for the image version.

_Required_: No

_Type_: String

_Pattern_: `(?!^[.-])^([a-zA-Z0-9-_.]+)$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Aliases`

A list of aliases for the image version.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BaseImage`

The container image that the SageMaker image version is based on.

_Required_: Yes

_Type_: String

_Pattern_: `.+`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Horovod`

Indicates whether the image version supports Horovod distributed training framework.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageName`

The name of the parent image.

_Length Constraints_: Minimum length of 1. Maximum length of 63.

_Pattern_: `^[a-zA-Z0-9]([-.]?[a-zA-Z0-9]){0,62}$`

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z0-9]([-.]?[A-Za-z0-9])*$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`JobType`

The job type that the image version supports (for example, TRAINING or INFERENCE).

_Required_: No

_Type_: String

_Allowed values_: `TRAINING | INFERENCE | NOTEBOOK_KERNEL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MLFramework`

The machine learning framework that the image version supports.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z]+ ?\d+\.\d+(\.\d+)?$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Processor`

The processor architecture that the image version supports (for example, CPU or GPU).

_Required_: No

_Type_: String

_Allowed values_: `CPU | GPU`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProgrammingLang`

The programming language that the image version supports.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z]+ ?\d+\.\d+(\.\d+)?$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReleaseNotes`

Release notes for the image version.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VendorGuidance`

Vendor guidance for the image version, such as stability or deprecation status.

_Required_: No

_Type_: String

_Allowed values_: `NOT_PROVIDED | STABLE | TO_BE_ARCHIVED | ARCHIVED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ImageVersionArn.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ContainerImage`

The URI of the container image version referenced by ImageVersion.

`ImageArn`

The Amazon Resource Name (ARN) of the parent Image.

`ImageVersionArn`

The Amazon Resource Name (ARN) of the image version.

_Type_: String

_Length Constraints_: Maximum length of 256.

_Pattern_:
`^arn:aws(-[\w]+)*:sagemaker:.+:[0-9]{12}:image-version/[a-z0-9]([-.]?[a-z0-9])*/[0-9]+$`

`Version`

The version of the image.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::SageMaker::InferenceComponent
