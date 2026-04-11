This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::GroupProfile

The details of a group profile in Amazon DataZone.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataZone::GroupProfile",
  "Properties" : {
      "DomainIdentifier" : String,
      "GroupIdentifier" : String,
      "Status" : String
    }
}

```

### YAML

```yaml

Type: AWS::DataZone::GroupProfile
Properties:
  DomainIdentifier: String
  GroupIdentifier: String
  Status: String

```

## Properties

`DomainIdentifier`

The identifier of the Amazon DataZone domain in which a group profile exists.

_Required_: Yes

_Type_: String

_Pattern_: `^dzd[-_][a-zA-Z0-9_-]{1,36}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GroupIdentifier`

The ID of the group of a project member.

_Required_: Yes

_Type_: String

_Pattern_: `(^([0-9a-f]{10}-|)[A-Fa-f0-9]{8}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{12}$|[\p{L}\p{M}\p{S}\p{N}\p{P}\t\n\r  ]+)`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Status`

The status of a group profile.

_Required_: No

_Type_: String

_Allowed values_: `ASSIGNED | NOT_ASSIGNED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a string containing pipe-separated `DomainId` and
`GroupProfileId` that uniquely identify the group profile. For example: `{ "Ref": "MyGroupProfile"
    }` for the resource with the logical ID `MyGroupProfile`, `Ref` returns
`DomainId|GroupProfileId`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DomainId`

The identifier of the Amazon DataZone domain in which a group profile exists.

`GroupName`

The name of a group profile.

`Id`

The ID of a group profile.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Model

AWS::DataZone::Owner

All content copied from https://docs.aws.amazon.com/.
