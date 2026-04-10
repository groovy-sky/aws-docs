This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::SchedulingPolicy ShareAttributes

Specifies the weights for the share identifiers for the fair-share policy. Share
identifiers that aren't included have a default weight of `1.0`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ShareIdentifier" : String,
  "WeightFactor" : Number
}

```

### YAML

```yaml

  ShareIdentifier: String
  WeightFactor: Number

```

## Properties

`ShareIdentifier`

A share identifier or share identifier prefix. If the string ends with an asterisk
(\*), this entry specifies the weight factor to use for share identifiers that start with
that prefix. The list of share identifiers in a fair-share policy can't overlap. For
example, you can't have one that specifies a `shareIdentifier` of `UserA*`
and another that specifies a `shareIdentifier` of `UserA1`.

There can be no more than 500 share identifiers active in a job queue.

The string is limited to 255 alphanumeric characters, and can be followed by an asterisk
(\*).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WeightFactor`

The weight factor for the share identifier. The default value is 1.0. A lower value has
a higher priority for compute resources. For example, jobs that use a share identifier with a
weight factor of 0.125 (1/8) get 8 times the compute resources of jobs that use a share
identifier with a weight factor of 1.

The smallest supported value is 0.0001, and the largest supported value is 999.9999.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [ShareAttributes](../../../../reference/batch/latest/apireference/api-shareattributes.md) in the _AWS Batch API Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

QuotaSharePolicy

AWS::Batch::ServiceEnvironment

All content copied from https://docs.aws.amazon.com/.
