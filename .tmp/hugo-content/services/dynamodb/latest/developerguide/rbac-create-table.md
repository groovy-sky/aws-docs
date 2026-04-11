# Create a table with a resource-based policy

You can add a resource-based policy while you create a table by using the DynamoDB console,
[CreateTable](../../../../reference/amazondynamodb/latest/apireference/api-createtable.md) API, AWS CLI, [AWS\
SDK](rbac-attach-resource-based-policy.md#rbac-attach-policy-java-sdk), or an CloudFormation template.

The following example creates a table named `MusicCollection`
using the `create-table` AWS CLI command. This command also includes the
`resource-policy` parameter that adds a resource-based policy to the table.
This policy allows the user `John` to perform the [RestoreTableToPointInTime](../../../../reference/amazondynamodb/latest/apireference/api-restoretabletopointintime.md), [GetItem](../../../../reference/amazondynamodb/latest/apireference/api-getitem.md), and [PutItem](../../../../reference/amazondynamodb/latest/apireference/api-putitem.md) API actions on the table.

Remember to replace the `italicized` text with your resource-specific information.

```nohighlight

aws dynamodb create-table \
    --table-name MusicCollection \
    --attribute-definitions AttributeName=Artist,AttributeType=S AttributeName=SongTitle,AttributeType=S \
    --key-schema AttributeName=Artist,KeyType=HASH AttributeName=SongTitle,KeyType=RANGE \
    --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 \
    --resource-policy \
        "{
            \"Version\": \"2012-10-17\",
            \"Statement\": [
              {
                    \"Effect\": \"Allow\",
                    \"Principal\": {
                        \"AWS\": \"arn:aws:iam::123456789012:user/John\"
                    },
                    \"Action\": [
                        \"dynamodb:RestoreTableToPointInTime\",
                        \"dynamodb:GetItem\",
                        \"dynamodb:DescribeTable\"
                    ],
                    \"Resource\": \"arn:aws:dynamodb:us-west-2:123456789012:table/MusicCollection\"
                }
            ]
        }"
```

1. Sign in to the AWS Management Console and open the DynamoDB console at
    [https://console.aws.amazon.com/dynamodb/](https://console.aws.amazon.com/dynamodb).

2. On the dashboard, choose **Create table**.

3. In **Table details**, enter the table name, partition key, and
    sort key details.

4. In **Table settings**, choose **Customize**
**settings**.

5. (Optional) Specify your options for **Table class**,
    **Capacity calculator**, **Read/write capacity**
**settings**, **Secondary indexes**, **Encryption at**
**rest**, and **Deletion protection**.

6. In **Resource-based policy**, add a policy to define the access
    permissions for the table and its indexes. In this policy, you specify who has access
    to these resources, and the actions they are allowed to perform on each resource. To
    add a policy, do one of the following:

- Type or paste a JSON policy document. For details about the IAM policy
language, see [Creating policies using the JSON editor](../../../iam/latest/userguide/access-policies-create-console.md#access_policies_create-json-editor) in the
_IAM User Guide_.

###### Tip

To see examples of resource-based policies in the Amazon DynamoDB Developer Guide, choose
**Policy examples**.

- Choose **Add new statement** to add a new statement and enter
the information in the provided fields. Repeat this step for as many statements as
you would like to add.

###### Important

Make sure that you resolve any security warnings, errors, or suggestions before
you save your policy.

The following IAM policy example allows the user
`John` to perform the [RestoreTableToPointInTime](../../../../reference/amazondynamodb/latest/apireference/api-restoretabletopointintime.md), [GetItem](../../../../reference/amazondynamodb/latest/apireference/api-getitem.md), and
[PutItem](../../../../reference/amazondynamodb/latest/apireference/api-putitem.md) API
actions on the table `MusicCollection`.

Remember to replace the `italicized` text with your resource-specific information.
JSON

```json

{
  "Version":"2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::123456789012:user/username"
      },
      "Action": [
        "dynamodb:RestoreTableToPointInTime",
        "dynamodb:GetItem",
        "dynamodb:PutItem"
      ],
      "Resource": "arn:aws:dynamodb:us-east-1:123456789012:table/MusicCollection"
    }
  ]
}

```

7. (Optional) Choose **Preview external access** in the lower-right
    corner to preview how your new policy affects public and cross-account access to your
    resource. Before you save your policy, you can check whether it introduces new
    IAM Access Analyzer findings or resolves existing findings. If you don’t see an active
    analyzer, choose **Go to Access Analyzer** to [create an account analyzer](../../../iam/latest/userguide/access-analyzer-getting-started.md#access-analyzer-enabling) in IAM Access Analyzer. For more information, see
    [Preview\
    access](../../../iam/latest/userguide/access-analyzer-access-preview.md).

8. Choose **Create table**.

Using the AWS::DynamoDB::Table resource

The following CloudFormation template creates a table with a stream using the [AWS::DynamoDB::Table](../../../cloudformation/latest/userguide/aws-resource-dynamodb-table.md) resource. This template also includes
resource-based policies that are attached to both the table and the stream.

```nohighlight

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "MusicCollectionTable": {
            "Type": "AWS::DynamoDB::Table",
            "Properties": {
                "AttributeDefinitions": [
                    {
                        "AttributeName": "Artist",
                        "AttributeType": "S"
                    }
                ],
                "KeySchema": [
                    {
                        "AttributeName": "Artist",
                        "KeyType": "HASH"
                    }
                ],
                "BillingMode": "PROVISIONED",
                "ProvisionedThroughput": {
                    "ReadCapacityUnits": 5,
                    "WriteCapacityUnits": 5
                },
                "StreamSpecification": {
                  "StreamViewType": "OLD_IMAGE",
                  "ResourcePolicy": {
                    "PolicyDocument": {
                      "Version": "2012-10-17",
                      "Statement": [
                        {
                            "Principal": {
                                "AWS": "arn:aws:iam::111122223333:user/John"
                            },
                            "Effect": "Allow",
                            "Action": [
                                "dynamodb:GetRecords",
                                "dynamodb:GetShardIterator",
                                "dynamodb:DescribeStream"
                            ],
                            "Resource": "*"
                        }
                      ]
                    }
                  }
                },
                "TableName": "MusicCollection",
                "ResourcePolicy": {
                    "PolicyDocument": {
                        "Version": "2012-10-17",
                        "Statement": [
                            {
                                "Principal": {
                                    "AWS": [
                                        "arn:aws:iam::111122223333:user/John"
                                    ]
                                },
                                "Effect": "Allow",
                                "Action": "dynamodb:GetItem",
                                "Resource": "*"
                            }
                        ]
                    }
                }
            }

        }
    }
}
```

Using the AWS::DynamoDB::GlobalTable resource

The following CloudFormation template creates a table with the [AWS::DynamoDB::GlobalTable](../../../cloudformation/latest/userguide/aws-resource-dynamodb-globaltable.md) resource and attaches a resource-based
policy to the table and its stream.

```nohighlight

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "GlobalMusicCollection": {
            "Type": "AWS::DynamoDB::GlobalTable",
            "Properties": {
                "TableName": "MusicCollection",
                "AttributeDefinitions": [{
                    "AttributeName": "Artist",
                    "AttributeType": "S"
                }],
                "KeySchema": [{
                    "AttributeName": "Artist",
                    "KeyType": "HASH"
                }],
                "BillingMode": "PAY_PER_REQUEST",
                "StreamSpecification": {
                    "StreamViewType": "NEW_AND_OLD_IMAGES"
                },
                "Replicas": [
                    {
                        "Region": "us-east-1",
                        "ResourcePolicy": {
                            "PolicyDocument": {
                                "Version": "2012-10-17",
                                "Statement": [{
                                    "Principal": {
                                        "AWS": [
                                            "arn:aws:iam::111122223333:user/John"
                                        ]
                                    },
                                    "Effect": "Allow",
                                    "Action": "dynamodb:GetItem",
                                    "Resource": "*"
                                }]
                            }
                        },
                        "ReplicaStreamSpecification": {
                            "ResourcePolicy": {
                                "PolicyDocument": {
                                    "Version": "2012-10-17",
                                    "Statement": [{
                                        "Principal": {
                                            "AWS": "arn:aws:iam::111122223333:user/John"
                                        },
                                        "Effect": "Allow",
                                        "Action": [
                                            "dynamodb:GetRecords",
                                            "dynamodb:GetShardIterator",
                                            "dynamodb:DescribeStream"
                                        ],
                                        "Resource": "*"
                                    }]
                                }
                            }
                        }
                    }
                ]
            }
        }
    }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Resource-based
policies

Attach resource-based
policy

All content copied from https://docs.aws.amazon.com/.
