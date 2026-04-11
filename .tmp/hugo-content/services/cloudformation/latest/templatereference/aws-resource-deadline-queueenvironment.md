This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Deadline::QueueEnvironment

Creates an environment for a queue that defines how jobs in the queue run.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Deadline::QueueEnvironment",
  "Properties" : {
      "FarmId" : String,
      "Priority" : Integer,
      "QueueId" : String,
      "Template" : String,
      "TemplateType" : String
    }
}

```

### YAML

```yaml

Type: AWS::Deadline::QueueEnvironment
Properties:
  FarmId: String
  Priority: Integer
  QueueId: String
  Template: String
  TemplateType: String

```

## Properties

`FarmId`

The identifier assigned to the farm that contains the queue.

_Required_: Yes

_Type_: String

_Pattern_: `^farm-[0-9a-f]{32}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Priority`

The queue environment's priority.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueueId`

The unique identifier of the queue that contains the environment.

_Required_: Yes

_Type_: String

_Pattern_: `^queue-[0-9a-f]{32}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Template`

A JSON or YAML template that describes the processing environment for the queue.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `15000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateType`

Specifies whether the template for the queue environment is JSON or YAML.

_Required_: Yes

_Type_: String

_Allowed values_: `JSON | YAML`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a reference to a queue environment object.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Name`

The name of the queue environment.

`QueueEnvironmentId`

The queue environment ID.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WindowsUser

AWS::Deadline::QueueFleetAssociation

All content copied from https://docs.aws.amazon.com/.
