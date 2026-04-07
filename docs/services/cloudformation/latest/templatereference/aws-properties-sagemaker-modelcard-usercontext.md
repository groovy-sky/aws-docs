This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelCard UserContext

Information about the user who created or modified a SageMaker resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DomainId" : String,
  "UserProfileArn" : String,
  "UserProfileName" : String
}

```

### YAML

```yaml

  DomainId: String
  UserProfileArn: String
  UserProfileName: String

```

## Properties

`DomainId`

The domain associated with the user.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserProfileArn`

The Amazon Resource Name (ARN) of the user's profile.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserProfileName`

The name of the user's profile.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TrainingMetric

AWS::SageMaker::ModelExplainabilityJobDefinition
