# Aurora MySQL database engine updates 2019-12-23 (version 1.22.1) (Deprecated)

**Version:** 1.22.1

Aurora MySQL 1.22.1 is generally available.
Aurora MySQL 1.\* versions are compatible with MySQL 5.6 and Aurora MySQL 2.\* versions are compatible with MySQL 5.7.

This engine version is scheduled to be deprecated on February 28, 2023. For more information, see
[Preparing for Amazon Aurora MySQL-Compatible Edition version 1 end of life](../aurorauserguide/aurora-mysql56-eol.md).

Currently supported Aurora MySQL releases are 1.19.5, 1.19.6, 1.22.\*, 1.23.\*, 2.04.\*, 2.07.\*, 2.08.\*, 2.09.\*, 2.10.\*, 3.01.\* and 3.02.\*.

To create a cluster with an older version of Aurora MySQL, please specify the engine version through the AWS Management Console, the AWS CLI
or the RDS API. You have the option to upgrade existing Aurora MySQL 1.\* database clusters to Aurora MySQL 1.22.1.

###### Note

This version is currently not available in the following AWS Regions:
AWS GovCloud (US-East) \[us-gov-east-1\], AWS GovCloud (US-West) \[us-gov-west-1\],
China (Ningxia) \[cn-northwest-1\], Asia Pacific (Hong Kong) \[ap-east-1\], and
Middle East (Bahrain) \[me-south-1. There will be a separate announcement once it
is made available.

If you have any questions or concerns, AWS Support is available on the community forums and through
[AWS Support](https://aws.amazon.com/support). For more information, see
[Maintaining an Amazon Aurora DB cluster](../aurorauserguide/user-upgradedbinstance-maintenance.md) in the _Amazon Aurora User Guide_.

###### Note

The procedure to upgrade your DB cluster has changed. For more information, see [Upgrading the minor version or patch level of an Aurora MySQL DB cluster](../aurorauserguide/auroramysql-updates-patching.md) in the
_Amazon Aurora User Guide_.

## Improvements

**Critical fixes:**

- Fixed issues that prevented engine recovery involving table locks and temporary tables.

- Improved the stability of binary log when temporary tables are used.

**High priority fixes:**

- Fixed a slow memory leak in Aurora specific database tracing and logging sub-system that lowers the freeable memory.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Aurora MySQL updates: 2020-03-05 (version 1.22.2) (Deprecated)

Aurora MySQL updates: 2019-11-25 (version 1.22.0) (Deprecated)

All content copied from https://docs.aws.amazon.com/.
