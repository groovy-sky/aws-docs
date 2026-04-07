This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Organizations::Organization

Creates an AWS organization. The account whose user is calling the
[`CreateOrganization`](https://docs.aws.amazon.com/organizations/latest/APIReference/API_CreateOrganization.html) operation automatically becomes the
[management account](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#account) of the new organization.

This operation must be called using credentials from the account that is to become the
new organization's management account. The principal must also have the [relevant IAM\
permissions](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_create.html).

###### Important

- If you delete an organization, you can't recover it. If you created any
policies inside of the organization, they're also deleted and you can't
recover them.

- You can delete an organization only after you remove all member accounts
from the organization. If you created some of your member accounts using
AWS Organizations, you might be blocked from removing those
accounts. You can remove a member account only if it has all the information
that's required to operate as a standalone AWS account. For
more information about how to provide that information and then remove the
account, see [Leave an organization from your member account](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_leave-as-member.html) in the _AWS Organizations User Guide_.

- If you closed a member account before you remove it from the organization,
it enters a 'suspended' state for a period of time and you can't remove the
account from the organization until it is finally closed. This can take up
to 90 days and can prevent you from deleting the organization until all
member accounts are completely closed.

For more information, see [Deleting an\
organization](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_delete.html) in the _AWS Organizations User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Organizations::Organization",
  "Properties" : {
      "FeatureSet" : String
    }
}

```

### YAML

```yaml

Type: AWS::Organizations::Organization
Properties:
  FeatureSet: String

```

## Properties

`FeatureSet`

Specifies the feature set supported by the new organization. Each feature set supports
different levels of functionality.

- `ALL`– In addition to all the features supported by the
consolidated billing feature set, the management account gains access to
advanced features that give you more control over accounts in your organization. For more information, see [All features](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#feature-set-all) in the _AWS Organizations User Guide_.

- `CONSOLIDATED_BILLING`– All member accounts have their bills consolidated to and
paid by the management account. For more information, see [Consolidated billing](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#feature-set-cb-only) in the _AWS Organizations User Guide_.

###### Note

The consolidated billing feature feature set isn't available for organizations in
the AWS GovCloud (US) Region.

If you don't specify this property, the default value is `ALL`.

_Required_: No

_Type_: String

_Allowed values_: `ALL | CONSOLIDATED_BILLING`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `AccountId`. For example:
`123456789012`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of an organization.

`Id`

The unique identifier (ID) of an organization.

`ManagementAccountArn`

The Amazon Resource Name (ARN) of the account that is designated as the management
account for the organization.

`ManagementAccountEmail`

The email address that is associated with the AWS account that is
designated as the management account for the organization.

`ManagementAccountId`

The unique identifier (ID) of the management account of an organization.

`RootId`

The unique identifier (ID) for the root.

## Examples

- [Organization FeatureSet specified as ALL](#aws-resource-organizations-organization--examples--Organization_FeatureSet_specified_as_ALL)

- [Organization FeatureSet specified as CONSOLIDATED\_BILLING](#aws-resource-organizations-organization--examples--Organization_FeatureSet_specified_as_CONSOLIDATED_BILLING)

### Organization FeatureSet specified as ALL

This example illustrates how to specify the organization feature set as
`ALL` in `AWS::Organizations::Organization`.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Description": "AWS CloudFormation Organizations Template Example",
  "Resources": {
    "OrganizationTemplateExample": {
      "DeletionPolicy": "Retain",
      "Type": "AWS::Organizations::Organization",
      "Properties": {
        "FeatureSet": "ALL"
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
  OrganizationTemplateExample:
    DeletionPolicy: Retain
    Type: 'AWS::Organizations::Organization'
    Properties:
      FeatureSet: ALL

```

### Organization FeatureSet specified as CONSOLIDATED\_BILLING

This example illustrates how to specify the organization feature set as
`CONSOLIDATED_BILLING` in
`AWS::Organizations::Organization`.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Description": "AWS CloudFormation Organizations Template Example",
  "Resources": {
    "OrganizationTemplateExample": {
      "DeletionPolicy": "Retain",
      "Type": "AWS::Organizations::Organization",
      "Properties": {
        "FeatureSet": "CONSOLIDATED_BILLING"
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
  OrganizationTemplateExample:
    DeletionPolicy: Retain
    Type: 'AWS::Organizations::Organization'
    Properties:
      FeatureSet: CONSOLIDATED_BILLING

```

## See also

- [Creating an organization](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_create.html) in the _AWS Organizations_
_User Guide_.

- [CreateOrganization](https://docs.aws.amazon.com/organizations/latest/APIReference/API_CreateOrganization.html) in the _AWS Organizations API_
_Reference Guide_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::Organizations::OrganizationalUnit
