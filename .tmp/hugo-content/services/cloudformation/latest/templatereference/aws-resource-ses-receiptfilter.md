This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ReceiptFilter

Specify a new IP address filter. You use IP address filters when you receive email
with Amazon SES.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SES::ReceiptFilter",
  "Properties" : {
      "Filter" : Filter
    }
}

```

### YAML

```yaml

Type: AWS::SES::ReceiptFilter
Properties:
  Filter:
    Filter

```

## Properties

`Filter`

A data structure that describes the IP address filter to create, which consists of a
name, an IP address range, and whether to allow or block mail from it.

_Required_: Yes

_Type_: [Filter](aws-properties-ses-receiptfilter-filter.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Fn::GetAtt

## Examples

Specifies an IP address filter for incoming email.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "AWS SES ReceiptFilter Sample Template",
    "Parameters": {
        "FilterName": {
            "Type": "String"
        },
        "Policy": {
            "Type": "String"
        },
        "Cidr": {
            "Type": "String"
        }
    },
    "Resources": {
        "ReceiptFilter": {
            "Type": "AWS::SES::ReceiptFilter",
            "Properties": {
                "Filter": {
                    "Name": {
                        "Ref": "FilterName"
                    },
                    "IpFilter": {
                        "Policy": {
                            "Ref": "Policy"
                        },
                        "Cidr": {
                            "Ref": "Cidr"
                        }
                    }
                }
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: AWS SES ReceiptFilter Sample Template
Parameters:
  FilterName:
    Type: String
  Policy:
    Type: String
  Cidr:
    Type: String
Resources:
  ReceiptFilter:
    Type: 'AWS::SES::ReceiptFilter'
    Properties:
      Filter:
        Name: !Ref FilterName
        IpFilter:
          Policy: !Ref Policy
          Cidr: !Ref Cidr
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Filter

All content copied from https://docs.aws.amazon.com/.
