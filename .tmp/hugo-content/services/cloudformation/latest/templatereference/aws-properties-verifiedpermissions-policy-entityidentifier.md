This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VerifiedPermissions::Policy EntityIdentifier

Contains the identifier of an entity in a policy, including its ID and type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EntityId" : String,
  "EntityType" : String
}

```

### YAML

```yaml

  EntityId: String
  EntityType: String

```

## Properties

`EntityId`

The identifier of an entity.

`"entityId":"identifier"`

_Required_: Yes

_Type_: String

_Pattern_: `^.*$`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EntityType`

The type of an entity.

Example: `"entityType":"typeName"`

_Required_: Yes

_Type_: String

_Pattern_: `^.*$`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::VerifiedPermissions::Policy

PolicyDefinition

All content copied from https://docs.aws.amazon.com/.
