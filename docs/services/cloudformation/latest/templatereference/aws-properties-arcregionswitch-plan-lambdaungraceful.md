This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ARCRegionSwitch::Plan LambdaUngraceful

Configuration for handling failures when invoking Lambda functions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Behavior" :
}

```

### YAML

```yaml

  Behavior:

```

## Properties

`Behavior`

The ungraceful behavior for a Lambda function, which must be set to `skip`.

_Required_: No

_Type_:

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Lambdas

ParallelExecutionBlockConfiguration

All content copied from https://docs.aws.amazon.com/.
