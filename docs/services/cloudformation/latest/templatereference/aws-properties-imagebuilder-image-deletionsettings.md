This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::Image DeletionSettings

Contains deletion settings of underlying resources of an image when it is replaced or deleted, including its Amazon Machine Images (AMIs), snapshots, or containers.

###### Note

If you specify the `Retain` option in the [DeletionPolicy](aws-attribute-updatereplacepolicy.md) or [UpdateReplacePolicy](aws-attribute-deletionpolicy.md), the deletion of underlying resources will not be executed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExecutionRole" : String
}

```

### YAML

```yaml

  ExecutionRole: String

```

## Properties

`ExecutionRole`

The name or Amazon Resource Name (ARN) for the IAM role you create that grants Image Builder access to delete the image and its underlying resources.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ImageBuilder::Image

EcrConfiguration
