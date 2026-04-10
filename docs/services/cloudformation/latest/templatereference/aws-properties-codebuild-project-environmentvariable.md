This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeBuild::Project EnvironmentVariable

`EnvironmentVariable` is a property of the [AWS CodeBuild Project Environment](../userguide/aws-properties-codebuild-project-environment.md)
property type that specifies the name and value of an environment variable for an AWS CodeBuild
project environment. When you use the environment to run a build, these variables are available for your builds to use. `EnvironmentVariable`
contains a list of `EnvironmentVariable` property types.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Type" : String,
  "Value" : String
}

```

### YAML

```yaml

  Name: String
  Type: String
  Value: String

```

## Properties

`Name`

The name or key of the environment variable.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of environment variable. Valid values include:

- `PARAMETER_STORE`: An environment variable stored in Systems Manager
Parameter Store. For environment variables of this type, specify the name of the parameter as the `value` of the
EnvironmentVariable. The parameter value will be substituted for the name at runtime. You can also define Parameter
Store environment variables in the buildspec. To learn how to do so,
see [env/parameter-store](../../../codebuild/latest/userguide/build-spec-ref.md#build-spec.env.parameter-store) in the
_AWS CodeBuild User Guide_.

- `PLAINTEXT`: An environment variable in plain text format. This is
the default value.

- `SECRETS_MANAGER`: An environment variable stored in AWS Secrets Manager. For environment variables of this type,
specify the name of the secret as the `value` of the EnvironmentVariable. The secret value will be substituted for the
name at runtime. You can also define AWS Secrets Manager environment variables in the buildspec. To learn how to do so, see
[env/secrets-manager](../../../codebuild/latest/userguide/build-spec-ref.md#build-spec.env.secrets-manager) in the
_AWS CodeBuild User Guide_.

_Required_: No

_Type_: String

_Allowed values_: `PLAINTEXT | PARAMETER_STORE | SECRETS_MANAGER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of the environment variable.

###### Important

We strongly discourage the use of `PLAINTEXT` environment variables to
store sensitive values, especially AWS secret key IDs.
`PLAINTEXT` environment variables can be displayed in plain text
using the AWS CodeBuild console and the AWS CLI. For sensitive values, we recommend you use an
environment variable of type `PARAMETER_STORE` or
`SECRETS_MANAGER`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

#### JSON

```json

{
  "Project": {
    "Environment": {
      "EnvironmentVariables": [
        {
          "Name": "MY_VAR_1",
          "Type": "PLAINTEXT",
          "Value": "VAR_1_VALUE"
        },
        {
          "Name": "MY_VAR_2",
          "Type": "PLAINTEXT",
          "Value": "VAR_2_VALUE"
        }
      ]
    }
  }
}
```

#### YAML

```yaml

Project:
  Type: AWS::CodeBuild::Project
  Properties:
    Environment:
      EnvironmentVariables:
        - Name: MY_VAR_1
          Type: PLAINTEXT
          Value: VAR_1_VALUE
        - Name: MY_VAR_2
          Type: PLAINTEXT
          Value: VAR_2_VALUE

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Environment

GitSubmodulesConfig

All content copied from https://docs.aws.amazon.com/.
