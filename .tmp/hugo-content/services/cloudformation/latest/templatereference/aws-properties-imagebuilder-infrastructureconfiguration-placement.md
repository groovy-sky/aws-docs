This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::InfrastructureConfiguration Placement

By default, EC2 instances run on shared tenancy hardware. This means that multiple
AWS accounts might share the same physical hardware. When you use dedicated hardware,
the physical server that hosts your instances is dedicated to your AWS account.
Instance placement settings contain the details for the physical hardware where
instances that Image Builder launches during image creation will run.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AvailabilityZone" : String,
  "HostId" : String,
  "HostResourceGroupArn" : String,
  "Tenancy" : String
}

```

### YAML

```yaml

  AvailabilityZone: String
  HostId: String
  HostResourceGroupArn: String
  Tenancy: String

```

## Properties

`AvailabilityZone`

The Availability Zone where your build and test instances will launch.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HostId`

The ID of the Dedicated Host on which build and test instances run. This only
applies if `tenancy` is `host`. If you specify the host ID, you
must not specify the resource group ARN. If you specify both, Image Builder returns an error.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HostResourceGroupArn`

The Amazon Resource Name (ARN) of the host resource group in which to launch build and test instances.
This only applies if `tenancy` is `host`. If you specify the resource
group ARN, you must not specify the host ID. If you specify both, Image Builder returns an error.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tenancy`

The tenancy of the instance. An instance with a tenancy of `dedicated`
runs on single-tenant hardware. An instance with a tenancy of `host` runs
on a Dedicated Host.

If tenancy is set to `host`, then you can optionally specify one target
for placement – either host ID or host resource group ARN. If automatic placement
is enabled for your host, and you don't specify any placement target, Amazon EC2 will try to
find an available host for your build and test instances.

_Required_: No

_Type_: String

_Allowed values_: `default | dedicated | host`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Logging

S3Logs

All content copied from https://docs.aws.amazon.com/.
