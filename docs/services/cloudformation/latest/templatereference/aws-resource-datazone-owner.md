This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::Owner

The owner that you want to add to the entity.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataZone::Owner",
  "Properties" : {
      "DomainIdentifier" : String,
      "EntityIdentifier" : String,
      "EntityType" : String,
      "Owner" : OwnerProperties
    }
}

```

### YAML

```yaml

Type: AWS::DataZone::Owner
Properties:
  DomainIdentifier: String
  EntityIdentifier: String
  EntityType: String
  Owner:
    OwnerProperties

```

## Properties

`DomainIdentifier`

The ID of the domain in which you want to add the entity owner.

_Required_: Yes

_Type_: String

_Pattern_: `^dzd[-_][a-zA-Z0-9_-]{1,36}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EntityIdentifier`

The ID of the entity to which you want to add an owner.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EntityType`

The type of an entity.

_Required_: Yes

_Type_: String

_Allowed values_: `DOMAIN_UNIT`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Owner`

The owner that you want to add to the entity.

_Required_: Yes

_Type_: [OwnerProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-owner-ownerproperties.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a string containing pipe-separated `DomainId`, `EntityType`,
`EntityId`, `OwnerType`, and `OwnerId`, which uniquely identifies an owner. For
example: `{ "Ref": "MyOwner" }` for the resource with the logical ID MyOwner, Ref returns
`DomainId|EntityType|EntityId|OwnerType|OwnerId`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`OwnerIdentifier`

The ID of the entity to which you want to add an owner.

`OwnerType`

The owner that you want to add to the entity.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::DataZone::GroupProfile

OwnerGroupProperties
