This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::CapacityProvider ManagedInstancesLocalStorageConfiguration

The local storage configuration for Amazon ECS Managed Instances. This defines how
ECS uses and configures instance store volumes available on container instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "UseLocalStorage" : Boolean
}

```

### YAML

```yaml

  UseLocalStorage: Boolean

```

## Properties

`UseLocalStorage`

Use instance store volumes for data storage when available. EBS volumes are not
provisioned for data storage. If the container instance has multiple instance store volumes,
a single data volume is created. Consider defining instance store requirements using the
`localStorage`, `localStorageTypes` and `totalLocalStorageGB`
properties.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InstanceRequirementsRequest

ManagedInstancesNetworkConfiguration

All content copied from https://docs.aws.amazon.com/.
