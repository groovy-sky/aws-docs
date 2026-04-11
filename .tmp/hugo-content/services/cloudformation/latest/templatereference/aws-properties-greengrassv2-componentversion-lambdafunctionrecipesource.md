This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GreengrassV2::ComponentVersion LambdaFunctionRecipeSource

Contains information about an AWS Lambda function to import to create a
component.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComponentDependencies" : {Key: Value, ...},
  "ComponentLambdaParameters" : LambdaExecutionParameters,
  "ComponentName" : String,
  "ComponentPlatforms" : [ ComponentPlatform, ... ],
  "ComponentVersion" : String,
  "LambdaArn" : String
}

```

### YAML

```yaml

  ComponentDependencies:
    Key: Value
  ComponentLambdaParameters:
    LambdaExecutionParameters
  ComponentName: String
  ComponentPlatforms:
    - ComponentPlatform
  ComponentVersion: String
  LambdaArn: String

```

## Properties

`ComponentDependencies`

The component versions on which this Lambda function component
depends.

_Required_: No

_Type_: Object of [ComponentDependencyRequirement](aws-properties-greengrassv2-componentversion-componentdependencyrequirement.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ComponentLambdaParameters`

The system and runtime parameters for the Lambda function as it runs on the
AWS IoT Greengrass core device.

_Required_: No

_Type_: [LambdaExecutionParameters](aws-properties-greengrassv2-componentversion-lambdaexecutionparameters.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ComponentName`

The name of the component.

Defaults to the name of the Lambda function.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ComponentPlatforms`

The platforms that the component version supports.

_Required_: No

_Type_: Array of [ComponentPlatform](aws-properties-greengrassv2-componentversion-componentplatform.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ComponentVersion`

The version of the component.

Defaults to the version of the Lambda function as a semantic version. For
example, if your function version is `3`, the component version becomes
`3.0.0`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LambdaArn`

The ARN of the Lambda function. The ARN must include the version of the
function to import. You can't use version aliases like `$LATEST`.

_Required_: No

_Type_: String

_Pattern_: `^arn:[^:]*:lambda:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LambdaExecutionParameters

LambdaLinuxProcessParams

All content copied from https://docs.aws.amazon.com/.
