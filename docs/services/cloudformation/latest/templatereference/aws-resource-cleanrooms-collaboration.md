This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::Collaboration

Creates a new collaboration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CleanRooms::Collaboration",
  "Properties" : {
      "AllowedResultRegions" : [ String, ... ],
      "AnalyticsEngine" : String,
      "AutoApprovedChangeTypes" : [ String, ... ],
      "CreatorDisplayName" : String,
      "CreatorMemberAbilities" : [ String, ... ],
      "CreatorMLMemberAbilities" : MLMemberAbilities,
      "CreatorPaymentConfiguration" : PaymentConfiguration,
      "DataEncryptionMetadata" : DataEncryptionMetadata,
      "Description" : String,
      "IsMetricsEnabled" : Boolean,
      "JobLogStatus" : String,
      "Members" : [ MemberSpecification, ... ],
      "Name" : String,
      "QueryLogStatus" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CleanRooms::Collaboration
Properties:
  AllowedResultRegions:
    - String
  AnalyticsEngine: String
  AutoApprovedChangeTypes:
    - String
  CreatorDisplayName: String
  CreatorMemberAbilities:
    - String
  CreatorMLMemberAbilities:
    MLMemberAbilities
  CreatorPaymentConfiguration:
    PaymentConfiguration
  DataEncryptionMetadata:
    DataEncryptionMetadata
  Description: String
  IsMetricsEnabled: Boolean
  JobLogStatus: String
  Members:
    - MemberSpecification
  Name: String
  QueryLogStatus: String
  Tags:
    - Tag

```

## Properties

`AllowedResultRegions`

The AWS Regions where collaboration query results can be stored. Returns
the list of Region identifiers that were specified when the collaboration was created. This
list is used to enforce regional storage policies and compliance requirements.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AnalyticsEngine`

The analytics engine for the collaboration.

###### Note

After July 16, 2025, the `CLEAN_ROOMS_SQL` parameter will no longer be
available.

_Required_: No

_Type_: String

_Allowed values_: `CLEAN_ROOMS_SQL | SPARK`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutoApprovedChangeTypes`

The types of change requests that are automatically approved for this collaboration.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CreatorDisplayName`

A display name of the collaboration creator.

_Required_: Yes

_Type_: String

_Pattern_: `^(?!\s*$)[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDBFF-\uDC00\uDFFF\t]*$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CreatorMemberAbilities`

The abilities granted to the collaboration creator.

_Allowed values_ `CAN_QUERY` \| `CAN_RECEIVE_RESULTS` \| `CAN_RUN_JOB`

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CreatorMLMemberAbilities`

The ML member abilities for a collaboration member.

_Required_: No

_Type_: [MLMemberAbilities](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cleanrooms-collaboration-mlmemberabilities.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CreatorPaymentConfiguration`

An object representing the collaboration member's payment responsibilities set by the
collaboration creator.

_Required_: No

_Type_: [PaymentConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cleanrooms-collaboration-paymentconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataEncryptionMetadata`

The settings for client-side encryption for cryptographic computing.

_Required_: No

_Type_: [DataEncryptionMetadata](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cleanrooms-collaboration-dataencryptionmetadata.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

A description of the collaboration provided by the collaboration owner.

_Required_: Yes

_Type_: String

_Pattern_: `^(?!\s*$)[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDBFF-\uDC00\uDFFF\t\r\n]*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsMetricsEnabled`

An indicator as to whether metrics are enabled for the collaboration.

When `true`, collaboration members can opt in to Amazon CloudWatch metrics
for their membership queries.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JobLogStatus`

An indicator as to whether job logging has been enabled or disabled
for the collaboration.

When `ENABLED`, AWS Clean Rooms logs details about jobs run within this
collaboration and those logs can be viewed in Amazon CloudWatch Logs. The default value is
`DISABLED`.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Members`

A list of initial members, not including the creator. This list is immutable.

_Required_: No

_Type_: Array of [MemberSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cleanrooms-collaboration-memberspecification.html)

_Minimum_: `0`

_Maximum_: `9`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

A human-readable identifier provided by the collaboration owner. Display names are not
unique.

_Required_: Yes

_Type_: String

_Pattern_: `^(?!\s*$)[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDBFF-\uDC00\uDFFF\t]*$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryLogStatus`

An indicator as to whether query logging has been enabled or disabled for the
collaboration.

When `ENABLED`, AWS Clean Rooms logs details about queries run within this
collaboration and those logs can be viewed in Amazon CloudWatch Logs. The default value is
`DISABLED`.

_Required_: Yes

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An optional label that you can assign to a resource when you create it. Each tag
consists of a key and an optional value, both of which you define. When you use tagging,
you can also use tag-based access control in IAM policies to control access
to this resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cleanrooms-collaboration-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `CollaborationIdentifier`, such as
`a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`. For example:

`{ "Ref": "MyCollaboration" }`

For more information about using the `Ref` function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the Amazon Resource Name (ARN) of the specified collaboration.

Example:
`arn:aws:cleanrooms:us-east-1:111122223333:collaboration/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`

`CollaborationIdentifier`

Returns the unique identifier of the specified collaboration.

Example: `a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`

## Examples

### Create a collaboration

The following example creates a collaboration with the collaboration creator.

#### JSON

```json

"ExampleCollaboration": {
  {
    "Type": "AWS::CleanRooms::Collaboration",
    "Properties": {
      "Name": "Example Collaboration",
      "Description": "Example AWS Clean Rooms collaboration",
      "CreatorDisplayName": "Member 1",
      "CreatorMemberAbilities": ["CAN_QUERY", "CAN_RECEIVE_RESULTS"],
      "Members": [
        {
          "AccountId": "111122223333",
          "DisplayName": "Member 2",
          "MemberAbilities": []
        },
        {
          "AccountId": "444455556666",
          "DisplayName": "Member 3",
          "MemberAbilities": []
        }
      ],
      "QueryLogStatus": "ENABLED"
    }
  }
}
```

#### YAML

```yaml

ExampleCollaboration:
  Type: AWS::CleanRooms::Collaboration
  Properties:
    Name: Example Collaboration
    Description: Example AWS Clean Rooms collaboration
    CreatorDisplayName: Member 1
    CreatorMemberAbilities:
      - CAN_QUERY
      - CAN_RECEIVE_RESULTS
    Members:
      - AccountId: 111122223333
        DisplayName: Member 2
        MemberAbilities: []
      - AccountId: 444455556666
        DisplayName: Member 3
        MemberAbilities: []
    QueryLogStatus: ENABLED
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

DataEncryptionMetadata
