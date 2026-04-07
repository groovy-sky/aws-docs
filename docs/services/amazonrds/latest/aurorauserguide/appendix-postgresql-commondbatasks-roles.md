# Understanding PostgreSQL roles and permissions

When you create an Aurora PostgreSQL DB cluster
using the AWS Management Console, an
administrator account is created at the same time. By default, its name is
`postgres`, as shown in the following screenshot:

![The default login identity for Credentials in the Create database page is postgres.](https://docs.aws.amazon.com/images/AmazonRDS/latest/AuroraUserGuide/images/default-login-identity-apg-rpg.png)

You can choose another name rather than accept the default ( `postgres`). If you
do, the name you choose must start with a letter and be between 1 and 16 alphanumeric
characters. For simplicity's sake, we refer to this main user account by its default value
( `postgres`) throughout this guide.

If you use the `create-db-cluster` AWS CLI rather than the
AWS Management Console, you create the user name by passing it with the `master-username`
parameter. For more information, see [Step 2: Create an Aurora PostgreSQL DB cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_GettingStartedAurora.AuroraPostgreSQL.FullConfig.html#CHAP_GettingStarted.AuroraPostgreSQL.CreateDBCluster).

Whether you use the AWS Management Console, the AWS CLI, or the Amazon RDS API, and whether you use the default
`postgres` name or choose a different name, this first database user account is a
member of the `rds_superuser` group and has `rds_superuser`
privileges.

###### Topics

- [Understanding the rds\_superuser role](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Appendix.PostgreSQL.CommonDBATasks.Roles.rds_superuser.html)

- [Controlling user access to the PostgreSQL database](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Appendix.PostgreSQL.CommonDBATasks.Access.html)

- [Delegating and controlling user password management](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Appendix.PostgreSQL.CommonDBATasks.RestrictPasswordMgmt.html)

- [Using SCRAM for PostgreSQL password encryption](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/PostgreSQL_Password_Encryption_configuration.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Security with Aurora PostgreSQL

Understanding the rds\_superuser role
