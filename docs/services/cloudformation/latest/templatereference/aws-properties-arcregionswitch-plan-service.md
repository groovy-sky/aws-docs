This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ARCRegionSwitch::Plan Service

The service for a cross account role.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClusterArn" : String,
  "CrossAccountRole" : String,
  "ExternalId" : String,
  "ServiceArn" : String
}

```

### YAML

```yaml

  ClusterArn: String
  CrossAccountRole: String
  ExternalId: String
  ServiceArn: String

```

## Properties

`ClusterArn`

The cluster Amazon Resource Name (ARN) for a service.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[a-zA-Z-]*:ecs:[a-z0-9-]+:\d{12}:cluster/[a-zA-Z0-9_-]{1,255}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CrossAccountRole`

The cross account role for a service.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[a-zA-Z0-9-]*:iam::[0-9]{12}:role/.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExternalId`

The external ID (secret key) for the service.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceArn`

The Amazon Resource Name (ARN) for a service.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[a-zA-Z-]*:ecs:[a-z0-9-]+:\d{12}:service/[a-zA-Z0-9_-]+/[a-zA-Z0-9_-]{1,255}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3ReportOutputConfiguration

Step

All content copied from https://docs.aws.amazon.com/.
