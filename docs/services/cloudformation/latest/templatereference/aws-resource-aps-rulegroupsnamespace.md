This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::APS::RuleGroupsNamespace

The definition of a rule groups namespace in an Amazon Managed Service for Prometheus
workspace. A rule groups namespace is associated with exactly one rules file. A
workspace can have multiple rule groups namespaces. For more information about rules
files, see [Creating a rules\
file](../../../prometheus/latest/userguide/amp-ruler-rulesfile.md), in the _Amazon Managed Service for Prometheus User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::APS::RuleGroupsNamespace",
  "Properties" : {
      "Data" : String,
      "Name" : String,
      "Tags" : [ Tag, ... ],
      "Workspace" : String
    }
}

```

### YAML

```yaml

Type: AWS::APS::RuleGroupsNamespace
Properties:
  Data: String
  Name: String
  Tags:
    - Tag
  Workspace: String

```

## Properties

`Data`

The rules file used in the namespace.

For more details about the rules file, see [Creating a rules\
file](../../../prometheus/latest/userguide/amp-ruler-rulesfile.md) in the _Amazon Managed Service for Prometheus User Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the rule groups namespace.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The list of tag keys and values that are associated with the rule groups
namespace.

_Required_: No

_Type_: Array of [Tag](aws-properties-aps-rulegroupsnamespace-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Workspace`

The ID of the workspace to add the rule groups namespace.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws|aws-us-gov|aws-cn):aps:[a-z0-9-]+:[0-9]+:workspace/[a-zA-Z0-9-]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:

`{ "Ref": "Arn" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the rule groups namespace. For example,
`arn:aws:aps:<region>:123456789012:rulegroupsnamespace/ws-example1-1234-abcd-5678-ef90abcd1234/rulesfile1`.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::APS::ResourcePolicy

Tag

All content copied from https://docs.aws.amazon.com/.
