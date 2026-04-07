This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Crawler

The `AWS::Glue::Crawler` resource specifies an AWS Glue crawler. For more
information, see [Cataloging Tables with a Crawler](https://docs.aws.amazon.com/glue/latest/dg/add-crawler.html) and [Crawler Structure](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-crawler-crawling.html#aws-glue-api-crawler-crawling-Crawler) in the _AWS Glue Developer_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Glue::Crawler",
  "Properties" : {
      "Classifiers" : [ String, ... ],
      "Configuration" : String,
      "CrawlerSecurityConfiguration" : String,
      "DatabaseName" : String,
      "Description" : String,
      "LakeFormationConfiguration" : LakeFormationConfiguration,
      "Name" : String,
      "RecrawlPolicy" : RecrawlPolicy,
      "Role" : String,
      "Schedule" : Schedule,
      "SchemaChangePolicy" : SchemaChangePolicy,
      "TablePrefix" : String,
      "Tags" : [ Tag, ... ],
      "Targets" : Targets
    }
}

```

### YAML

```yaml

Type: AWS::Glue::Crawler
Properties:
  Classifiers:
    - String
  Configuration: String
  CrawlerSecurityConfiguration: String
  DatabaseName: String
  Description: String
  LakeFormationConfiguration:
    LakeFormationConfiguration
  Name: String
  RecrawlPolicy:
    RecrawlPolicy
  Role: String
  Schedule:
    Schedule
  SchemaChangePolicy:
    SchemaChangePolicy
  TablePrefix: String
  Tags:
    - Tag
  Targets:
    Targets

```

## Properties

`Classifiers`

A list of UTF-8 strings that specify the names of custom classifiers that are associated
with the crawler.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Configuration`

Crawler configuration information. This versioned JSON string allows users to specify
aspects of a crawler's behavior. For more information, see [Configuring a\
Crawler](https://docs.aws.amazon.com/glue/latest/dg/crawler-configuration.html).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CrawlerSecurityConfiguration`

The name of the `SecurityConfiguration` structure to be used by this
crawler.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabaseName`

The name of the database in which the crawler's output is stored.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the crawler.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LakeFormationConfiguration`

Specifies whether the crawler should use AWS Lake Formation credentials for the crawler instead of the IAM role credentials.

_Required_: No

