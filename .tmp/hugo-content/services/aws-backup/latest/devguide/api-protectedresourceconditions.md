# ProtectedResourceConditions

The conditions that you define for resources in
your restore testing plan using tags.

## Contents

**StringEquals**

Filters the values of your tagged resources for only
those resources that you tagged with the same value.
Also called "exact matching."

Type: Array of [KeyValue](api-keyvalue.md) objects

Required: No

**StringNotEquals**

Filters the values of your tagged resources for only
those resources that you tagged that do not have the same value.
Also called "negated matching."

Type: Array of [KeyValue](api-keyvalue.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/protectedresourceconditions.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/protectedresourceconditions.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/protectedresourceconditions.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProtectedResource

RecoveryPointByBackupVault

All content copied from https://docs.aws.amazon.com/.
