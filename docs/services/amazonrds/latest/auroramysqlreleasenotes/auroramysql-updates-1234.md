# Aurora MySQL database engine updates 2021-09-30 (version 1.23.4) (Deprecated)

**Version:** 1.23.4

Aurora MySQL 1.23.4 is generally available.
Aurora MySQL 2.\* versions are compatible with MySQL 5.7 and
Aurora MySQL 1.\* versions are compatible with MySQL 5.6.

This engine version is scheduled to be deprecated on February 28, 2023. For more information, see
[Preparing for Amazon Aurora MySQL-Compatible Edition version 1 end of life](../aurorauserguide/aurora-mysql56-eol.md).

Currently supported Aurora MySQL releases are 1.19.5, 1.19.6, 1.22.\*, 1.23.\*, 2.04.\*, 2.07.\*, 2.08.\*, 2.09.\*, 2.10.\*, 3.01.\* and 3.02.\*.

To create a cluster with an older version of Aurora MySQL, specify the engine version through the
RDS Console, the AWS CLI, or the Amazon RDS API.

If you have any questions or concerns, AWS Support is available on the community forums and through
[AWS Support](https://aws.amazon.com/support). For more information, see
[Maintaining an Amazon Aurora DB cluster](../aurorauserguide/user-upgradedbinstance-maintenance.md) in the _Amazon Aurora User Guide_.

## Improvements

**General improvements:**

- Fixed an issue that can cause high CPU consumption on the
reader instances due to excessive logging of informational messages
in internal diagnostic log files.

**High-priority fixes:**

- [CVE-2021-2307](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2021-2307)

- [CVE-2021-2226](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2021-2226)

- [CVE-2021-2160](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2021-2160)

- [CVE-2021-2154](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2021-2154)

- [CVE-2021-2060](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2021-2060)

- [CVE-2021-2032](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2021-2032)

- [CVE-2021-2001](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2021-2001)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Aurora MySQL version 1 (Deprecated)

Aurora MySQL updates: 2021-06-28 (version 1.23.3) (Deprecated)

All content copied from https://docs.aws.amazon.com/.
