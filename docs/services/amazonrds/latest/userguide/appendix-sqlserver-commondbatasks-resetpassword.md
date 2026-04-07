# Resetting the db\_owner role membership for master user for Amazon RDS for SQL Server

If you lock your master user out of the `db_owner` role membership on your
RDS for SQL Server database and no other database user can grant the membership, you can restore
lost membership by modifying the DB instance master user password.

By changing the DB instance master user password, RDS grants the `db_owner`
membership to the databases in the DB instance that might have been accidentally
revoked. You can change the DB instance password by using the Amazon RDS console, the AWS CLI
command [modify-db-instance](https://docs.aws.amazon.com/cli/latest/reference/rds/modify-db-instance.html), or by using the [ModifyDBInstance](../../../../reference/amazonrds/latest/apireference/api-modifydbinstance.md) API operation.
For more information about modifying a DB instance, see [Modifying an Amazon RDS DB instance](overview-dbinstance-modifying.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Renaming a Multi-AZ database

Restoring
license-terminated DB instances
