This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard LinkSharingConfiguration

A structure that contains the configuration of a shareable link to the
dashboard.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Permissions" : [ ResourcePermission, ... ]
}

```

### YAML

```yaml

  Permissions:
    - ResourcePermission

```

## Properties

`Permissions`

A structure that contains the permissions of a shareable link.

_Required_: No

_Type_: Array of [ResourcePermission](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-dashboard-resourcepermission.html)

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LineSeriesAxisDisplayOptions

ListControlDisplayOptions
