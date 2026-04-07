This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AccessAnalyzer::Analyzer AnalyzerConfiguration

Contains information about the configuration of an analyzer for an AWS organization or
account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InternalAccessConfiguration" : InternalAccessConfiguration,
  "UnusedAccessConfiguration" : UnusedAccessConfiguration
}

```

### YAML

```yaml

  InternalAccessConfiguration:
    InternalAccessConfiguration
  UnusedAccessConfiguration:
    UnusedAccessConfiguration

```

## Properties

`InternalAccessConfiguration`

Specifies the configuration of an internal access analyzer for an AWS
organization or account. This configuration determines how the analyzer evaluates access
within your AWS environment.

_Required_: No

_Type_: [InternalAccessConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-accessanalyzer-analyzer-internalaccessconfiguration.html)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`UnusedAccessConfiguration`

Specifies the configuration of an unused access analyzer for an AWS organization or
account.

_Required_: No

_Type_: [UnusedAccessConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-accessanalyzer-analyzer-unusedaccessconfiguration.html)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

## Examples

### Declare an AnalyzerConfiguration Resource

The following example shows how to declare a IAM Access Analyzer
`AnalyzerConfiguration` resource:

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "Analyzer": {
      "Properties": {
        "AnalyzerName": "DevUnusedAccessAccountAnalyzer",
        "AnalyzerConfiguration": {
          "UnusedAccessConfiguration": {
            "UnusedAccessAge": 90,
            "AnalysisRule": {
              "Exclusions": [
                {
                  "ResourceTags": [
                    [
                      {
                        "Key": "Kind",
                        "Value": "Dev"
                      }
                    ],
                    [
                      {
                        "Key": "AnotherKey"
                      }
                    ]
                  ]
                }
              ]
            }
          }
        },
        "ArchiveRules": [
          {
            "Filter": [
              {
                "Eq": [
                  "123456789012"
                ],
                "Property": "resourceOwnerAccount"
              }
            ],
            "RuleName": "ArchiveTrustedAccountAccess"
          },
          {
            "Filter": [
              {
                "Contains": [
                  "arn:aws:s3:::amzn-s3-demo-logging-bucket",
                  "arn:aws:s3:::amzn-s3-demo-website-bucket"
                ],
                "Property": "resource"
              }
            ],
            "RuleName": "ArchivePublicS3BucketsAccess"
          }
        ],
        "Tags": [
          {
            "Key": "Kind",
            "Value": "Dev"
          }
        ],
        "Type": "ACCOUNT_UNUSED_ACCESS"
      },
      "Type": "AWS::AccessAnalyzer::Analyzer"
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Resources:
  Analyzer:
    Properties:
      AnalyzerName: DevUnusedAccessAccountAnalyzer
      AnalyzerConfiguration:
        UnusedAccessConfiguration:
          UnusedAccessAge: 90
          AnalysisRule:
            Exclusions:
            - ResourceTags:
              - - Key: Kind
                  Value: Dev
              - - Key: AnotherKey
      ArchiveRules:
      - Filter:
        - Eq:
          - '123456789012'
          Property: resourceOwnerAccount
        RuleName: ArchiveTrustedAccountAccess
      - Filter:
        - Contains:
          - arn:aws:s3:::amzn-s3-demo-logging-bucket
          - arn:aws:s3:::amzn-s3-demo-website-bucket
          Property: resource
        RuleName: ArchivePublicS3BucketsAccess
      Tags:
      - Key: Kind
        Value: Dev
      Type: ACCOUNT_UNUSED_ACCESS
    Type: AWS::AccessAnalyzer::Analyzer
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AnalysisRuleCriteria

ArchiveRule
