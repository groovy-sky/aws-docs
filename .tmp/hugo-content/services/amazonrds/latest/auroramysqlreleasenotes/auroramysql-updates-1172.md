# Aurora MySQL database engine updates 2018-04-27 (version 1.17.2) (Deprecated)

**Version:** 1.17.2

Aurora MySQL 1.17.2 is generally available. All new Aurora MySQL database clusters with MySQL 5.6
compatibility, including those restored
from snapshots, will be created in Aurora MySQL 1.17.2. You have the option, but are not required, to
upgrade existing database clusters to Aurora MySQL 1.17.2. You can create new database clusters in Aurora MySQL
1.14.4, Aurora MySQL 1.15.1, or Aurora MySQL 1.16. You can do so using the AWS CLI or the Amazon RDS API and specifying
the engine version.

With version 1.17.2 of Aurora MySQL, we are using a cluster patching model where all nodes in an Aurora DB
cluster are patched at the same time.

If you have any questions or concerns, AWS Support is available on the community forums and through
[AWS Support](https://aws.amazon.com/support). For more information, see
[Maintaining an Amazon Aurora DB cluster](../aurorauserguide/user-upgradedbinstance-maintenance.md) in the _Amazon Aurora User Guide_.

## Improvements

- Fixed an issue which was causing restarts during certain DDL partition operations.

- Fixed an issue which was causing support for invocation of AWS Lambda functions via native Aurora MySQL functions to be disabled.

- Fixed an issue with cache invalidation which was causing restarts on Aurora Replicas.

- Fixed an issue in lock manager which was causing restarts.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Aurora MySQL updates: 2018-06-05 (version 1.17.3) (Deprecated)

Aurora MySQL updates: 2018-03-23 (version 1.17.1) (Deprecated)

All content copied from https://docs.aws.amazon.com/.
