This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::StackSet DeploymentTargets

The AWS Organizations accounts or AWS accounts to deploy stacks to
in the specified Regions.

When deploying to AWS Organizations accounts with `SERVICE_MANAGED`
permissions:

- You must specify the `OrganizationalUnitIds` property.

- If you specify organizational units (OUs) for
`OrganizationalUnitIds` and use either the `Accounts`
or `AccountsUrl` property, you must also specify the
`AccountFilterType` property.

When deploying to AWS accounts with `SELF_MANAGED`
permissions:

- You must specify either the `Accounts` or `AccountsUrl`
property, but not both.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccountFilterType" : String,
  "Accounts" : [ String, ... ],
  "AccountsUrl" : String,
  "OrganizationalUnitIds" : [ String, ... ]
}

```

### YAML

```yaml

  AccountFilterType: String
  Accounts:
    - String
  AccountsUrl: String
  OrganizationalUnitIds:
    - String

```

## Properties

`AccountFilterType`

Refines which accounts to deploy stacks to by specifying how to use the
`Accounts` and `OrganizationalUnitIds` properties
together.

The following values determine how CloudFormation selects target
accounts:

- `INTERSECTION`: StackSet deploys to the accounts specified in the
`Accounts` property.

- `DIFFERENCE`: StackSet deploys to the OU, excluding the accounts
specified in the `Accounts` property.

- `UNION`: StackSet deploys to the OU, and the accounts specified in
the `Accounts` property. `UNION` is not supported for
create operations when using StackSet as a resource or the
`CreateStackInstances` API.

_Required_: No

_Type_: String

_Allowed values_: `NONE | UNION | INTERSECTION | DIFFERENCE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Accounts`

The account IDs of the AWS accounts. If you have many account numbers,
you can provide those accounts using the `AccountsUrl` property
instead.

_Pattern_: `^[0-9]{12}$`

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AccountsUrl`

The Amazon S3 URL path to a file that contains a list of AWS account IDs. The file
format must be either `.csv` or `.txt`, and the data can be comma-separated
or new-line-separated. There is currently a 10MB limit for the data (approximately 800,000
accounts).

This property serves the same purpose as `Accounts` but allows you to specify a
large number of accounts.

_Required_: No

_Type_: String

_Pattern_: `(s3://|http(s?)://).+`

_Minimum_: `1`

_Maximum_: `5120`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrganizationalUnitIds`

The organization root ID or organizational unit (OU) IDs.

_Pattern_:
`^(ou-[a-z0-9]{4,32}-[a-z0-9]{8,32}|r-[a-z0-9]{4,32})$`

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AutoDeployment

ManagedExecution

All content copied from https://docs.aws.amazon.com/.
