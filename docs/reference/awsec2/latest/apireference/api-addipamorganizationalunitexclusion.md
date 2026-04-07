# AddIpamOrganizationalUnitExclusion

Add an Organizational Unit (OU) exclusion to your IPAM. If your IPAM is integrated with AWS Organizations and you add an organizational unit (OU) exclusion, IPAM will not manage the IP addresses in accounts in that OU exclusion. There is a limit on the number of exclusions you can create. For more information, see [Quotas for your IPAM](../../../../services/vpc/latest/ipam/quotas-ipam.md) in the _Amazon VPC IPAM User Guide_.

## Contents

**OrganizationsEntityPath**

An AWS Organizations entity path. Build the path for the OU(s) using AWS Organizations IDs separated by a `/`. Include all child OUs by ending the path with `/*`.

- Example 1

- Path to a child OU: `o-a1b2c3d4e5/r-f6g7h8i9j0example/ou-ghi0-awsccccc/ou-jkl0-awsddddd/`

- In this example, `o-a1b2c3d4e5` is the organization ID, `r-f6g7h8i9j0example` is the root ID , `ou-ghi0-awsccccc` is an OU ID, and `ou-jkl0-awsddddd` is a child OU ID.

- IPAM will not manage the IP addresses in accounts in the child OU.

- Example 2

- Path where all child OUs will be part of the exclusion: `o-a1b2c3d4e5/r-f6g7h8i9j0example/ou-ghi0-awsccccc/*`

- In this example, IPAM will not manage the IP addresses in accounts in the OU ( `ou-ghi0-awsccccc`) or in accounts in any OUs that are children of the OU.

For more information on how to construct an entity path, see [Understand the AWS Organizations entity path](../../../../services/iam/latest/userguide/access-policies-last-accessed-view-data-orgs.md#access_policies_access-advisor-viewing-orgs-entity-path) in the _AWS Identity and Access Management User Guide_.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/addipamorganizationalunitexclusion.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/addipamorganizationalunitexclusion.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/addipamorganizationalunitexclusion.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AddIpamOperatingRegion

AdditionalDetail
