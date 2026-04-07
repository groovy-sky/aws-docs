This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::ProjectMembership Member

The details about a project member.

Important - this data type is a UNION, so only one of the following members can be specified when used or
returned.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GroupIdentifier" : String,
  "UserIdentifier" : String
}

```

### YAML

```yaml

  GroupIdentifier: String
  UserIdentifier: String

```

## Properties

`GroupIdentifier`

The ID of the group of a project member.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UserIdentifier`

The user ID of a project member.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::DataZone::ProjectMembership

AWS::DataZone::ProjectProfile
