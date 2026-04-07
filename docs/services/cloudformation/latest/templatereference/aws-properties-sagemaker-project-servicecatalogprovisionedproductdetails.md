This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Project ServiceCatalogProvisionedProductDetails

Details of a provisioned service catalog product. For information about service catalog,
see [What is AWS Service\
Catalog](https://docs.aws.amazon.com/servicecatalog/latest/adminguide/introduction.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ProvisionedProductId" : String,
  "ProvisionedProductStatusMessage" : String
}

```

### YAML

```yaml

  ProvisionedProductId: String
  ProvisionedProductStatusMessage: String

```

## Properties

`ProvisionedProductId`

The ID of the provisioned product.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*$`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProvisionedProductStatusMessage`

The current status of the product.

- `AVAILABLE` \- Stable state, ready to perform any operation. The most recent operation succeeded and completed.

- `UNDER_CHANGE` \- Transitive state. Operations performed might not have valid results. Wait for an AVAILABLE status before performing operations.

- `TAINTED` \- Stable state, ready to perform any operation. The stack has completed the requested operation but is not exactly what was requested. For example, a request to update to a new version failed and the stack rolled back to the current version.

- `ERROR` \- An unexpected error occurred. The provisioned product exists but the stack is not running. For example, CloudFormation received a parameter value that was not valid and could not launch the stack.

- `PLAN_IN_PROGRESS` \- Transitive state. The plan operations were performed to provision a new product, but resources have not yet been created. After reviewing the list of resources to be created, execute the plan. Wait for an AVAILABLE status before performing operations.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ProvisioningParameter

ServiceCatalogProvisioningDetails
