This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceCatalog::CloudFormationProduct SourceConnection

A top level `ProductViewDetail` response containing details about the product’s connection.
AWS Service Catalog returns this field for the `CreateProduct`, `UpdateProduct`,
`DescribeProductAsAdmin`, and `SearchProductAsAdmin` APIs.
This response contains the same fields as the `ConnectionParameters` request, with the
addition of the `LastSync` response.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConnectionParameters" : ConnectionParameters,
  "Type" : String
}

```

### YAML

```yaml

  ConnectionParameters:
    ConnectionParameters
  Type: String

```

## Properties

`ConnectionParameters`

The connection details based on the connection `Type`.

_Required_: Yes

_Type_: [ConnectionParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-servicecatalog-cloudformationproduct-connectionparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The only supported `SourceConnection` type is Codestar.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ProvisioningArtifactProperties

Tag
