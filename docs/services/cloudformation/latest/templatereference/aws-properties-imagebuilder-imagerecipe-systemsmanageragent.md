This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::ImageRecipe SystemsManagerAgent

Contains settings for the Systems Manager agent on your build instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "UninstallAfterBuild" : Boolean
}

```

### YAML

```yaml

  UninstallAfterBuild: Boolean

```

## Properties

`UninstallAfterBuild`

Controls whether the Systems Manager agent is removed from your final build image, prior to
creating the new AMI. If this is set to true, then the agent is removed from the final
image. If it's set to false, then the agent is left in, so that it is included in the
new AMI. default value is false.

The default behavior of uninstallAfterBuild is to remove the SSM Agent if it was installed by EC2 Image Builder

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LatestVersion

AWS::ImageBuilder::InfrastructureConfiguration

All content copied from https://docs.aws.amazon.com/.
