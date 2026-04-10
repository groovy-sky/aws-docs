This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::Cluster ElasticLoadBalancing

Indicates the current configuration of the load balancing capability on your EKS Auto
Mode cluster. For example, if the capability is enabled or disabled. For more
information, see EKS Auto Mode load balancing capability in the _Amazon EKS User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean
}

```

### YAML

```yaml

  Enabled: Boolean

```

## Properties

`Enabled`

Indicates if the load balancing capability is enabled on your EKS Auto Mode cluster.
If the load balancing capability is enabled, EKS Auto Mode will create and delete load
balancers in your AWS account.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ControlPlaneScalingConfig

EncryptionConfig

All content copied from https://docs.aws.amazon.com/.
