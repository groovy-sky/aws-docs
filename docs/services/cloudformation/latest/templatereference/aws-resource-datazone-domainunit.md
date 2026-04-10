This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::DomainUnit

The summary of the domain unit.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataZone::DomainUnit",
  "Properties" : {
      "Description" : String,
      "DomainIdentifier" : String,
      "Name" : String,
      "ParentDomainUnitIdentifier" : String
    }
}

```

### YAML

```yaml

Type: AWS::DataZone::DomainUnit
Properties:
  Description: String
  DomainIdentifier: String
  Name: String
  ParentDomainUnitIdentifier: String

```

## Properties

`Description`

The description of the domain unit.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainIdentifier`

The ID of the domain where you want to crate a domain unit.

_Required_: Yes

_Type_: String

_Pattern_: `^dzd[-_][a-zA-Z0-9_-]{1,36}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the domain unit.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w -]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParentDomainUnitIdentifier`

The ID of the parent domain unit.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a string containing pipe-separated `DomainId` and
`DomainUnitId`, which uniquely identifies a domain unit. For example: { `"Ref":
    "MyDomainUnit"` } for the resource with the logical ID MyDomainUnit, Ref returns
`DomainId|DomainUnitId`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The time stamp at which the domain unit was created.

`DomainId`

The ID of the domain in which the domain unit lives.

`Id`

The ID of the domain unit.

`Identifier`

The identifier of the domain unit that you want to get.

`LastUpdatedAt`

The timestamp at which the domain unit was last updated.

`ParentDomainUnitId`

The ID of the parent domain unit.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::DataZone::Environment

All content copied from https://docs.aws.amazon.com/.
