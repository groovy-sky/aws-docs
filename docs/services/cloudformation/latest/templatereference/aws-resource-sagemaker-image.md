This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Image

Creates a custom SageMaker AI image. A SageMaker AI image is a set of image versions. Each image
version represents a container image stored in Amazon ECR. For more information, see
[Bring your own SageMaker AI image](https://docs.aws.amazon.com/sagemaker/latest/dg/studio-byoi.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SageMaker::Image",
  "Properties" : {
      "ImageDescription" : String,
      "ImageDisplayName" : String,
      "ImageName" : String,
      "ImageRoleArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SageMaker::Image
Properties:
  ImageDescription: String
  ImageDisplayName: String
  ImageName: String
  ImageRoleArn: String
  Tags:
    - Tag

```

## Properties

`ImageDescription`

The description of the image.

_Required_: No

_Type_: String

_Pattern_: `.+`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageDisplayName`

The display name of the image.

_Length Constraints_: Minimum length of 1. Maximum length of 128.

_Pattern_: `^\S(.*\S)?$`

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9 -_]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageName`

The name of the Image. Must be unique by region in your account.

_Length Constraints_: Minimum length of 1. Maximum length of 63.

_Pattern_: `^[a-zA-Z0-9]([-.]?[a-zA-Z0-9]){0,62}$`

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9]([-.]?[a-zA-Z0-9])*$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ImageRoleArn`

The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on your behalf.

_Length Constraints_: Minimum length of 20. Maximum length of 2048.

_Pattern_: `^arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(-[\w]+)*:iam::[0-9]{12}:role/.*$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of key-value pairs to apply to this resource.

_Array Members_: Minimum number of 0 items. Maximum number of 50 items.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-image-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ImageArn.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ImageArn`

The Amazon Resource Name (ARN) of the image.

_Type_: String

_Length Constraints_: Maximum length of 256.

_Pattern_:
`^arn:aws(-[\w]+)*:sagemaker:.+:[0-9]{12}:image/[a-z0-9]([-.]?[a-z0-9])*$`

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TtlDuration

Tag
