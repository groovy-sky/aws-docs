This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ARCRegionSwitch::Plan KubernetesResourceType

Defines the type of Kubernetes resource to scale in an Amazon EKS cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApiVersion" : String,
  "Kind" : String
}

```

### YAML

```yaml

  ApiVersion: String
  Kind: String

```

## Properties

`ApiVersion`

The API version type for the Kubernetes resource.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Kind`

The kind for the Kubernetes resource.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GlobalAuroraUngraceful

KubernetesScalingResource

All content copied from https://docs.aws.amazon.com/.
