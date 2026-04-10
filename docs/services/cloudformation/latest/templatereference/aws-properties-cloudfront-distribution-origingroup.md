This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution OriginGroup

An origin group includes two origins (a primary origin and a secondary origin to failover
to) and a failover criteria that you specify. You create an origin group to support
origin failover in CloudFront. When you create or update a distribution, you can specify the
origin group instead of a single origin, and CloudFront will failover from the primary origin
to the secondary origin under the failover conditions that you've chosen.

Optionally, you can choose selection criteria for your origin group to specify how your origins are selected when your distribution routes viewer requests.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FailoverCriteria" : OriginGroupFailoverCriteria,
  "Id" : String,
  "Members" : OriginGroupMembers,
  "SelectionCriteria" : String
}

```

### YAML

```yaml

  FailoverCriteria:
    OriginGroupFailoverCriteria
  Id: String
  Members:
    OriginGroupMembers
  SelectionCriteria: String

```

## Properties

`FailoverCriteria`

A complex type that contains information about the failover criteria for an origin
group.

_Required_: Yes

_Type_: [OriginGroupFailoverCriteria](aws-properties-cloudfront-distribution-origingroupfailovercriteria.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

The origin group's ID.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Members`

A complex type that contains information about the origins in an origin group.

_Required_: Yes

_Type_: [OriginGroupMembers](aws-properties-cloudfront-distribution-origingroupmembers.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectionCriteria`

The selection criteria for the origin group. For more information, see [Create an origin group](../../../amazoncloudfront/latest/developerguide/high-availability-origin-failover.md#concept_origin_groups.creating) in the _Amazon CloudFront_
_Developer Guide_.

_Required_: No

_Type_: String

_Allowed values_: `default | media-quality-based`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OriginCustomHeader

OriginGroupFailoverCriteria

All content copied from https://docs.aws.amazon.com/.
