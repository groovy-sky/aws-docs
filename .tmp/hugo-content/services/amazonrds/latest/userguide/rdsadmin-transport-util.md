# Transporting tablespaces

Use the Amazon RDS package `rdsadmin.rdsadmin_transport_util` to copy a set of
tablespaces from an on-premises Oracle database to an RDS for Oracle DB instance. At the physical level,
the transportable tablespace feature incrementally copies source data files and metadata
files to your target instance. You can transfer the files using either Amazon EFS or Amazon S3. For
more information, see [Migrating using Oracle transportable tablespaces](oracle-migrating-tts.md).

###### Topics

- [Importing transported tablespaces to your DB instance](rdsadmin-transport-util-import-xtts-tablespaces.md)

- [Importing transportable tablespace metadata into your DB instance](rdsadmin-transport-util-import-xtts-metadata.md)

- [Listing orphaned files after a tablespace import](rdsadmin-transport-util-list-xtts-orphan-files.md)

- [Deleting orphaned data files after a tablespace import](rdsadmin-transport-util-cleanup-incomplete-xtts-import.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Other
tasks

Importing transported tablespaces

All content copied from https://docs.aws.amazon.com/.
