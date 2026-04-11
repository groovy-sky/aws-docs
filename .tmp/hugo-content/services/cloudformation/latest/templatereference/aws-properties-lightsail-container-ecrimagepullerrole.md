This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Container EcrImagePullerRole

Describes the IAM role that you can use to grant a Lightsail container service access to Amazon ECR private repositories.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IsActive" : Boolean,
  "PrincipalArn" : String
}

```

### YAML

```yaml

  IsActive: Boolean
  PrincipalArn: String

```

## Properties

`IsActive`

A boolean value that indicates whether the `ECRImagePullerRole` is active.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrincipalArn`

The principle Amazon Resource Name (ARN) of the role. This property is read-only.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContainerServiceDeployment

EnvironmentVariable

All content copied from https://docs.aws.amazon.com/.
