This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::Listener DefaultAction

The action for the default rule. Each listener has a default rule. The default rule is used
if no other rules match.

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

Describes an action that returns a custom HTTP response.

_Required_: No

_Type_: [FixedResponse](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-vpclattice-listener-fixedresponse.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Forward`

Describes a forward action. You can use forward actions to route requests to one or more
target groups.

_Required_: No

_Type_: [Forward](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-vpclattice-listener-forward.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::VpcLattice::Listener

FixedResponse
