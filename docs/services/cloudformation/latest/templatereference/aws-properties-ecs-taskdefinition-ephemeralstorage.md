This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::TaskDefinition EphemeralStorage

The amount of ephemeral storage to allocate for the task. This parameter is used to
expand the total amount of ephemeral storage available, beyond the default amount, for
tasks hosted on AWS Fargate. For more information, see [Using data volumes in\
tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_data_volumes.html) in the _Amazon ECS Developer Guide;_.

###### Note

For tasks using the Fargate launch type, the task requires the following
platforms:

- Linux platform version `1.4.0` or later.

- Windows platform version `1.0.0` or later.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SizeInGiB" : Integer
}

```

### YAML

```yaml

  SizeInGiB: Integer

```

## Properties

`SizeInGiB`

The total amount, in GiB, of ephemeral storage to set for the task. The minimum
supported value is `21` GiB and the maximum supported value is
`200` GiB.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EnvironmentFile

FirelensConfiguration
