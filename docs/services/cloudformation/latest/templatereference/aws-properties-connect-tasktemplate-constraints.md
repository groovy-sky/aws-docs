This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::TaskTemplate Constraints

Describes constraints that apply to the template fields.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InvisibleFields" : [ InvisibleFieldInfo, ... ],
  "ReadOnlyFields" : [ ReadOnlyFieldInfo, ... ],
  "RequiredFields" : [ RequiredFieldInfo, ... ]
}

```

### YAML

```yaml

  InvisibleFields:
    - InvisibleFieldInfo
  ReadOnlyFields:
    - ReadOnlyFieldInfo
  RequiredFields:
    - RequiredFieldInfo

```

## Properties

`InvisibleFields`

Lists the fields that are invisible to agents.

_Required_: No

_Type_: Array of [InvisibleFieldInfo](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-tasktemplate-invisiblefieldinfo.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReadOnlyFields`

Lists the fields that are read-only to agents, and cannot be edited.

_Required_: No

_Type_: Array of [ReadOnlyFieldInfo](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-tasktemplate-readonlyfieldinfo.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequiredFields`

Lists the fields that are required to be filled by agents.

_Required_: No

_Type_: Array of [RequiredFieldInfo](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-tasktemplate-requiredfieldinfo.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Connect::TaskTemplate

DefaultFieldValue
