This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53::RecordSet CidrRoutingConfig

The object that is specified in resource record set object when you are linking a
resource record set to a CIDR location.

A `LocationName` with an asterisk “\*” can be used to create a default CIDR
record. `CollectionId` is still required for default record.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CollectionId" : String,
  "LocationName" : String
}

```

### YAML

```yaml

  CollectionId: String
  LocationName: String

```

## Properties

`CollectionId`

The CIDR collection ID.

_Required_: Yes

_Type_: String

_Pattern_: `[0-9a-f]{8}-(?:[0-9a-f]{4}-){3}[0-9a-f]{12}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LocationName`

The CIDR collection location name.

_Required_: Yes

_Type_: String

_Pattern_: `[0-9A-Za-z_\-\*]+`

_Minimum_: `1`

_Maximum_: `16`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AliasTarget

Coordinates
