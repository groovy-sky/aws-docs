This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeCommit::Repository RepositoryTrigger

Information about a trigger for a repository.

###### Note

If you want to receive notifications about repository events, consider using notifications instead of
triggers. For more information, see [Configuring\
notifications for repository events](../../../codecommit/latest/userguide/how-to-repository-email.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Branches" : [ String, ... ],
  "CustomData" : String,
  "DestinationArn" : String,
  "Events" : [ String, ... ],
  "Name" : String
}

```

### YAML

```yaml

  Branches:
    - String
  CustomData: String
  DestinationArn: String
  Events:
    - String
  Name: String

```

## Properties

`Branches`

The branches to be included in the trigger configuration. If you specify an empty
array, the trigger applies to all branches.

###### Note

Although no content is required in the array, you must include the array itself.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomData`

Any custom data associated with the trigger to be included in the information sent to
the target of the trigger.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationArn`

The ARN of the resource that is the target for a trigger (for example, the ARN of a
topic in Amazon SNS).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Events`

The repository events that cause the trigger to run actions in another service, such
as sending a notification through Amazon SNS.

###### Note

The valid value "all" cannot be used with any other values.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the trigger.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Code

S3

All content copied from https://docs.aws.amazon.com/.
