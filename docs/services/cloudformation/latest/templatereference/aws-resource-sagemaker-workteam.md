This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Workteam

Creates a new work team for labeling your data. A work team is defined by one or more
Amazon Cognito user pools. You must first create the user pools before you can create a work
team.

You cannot create more than 25 work teams in an account and region.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SageMaker::Workteam",
  "Properties" : {
      "Description" : String,
      "MemberDefinitions" : [ MemberDefinition, ... ],
      "NotificationConfiguration" : NotificationConfiguration,
      "Tags" : [ Tag, ... ],
      "WorkforceName" : String,
      "WorkteamName" : String
    }
}

```

### YAML

```yaml

Type: AWS::SageMaker::Workteam
Properties:
  Description: String
  MemberDefinitions:
    - MemberDefinition
  NotificationConfiguration:
    NotificationConfiguration
  Tags:
    - Tag
  WorkforceName: String
  WorkteamName: String

```

## Properties

`Description`

A description of the work team.

_Required_: No

_Type_: String

_Pattern_: `.+`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MemberDefinitions`

A list of `MemberDefinition` objects that contains objects that identify
the workers that make up the work team.

Workforces can be created using Amazon Cognito or your own OIDC Identity Provider (IdP).
For private workforces created using Amazon Cognito use
`CognitoMemberDefinition`. For workforces created using your own OIDC identity
provider (IdP) use `OidcMemberDefinition`.

_Required_: No

_Type_: Array of [MemberDefinition](aws-properties-sagemaker-workteam-memberdefinition.md)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotificationConfiguration`

Configures SNS notifications of available or expiring work items for work
teams.

_Required_: No

_Type_: [NotificationConfiguration](aws-properties-sagemaker-workteam-notificationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs.

_Required_: No

_Type_: Array of [Tag](aws-properties-sagemaker-workteam-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkforceName`

The name of the workforce.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9]([a-zA-Z0-9\-]){0,62}`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`WorkteamName`

The name of the work team.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The name of the work team.

`WorkteamName`

The name of the work team.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UserSettings

CognitoMemberDefinition

All content copied from https://docs.aws.amazon.com/.
