This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution OriginGroups

A complex data type for the origin groups specified for a distribution.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Items" : [ OriginGroup, ... ],
  "Quantity" : Integer
}

```

### YAML

```yaml

  Items:
    - OriginGroup
  Quantity: Integer

```

## Properties

`Items`

The items (origin groups) in a distribution.

_Required_: No

_Type_: Array of [OriginGroup](aws-properties-cloudfront-distribution-origingroup.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Quantity`

The number of origin groups.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OriginGroupMembers

OriginMtlsConfig

All content copied from https://docs.aws.amazon.com/.
