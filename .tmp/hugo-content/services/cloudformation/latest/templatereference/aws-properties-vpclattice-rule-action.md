This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::Rule Action

Describes the action for a rule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FixedResponse" : FixedResponse,
  "Forward" : Forward
}

```

### YAML

```yaml

  FixedResponse:
    FixedResponse
  Forward:
    Forward

```

## Properties

`FixedResponse`

The fixed response action. The rule returns a custom HTTP response.

_Required_: No

_Type_: [FixedResponse](aws-properties-vpclattice-rule-fixedresponse.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Forward`

The forward action. Traffic that matches the rule is forwarded to the specified target
groups.

_Required_: No

_Type_: [Forward](aws-properties-vpclattice-rule-forward.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::VpcLattice::Rule

FixedResponse

All content copied from https://docs.aws.amazon.com/.
