# LatestRevokeRequest

Contains information about the latest request to revoke access to a backup vault.

## Contents

**ExpiryDate**

The date and time when the revoke request will expire.

Type: Timestamp

Required: No

**InitiationDate**

The date and time when the revoke request was initiated.

Type: Timestamp

Required: No

**MpaSessionArn**

The ARN of the MPA session associated with this revoke request.

Type: String

Required: No

**Status**

The current status of the revoke request.

Type: String

Valid Values: `PENDING | FAILED`

Required: No

**StatusMessage**

A message describing the current status of the revoke request.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/latestrevokerequest.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/latestrevokerequest.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/latestrevokerequest.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LatestMpaApprovalTeamUpdate

LegalHold

All content copied from https://docs.aws.amazon.com/.
