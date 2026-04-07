This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::Service ServiceVolumeConfiguration

The configuration for a volume specified in the task definition as a volume that is
configured at launch time. Currently, the only supported volume type is an Amazon EBS
volume.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ManagedEBSVolume" : ServiceManagedEBSVolumeConfiguration,
  "Name" : String
}

```

### YAML

```yaml

  ManagedEBSVolume:
    ServiceManagedEBSVolumeConfiguration
  Name: String

```

## Properties

`ManagedEBSVolume`

The configuration for the Amazon EBS volume that Amazon ECS creates and manages on
your behalf. These settings are used to create each Amazon EBS volume, with one volume
created for each task in the service. The Amazon EBS volumes are visible in your account
in the Amazon EC2 console once they are created.

_Required_: No

_Type_: [ServiceManagedEBSVolumeConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ecs-service-servicemanagedebsvolumeconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the volume. This value must match the volume name from the
`Volume` object in the task definition.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ServiceRegistry

Tag
