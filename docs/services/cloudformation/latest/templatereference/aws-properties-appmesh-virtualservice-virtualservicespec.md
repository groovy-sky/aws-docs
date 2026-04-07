This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualService VirtualServiceSpec

An object that represents the specification of a virtual service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Provider" : VirtualServiceProvider
}

```

### YAML

```yaml

  Provider:
    VirtualServiceProvider

```

## Properties

`Provider`

The App Mesh object that is acting as the provider for a virtual service. You
can specify a single virtual node or virtual router.

_Required_: No

_Type_: [VirtualServiceProvider](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appmesh-virtualservice-virtualserviceprovider.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VirtualServiceProvider

Next
