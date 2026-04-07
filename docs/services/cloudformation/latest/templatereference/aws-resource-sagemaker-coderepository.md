This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::CodeRepository

Creates a Git repository as a resource in your SageMaker AI account. You can
associate the repository with notebook instances so that you can use Git source control
for the notebooks you create. The Git repository is a resource in your SageMaker AI
account, so it can be associated with more than one notebook instance, and it persists
independently from the lifecycle of any notebook instances it is associated with.

The repository can be hosted either in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html)
or in any other Git repository.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SageMaker::CodeRepository",
  "Properties" : {
      "CodeRepositoryName" : String,
      "GitConfig" : GitConfig,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SageMaker::CodeRepository
Properties:
  CodeRepositoryName: String
  GitConfig:
    GitConfig
  Tags:
    - Tag

```

## Properties

`CodeRepositoryName`

The name of the Git repository.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GitConfig`

Configuration details for the Git repository, including the URL where it is located
and the ARN of the AWS Secrets Manager secret that contains the
credentials used to access the repository.

_Required_: Yes

_Type_: [GitConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-coderepository-gitconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

List of tags for Code Repository.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-coderepository-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the code repository.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CodeRepositoryName`

The name of the code repository, such as `myCodeRepo`.

`Id`

The name of the code repository.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VpcConfig

GitConfig
