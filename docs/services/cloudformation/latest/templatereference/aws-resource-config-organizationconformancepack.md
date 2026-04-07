This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Config::OrganizationConformancePack

OrganizationConformancePack deploys conformance packs across member accounts in an AWS Organizations.
OrganizationConformancePack enables organization service access for `config-multiaccountsetup.amazonaws.com` through the `EnableAWSServiceAccess` action and
creates a service linked role in the master account of your organization.
The service linked role is created only when the role does not exist in the master account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Config::OrganizationConformancePack",
  "Properties" : {
      "ConformancePackInputParameters" : [ ConformancePackInputParameter, ... ],
      "DeliveryS3Bucket" : String,
      "DeliveryS3KeyPrefix" : String,
      "ExcludedAccounts" : [ String, ... ],
      "OrganizationConformancePackName" : String,
      "TemplateBody" : String,
      "TemplateS3Uri" : String
    }
}

```

### YAML

```yaml

Type: AWS::Config::OrganizationConformancePack
Properties:
  ConformancePackInputParameters:
    - ConformancePackInputParameter
  DeliveryS3Bucket: String
  DeliveryS3KeyPrefix: String
  ExcludedAccounts:
    - String
  OrganizationConformancePackName: String
  TemplateBody: String
  TemplateS3Uri: String

```

## Properties

`ConformancePackInputParameters`

A list of `ConformancePackInputParameter` objects.

_Required_: No

_Type_: Array of [ConformancePackInputParameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-config-organizationconformancepack-conformancepackinputparameter.html)

_Minimum_: `0`

_Maximum_: `60`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeliveryS3Bucket`

The name of the Amazon S3 bucket where AWS Config stores conformance pack templates.

###### Note

This field is optional.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeliveryS3KeyPrefix`

Any folder structure you want to add to an Amazon S3 bucket.

###### Note

This field is optional.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExcludedAccounts`

A comma-separated list of accounts excluded from organization conformance pack.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrganizationConformancePackName`

The name you assign to an organization conformance pack.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z][-a-zA-Z0-9]*`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TemplateBody`

A string containing full conformance pack template body. Structure containing the template body with a minimum length of 1 byte and a maximum length of 51,200 bytes.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `51200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateS3Uri`

Location of file containing the template body. The uri must point to the conformance pack template (max size: 300 KB).

_Required_: No

_Type_: String

_Pattern_: `s3://.*`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of organization conformance pack.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Organization Conformance Pack

The following example creates an organization conformance pack.

#### JSON

```json

{
    "Resources": {
        "OrganizationConformancePack": {
            "Type": "AWS::Config::OrganizationConformancePack",
            "Properties": {
                "OrganizationConformancePackName": "OrganizationConformancePackName",
                "DeliveryS3Bucket": "DeliveryS3Bucket",
                "TemplateS3Uri": "s3://bucketname/prefix"
             }
         }
     }
}

```

#### YAML

```yaml

---
Resources:
    OrganizationConformancePack:
        Type: AWS::Config::OrganizationConformancePack
        Properties:
            OrganizationConformancePackName: OrganizationConformancePackName
            DeliveryS3Bucket: DeliveryS3Bucket
            TemplateS3Uri: s3://bucketname/prefix

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

OrganizationManagedRuleMetadata

ConformancePackInputParameter
