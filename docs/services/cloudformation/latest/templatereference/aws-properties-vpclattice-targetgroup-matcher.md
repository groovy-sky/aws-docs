This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::TargetGroup Matcher

Describes the codes to use when checking for a successful response from a target for health
checks.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HttpCode" : String
}

```

### YAML

```yaml

  HttpCode: String

```

## Properties

`HttpCode`

The HTTP code to use when checking for a successful response from a target.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9-,]+$`

_Minimum_: `3`

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HealthCheckConfig

Tag

All content copied from https://docs.aws.amazon.com/.
