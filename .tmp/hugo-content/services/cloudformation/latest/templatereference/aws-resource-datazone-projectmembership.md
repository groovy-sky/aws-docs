This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::ProjectMembership

The `AWS::DataZone::ProjectMembership` resource adds a member to an Amazon DataZone project. Project
members consume assets from the Amazon DataZone catalog and produce new assets using one or more analytical
workflows.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataZone::ProjectMembership",
  "Properties" : {
      "Designation" : String,
      "DomainIdentifier" : String,
      "Member" : Member,
      "ProjectIdentifier" : String
    }
}

```

### YAML

```yaml

Type: AWS::DataZone::ProjectMembership
Properties:
  Designation: String
  DomainIdentifier: String
  Member:
    Member
  ProjectIdentifier: String

```

## Properties

`Designation`

The designated role of a project member.

_Required_: Yes

_Type_: String

_Allowed values_: `PROJECT_OWNER | PROJECT_CONTRIBUTOR | PROJECT_CATALOG_VIEWER | PROJECT_CATALOG_CONSUMER | PROJECT_CATALOG_STEWARD`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainIdentifier`

The ID of the Amazon DataZone domain in which project membership is created.

_Required_: Yes

_Type_: String

_Pattern_: `^dzd[-_][a-zA-Z0-9_-]{1,36}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Member`

The details about a project member.

_Required_: Yes

_Type_: [Member](aws-properties-datazone-projectmembership-member.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProjectIdentifier`

The ID of the project for which this project membership was created.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]{1,36}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a string containing pipe-separated `DomainId`, `MemberId`,
`MemberType`, and `ProjectId` that uniquely identify the project membership. For example:
`{ "Ref": "MyProjectMembership" }` for the resource with the logical ID `MyProjectMembership`,
`Ref` returns `DomainId|MemberId|MemberType|ProjectId`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`MemberIdentifier`

The identifier of the project member.

`MemberIdentifierType`

The type of the project member identifier.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceTag

Member

All content copied from https://docs.aws.amazon.com/.
