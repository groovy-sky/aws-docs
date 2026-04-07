This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Organizations::Policy

Creates a policy of a specified type that you can attach to a root, an organizational
unit (OU), or an individual AWS account.

For more information about policies and their use, see [Managing AWS Organizations\
policies](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies.html).

If the request includes tags, then the requester must have the
`organizations:TagResource` permission.

This operation can be called only from the organization's management account or a member account designated as a delegated administrator.

###### Note

Before you can create a policy of a given type, you must first [enable that policy type](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_enable-disable.html) in your organization.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Organizations::Policy",
  "Properties" : {
      "Content" : Json,
      "Description" : String,
      "Name" : String,
      "Tags" : [ Tag, ... ],
      "TargetIds" : [ String, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::Organizations::Policy
Properties:
  Content: Json
  Description: String
  Name: String
  Tags:
    - Tag
  TargetIds:
    - String
  Type: String

```

## Properties

`Content`

The policy text content. You can specify the policy content as a JSON object or a JSON
string.

###### Important

When you specify the policy content as a JSON string, you can't perform drift
detection on the CloudFormation stack. For this reason, we recommend
specifying the policy content as a JSON object instead.

The text that you supply must adhere to the rules of the policy type you specify in
the `Type` parameter. The following AWS Organizations quotas are enforced
for the maximum size of a policy document:

- Service control policies: 5,120 characters

- Resource control policies: 5,120 characters

- Declarative policies: 10,000 characters

- Backup policies: 10,000 characters

- Tag policies: 10,000 characters

- Chat applications policies: 10,000 characters

- AI services opt-out policies: 2,500 characters

- Security Hub policies: 10,000 characters

- Amazon Inspector policies: 10,000 characters

- Amazon Bedrock policies: 10,000 characters

- Upgrade rollout policies: 10,000 characters

For more information about Organizations service quotas, see [Quotas for AWS Organizations](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_reference_limits.html) in the _AWS Organizations_
_User Guide_.

_Required_: Yes

_Type_: Json

_Pattern_: `[\s\S]*`

_Minimum_: `1`

_Maximum_: `1000000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

Human readable description of the policy.

_Required_: No

_Type_: String

_Pattern_: `[\s\S]*`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Name of the policy.

The [regex pattern](http://wikipedia.org/wiki/regex) that is used to
validate this parameter is a string of any of the characters in the ASCII character
range.

_Required_: Yes

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of tags that you want to attach to the newly created policy. For each tag in
the list, you must specify both a tag key and a value. You can set the value to an empty
string, but you can't set it to `null`. For more information about tagging,
see [Tagging AWS Organizations\
resources](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_tagging.html) in the AWS Organizations User Guide.

###### Note

If any one of the tags is not valid or if you exceed the allowed number of tags
for a policy, then the entire request fails and the policy is not created.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-organizations-policy-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetIds`

List of unique identifiers (IDs) of the root, OU, or account that you want to attach
the policy to. You can get the ID by calling the [ListRoots](https://docs.aws.amazon.com/organizations/latest/APIReference/API_ListRoots.html), [ListOrganizationalUnitsForParent](https://docs.aws.amazon.com/organizations/latest/APIReference/API_ListOrganizationalUnitsForParent.html), or [ListAccounts](https://docs.aws.amazon.com/organizations/latest/APIReference/API_ListAccounts.html)
operations. If you don't specify this parameter, the policy is created but not attached
to any organization resource.

The [regex pattern](http://wikipedia.org/wiki/regex) for a target ID
string requires one of the following:

- **Root** \- A string that begins with "r-" followed
by from 4 to 32 lowercase letters or digits.

- **Account** \- A string that consists of exactly 12
digits.

- **Organizational unit (OU)** \- A string that begins
with "ou-" followed by from 4 to 32 lowercase letters or digits (the ID of the
root that the OU is in). This string is followed by a second "-" dash and from 8
to 32 additional lowercase letters or digits.

_Required_: No

_Type_: Array of String

_Pattern_: `^(r-[0-9a-z]{4,32})|(\d{12})|(ou-[0-9a-z]{4,32}-[a-z0-9]{8,32})$`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of policy to create.

_Required_: Yes

_Type_: String

_Allowed values_: `AISERVICES_OPT_OUT_POLICY | BACKUP_POLICY | BEDROCK_POLICY | CHATBOT_POLICY | DECLARATIVE_POLICY_EC2 | INSPECTOR_POLICY | NETWORK_SECURITY_DIRECTOR_POLICY | RESOURCE_CONTROL_POLICY | S3_POLICY | SECURITYHUB_POLICY | SERVICE_CONTROL_POLICY | TAG_POLICY | UPGRADE_ROLLOUT_POLICY`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `Id`. For example:
`p-examplepolicyid111`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the Amazon Resource Name (ARN) of the policy. For example:
`arn:aws:organizations::111111111111:policy/o-exampleorgid/service_control_policy/p-examplepolicyid111`.

`AwsManaged`

Returns a boolean value that indicates whether the specified policy is an AWS managed
policy. If true, then you can attach the policy to roots, OUs, or accounts, but you
cannot edit it. For example: `true | false`.

`Id`

Returns the unique identifier (ID) of the policy. For example:
`p-examplepolicyid111`.

## Examples

- [Organization Policy Content Specified as a JSON Object](#aws-resource-organizations-policy--examples--Organization_Policy_Content_Specified_as_a_JSON_Object)

- [Organization Policy Content Specified as a JSON String](#aws-resource-organizations-policy--examples--Organization_Policy_Content_Specified_as_a_JSON_String)

### Organization Policy Content Specified as a JSON Object

This example illustrates how to specify the organization policy content as a
JSON object in `AWS::Organizations::Policy`. The organization policy
is specified inline as a JSON object in the `Content` property of
`AWS::Organizations::Policy`.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Description": "AWS CloudFormation Organizations Template Example",
  "Resources": {
    "PolicyTestTemplate": {
      "DeletionPolicy": "Retain",
      "Type": "AWS::Organizations::Policy",
      "Properties": {
        "Type": "SERVICE_CONTROL_POLICY",
        "Name": "SCPDenyLeaveOrganization",
        "Content": {
          "Version": "2012-10-17",
          "Statement": [
            {
              "Sid": "SCPDenyLeaveOrganization",
              "Effect": "Deny",
              "Action": [
                "organizations:LeaveOrganization"
              ],
              "Resource": "*"
            }
          ]
        }
      }
    }
  }
}

```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: AWS CloudFormation Organizations Template Example
Resources:
  PolicyTestTemplate:
    DeletionPolicy: Retain
    Type: AWS::Organizations::Policy
    Properties:
      Type: SERVICE_CONTROL_POLICY
      Name: SCPDenyLeaveOrganization
      Content:
        Version: 2012-10-17
        Statement:
          - Sid: SCPDenyLeaveOrganization
            Effect: Deny
            Action:
              - 'organizations:LeaveOrganization'
            Resource: '*'

```

### Organization Policy Content Specified as a JSON String

This example illustrates how to specify the organization policy content as a
JSON string in `AWS::Organizations::Policy`. The organization policy
is specified inline as a JSON string in the `Content` property of
`AWS::Organizations::Policy`.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Description": "AWS CloudFormation Organizations Template Example",
  "Resources": {
    "PolicyTestTemplate": {
      "DeletionPolicy": "Retain",
      "Type": "AWS::Organizations::Policy",
      "Properties": {
        "Type": "SERVICE_CONTROL_POLICY",
        "Name": "SCPDenyLeaveOrganization",
        "Content": "{\"Version\":\"2012-10-17\",\"Statement\":[{\"Sid\":\"SCPDenyLeaveOrganization\",\"Effect\":\"Deny\",\"Action\":[\"organizations:LeaveOrganization\"],\"Resource\":\"*\"}]}"
      }
    }
  }
}

```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: AWS CloudFormation Organizations Template Example
Resources:
  PolicyTestTemplate:
    DeletionPolicy: Retain
    Type: AWS::Organizations::Policy
    Properties:
      Type: SERVICE_CONTROL_POLICY
      Name: SCPDenyLeaveOrganization
      Content: >-
        {"Version":"2012-10-17","Statement":[{"Sid":"SCPDenyLeaveOrganization","Effect":"Deny","Action":["organizations:LeaveOrganization"],"Resource":"*"}]}

```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: AWS CloudFormation Organizations Template Example
Resources:
  PolicyTestTemplate:
    DeletionPolicy: Retain
    Type: AWS::Organizations::Policy
    Properties:
      Type: SERVICE_CONTROL_POLICY
      Name: SCPDenyLeaveOrganization
      Content: >-
        {
          "Version": "2012-10-17",
          "Statement": [
            {
              "Sid": "SCPDenyLeaveOrganization",
              "Effect": "Deny",
              "Action": [
                "organizations:LeaveOrganization"
              ],
              "Resource": "*"
            }
          ]
        }

```

## See also

- [Managing\
AWS Organizations policies](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies.html) in the _AWS Organizations User Guide_.

- [CreatePolicy](https://docs.aws.amazon.com/organizations/latest/APIReference/API_CreatePolicy.html) in the _AWS Organizations API_
_Reference Guide_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
