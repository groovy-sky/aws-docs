This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::Project EnvironmentConfigurationUserParameter

The environment configuration user parameters.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnvironmentConfigurationName" : String,
  "EnvironmentId" : String,
  "EnvironmentParameters" : [ EnvironmentParameter, ... ]
}

```

### YAML

```yaml

  EnvironmentConfigurationName: String
  EnvironmentId: String
  EnvironmentParameters:
    - EnvironmentParameter

```

## Properties

`EnvironmentConfigurationName`

The environment configuration name.

_Required_: No

_Type_: String

_Pattern_: `^[\w -]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentId`

The ID of the environment.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]{1,36}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentParameters`

The environment parameters.

_Required_: No

_Type_: Array of [EnvironmentParameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-project-environmentparameter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::DataZone::Project

EnvironmentParameter
