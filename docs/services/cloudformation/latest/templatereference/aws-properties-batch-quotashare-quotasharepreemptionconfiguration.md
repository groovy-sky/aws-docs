This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::QuotaShare QuotaSharePreemptionConfiguration

Specifies the preemption behavior for jobs in a quota share.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InSharePreemption" : String
}

```

### YAML

```yaml

  InSharePreemption: String

```

## Properties

`InSharePreemption`

Specifies whether jobs within a quota share can be preempted by another, higher priority job in the same quota share.

_Required_: Yes

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

QuotaShareCapacityLimit

QuotaShareResourceSharingConfiguration

All content copied from https://docs.aws.amazon.com/.
