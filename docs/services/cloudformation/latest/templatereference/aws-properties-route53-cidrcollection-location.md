This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53::CidrCollection Location

Specifies the list of CIDR blocks for a CIDR location.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CidrList" : [ String, ... ],
  "LocationName" : String
}

```

### YAML

```yaml

  CidrList:
    - String
  LocationName: String

```

## Properties

`CidrList`

List of CIDR blocks.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LocationName`

The CIDR collection location name.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `16`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [LocationSummary](https://docs.aws.amazon.com/Route53/latest/APIReference/API_LocationSummary.html) in the _Amazon Route 53 API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Route53::CidrCollection

AWS::Route53::DNSSEC
