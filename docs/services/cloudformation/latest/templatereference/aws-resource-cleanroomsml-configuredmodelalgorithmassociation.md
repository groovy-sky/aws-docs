This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRoomsML::ConfiguredModelAlgorithmAssociation

Associates a configured model algorithm to a collaboration for use by any member of the
collaboration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CleanRoomsML::ConfiguredModelAlgorithmAssociation",
  "Properties" : {
      "ConfiguredModelAlgorithmArn" : String,
      "Description" : String,
      "MembershipIdentifier" : String,
      "Name" : String,
      "PrivacyConfiguration" : PrivacyConfiguration,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CleanRoomsML::ConfiguredModelAlgorithmAssociation
Properties:
  ConfiguredModelAlgorithmArn: String
  Description: String
  MembershipIdentifier: String
  Name: String
  PrivacyConfiguration:
    PrivacyConfiguration
  Tags:
    - Tag

```

## Properties

`ConfiguredModelAlgorithmArn`

The Amazon Resource Name (ARN) of the configured model algorithm that is associated to
the collaboration.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z]*:cleanrooms-ml:[-a-z0-9]+:[0-9]{12}:configured-model-algorithm/[-a-zA-Z0-9_/.]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description of the configured model algorithm association.

_Required_: No

_Type_: String

_Pattern_: `^[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDBFF-\uDC00\uDFFF\t\r\n]*$`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MembershipIdentifier`

The membership ID of the member that created the configured model algorithm
association.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the configured model algorithm association.

_Required_: Yes

_Type_: String

_Pattern_: `^(?!\s*$)[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDBFF-\uDC00\uDFFF\t]*$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PrivacyConfiguration`

Information about the privacy configuration for a configured model algorithm
association.

_Required_: No

_Type_: [PrivacyConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-privacyconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The optional metadata that you apply to the resource to help you categorize and organize them. Each tag consists of a key and an optional value, both of which you define.

The following basic restrictions apply to tags:

- Maximum number of tags per resource - 50.

- For each resource, each tag key must be unique, and each tag key can have only one value.

- Maximum key length - 128 Unicode characters in UTF-8.

- Maximum value length - 256 Unicode characters in UTF-8.

- If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . \_ : / @.

- Tag keys and values are case sensitive.

- Do not use `aws:`, `AWS:`, or any upper or lowercase combination of such as a prefix for keys as it is reserved. You cannot edit or delete tag keys with this prefix. Values can have this prefix. If a tag value has `aws` as its prefix but the key does not, then Clean Rooms ML considers it to be a user tag and will count against the limit of 50 tags. Tags with only the key prefix of `aws` do not count against your tags per resource limit.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource ARN. For example:

`{ "Ref": "myConfiguredModelAlgorithmAssociation" }`

For the Clean Rooms ML configured model algorithm association, `Ref` returns the ARN of the configured model algorithm association.

Example: `arn:aws:cleanrooms-ml:us-east-1:123456789012:membership/a1b2c3d4-e5f6-7890-abcd-ef1234567890/configured-model-algorithm-association/a1b2c3d4e5f6`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CollaborationIdentifier`

The collaboration ID of the collaboration that contains the configured model algorithm
association.

`ConfiguredModelAlgorithmAssociationArn`

The Amazon Resource Name (ARN) of the configured model algorithm association.

## Examples

### Create a configured model algorithm association

The following example creates a configured model algorithm association with privacy configuration policies.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "MyConfiguredModelAlgorithmAssociation": {
      "Type": "AWS::CleanRoomsML::ConfiguredModelAlgorithmAssociation",
      "Properties": {
        "Name": "MyMLAssociation",
        "MembershipIdentifier": "12345678-1234-1234-1234-123456789012",
        "ConfiguredModelAlgorithmArn": "arn:aws:cleanrooms-ml:us-east-1:123456789012:configured-model-algorithm/my-algorithm",
        "Description": "Association for collaborative ML training",
        "PrivacyConfiguration": {
          "Policies": {
            "TrainedModels": {
              "ContainerLogs": [
                {
                  "AllowedAccountIds": ["123456789012"],
                  "LogType": "ERROR_SUMMARY"
                }
              ],
              "ContainerMetrics": {
                "NoiseLevel": "HIGH"
              }
            },
            "TrainedModelExports": {
              "MaxSize": {
                "Unit": "GB",
                "Value": 10
              },
              "FilesToExport": ["MODEL", "OUTPUT"]
            }
          }
        },
        "Tags": [
          {
            "Key": "Project",
            "Value": "CollaborativeML"
          }
        ]
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Resources:
  MyConfiguredModelAlgorithmAssociation:
    Type: AWS::CleanRoomsML::ConfiguredModelAlgorithmAssociation
    Properties:
      Name: MyMLAssociation
      MembershipIdentifier: 12345678-1234-1234-1234-123456789012
      ConfiguredModelAlgorithmArn: arn:aws:cleanrooms-ml:us-east-1:123456789012:configured-model-algorithm/my-algorithm
      Description: Association for collaborative ML training
      PrivacyConfiguration:
        Policies:
          TrainedModels:
            ContainerLogs:
              - AllowedAccountIds:
                  - '123456789012'
                LogType: ERROR_SUMMARY
            ContainerMetrics:
              NoiseLevel: HIGH
          TrainedModelExports:
            MaxSize:
              Unit: GB
              Value: 10
            FilesToExport:
              - MODEL
              - OUTPUT
      Tags:
        - Key: Project
          Value: CollaborativeML
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

CustomEntityConfig
