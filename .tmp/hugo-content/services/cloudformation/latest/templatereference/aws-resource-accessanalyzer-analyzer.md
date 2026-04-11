This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AccessAnalyzer::Analyzer

The `AWS::AccessAnalyzer::Analyzer` resource specifies a new analyzer. The
analyzer is an object that represents the IAM Access Analyzer feature. An analyzer is
required for Access Analyzer to become operational.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AccessAnalyzer::Analyzer",
  "Properties" : {
      "AnalyzerConfiguration" : AnalyzerConfiguration,
      "AnalyzerName" : String,
      "ArchiveRules" : [ ArchiveRule, ... ],
      "Tags" : [ Tag, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::AccessAnalyzer::Analyzer
Properties:
  AnalyzerConfiguration:
    AnalyzerConfiguration
  AnalyzerName: String
  ArchiveRules:
    - ArchiveRule
  Tags:
    - Tag
  Type: String

```

## Properties

`AnalyzerConfiguration`

Contains information about the configuration of an analyzer for an AWS organization or
account.

_Required_: No

_Type_: [AnalyzerConfiguration](aws-properties-accessanalyzer-analyzer-analyzerconfiguration.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`AnalyzerName`

The name of the analyzer.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ArchiveRules`

Specifies the archive rules to add for the analyzer. Archive rules automatically archive
findings that meet the criteria you define for the rule.

_Required_: No

_Type_: Array of [ArchiveRule](aws-properties-accessanalyzer-analyzer-archiverule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to the analyzer. You can use the set of Unicode
letters, digits, whitespace, `_`, `.`, `/`,
`=`, `+`, and `-`.

For the tag key, you can specify a value that is 1 to 128 characters in length and
cannot be prefixed with `aws:`.

For the tag value, you can specify a value that is 0 to 256 characters in length.

_Required_: No

_Type_: Array of [Tag](aws-properties-accessanalyzer-analyzer-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type represents the zone of trust for the analyzer.

_Allowed Values_: ACCOUNT \| ORGANIZATION \| ACCOUNT\_UNUSED\_ACCESS \|
ACCOUNT\_INTERNAL\_ACCESS \| ORGANIZATION\_INTERNAL\_ACCESS \|
ORGANIZATION\_UNUSED\_ACCESS

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the analyzer created.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`Arn`

The ARN of the analyzer that was created.

## Examples

### Declare an Analyzer Resource

The following example shows how to declare a IAM Access Analyzer
`Analyzer` resource:

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "Analyzer": {
      "Properties": {
        "AnalyzerName": "DevAccountAnalyzer",
        "ArchiveRules": [
          {
            "Filter": [
              {
                "Eq": [
                  "123456789012"
                ],
                "Property": "principal.AWS"
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
        "Type": "ACCOUNT"
      },
      "Type": "AWS::AccessAnalyzer::Analyzer"
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  Analyzer:
    Type: 'AWS::AccessAnalyzer::Analyzer'
    Properties:
      AnalyzerName: MyAccountAnalyzer
      Type: ACCOUNT
      Tags:
        -
          Key: Kind
          Value: Dev
      ArchiveRules:
        -
          # Archive findings for a trusted AWS account
          RuleName: ArchiveTrustedAccountAccess
          Filter:
            -
              Property: 'principal.AWS'
              Eq:
                - '123456789012'
        -
          # Archive findings for known public S3 buckets
          RuleName: ArchivePublicS3BucketsAccess
          Filter:
            -
              Property: 'resource'
              Contains:
                - 'arn:aws:s3:::amzn-s3-demo-logging-bucket'
                - 'arn:aws:s3:::amzn-s3-demo-website-bucket'
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Identity and Access Management Access Analyzer

AnalysisRule

All content copied from https://docs.aws.amazon.com/.
