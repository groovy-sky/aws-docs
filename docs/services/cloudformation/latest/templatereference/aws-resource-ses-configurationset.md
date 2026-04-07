This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ConfigurationSet

Configuration sets let you create groups of rules that you can apply to the emails you
send using Amazon SES. For more information about using configuration sets, see [Using Amazon\
SES Configuration Sets](https://docs.aws.amazon.com/ses/latest/dg/using-configuration-sets.html) in the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/dg).

###### Note

**Required permissions:**

To apply any of the resource options, you will need to have the corresponding
AWS Identity and Access Management (IAM) SES API v2
permissions:

- `ses:GetConfigurationSet`

- (This permission is replacing the v1
_ses:DescribeConfigurationSet_ permission
which will not work with these v2 resource options.)

- `ses:PutConfigurationSetDeliveryOptions`

- `ses:PutConfigurationSetReputationOptions`

- `ses:PutConfigurationSetSendingOptions`

- `ses:PutConfigurationSetSuppressionOptions`

- `ses:PutConfigurationSetTrackingOptions`

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SES::ConfigurationSet",
  "Properties" : {
      "ArchivingOptions" : ArchivingOptions,
      "DeliveryOptions" : DeliveryOptions,
      "Name" : String,
      "ReputationOptions" : ReputationOptions,
      "SendingOptions" : SendingOptions,
      "SuppressionOptions" : SuppressionOptions,
      "Tags" : [ Tag, ... ],
      "TrackingOptions" : TrackingOptions,
      "VdmOptions" : VdmOptions
    }
}

```

### YAML

```yaml

Type: AWS::SES::ConfigurationSet
Properties:
  ArchivingOptions:
    ArchivingOptions
  DeliveryOptions:
    DeliveryOptions
  Name: String
  ReputationOptions:
    ReputationOptions
  SendingOptions:
    SendingOptions
  SuppressionOptions:
    SuppressionOptions
  Tags:
    - Tag
  TrackingOptions:
    TrackingOptions
  VdmOptions:
    VdmOptions

```

## Properties

`ArchivingOptions`

An object that defines the MailManager archive where sent emails are archived that you send
using the configuration set.

_Required_: No

_Type_: [ArchivingOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ses-configurationset-archivingoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeliveryOptions`

Specifies the name of the dedicated IP pool to associate with the configuration set
and whether messages that use the configuration set are required to use Transport Layer
Security (TLS).

_Required_: No

_Type_: [DeliveryOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ses-configurationset-deliveryoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the configuration set. The name must meet the following
requirements:

- Contain only letters (a-z, A-Z), numbers (0-9), underscores (\_), or dashes
(-).

- Contain 64 characters or fewer.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]{1,64}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReputationOptions`

An object that defines whether or not Amazon SES collects reputation metrics for the emails
that you send that use the configuration set.

_Required_: No

_Type_: [ReputationOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ses-configurationset-reputationoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SendingOptions`

An object that defines whether or not Amazon SES can send email that you send using the
configuration set.

_Required_: No

_Type_: [SendingOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ses-configurationset-sendingoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SuppressionOptions`

An object that contains information about the suppression list preferences for your
account.

_Required_: No

_Type_: [SuppressionOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ses-configurationset-suppressionoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of objects that define the tags (keys and values) that are associated with
the configuration set.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ses-configurationset-tag.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrackingOptions`

An object that defines the open and click tracking options for emails that you send
using the configuration set.

_Required_: No

_Type_: [TrackingOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ses-configurationset-trackingoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VdmOptions`

The Virtual Deliverability Manager (VDM) options that apply to the configuration
set.

_Required_: No

_Type_: [VdmOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ses-configurationset-vdmoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

Specifies a configuration set.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "AWS SES ConfigurationSet Sample Template",
    "Parameters": {
        "ConfigSetName": {
            "Type": "String"
        }
    },
    "Resources": {
        "ConfigSet": {
            "Type": "AWS::SES::ConfigurationSet",
            "Properties": {
                "Name": {
                    "Ref": "ConfigSetName"
                }
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: AWS SES ConfigurationSet Sample Template
Parameters:
  ConfigSetName:
    Type: String
Resources:
  ConfigSet:
    Type: 'AWS::SES::ConfigurationSet'
    Properties:
      Name: !Ref ConfigSetName
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Simple Email Service

ArchivingOptions
