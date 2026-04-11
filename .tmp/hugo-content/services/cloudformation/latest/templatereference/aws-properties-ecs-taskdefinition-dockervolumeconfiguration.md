This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::TaskDefinition DockerVolumeConfiguration

The `DockerVolumeConfiguration` property specifies a Docker volume
configuration and is used when you use Docker volumes. Docker volumes are only supported
when you are using the EC2 launch type. Windows containers only support the use of the
`local` driver. To use bind mounts, specify a `host`
instead.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Autoprovision" : Boolean,
  "Driver" : String,
  "DriverOpts" : {Key: Value, ...},
  "Labels" : {Key: Value, ...},
  "Scope" : String
}

```

### YAML

```yaml

  Autoprovision: Boolean
  Driver: String
  DriverOpts:
    Key: Value
  Labels:
    Key: Value
  Scope: String

```

## Properties

`Autoprovision`

If this value is `true`, the Docker volume is created if it doesn't already
exist.

###### Note

This field is only used if the `scope` is `shared`.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Driver`

The Docker volume driver to use. The driver value must match the driver name provided
by Docker because it is used for task placement. If the driver was installed using the
Docker plugin CLI, use `docker plugin ls` to retrieve the driver name from
your container instance. If the driver was installed using another method, use Docker
plugin discovery to retrieve the driver name. This parameter maps to `Driver`
in the docker container create command and the `xxdriver` option to docker
volume create.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DriverOpts`

A map of Docker driver-specific options passed through. This parameter maps to
`DriverOpts` in the docker create-volume command and the
`xxopt` option to docker volume create.

_Required_: No

_Type_: Object of String

_Pattern_: `.{1,}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Labels`

Custom metadata to add to your Docker volume. This parameter maps to
`Labels` in the docker container create command and the
`xxlabel` option to docker volume create.

_Required_: No

_Type_: Object of String

_Pattern_: `.{1,}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Scope`

The scope for the Docker volume that determines its lifecycle. Docker volumes that are
scoped to a `task` are automatically provisioned when the task starts and
destroyed when the task stops. Docker volumes that are scoped as `shared`
persist after the task stops.

_Required_: No

_Type_: String

_Allowed values_: `task | shared`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Device

EFSVolumeConfiguration

All content copied from https://docs.aws.amazon.com/.
