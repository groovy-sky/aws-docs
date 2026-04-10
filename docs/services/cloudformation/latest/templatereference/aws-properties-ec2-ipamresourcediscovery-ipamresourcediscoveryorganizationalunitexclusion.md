This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::IPAMResourceDiscovery IpamResourceDiscoveryOrganizationalUnitExclusion

If your IPAM is integrated with AWS Organizations,
you can exclude an [organizational unit (OU)](../../../organizations/latest/userguide/orgs-getting-started-concepts.md#organizationalunit) from being managed by IPAM. When you exclude an OU, IPAM will not manage the IP addresses in accounts in that OU. For more information, see [Exclude organizational units from IPAM](../../../vpc/latest/ipam/exclude-ous.md) in the _Amazon Virtual Private Cloud IP Address Manager User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OrganizationsEntityPath" : String
}

```

### YAML

```yaml

  OrganizationsEntityPath: String

```

## Properties

`OrganizationsEntityPath`

An AWS Organizations entity path. For more information on the entity path, see [Understand the AWS Organizations entity path](../../../iam/latest/userguide/access-policies-last-accessed-view-data-orgs.md#access_policies_access-advisor-viewing-orgs-entity-path) in the _AWS Identity and Access Management User Guide_.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IpamOperatingRegion

Tag

All content copied from https://docs.aws.amazon.com/.
