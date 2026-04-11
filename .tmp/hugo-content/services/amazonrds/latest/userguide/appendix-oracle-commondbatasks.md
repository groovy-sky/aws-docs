# Administering your RDS for Oracle DB instance

Following are the common management tasks that you perform with an RDS for Oracle DB instance. Some
tasks are the same for all RDS DB instances. Other tasks are specific to RDS for Oracle.

The following tasks are common to all RDS databases, but Oracle Database has special
considerations. For example, you connect to an Oracle database using the Oracle clients
SQL\*Plus and SQL Developer.

Task areaRelevant documentation

**Instance classes, storage, and**
**PIOPS**

If you are creating a production instance, learn how instance classes,
storage types, and Provisioned IOPS work in Amazon RDS.

[RDS for Oracle DB instance classes](oracle-concepts-instanceclasses.md)

[Amazon RDS storage types](chap-storage.md#Concepts.Storage)

**Multi-AZ deployments**

A production DB instance should use Multi-AZ deployments. Multi-AZ
deployments provide increased availability, data durability, and fault
tolerance for DB instances.

[Configuring and managing a Multi-AZ deployment for Amazon RDS](concepts-multiaz.md)

**Amazon VPC**

If your AWS account has a default virtual private cloud (VPC), then your DB instance is
automatically created inside the default VPC. If your account
doesn't have a default VPC, and you want the DB instance in a VPC,
create the VPC and subnet groups before you create the instance.

[Working with a DB instance in a VPC](user-vpc-workingwithrdsinstanceinavpc.md)

**Security groups**

By default, DB instances use a firewall that prevents access. Make
sure that you create a security group with the correct IP addresses and
network configuration to access the DB instance.

[Controlling access with security groups](overview-rdssecuritygroups.md)

**Parameter groups**

If your DB instance is going to require specific database parameters,
create a parameter group before you create the DB instance.

[Parameter groups for Amazon RDS](user-workingwithparamgroups.md)

**Option groups**

If your DB instance requires specific database options, create an
option group before you create the DB instance.

[Adding options to Oracle DB instances](appendix-oracle-options.md)

**Connecting to your DB instance**

After creating a security group and associating it to a DB instance,
you can connect to the DB instance using any standard SQL client
application such as Oracle SQL\*Plus.

[Connecting to your Oracle DB instance](user-connecttooracleinstance.md)

**Backup and restore**

You can configure your DB instance to take automated backups, or take
manual snapshots, and then restore instances from the backups or
snapshots.

[Backing up, restoring, and exporting data](chap-commontasks-backuprestore.md)

**Monitoring**

You can monitor an Oracle DB instance by using CloudWatch Amazon RDS metrics,
events, and enhanced monitoring.

[Viewing metrics in the Amazon RDS console](user-monitoring.md)

[Viewing Amazon RDS events](user-listevents.md)

**Log files**

You can access the log files for your Oracle DB instance.

[Monitoring Amazon RDS log files](user-logaccess.md)

Following, you can find a description for Amazon RDS–specific implementations of common
DBA tasks for RDS Oracle. To deliver a managed service experience, Amazon RDS doesn't provide
shell access to DB instances. Also, RDS restricts access to certain system procedures and
tables that require advanced privileges. In many of the tasks, you run the
`rdsadmin` package, which is an Amazon RDS–specific tool that enables you to
administer your database.

The following are common DBA tasks for DB instances running Oracle:

- [System\
tasks](appendix-oracle-commondbatasks-system.md)

[Disconnecting a session](appendix-oracle-commondbatasks-disconnectingsession.md)

Amazon RDS method:
`rdsadmin.rdsadmin_util.disconnect`

Oracle method: `alter system disconnect
                                      session`

[Terminating a session](appendix-oracle-commondbatasks-killingsession.md)

Amazon RDS method: `rdsadmin.rdsadmin_util.kill`

Oracle method: `alter system kill session`

[Canceling a SQL statement in a session](appendix-oracle-commondbatasks-cancellingsql.md)

Amazon RDS method:
`rdsadmin.rdsadmin_util.cancel`

Oracle method: `alter system cancel sql`

[Enabling and disabling restricted sessions](appendix-oracle-commondbatasks-restrictedsession.md)

Amazon RDS method:
`rdsadmin.rdsadmin_util.restricted_session`

Oracle method: `alter system enable restricted
                                          session`

[Flushing the shared pool](appendix-oracle-commondbatasks-flushingsharedpool.md)

Amazon RDS method:
`rdsadmin.rdsadmin_util.flush_shared_pool`

Oracle method: `alter system flush
                                      shared_pool`

[Flushing the buffer cache](appendix-oracle-commondbatasks-flushingsharedpool.md#Appendix.Oracle.CommonDBATasks.FlushingBufferCache)

Amazon RDS method:
`rdsadmin.rdsadmin_util.flush_buffer_cache`

Oracle method: `alter system flush
                                      buffer_cache`

[Granting SELECT or EXECUTE privileges to SYS objects](appendix-oracle-commondbatasks-transferprivileges.md)

Amazon RDS method:
`rdsadmin.rdsadmin_util.grant_sys_object`

Oracle method: `grant`

[Revoking SELECT or EXECUTE privileges on SYS objects](appendix-oracle-commondbatasks-revokeprivileges.md)

Amazon RDS method:
`rdsadmin.rdsadmin_util.revoke_sys_object`

Oracle method: `revoke`

[Managing RDS\_X$ views for Oracle DB instances](appendix-oracle-commondbatasks-x-dollar.md)

Amazon RDS method:
`rdsadmin.rdsadmin_util.create_sys_x$_view`

Oracle method: `CREATE VIEW`

[Granting privileges to non-master users](appendix-oracle-commondbatasks-permissionsnonmasters.md)

Amazon RDS method: `grant`

[Creating custom functions to verify passwords](appendix-oracle-commondbatasks-custompassword.md)

Amazon RDS method:
`rdsadmin.rdsadmin_password_verify.create_verify_function`

Amazon RDS method:
`rdsadmin.rdsadmin_password_verify.create_passthrough_verify_fcn`

[Setting up a custom DNS server](appendix-oracle-commondbatasks-system.md#Appendix.Oracle.CommonDBATasks.CustomDNS)

—

[Listing allowed system diagnostic events](appendix-oracle-commondbatasks-systemevents.md#Appendix.Oracle.CommonDBATasks.SystemEvents.listing)

Amazon RDS method:
`rdsadmin.rdsadmin_util.list_allowed_system_events`

Oracle method: —

[Setting system diagnostic events](appendix-oracle-commondbatasks-systemevents.md#Appendix.Oracle.CommonDBATasks.SystemEvents.setting)

Amazon RDS method:
`rdsadmin.rdsadmin_util.set_allowed_system_events`

Oracle method: `ALTER SYSTEM SET EVENTS
                                              'set_event_clause'`

[Listing system diagnostic events that are set](appendix-oracle-commondbatasks-systemevents.md#Appendix.Oracle.CommonDBATasks.SystemEvents.listing-set)

Amazon RDS method:
`rdsadmin.rdsadmin_util.list_set_system_events`

Oracle method: `ALTER SESSION SET EVENTS 'IMMEDIATE
                                          EVENTDUMP(SYSTEM)'`

[Unsetting system diagnostic events](appendix-oracle-commondbatasks-systemevents.md#Appendix.Oracle.CommonDBATasks.SystemEvents.unsetting)

Amazon RDS method:
`rdsadmin.rdsadmin_util.unset_system_event`

Oracle method: `ALTER SYSTEM SET EVENTS
                                              'unset_event_clause'`

- [Database\
tasks](appendix-oracle-commondbatasks-database.md)

[Changing the global name of a database](appendix-oracle-commondbatasks-renamingglobalname.md)

Amazon RDS method:
`rdsadmin.rdsadmin_util.rename_global_name`

Oracle method: `alter database rename`

[Creating and sizing tablespaces in RDS for Oracle](appendix-oracle-commondbatasks-tablespacesanddatafiles.md#Appendix.Oracle.CommonDBATasks.CreatingTablespacesAndDatafiles)

Amazon RDS method: `create tablespace`

Oracle method: `alter database`

[Setting the default tablespace in RDS for Oracle](appendix-oracle-commondbatasks-tablespacesanddatafiles.md#Appendix.Oracle.CommonDBATasks.SettingDefaultTablespace)

Amazon RDS method:
`rdsadmin.rdsadmin_util.alter_default_tablespace`

Oracle method: `alter database default tablespace
                                      `

[Setting the default temporary tablespace in RDS for Oracle](appendix-oracle-commondbatasks-tablespacesanddatafiles.md#Appendix.Oracle.CommonDBATasks.SettingDefTempTablespace)

Amazon RDS method:
`rdsadmin.rdsadmin_util.alter_default_temp_tablespace`

Oracle method: `alter database default temporary
                                          tablespace `

[Creating a temporary tablespace on the instance store](appendix-oracle-commondbatasks-tablespacesanddatafiles.md#Appendix.Oracle.CommonDBATasks.creating-tts-instance-store)

Amazon RDS method: `rdsadmin.rdsadmin_util.create_inst_store_tmp_tblspace`

Oracle method: `create temporary tablespace`

[Checkpointing a database](appendix-oracle-commondbatasks-checkpointingdatabase.md)

Amazon RDS method:
`rdsadmin.rdsadmin_util.checkpoint`

Oracle method: `alter system checkpoint`

[Setting distributed recovery](appendix-oracle-commondbatasks-settingdistributedrecovery.md)

Amazon RDS method:
`rdsadmin.rdsadmin_util.enable_distr_recovery`

Oracle method: `alter system enable distributed
                                          recovery`

[Setting the database time zone](appendix-oracle-commondbatasks-timezonesupport.md)

Amazon RDS method:
`rdsadmin.rdsadmin_util.alter_db_time_zone`

Oracle method: `alter database set
                                      time_zone`

[Working with external tables in RDS for Oracle](appendix-oracle-commondbatasks-external-tables.md)

—

[Generating performance reports with Automatic Workload Repository (AWR)](appendix-oracle-commondbatasks-awr.md)

Amazon RDS method: `rdsadmin.rdsadmin_diagnostic_util`
procedures

Oracle method: `dbms_workload_repository`
package

[Adjusting database links for use with DB instances in a VPC](appendix-oracle-commondbatasks-dblinks.md)

—

[Setting the default edition for a DB instance](appendix-oracle-commondbatasks-defaultedition.md)

Amazon RDS method:
`rdsadmin.rdsadmin_util.alter_default_edition`

Oracle method: `alter database default
                                      edition`

[Enabling auditing for the SYS.AUD$ table](appendix-oracle-commondbatasks-enablingauditing.md)

Amazon RDS method:
`rdsadmin.rdsadmin_master_util.audit_all_sys_aud_table`

Oracle method: `audit`

[Disabling auditing for the SYS.AUD$ table](appendix-oracle-commondbatasks-disablingauditing.md)

Amazon RDS method:
`rdsadmin.rdsadmin_master_util.noaudit_all_sys_aud_table`

Oracle method: `noaudit`

[Cleaning up interrupted online index builds](appendix-oracle-commondbatasks-cleanupindex.md)

Amazon RDS method:
`rdsadmin.rdsadmin_dbms_repair.online_index_clean`

Oracle method:
`dbms_repair.online_index_clean`

[Skipping corrupt blocks](appendix-oracle-commondbatasks-skippingcorruptblocks.md)

Amazon RDS method: Several
`rdsadmin.rdsadmin_dbms_repair` procedures

Oracle method: `dbms_repair` package

[Resizing tablespaces, data files, and tempfiles in RDS for Oracle](appendix-oracle-commondbatasks-resizetempspacereadreplica.md)

Amazon RDS method:
`rdsadmin.rdsadmin_util.resize_temp_tablespace`,
`rdsadmin.rdsadmin_util.resize_tempfile`, or
`rdsadmin.rdsadmin_util.autoextend_tempfile`
procedures

`rdsadmin.rdsadmin_util.resize_datafile` or
`rdsadmin.rdsadmin_util.autoextend_datafile`
procedure

Oracle method: —

[Purging the recycle bin](#Appendix.Oracle.CommonDBATasks.PurgeRecycleBin)

Amazon RDS method: `EXEC
                                          rdsadmin.rdsadmin_util.purge_dba_recyclebin`

Oracle method: `purge dba_recyclebin`

[Setting the default displayed values for full redaction](appendix-oracle-commondbatasks-fullredaction.md)

Amazon RDS method: `EXEC
                                          rdsadmin.rdsadmin_util.dbms_redact_upd_full_rdct_val`

Oracle method: `exec
                                          dbms_redact.UPDATE_FULL_REDACTION_VALUES`

- [Log tasks](appendix-oracle-commondbatasks-log.md)

[Setting force logging](appendix-oracle-commondbatasks-log.md#Appendix.Oracle.CommonDBATasks.SettingForceLogging)

Amazon RDS method:
`rdsadmin.rdsadmin_util.force_logging`

Oracle method: `alter database force
                                      logging`

[Setting supplemental logging](appendix-oracle-commondbatasks-log.md#Appendix.Oracle.CommonDBATasks.AddingSupplementalLogging)

Amazon RDS method:
`rdsadmin.rdsadmin_util.alter_supplemental_logging`

Oracle method: `alter database add supplemental
                                          log`

[Switching online log files](appendix-oracle-commondbatasks-log.md#Appendix.Oracle.CommonDBATasks.SwitchingLogfiles)

Amazon RDS method:
`rdsadmin.rdsadmin_util.switch_logfile`

Oracle method: `alter system switch logfile`

[Adding online redo logs](appendix-oracle-commondbatasks-log.md#Appendix.Oracle.CommonDBATasks.RedoLogs)

Amazon RDS method:
`rdsadmin.rdsadmin_util.add_logfile`

[Dropping online redo logs](appendix-oracle-commondbatasks-log.md#Appendix.Oracle.CommonDBATasks.DroppingRedoLogs)

Amazon RDS method:
`rdsadmin.rdsadmin_util.drop_logfile`

[Resizing online redo logs](appendix-oracle-commondbatasks-resizingredologs.md)

—

[Retaining archived redo logs](appendix-oracle-commondbatasks-retainredologs.md)

Amazon RDS method:
`rdsadmin.rdsadmin_util.set_configuration`

[Downloading archived redo logs from Amazon S3](appendix-oracle-commondbatasks-download-redo-logs.md)

Amazon RDS method:
`rdsadmin.rdsadmin_archive_log_download.download_log_with_seqnum`

Amazon RDS method:
`rdsadmin.rdsadmin_archive_log_download.download_logs_in_seqnum_range`

[Accessing online and archived redo logs](appendix-oracle-commondbatasks-log-download.md)

Amazon RDS method:
`rdsadmin.rdsadmin_master_util.create_archivelog_dir`

Amazon RDS method:
`rdsadmin.rdsadmin_master_util.create_onlinelog_dir`

- [RMAN\
tasks](appendix-oracle-commondbatasks-rman.md)

[Validating database files in RDS for Oracle](appendix-oracle-commondbatasks-validatedbfiles.md)

Amazon RDS method:
`rdsadmin_rman_util.procedure`

Oracle method: `RMAN VALIDATE`

[Enabling and disabling block change tracking](appendix-oracle-commondbatasks-blockchangetracking.md)

Amazon RDS method:
`rdsadmin_rman_util.procedure`

Oracle method: `ALTER DATABASE`

[Crosschecking archived redo logs](appendix-oracle-commondbatasks-crosscheck.md)

Amazon RDS method:
`rdsadmin_rman_util.crosscheck_archivelog`

Oracle method: `RMAN BACKUP`

[Backing up archived redo log files](appendix-oracle-commondbatasks-backuparchivedlogs.md)

Amazon RDS method:
`rdsadmin_rman_util.procedure`

Oracle method: `RMAN BACKUP`

[Performing a full database backup](appendix-oracle-commondbatasks-backupdatabasefull.md)

Amazon RDS method:
`rdsadmin_rman_util.backup_database_full`

Oracle method: `RMAN BACKUP`

[Performing an incremental database backup](appendix-oracle-commondbatasks-backupdatabaseincremental.md)

Amazon RDS method:
`rdsadmin_rman_util.backup_database_incremental`

Oracle method: `RMAN BACKUP`

[Backing up a tablespace](appendix-oracle-commondbatasks-backuptablespace.md)

Amazon RDS method:
`rdsadmin_rman_util.backup_database_tablespace`

Oracle method: `RMAN BACKUP`

- [Oracle Scheduler\
tasks](appendix-oracle-commondbatasks-scheduler.md)

[Modifying DBMS\_SCHEDULER jobs](appendix-oracle-commondbatasks-scheduler.md#Appendix.Oracle.CommonDBATasks.ModifyScheduler)

Amazon RDS method: `dbms_scheduler.set_attribute`

Oracle method:
`dbms_scheduler.set_attribute`

[Modifying AutoTask maintenance windows](appendix-oracle-commondbatasks-scheduler.md#Appendix.Oracle.CommonDBATasks.Scheduler.maintenance-windows)

Amazon RDS method: `dbms_scheduler.set_attribute`

Oracle method:
`dbms_scheduler.set_attribute`

[Setting the time zone for Oracle Scheduler jobs](appendix-oracle-commondbatasks-scheduler.md#Appendix.Oracle.CommonDBATasks.Scheduler.TimeZone)

Amazon RDS method:
`dbms_scheduler.set_scheduler_attribute`

Oracle method:
`dbms_scheduler.set_scheduler_attribute`

[Turning off Oracle Scheduler jobs owned by SYS](appendix-oracle-commondbatasks-scheduler.md#Appendix.Oracle.CommonDBATasks.Scheduler.Disabling)

Amazon RDS method:
`rdsadmin.rdsadmin_dbms_scheduler.disable`

Oracle method: `dbms_scheduler.disable`

[Turning on Oracle Scheduler jobs owned by SYS](appendix-oracle-commondbatasks-scheduler.md#Appendix.Oracle.CommonDBATasks.Scheduler.Enabling)

Amazon RDS method:
`rdsadmin.rdsadmin_dbms_scheduler.enable`

Oracle method: `dbms_scheduler.enable`

[Modifying the Oracle Scheduler repeat interval for jobs of CALENDAR type](appendix-oracle-commondbatasks-scheduler.md#Appendix.Oracle.CommonDBATasks.Scheduler.Modifying_Calendar)

Amazon RDS method:
`rdsadmin.rdsadmin_dbms_scheduler.set_attribute`

Oracle method:
`dbms_scheduler.set_attribute`

[Modifying the Oracle Scheduler repeat interval for jobs of NAMED type](appendix-oracle-commondbatasks-scheduler.md#Appendix.Oracle.CommonDBATasks.Scheduler.Modifying_Named)

Amazon RDS method:
`rdsadmin.rdsadmin_dbms_scheduler.set_attribute`

Oracle method:
`dbms_scheduler.set_attribute`

[Turning off autocommit for Oracle Scheduler job creation](appendix-oracle-commondbatasks-scheduler.md#Appendix.Oracle.CommonDBATasks.Scheduler.autocommit)

Amazon RDS method:
`rdsadmin.rdsadmin_dbms_scheduler.set_no_commit_flag`

Oracle method:
`dbms_isched.set_no_commit_flag`

- [Diagnosing\
problems](appendix-oracle-commondbatasks-diagnostics.md)

[Listing incidents](appendix-oracle-commondbatasks-diagnostics.md#Appendix.Oracle.CommonDBATasks.Incidents)

Amazon RDS method:
`rdsadmin.rdsadmin_adrci_util.list_adrci_incidents`

Oracle method: ADRCI command `show incident`

[Listing problems](appendix-oracle-commondbatasks-diagnostics.md#Appendix.Oracle.CommonDBATasks.Problems)

Amazon RDS method:
`rdsadmin.rdsadmin_adrci_util.list_adrci_problem`

Oracle method: ADRCI command `show problem`

[Creating incident packages](appendix-oracle-commondbatasks-diagnostics.md#Appendix.Oracle.CommonDBATasks.IncPackages)

Amazon RDS method:
`rdsadmin.rdsadmin_adrci_util.create_adrci_package`

Oracle method: ADRCI command `ips create
                                      package`

[Showing trace files](appendix-oracle-commondbatasks-diagnostics.md#Appendix.Oracle.CommonDBATasks.ShowTrace)

Amazon RDS method:
`rdsadmin.rdsadmin_adrci_util.show_adrci_tracefile`

Oracle method: ADRCI command `show
                                      tracefile`

- [Other\
tasks](appendix-oracle-commondbatasks-misc.md)

[Creating and dropping directories in the main data storage space](appendix-oracle-commondbatasks-misc.md#Appendix.Oracle.CommonDBATasks.NewDirectories)

Amazon RDS method:
`rdsadmin.rdsadmin_util.create_directory`

Oracle method: `CREATE DIRECTORY`

Amazon RDS method:
`rdsadmin.rdsadmin_util.drop_directory`

Oracle method: `DROP DIRECTORY`

[Listing files in a DB instance directory](appendix-oracle-commondbatasks-misc.md#Appendix.Oracle.CommonDBATasks.ListDirectories)

Amazon RDS method:
`rdsadmin.rds_file_util.listdir`

Oracle method: —

[Reading files in a DB instance directory](appendix-oracle-commondbatasks-misc.md#Appendix.Oracle.CommonDBATasks.ReadingFiles)

Amazon RDS method:
`rdsadmin.rds_file_util.read_text_file`

Oracle method: —

[Accessing Opatch files](appendix-oracle-commondbatasks-misc.md#Appendix.Oracle.CommonDBATasks.accessing-opatch-files)

Amazon RDS method:
`rdsadmin.rds_file_util.read_text_file` or
`rdsadmin.tracefile_listing`

Oracle method: `opatch`

[Setting parameters for advisor tasks](appendix-oracle-commondbatasks-misc.md#Appendix.Oracle.CommonDBATasks.setting-task-parameters)

Amazon RDS method:
`rdsadmin.rdsadmin_util.advisor_task_set_parameter`

Oracle method: Various stored package procedures

[Disabling AUTO\_STATS\_ADVISOR\_TASK](appendix-oracle-commondbatasks-misc.md#Appendix.Oracle.CommonDBATasks.dropping-advisor-task)

Amazon RDS method:
`rdsadmin.rdsadmin_util.advisor_task_drop`

Oracle method: —

[Re-enabling AUTO\_STATS\_ADVISOR\_TASK](appendix-oracle-commondbatasks-misc.md#Appendix.Oracle.CommonDBATasks.recreating-advisor-task)

Amazon RDS method:
`rdsadmin.rdsadmin_util.dbms_stats_init`

Oracle method: —

You can also use Amazon RDS procedures for Amazon S3 integration with Oracle and for running OEM
Management Agent database tasks. For more information, see [Amazon S3 integration](oracle-s3-integration.md) and [Performing database tasks with the Management Agent](oracle-options-oemagent.md#Oracle.Options.OEMAgent.DBTasks).

## Purging the recycle bin

When you drop a table, your Oracle database doesn't immediately remove its
storage space. The database renames the table and places it and any associated
objects in a recycle bin. Purging the recycle bin removes these items and releases
their storage space.

To purge the entire recycle bin, use the Amazon RDS procedure
`rdsadmin.rdsadmin_util.purge_dba_recyclebin`. However, this
procedure can't purge the recycle bin of `SYS` and
`RDSADMIN` objects. If you need to purge these objects, contact AWS
Support.

The following example purges the entire recycle bin.

```sql

EXEC rdsadmin.rdsadmin_util.purge_dba_recyclebin;
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Upgrading your CDB

System
tasks

All content copied from https://docs.aws.amazon.com/.
