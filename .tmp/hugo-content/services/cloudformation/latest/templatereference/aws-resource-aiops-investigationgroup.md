This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AIOps::InvestigationGroup

Creates an _investigation group_ in your account. Creating an
investigation group is a one-time setup task for each Region in your account. It is a
necessary task to be able to perform investigations.

Settings in the investigation group help you centrally manage the common properties of
your investigations, such as the following:

- Who can access the investigations

- Whether investigation data is encrypted with a customer managed AWS Key Management Service key.

- How long investigations and their data are retained by default.

Currently, you can have one investigation group in each Region in your account. Each
investigation in a Region is a part of the investigation group in that Region

To create an investigation group and set up CloudWatch investigations, you must be signed in to an IAM principal that has either the `AIOpsConsoleAdminPolicy` or the
`AdministratorAccess`
IAM policy attached, or to an account that has similar permissions.

###### Important

You can configure CloudWatch alarms to start investigations and add events to
investigations. If you create your investigation group with
`CreateInvestigationGroup` and you want to enable alarms to do this, you
must use `PutInvestigationGroupPolicy` to create a resource policy that grants this
permission to CloudWatch alarms.

For more information about configuring CloudWatch alarms,
see [Using Amazon CloudWatch alarms](../../../amazoncloudwatch/latest/monitoring/alarmthatsendsemail.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AIOps::InvestigationGroup",
  "Properties" : {
      "ChatbotNotificationChannels" : [ ChatbotNotificationChannel, ... ],
      "CrossAccountConfigurations" : [ CrossAccountConfiguration, ... ],
      "EncryptionConfig" : EncryptionConfigMap,
      "InvestigationGroupPolicy" : String,
      "IsCloudTrailEventHistoryEnabled" : Boolean,
      "Name" : String,
      "RetentionInDays" : Integer,
      "RoleArn" : String,
      "TagKeyBoundaries" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::AIOps::InvestigationGroup
Properties:
  ChatbotNotificationChannels:
    - ChatbotNotificationChannel
  CrossAccountConfigurations:
    - CrossAccountConfiguration
  EncryptionConfig:
    EncryptionConfigMap
  InvestigationGroupPolicy: String
  IsCloudTrailEventHistoryEnabled: Boolean
  Name: String
  RetentionInDays: Integer
  RoleArn: String
  TagKeyBoundaries:
    - String
  Tags:
    - Tag

```

## Properties

`ChatbotNotificationChannels`

Use this property to integrate CloudWatch investigations with chat applications. This
property is an array. For the first string, specify the ARN of an Amazon SNS
topic. For the array of strings, specify the ARNs of one or more chat applications
configurations that you want to associate with that topic. For more information about
these configuration ARNs, see [Getting started with Amazon Q in chat applications](../../../chatbot/latest/adminguide/getting-started.md) and [Resource type defined by AWS Chatbot](../../../service-authorization/latest/reference/list-awschatbot.md#awschatbot-resources-for-iam-policies).

_Required_: No

_Type_: Array of [ChatbotNotificationChannel](aws-properties-aiops-investigationgroup-chatbotnotificationchannel.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CrossAccountConfigurations`

List of `sourceRoleArn` values that have been configured for cross-account access.

_Required_: No

_Type_: Array of [CrossAccountConfiguration](aws-properties-aiops-investigationgroup-crossaccountconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionConfig`

Specifies the customer managed AWS KMS key that the investigation group uses
to encrypt data, if there is one. If not, the investigation group uses an AWS key to encrypt the data.

_Required_: No

_Type_: [EncryptionConfigMap](aws-properties-aiops-investigationgroup-encryptionconfigmap.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InvestigationGroupPolicy`

Returns the JSON of the IAM resource policy associated with the specified
investigation group in a string. For example, `{\"Version\":\"2012-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":{\"Service\":\"aiops.alarms.cloudwatch.amazonaws.com\"},\"Action\":[\"aiops:CreateInvestigation\",\"aiops:CreateInvestigationEvent\"],\"Resource\":\"*\",\"Condition\":{\"StringEquals\":{\"aws:SourceAccount\":\"111122223333\"},\"ArnLike\":{\"aws:SourceArn\":\"arn:aws:cloudwatch:us-east-1:111122223333:alarm:*\"}}}]}`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsCloudTrailEventHistoryEnabled`

Specify `true` to enable CloudWatch investigations to have access to change events that are
recorded by CloudTrail. The default is `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Specify either the name or the ARN of the investigation group that you want to
view. This is used to set the name of the investigation group.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RetentionInDays`

Specifies how long that investigation data is kept.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The ARN of the IAM role that the investigation group uses for permissions
to gather data.

_Required_: No

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TagKeyBoundaries`

Displays the custom tag keys for custom applications in your system that you have
specified in the investigation group. Resource tags help CloudWatch investigations narrow the
search space when it is unable to discover definite relationships between resources.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The list of key-value pairs to associate with the resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-aiops-investigationgroup-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the investigation group. For example,
`arn:aws:aiops:Region:account-id:investigation-group:investigation-group-id`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the investigation group.

`CreatedAt`

The date and time that the investigation group was created.

`CreatedBy`

The name of the user who created the investigation group.

`LastModifiedAt`

The date and time that the investigation group was most recently modified.

`LastModifiedBy`

The name of the user who created the investigation group.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudWatch investigations

ChatbotNotificationChannel

All content copied from https://docs.aws.amazon.com/.
