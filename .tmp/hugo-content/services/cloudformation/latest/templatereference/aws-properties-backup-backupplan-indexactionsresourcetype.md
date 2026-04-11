This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::BackupPlan IndexActionsResourceType

Specifies index actions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ResourceTypes" : [ String, ... ]
}

```

### YAML

```yaml

  ResourceTypes:
    - String

```

## Properties

`ResourceTypes`

0 or 1 index action will be accepted for each BackupRule.

Valid values:

- `EBS` for Amazon Elastic Block Store

- `S3` for Amazon Simple Storage Service (Amazon S3)

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CopyActionResourceType

LifecycleResourceType

All content copied from https://docs.aws.amazon.com/.
