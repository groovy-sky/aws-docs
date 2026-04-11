# Setting the database time zone

You can set the time zone of your Amazon RDS Oracle database in the following ways:

- The `Timezone` option

The `Timezone` option changes the time zone at the host level
and affects all date columns and values such as `SYSDATE`. For
more information, see [Oracle time zone](appendix-oracle-options-timezone.md).

- The Amazon RDS procedure
`rdsadmin.rdsadmin_util.alter_db_time_zone`

The `alter_db_time_zone` procedure changes the time zone for
only certain data types, and doesn't change `SYSDATE`. There are
additional restrictions on setting the time zone listed in the [Oracle documentation](http://docs.oracle.com/cd/B19306_01/server.102/b14225/ch4datetime.htm).

###### Note

You can also set the default time zone for Oracle Scheduler. For more
information, see [Setting the time zone for Oracle Scheduler jobs](appendix-oracle-commondbatasks-scheduler.md#Appendix.Oracle.CommonDBATasks.Scheduler.TimeZone).

The `alter_db_time_zone` procedure has the following parameters.

Parameter nameData typeDefaultRequiredDescription

`p_new_tz`

varchar2

—

Yes

The new time zone as a named region or an absolute offset from
Coordinated Universal Time (UTC). Valid offsets range from
 -12:00 to +14:00.

The following example changes the time zone to UTC plus three hours.

```sql

EXEC rdsadmin.rdsadmin_util.alter_db_time_zone(p_new_tz => '+3:00');
```

The following example changes the time zone to the Africa/Algiers time zone.

```sql

EXEC rdsadmin.rdsadmin_util.alter_db_time_zone(p_new_tz => 'Africa/Algiers');
```

After you alter the time zone by using the `alter_db_time_zone`
procedure, reboot your DB instance for the change to take effect. For more
information, see [Rebooting a DB instance](user-rebootinstance.md). For information about upgrading time
zones, see [Time zone considerations](user-upgradedbinstance-oracle-ogpg.md#USER_UpgradeDBInstance.Oracle.OGPG.DST).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Setting distributed recovery

Working with
AWR

All content copied from https://docs.aws.amazon.com/.
