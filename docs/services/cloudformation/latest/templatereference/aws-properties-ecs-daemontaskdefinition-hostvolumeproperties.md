This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::DaemonTaskDefinition HostVolumeProperties

Details on a container instance bind mount host volume.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SourcePath" : String
}

```

### YAML

```yaml

  SourcePath: String

```

## Properties

`SourcePath`

When the `host` parameter is used, specify a `sourcePath` to
declare the path on the host container instance that's presented to the container. If
this parameter is empty, then the Docker daemon has assigned a host path for you. If the
`host` parameter contains a `sourcePath` file location, then
the data volume persists at the specified location on the host container instance until
you delete it manually. If the `sourcePath` value doesn't exist on the host
container instance, the Docker daemon creates it. If the location does exist, the
contents of the source path folder are exported.

If you're using the Fargate launch type, the `sourcePath` parameter is not
supported.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HealthCheck

KernelCapabilities

All content copied from https://docs.aws.amazon.com/.
