This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition ConsumableResourceRequirement

Information about a consumable resource required to run a job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConsumableResource" : String,
  "Quantity" : Integer
}

```

### YAML

```yaml

  ConsumableResource: String
  Quantity: Integer

```

## Properties

`ConsumableResource`

The name or ARN of the consumable resource.

_Required_: Yes

_Type_: String

_Pattern_: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Quantity`

The quantity of the consumable resource that is needed.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConsumableResourceProperties

ContainerProperties

All content copied from https://docs.aws.amazon.com/.
