This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution OriginGroupMembers

A complex data type for the origins included in an origin group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Items" : [ OriginGroupMember, ... ],
  "Quantity" : Integer
}

```

### YAML

```yaml

  Items:
    - OriginGroupMember
  Quantity: Integer

```

## Properties

`Items`

Items (origins) in an origin group.

_Required_: Yes

_Type_: Array of [OriginGroupMember](aws-properties-cloudfront-distribution-origingroupmember.md)

_Minimum_: `2`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Quantity`

The number of origins in an origin group.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OriginGroupMember

OriginGroups

All content copied from https://docs.aws.amazon.com/.
