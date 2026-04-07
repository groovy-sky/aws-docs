This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Workteam MemberDefinition

Defines an Amazon Cognito or your own OIDC IdP user group that is part of a work team.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CognitoMemberDefinition" : CognitoMemberDefinition,
  "OidcMemberDefinition" : OidcMemberDefinition
}

```

### YAML

```yaml

  CognitoMemberDefinition:
    CognitoMemberDefinition
  OidcMemberDefinition:
    OidcMemberDefinition

```

## Properties

`CognitoMemberDefinition`

The Amazon Cognito user group that is part of the work team.

_Required_: No

_Type_: [CognitoMemberDefinition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-workteam-cognitomemberdefinition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OidcMemberDefinition`

A list user groups that exist in your OIDC Identity Provider (IdP).
One to ten groups can be used to create a single private work team.
When you add a user group to the list of `Groups`, you can add that user group to one or more
private work teams. If you add a user group to a private work team, all workers in that user group
are added to the work team.

_Required_: No

_Type_: [OidcMemberDefinition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-workteam-oidcmemberdefinition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CognitoMemberDefinition

NotificationConfiguration
