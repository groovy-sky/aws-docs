# Attach a resource-based policy to a DynamoDB stream

You can attach a resource-based policy to an existing table's stream or modify an existing
policy by using the DynamoDB console, [PutResourcePolicy](../../../../reference/amazondynamodb/latest/apireference/api-putresourcepolicy.md) API, the AWS CLI, AWS SDK, or an [CloudFormation template](rbac-create-table.md#rbac-create-table-cfn).

###### Note

You can't attach a policy to a stream while creating it using the [CreateTable](../../../../reference/amazondynamodb/latest/apireference/api-createtable.md) or [UpdateTable](../../../../reference/amazondynamodb/latest/apireference/api-updatetable.md) APIs.
However, you can modify or delete a policy after a table is deleted. You can also modify or
delete the policy of a disabled stream.

The following IAM policy example uses the `put-resource-policy` AWS CLI
command to attach a resource-based policy to the stream of a table named
`MusicCollection`. This example allows the user
`John` to perform the [GetRecords](../../../../reference/amazondynamodb/latest/apireference/api-streams-getrecords.md),
[GetShardIterator](../../../../reference/amazondynamodb/latest/apireference/api-streams-getsharditerator.md), and [DescribeStream](../../../../reference/amazondynamodb/latest/apireference/api-streams-describestream.md) API actions on the stream.

Remember to replace the `italicized` text with your resource-specific information.

```nohighlight

aws dynamodb put-resource-policy \
    --resource-arn arn:aws:dynamodb:us-west-2:123456789012:table/MusicCollection/stream/2024-02-12T18:57:26.492 \
    --policy \
        "{
            \"Version\": \"2012-10-17\",
            \"Statement\": [
              {
                    \"Effect\": \"Allow\",
                    \"Principal\": {
                        \"AWS\": \"arn:aws:iam::111122223333:user/John\"
                    },
                    \"Action\": [
                        \"dynamodb:GetRecords\",
                        \"dynamodb:GetShardIterator\",
                        \"dynamodb:DescribeStream\"
                    ],
                    \"Resource\": \"arn:aws:dynamodb:us-west-2:123456789012:table/MusicCollection/stream/2024-02-12T18:57:26.492\"
                }
            ]
        }"
```

1. Sign in to the AWS Management Console and open the DynamoDB console at
    [https://console.aws.amazon.com/dynamodb/](https://console.aws.amazon.com/dynamodb).

2. On the DynamoDB console dashboard, choose **Tables** and then select
    an existing table.

Make sure the table you select has streams turned on. For information about
    turning on streams for a table, see [Enabling a stream](streams.md#Streams.Enabling).

3. Choose the **Permissions** tab.

4. In **Resource-based policy for active stream**, choose
    **Create stream policy**.

5. In the **Resource-based policy** editor, add a policy to define
    the access permissions for the stream. In this policy, you specify who has access to
    the stream and the actions they are allowed to perform on the stream. To add a policy,
    do one of the following:

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

6. (Optional) Choose **Preview external access** in the lower-right
    corner to preview how your new policy affects public and cross-account access to your
    resource. Before you save your policy, you can check whether it introduces new
    IAM Access Analyzer findings or resolves existing findings. If you don’t see an active
    analyzer, choose **Go to Access Analyzer** to [create an account analyzer](../../../iam/latest/userguide/access-analyzer-getting-started.md#access-analyzer-enabling) in IAM Access Analyzer. For more information, see
    [Preview\
    access](../../../iam/latest/userguide/access-analyzer-access-preview.md).

7. Choose **Create policy**.

The following IAM policy example attaches a resource-based policy to the stream of a
table named `MusicCollection`. This example allows the user
`John` to perform the [GetRecords](../../../../reference/amazondynamodb/latest/apireference/api-streams-getrecords.md),
[GetShardIterator](../../../../reference/amazondynamodb/latest/apireference/api-streams-getsharditerator.md), and [DescribeStream](../../../../reference/amazondynamodb/latest/apireference/api-streams-describestream.md) API actions on the stream.

Remember to replace the `italicized` text with your resource-specific information.

JSON

```json

{
  "Version":"2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::111122223333:user/username"
      },
      "Action": [
        "dynamodb:GetRecords",
        "dynamodb:GetShardIterator",
        "dynamodb:DescribeStream"
      ],
      "Resource": [
        "arn:aws:dynamodb:us-east-1:123456789012:table/MusicCollection/stream/2024-02-12T18:57:26.492"
      ]
    }
  ]
}

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Attach resource-based
policy

Remove resource-based
policy

All content copied from https://docs.aws.amazon.com/.
