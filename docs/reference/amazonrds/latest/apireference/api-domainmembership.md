# DomainMembership

An Active Directory Domain membership record associated with the DB instance or cluster.

## Contents

###### Note

In the following list, the required parameters are described first.

**AuthSecretArn**

The ARN for the Secrets Manager secret with the credentials for the user that's a member of the domain.

Type: String

Required: No

**DnsIps.member.N**

The IPv4 DNS IP addresses of the primary and secondary Active Directory domain controllers.

Type: Array of strings

Required: No

**Domain**

The identifier of the Active Directory Domain.

Type: String

Required: No

**FQDN**

The fully qualified domain name (FQDN) of the Active Directory Domain.

Type: String

Required: No

**IAMRoleName**

The name of the IAM role used when making API calls to the Directory Service.

Type: String

Required: No

**OU**

The Active Directory organizational unit for the DB instance or cluster.

Type: String

Required: No

**Status**

The status of the Active Directory Domain membership for the DB instance or cluster. Values include `joined`, `pending-join`, `failed`, and so on.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/domainmembership.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/domainmembership.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/domainmembership.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocLink

DoubleRange

All content copied from https://docs.aws.amazon.com/.
