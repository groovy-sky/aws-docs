This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::ProjectProfile EnvironmentConfigurationParametersDetails

The details of the environment configuration parameter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ParameterOverrides" : [ EnvironmentConfigurationParameter, ... ],
  "ResolvedParameters" : [ EnvironmentConfigurationParameter, ... ],
  "SsmPath" : String
}

```

### YAML

```yaml

  ParameterOverrides:
    - EnvironmentConfigurationParameter
  ResolvedParameters:
    - EnvironmentConfigurationParameter
  SsmPath: String

```

## Properties

`ParameterOverrides`

The parameter overrides.

_Required_: No

_Type_: Array of [EnvironmentConfigurationParameter](aws-properties-datazone-projectprofile-environmentconfigurationparameter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResolvedParameters`

The resolved environment configuration parameters.

_Required_: No

_Type_: Array of [EnvironmentConfigurationParameter](aws-properties-datazone-projectprofile-environmentconfigurationparameter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SsmPath`

Ssm path environment configuration parameters.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EnvironmentConfigurationParameter

Region

All content copied from https://docs.aws.amazon.com/.
