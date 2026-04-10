This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::Service EBSTagSpecification

The tag specifications of an Amazon EBS volume.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PropagateTags" : String,
  "ResourceType" : String,
  "Tags" : [ Tag, ... ]
}

```

### YAML

```yaml

  PropagateTags: String
  ResourceType: String
  Tags:
    - Tag

```

## Properties

`PropagateTags`

Determines whether to propagate the tags from the task definition to the
Amazon EBS volume. Tags can only propagate to a `SERVICE` specified in
`ServiceVolumeConfiguration`. If no value is specified, the tags
aren't propagated.

_Required_: No

_Type_: String

_Allowed values_: `SERVICE | TASK_DEFINITION`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceType`

The type of volume resource.

_Required_: Yes

_Type_: String

_Allowed values_: `volume`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags applied to this Amazon EBS volume. `AmazonECSCreated` and
`AmazonECSManaged` are reserved tags that can't be used.

_Required_: No

_Type_: Array of [Tag](aws-properties-ecs-service-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeploymentLifecycleHook

ForceNewDeployment

All content copied from https://docs.aws.amazon.com/.
