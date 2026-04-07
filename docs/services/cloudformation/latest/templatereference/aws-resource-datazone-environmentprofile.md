This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::EnvironmentProfile

The details of an environment profile.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataZone::EnvironmentProfile",
  "Properties" : {
      "AwsAccountId" : String,
      "AwsAccountRegion" : String,
      "Description" : String,
      "DomainIdentifier" : String,
      "EnvironmentBlueprintIdentifier" : String,
      "Name" : String,
      "ProjectIdentifier" : String,
      "UserParameters" : [ EnvironmentParameter, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DataZone::EnvironmentProfile
Properties:
  AwsAccountId: String
  AwsAccountRegion: String
  Description: String
  DomainIdentifier: String
  EnvironmentBlueprintIdentifier: String
  Name: String
  ProjectIdentifier: String
  UserParameters:
    - EnvironmentParameter

```

## Properties

`AwsAccountId`

The identifier of an AWS account in which an environment profile exists.

_Required_: Yes

_Type_: String

_Pattern_: `^\d{12}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AwsAccountRegion`

The AWS Region in which an environment profile exists.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z]{2}-[a-z]{4,10}-\d$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the environment profile.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainIdentifier`

The identifier of the Amazon DataZone domain in which the environment profile exists.

_Required_: Yes

_Type_: String

_Pattern_: `^dzd[-_][a-zA-Z0-9_-]{1,36}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnvironmentBlueprintIdentifier`

The identifier of a blueprint with which an environment profile is created.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]{1,36}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the environment profile.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w -]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProjectIdentifier`

The identifier of a project in which an environment profile exists.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]{1,36}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UserParameters`

The user parameters of this Amazon DataZone environment profile.

_Required_: No

_Type_: Array of [EnvironmentParameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-environmentprofile-environmentparameter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a string containing pipe-separated `DomainId` and the
`EnvironmentProfileId`, which uniquely identifies the environment profile. For example: `{ "Ref":
    "MyEnvironmentProfile" }` for the resource with the logical `ID MyEnvironmentProfile`,
`Ref` returns `DomainId|EnvironmentProfileId`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The timestamp of when an environment profile was created.

`CreatedBy`

The Amazon DataZone user who created the environment profile.

`DomainId`

The identifier of the Amazon DataZone domain in which the environment profile exists.

`EnvironmentBlueprintId`

The identifier of a blueprint with which an environment profile is created.

`Id`

The identifier of the environment profile.

`ProjectId`

The identifier of a project in which an environment profile exists.

`UpdatedAt`

The timestamp of when the environment profile was updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RegionalParameter

EnvironmentParameter
