This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LicenseManager::License

Specifies a granted license.

Granted licenses are licenses for products that your organization purchased from AWS Marketplace
or directly from a seller who integrated their software with managed entitlements. For more information,
see [Granted \
licenses](https://docs.aws.amazon.com/license-manager/latest/userguide/granted-licenses.html) in the _AWS License Manager User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::LicenseManager::License",
  "Properties" : {
      "Beneficiary" : String,
      "ConsumptionConfiguration" : ConsumptionConfiguration,
      "Entitlements" : [ Entitlement, ... ],
      "HomeRegion" : String,
      "Issuer" : IssuerData,
      "LicenseMetadata" : [ Metadata, ... ],
      "LicenseName" : String,
      "ProductName" : String,
      "ProductSKU" : String,
      "Status" : String,
      "Tags" : [ Tag, ... ],
      "Validity" : ValidityDateFormat
    }
}

```

### YAML

```yaml

Type: AWS::LicenseManager::License
Properties:
  Beneficiary: String
  ConsumptionConfiguration:
    ConsumptionConfiguration
  Entitlements:
    - Entitlement
  HomeRegion: String
  Issuer:
    IssuerData
  LicenseMetadata:
    - Metadata
  LicenseName: String
  ProductName: String
  ProductSKU: String
  Status: String
  Tags:
    - Tag
  Validity:
    ValidityDateFormat

```

## Properties

`Beneficiary`

License beneficiary.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConsumptionConfiguration`

Configuration for consumption of the license.

_Required_: Yes

_Type_: [ConsumptionConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-licensemanager-license-consumptionconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Entitlements`

License entitlements.

_Required_: Yes

_Type_: Array of [Entitlement](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-licensemanager-license-entitlement.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HomeRegion`

Home Region of the license.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Issuer`

License issuer.

_Required_: Yes

_Type_: [IssuerData](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-licensemanager-license-issuerdata.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LicenseMetadata`

License metadata.

_Required_: No

_Type_: Array of [Metadata](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-licensemanager-license-metadata.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LicenseName`

License name.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProductName`

Product name.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProductSKU`

Product SKU.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

License status.

_Required_: No

_Type_: String

_Allowed values_: `AVAILABLE | PENDING_AVAILABLE | DEACTIVATED | SUSPENDED | EXPIRED | PENDING_DELETE | DELETED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

One or more tags.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-licensemanager-license-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Validity`

Date and time range during which the license is valid, in ISO8601-UTC format.

_Required_: Yes

_Type_: [ValidityDateFormat](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-licensemanager-license-validitydateformat.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the license.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`LicenseArn`

The Amazon Resource Name (ARN) of the license.

`Version`

The license version.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

BorrowConfiguration
