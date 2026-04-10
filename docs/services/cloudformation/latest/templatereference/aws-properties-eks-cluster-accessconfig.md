This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::Cluster AccessConfig

The access configuration for the cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthenticationMode" : String,
  "BootstrapClusterCreatorAdminPermissions" : Boolean
}

```

### YAML

```yaml

  AuthenticationMode: String
  BootstrapClusterCreatorAdminPermissions: Boolean

```

## Properties

`AuthenticationMode`

The desired authentication mode for the cluster. If you create a cluster by using the
EKS API, AWS SDKs, or AWS CloudFormation, the default is `CONFIG_MAP`. If you create
the cluster by using the AWS Management Console, the default value is
`API_AND_CONFIG_MAP`.

_Required_: No

_Type_: String

_Allowed values_: `CONFIG_MAP | API_AND_CONFIG_MAP | API`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BootstrapClusterCreatorAdminPermissions`

Specifies whether or not the cluster creator IAM principal was set as a cluster
admin access entry during cluster creation time. The default value is
`true`.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EKS::Cluster

BlockStorage

All content copied from https://docs.aws.amazon.com/.