_Type_: [LakeFormationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-glue-crawler-lakeformationconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the crawler.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RecrawlPolicy`

A policy that specifies whether to crawl the entire dataset again, or to crawl only folders that were added since the last crawler run.

_Required_: No

_Type_: [RecrawlPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-glue-crawler-recrawlpolicy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Role`

The Amazon Resource Name (ARN) of an IAM role that's used to access customer resources,
such as Amazon Simple Storage Service (Amazon S3) data.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Schedule`

For scheduled crawlers, the schedule when the crawler runs.

_Required_: No

_Type_: [Schedule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-glue-crawler-schedule.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SchemaChangePolicy`

The policy that specifies update and delete behaviors for the crawler. The policy tells the crawler what to do in the event that it detects a change in a table that already exists in the customer's database at the time of the crawl. The `SchemaChangePolicy` does not affect whether or how new tables and partitions are added. New tables and partitions are always created regardless of the `SchemaChangePolicy` on a crawler.

The SchemaChangePolicy consists of two components, `UpdateBehavior` and `DeleteBehavior`.

_Required_: No

_Type_: [SchemaChangePolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-glue-crawler-schemachangepolicy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TablePrefix`

The prefix added to the names of tables that are created.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to use with this crawler.

_Required_: No

_Type_: Array of [`Tag`](aws-properties-resource-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Targets`

A collection of targets to crawl.

_Required_: Yes

_Type_: [Targets](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-glue-crawler-targets.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the crawler name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

- [Create a crawler](#aws-resource-glue-crawler--examples--Create_a_crawler)

- [Crawler Configuration](#aws-resource-glue-crawler--examples--Crawler_Configuration)

### Create a crawler

The following example creates a crawler for an Amazon S3 target.

#### JSON

```json

{
    "Description": "AWS Glue crawler test",
    "Resources": {
        "MyRole": {
            "Type": "AWS::IAM::Role",
            "Properties": {
                "AssumeRolePolicyDocument": {
                    "Version": "2012-10-17",
                    "Statement": [
                        {
                            "Effect": "Allow",
                            "Principal": {
                                "Service": [
                                    "glue.amazonaws.com"
                                ]
                            },
                            "Action": [
                                "sts:AssumeRole"
                            ]
                        }
                    ]
                },
                "Path": "/",
                "ManagedPolicyArns": ["arn:aws:iam::aws:policy/service-role/AWSGlueServiceRole"],
                "Policies": [
                    {
                        "PolicyName": "S3BucketAccessPolicy",
                        "PolicyDocument": {
                            "Version": "2012-10-17",
                            "Statement": [
                                {
                                    "Effect": "Allow",
                                    "Action": [
                                        "s3:GetObject",
                                        "s3:PutObject"
                                    ],
                                    "Resource": {
                                        "Fn::Join": [
                                            "",
                                            [
                                                {
                                                    "Fn::GetAtt": ["MyS3Bucket", "Arn"]
                                                },
                                                "*"
                                            ]
                                        ]
                                    }
                                }
                            ]
                        }
                    }
                ]
            }
        },
        "MyDatabase": {
            "Type": "AWS::Glue::Database",
            "Properties": {
                "CatalogId": {
                    "Ref": "AWS::AccountId"
                },
                "DatabaseInput": {
                    "Name": "dbcrawler",
                    "Description": "TestDatabaseDescription",
                    "LocationUri": "TestLocationUri",
                    "Parameters": {
                        "key1": "value1",
                        "key2": "value2"
                    }
                }
            }
        },
        "MyClassifier": {
            "Type": "AWS::Glue::Classifier",
            "Properties": {
                "GrokClassifier": {
                    "Name": "CrawlerClassifier",
                    "Classification": "wikiData",
                    "GrokPattern": "%{NOTSPACE:language} %{NOTSPACE:page_title} %{NUMBER:hits:long} %{NUMBER:retrieved_size:long}"
                }
            }
        },
        "MyS3Bucket": {
            "Type": "AWS::S3::Bucket",
            "Properties": {
                "BucketName": "crawlertesttarget",
                "AccessControl": "BucketOwnerFullControl"
            }
        },
        "MyCrawler2": {
            "Type": "AWS::Glue::Crawler",
            "Properties": {
                "Name": "testcrawler1",
                "Role": {
                    "Fn::GetAtt": [
                        "MyRole",
                        "Arn"
                    ]
                },
                "DatabaseName": {
                    "Ref": "MyDatabase"
                },
                "Classifiers": [
                    {
                        "Ref": "MyClassifier"
                    }
                ],
                "Targets": {
                    "S3Targets": [
                        {
                            "Path": {
                                "Ref": "MyS3Bucket"
                            }
                        }
                    ]
                },
                "SchemaChangePolicy": {
                    "UpdateBehavior": "UPDATE_IN_DATABASE",
                    "DeleteBehavior": "LOG"
                },
                "Tags": {
                    "key1": "value1"
                },
                "Schedule": {
                    "ScheduleExpression": "cron(0/10 * ? * MON-FRI *)"
                }
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  MyRole:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Version: "2012-10-17"
        Statement:
          -
            Effect: "Allow"
            Principal:
              Service:
                - "glue.amazonaws.com"
            Action:
              - "sts:AssumeRole"
      Path: "/"
      ManagedPolicyArns:
        ['arn:aws:iam::aws:policy/service-role/AWSGlueServiceRole']
      Policies:
        -
          PolicyName: "S3BucketAccessPolicy"
          PolicyDocument:
            Version: "2012-10-17"
            Statement:
              -
                Effect: "Allow"
                Action:
                  - "s3:GetObject"
                  - "s3:PutObject"
                Resource:
                  !Join
                    - ''
                    - - !GetAtt MyS3Bucket.Arn
                      - "*"

  MyDatabase:
    Type: AWS::Glue::Database
    Properties:
      CatalogId: !Ref AWS::AccountId
      DatabaseInput:
        Name: "dbcrawler"
        Description: "TestDatabaseDescription"
        LocationUri: "TestLocationUri"
        Parameters:
          key1 : "value1"
          key2 : "value2"

  MyClassifier:
    Type: AWS::Glue::Classifier
    Properties:
      GrokClassifier:
        Name: "CrawlerClassifier"
        Classification: "wikiData"
        GrokPattern: "%{NOTSPACE:language} %{NOTSPACE:page_title} %{NUMBER:hits:long} %{NUMBER:retrieved_size:long}"

  MyS3Bucket:
    Type: AWS::S3::Bucket
    Properties:
      BucketName: "crawlertesttarget"
      AccessControl: "BucketOwnerFullControl"

  MyCrawler2:
    Type: AWS::Glue::Crawler
    Properties:
      Name: "testcrawler1"
      Role: !GetAtt MyRole.Arn
      DatabaseName: !Ref MyDatabase
      Classifiers:
        - !Ref MyClassifier
      Targets:
        S3Targets:
          - Path: !Ref MyS3Bucket
      SchemaChangePolicy:
        UpdateBehavior: "UPDATE_IN_DATABASE"
        DeleteBehavior: "LOG"
      Tags:
        "Key1": "Value1"
      Schedule:
        ScheduleExpression: "cron(0/10 * ? * MON-FRI *)"
```

### Crawler Configuration

The following example specifies a configuration that controls a crawler's
behavior.

#### JSON

```json

{
    "Type": "AWS::Glue::Crawler",
    "Properties": {
        "Role": "role1",
        "Classifiers": [],
        "Description": "example classifier",
        "SchemaChangePolicy": "",
        "Schedule": "Schedule",
        "DatabaseName": "test",
        "Targets": [],
        "TablePrefix": "test-",
        "Name": "my-crawler",
        "Configuration": "{\"Version\":1.0,\"CrawlerOutput\":{\"Partitions\":{\"AddOrUpdateBehavior\":\"InheritFromTable\"},\"Tables\":{\"AddOrUpdateBehavior\":\"MergeNewColumns\"}}}"
    }
}
```

#### YAML

```yaml

Type: AWS::Glue::Crawler
Properties:
  Role: role1
  Classifiers:
    - ''
  Description: example classifier
  SchemaChangePolicy: ''
  Schedule: Schedule
  DatabaseName: test
  Targets:
    - ''
  TablePrefix: test-
  Name: my-crawler
  Configuration: "{\"Version\":1.0,\"CrawlerOutput\":{\"Partitions\":{\"AddOrUpdateBehavior\":\"InheritFromTable\"},\"Tables\":{\"AddOrUpdateBehavior\":\"MergeNewColumns\"}}}"

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PhysicalConnectionRequirements

CatalogTarget
