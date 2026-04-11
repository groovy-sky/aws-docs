This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis CustomActionSetParametersOperation

The set parameter operation that sets parameters in custom action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ParameterValueConfigurations" : [ SetParameterValueConfiguration, ... ]
}

```

### YAML

```yaml

  ParameterValueConfigurations:
    - SetParameterValueConfiguration

```

## Properties

`ParameterValueConfigurations`

The parameter that determines the value configuration.

_Required_: Yes

_Type_: Array of [SetParameterValueConfiguration](aws-properties-quicksight-analysis-setparametervalueconfiguration.md)

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomActionNavigationOperation

CustomActionURLOperation

All content copied from https://docs.aws.amazon.com/.
