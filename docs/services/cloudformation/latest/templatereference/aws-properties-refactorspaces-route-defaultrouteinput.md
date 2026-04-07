This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RefactorSpaces::Route DefaultRouteInput

The configuration for the default route type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActivationState" : String
}

```

### YAML

```yaml

  ActivationState: String

```

## Properties

`ActivationState`

If set to `ACTIVE`, traffic is forwarded to this route’s service after the
route is created.

_Required_: Yes

_Type_: String

_Allowed values_: `INACTIVE | ACTIVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::RefactorSpaces::Route

Tag
