This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template ResourcePermission

Permission for the resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Actions" : [ String, ... ],
  "Principal" : String
}

```

### YAML

```yaml

  Actions:
    - String
  Principal: String

```

## Properties

`Actions`

The IAM action to grant or revoke permissions on.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Principal`

The Amazon Resource Name (ARN) of the principal. This can be one of the following:

- The ARN of an Amazon Quick user or group associated with a data source or dataset. (This is
common.)

- The ARN of an Amazon Quick user, group, or namespace associated with an analysis, dashboard,
template, or theme. (This is common.)

- The ARN of an AWS account root: This is an IAM ARN rather than a Quick ARN. Use this option only to share resources (templates) across AWS accounts. (This is
less common.)

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RelativeDateTimeControlDisplayOptions

RollingDateConfiguration
