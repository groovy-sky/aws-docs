This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RTBFabric::ResponderGateway EksEndpointsConfiguration

Describes the configuration of an Amazon Elastic Kubernetes Service endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClusterApiServerCaCertificateChain" : String,
  "ClusterApiServerEndpointUri" : String,
  "ClusterName" : String,
  "EndpointsResourceName" : String,
  "EndpointsResourceNamespace" : String,
  "RoleArn" : String
}

```

### YAML

```yaml

  ClusterApiServerCaCertificateChain: String
  ClusterApiServerEndpointUri: String
  ClusterName: String
  EndpointsResourceName: String
  EndpointsResourceNamespace: String
  RoleArn: String

```

## Properties

`ClusterApiServerCaCertificateChain`

The CA certificate chain of the cluster API server.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2097152`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`ClusterApiServerEndpointUri`

The URI of the cluster API server endpoint.

_Required_: Yes

_Type_: String

_Pattern_: `^(https|http)://(?:[A-Za-z0-9](?:[A-Za-z0-9-]{0,61}[A-Za-z0-9])?)(?:\.(?:[A-Za-z0-9](?:[A-Za-z0-9-]{0,61}[A-Za-z0-9])?))+$`

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`ClusterName`

The name of the cluster.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z0-9][A-Za-z0-9-]{0,61}[A-Za-z0-9]$`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`EndpointsResourceName`

The name of the endpoint resource.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z0-9][A-Za-z0-9-]{0,61}[A-Za-z0-9]$`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`EndpointsResourceNamespace`

The namespace of the endpoint resource.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z0-9][A-Za-z0-9-]{0,61}[A-Za-z0-9]$`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`RoleArn`

The role ARN for the cluster.

_Required_: Yes

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AutoScalingGroupsConfiguration

ManagedEndpointConfiguration

All content copied from https://docs.aws.amazon.com/.
