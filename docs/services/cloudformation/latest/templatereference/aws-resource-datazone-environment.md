This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::Environment

The `AWS::DataZone::Environment` resource specifies an Amazon DataZone environment, which is a
collection of zero or more configured resources with a given set of IAM principals who can operate on those
resources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataZone::Environment",
  "Properties" : {
      "DeploymentOrder" : Integer,
      "Description" : String,
      "DomainIdentifier" : String,
      "EnvironmentAccountIdentifier" : String,
      "EnvironmentAccountRegion" : String,
      "EnvironmentBlueprintIdentifier" : String,
      "EnvironmentConfigurationId" : String,
      "EnvironmentProfileIdentifier" : String,
      "EnvironmentRoleArn" : String,
      "GlossaryTerms" : [ String, ... ],
      "Name" : String,
      "ProjectIdentifier" : String,
      "UserParameters" : [ EnvironmentParameter, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DataZone::Environment
Properties:
  DeploymentOrder: Integer
  Description: String
  DomainIdentifier: String
  EnvironmentAccountIdentifier: String
  EnvironmentAccountRegion: String
  EnvironmentBlueprintIdentifier: String
  EnvironmentConfigurationId: String
  EnvironmentProfileIdentifier: String
  EnvironmentRoleArn: String
  GlossaryTerms:
    - String
  Name: String
  ProjectIdentifier: String
  UserParameters:
    - EnvironmentParameter

```

## Properties

`DeploymentOrder`

The deployment order of the environment.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description of the environment.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainIdentifier`

The identifier of the Amazon DataZone domain in which the environment is created.

_Required_: Yes

_Type_: String

_Pattern_: `^dzd[-_][a-zA-Z0-9_-]{1,36}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnvironmentAccountIdentifier`

The identifier of the AWS account in which an environment exists.

_Required_: No

_Type_: String

_Pattern_: `^\d{12}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnvironmentAccountRegion`

The AWS Region in which an environment exists.

_Required_: No

_Type_: String

_Pattern_: `^[a-z]{2}-[a-z]{4,10}-\d$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnvironmentBlueprintIdentifier`

The ID of the blueprint with which the environment is being created.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnvironmentConfigurationId`

The configuration ID with which the environment is created.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnvironmentProfileIdentifier`

The identifier of the environment profile that is used to create this Amazon DataZone
environment.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]{0,36}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnvironmentRoleArn`

The ARN of the environment role.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlossaryTerms`

The glossary terms that can be used in this Amazon DataZone environment.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the Amazon DataZone environment.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProjectIdentifier`

The identifier of the Amazon DataZone project in which this environment is created.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]{1,36}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UserParameters`

The user parameters of this Amazon DataZone environment.

_Required_: No

_Type_: Array of [EnvironmentParameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-environment-environmentparameter.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a string containing pipe-separated `DomainId` and
`EnvironmentId`, which uniquely identifies the environment. For example: `{ "Ref": "MyEnvironment"
    }` for the resource with the logical ID `MyEnvironment`, `Ref` returns
`DomainId|EnvironmentId`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AwsAccountId`

The identifier of the AWS account in which an environment exists.

`AwsAccountRegion`

The AWS Region in which an environment exists.

`CreatedAt`

The timestamp of when the environment was created.

`CreatedBy`

The Amazon DataZone user who created the environment.

`DomainId`

The identifier of the Amazon DataZone domain in which the environment exists.

`EnvironmentBlueprintId`

The identifier of a blueprint with which an environment profile is created.

`EnvironmentProfileId`

The identifier of the environment profile with which the environment was created.

`Id`

The identifier of the environment.

`ProjectId`

The identifier of the project in which the environment exists.

`Provider`

The provider of the environment.

`Status`

The status of the environment.

`UpdatedAt`

The timestamp of when the environment was updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::DataZone::DomainUnit

EnvironmentParameter
