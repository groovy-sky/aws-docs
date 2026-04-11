This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::Listener FixedResponse

Describes an action that returns a custom HTTP response.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "StatusCode" : Integer
}

```

### YAML

```yaml

  StatusCode: Integer

```

## Properties

`StatusCode`

The HTTP response code. Only `404` and `500` status codes are supported.

_Required_: Yes

_Type_: Integer

_Minimum_: `100`

_Maximum_: `599`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DefaultAction

Forward

All content copied from https://docs.aws.amazon.com/.
