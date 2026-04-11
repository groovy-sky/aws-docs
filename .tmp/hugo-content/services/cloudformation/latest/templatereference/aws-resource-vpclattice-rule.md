This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::Rule

Creates a listener rule. Each listener has a default rule for checking connection requests,
but you can define additional rules. Each rule consists of a priority, one or more actions, and
one or more conditions. For more information, see [Listener rules](../../../vpc-lattice/latest/ug/listeners.md#listener-rules) in the
_Amazon VPC Lattice User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::VpcLattice::Rule",
  "Properties" : {
      "Action" : Action,
      "ListenerIdentifier" : String,
      "Match" : Match,
      "Name" : String,
      "Priority" : Integer,
      "ServiceIdentifier" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::VpcLattice::Rule
Properties:
  Action:
    Action
  ListenerIdentifier: String
  Match:
    Match
  Name: String
  Priority: Integer
  ServiceIdentifier: String
  Tags:
    - Tag

```

## Properties

`Action`

Describes the action for a rule.

_Required_: Yes

_Type_: [Action](aws-properties-vpclattice-rule-action.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ListenerIdentifier`

The ID or ARN of the listener.

_Required_: No

_Type_: String

_Pattern_: `^((listener-[0-9a-z]{17})|(arn(:[a-z0-9]+([.-][a-z0-9]+)*){2}(:([a-z0-9]+([.-][a-z0-9]+)*)?){2}:service/svc-[0-9a-z]{17}/listener/listener-[0-9a-z]{17}))$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Match`

The rule match.

_Required_: Yes

_Type_: [Match](aws-properties-vpclattice-rule-match.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the rule. The name must be unique within the listener. The valid characters
are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last character, or
immediately after another hyphen.

If you don't specify a name, CloudFormation generates one. However, if
you specify a name, and later want to replace the resource, you must specify a new
name.

_Required_: No

_Type_: String

_Pattern_: `^(?!rule-)(?![-])(?!.*[-]$)(?!.*[-]{2})[a-z0-9-]+$`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Priority`

The priority assigned to the rule. Each rule for a specific listener must have a unique
priority. The lower the priority number the higher the priority.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceIdentifier`

The ID or ARN of the service.

_Required_: No

_Type_: String

_Pattern_: `^((svc-[0-9a-z]{17})|(arn(:[a-z0-9]+([.-][a-z0-9]+)*){2}(:([a-z0-9]+([.-][a-z0-9]+)*)?){2}:service/svc-[0-9a-z]{17}))$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags for the rule.

_Required_: No

_Type_: Array of [Tag](aws-properties-vpclattice-rule-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the rule.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the rule.

`Id`

The ID of the listener.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::VpcLattice::ResourcePolicy

Action

All content copied from https://docs.aws.amazon.com/.
