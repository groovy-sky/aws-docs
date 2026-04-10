This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::Framework ControlScope

A framework consists of one or more controls. Each control has its own control scope.
The control scope can include one or more resource types, a combination of a tag key and
value, or a combination of one resource type and one resource ID. If no scope is specified,
evaluations for the rule are triggered when any resource in your recording group changes in
configuration.

###### Note

To set a control scope that includes all of a particular resource, leave the
`ControlScope` empty or do not pass it when calling
`CreateFramework`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComplianceResourceIds" : [ String, ... ],
  "ComplianceResourceTypes" : [ String, ... ],
  "Tags" : [ Tag, ... ]
}

```

### YAML

```yaml

  ComplianceResourceIds:
    - String
  ComplianceResourceTypes:
    - String
  Tags:
    - Tag

```

## Properties

`ComplianceResourceIds`

The ID of the only AWS resource that you want your control scope to
contain.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComplianceResourceTypes`

Describes whether the control scope includes one or more types of resources, such as
`EFS` or `RDS`.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tag key-value pair applied to those AWS resources that you want to
trigger an evaluation for a rule. A maximum of one key-value pair can be provided. The tag
value is optional, but it cannot be an empty string if you are creating or editing a
framework from the console (though the value can be an empty string when included
in a CloudFormation template).

The structure to assign a tag is:
`[{"Key":"string","Value":"string"}]`.

_Required_: No

_Type_: Array of [Tag](aws-properties-backup-framework-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ControlInputParameter

FrameworkControl

All content copied from https://docs.aws.amazon.com/.
