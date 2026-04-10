This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ResourceGroups::TagSyncTask

Onboards and syncs resources tagged with a specific tag key-value pair to an application.

**Minimum permissions**

To run this command, you must have the following permissions:

- `resource-groups:StartTagSyncTask`

- `resource-groups:CreateGroup`

- `iam:PassRole` for the role you provide to create a tag-sync task

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ResourceGroups::TagSyncTask",
  "Properties" : {
      "Group" : String,
      "RoleArn" : String,
      "TagKey" : String,
      "TagValue" : String
    }
}

```

### YAML

```yaml

Type: AWS::ResourceGroups::TagSyncTask
Properties:
  Group: String
  RoleArn: String
  TagKey: String
  TagValue: String

```

## Properties

`Group`

The Amazon resource name (ARN) or name of the application group for which you want to create a tag-sync task.

_Required_: Yes

_Type_: String

_Pattern_: `([a-zA-Z0-9_\\.-]{1,150}/[a-z0-9]{26})|(arn:aws(-[a-z]+)*:resource-groups(-(test|beta|gamma))?:[a-z]{2}(-[a-z]+)+-\d{1}:[0-9]{12}:group/[a-zA-Z0-9_\\.-]{1,150}/[a-z0-9]{26})`

_Minimum_: `12`

_Maximum_: `1600`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The Amazon resource name (ARN) of the role assumed by the service to tag and untag resources on your behalf.

_Required_: Yes

_Type_: String

_Pattern_: `arn:(aws[a-zA-Z-]*)?:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TagKey`

The tag key.

_Required_: Yes

_Type_: String

_Pattern_: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TagValue`

The tag value.

_Required_: Yes

_Type_: String

_Pattern_: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

The `TaskArn`.

### Fn::GetAtt

`GroupArn`

The Amazon resource name (ARN) of the application group.

`GroupName`

The name of the application group.

`Status`

The status of the tag-sync task.

Valid values include:

- `ACTIVE` \- The tag-sync task is actively managing resources in
the application by adding or removing the `awsApplication` tag from resources
when they are tagged or untagged with the specified tag key-value pair.

- `ERROR` \- The tag-sync task is not actively managing resources
in the application. Review the `ErrorMessage` for more information about
resolving the error.

`TaskArn`

The Amazon resource name (ARN) of the tag-sync task.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TagFilter

Next

All content copied from https://docs.aws.amazon.com/.
