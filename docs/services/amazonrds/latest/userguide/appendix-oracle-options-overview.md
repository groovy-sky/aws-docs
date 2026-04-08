# Overview of Oracle DB options

To enable options for your Oracle database, add them to an option group, and then associate the option group
with your DB instance. For more information, see [Working with option groups](user-workingwithoptiongroups.md).

###### Topics

- [Summary of Oracle Database options](#Appendix.Oracle.Options.summary)

- [Options supported for different editions](#Appendix.Oracle.Options.editions)

- [Memory requirements for specific options](#Appendix.Oracle.Options.memory)

## Summary of Oracle Database options

You can add the following options for Oracle DB instances.

OptionOption ID

[Amazon S3 integration](oracle-s3-integration.md)

`S3_INTEGRATION`

[Oracle Application Express (APEX)](appendix-oracle-options-apex.md)

`APEX`

`APEX-DEV`

[Oracle Enterprise Manager](oracle-options-oem.md)

`OEM`

`OEM_AGENT`

[Oracle Java virtual machine](oracle-options-java.md)

`JVM`

[Oracle Label Security](oracle-options-ols.md)

`OLS`

[Oracle Locator](oracle-options-locator.md)

`LOCATOR`

[Oracle native network encryption](appendix-oracle-options-networkencryption.md)

`NATIVE_NETWORK_ENCRYPTION`

[Oracle OLAP](oracle-options-olap.md)

`OLAP`

[Oracle Secure Sockets Layer](appendix-oracle-options-ssl.md)

`SSL`

[Oracle Spatial](oracle-options-spatial.md)

`SPATIAL`

[Oracle SQLT](oracle-options-sqlt.md)

`SQLT`

[Oracle Statspack](appendix-oracle-options-statspack.md)

`STATSPACK`

[Oracle time zone](appendix-oracle-options-timezone.md)

`Timezone`

[Oracle time zone file autoupgrade](appendix-oracle-options-timezone-file-autoupgrade.md)

`TIMEZONE_FILE_AUTOUPGRADE`

[Oracle Transparent Data Encryption](appendix-oracle-options-advsecurity.md)

`TDE`

[Oracle UTL\_MAIL](oracle-options-utlmail.md)

`UTL_MAIL`

[Oracle XML DB](appendix-oracle-options-xmldb.md)

`XMLDB`

## Options supported for different editions

RDS for Oracle prevents you from adding options to an edition if they aren't supported. To find out which
RDS options are supported in different Oracle Database editions, use the command `aws rds
                describe-option-group-options`. The following example lists supported options for Oracle Database
19c Enterprise Edition.

```

aws rds describe-option-group-options \
    --engine-name oracle-ee \
    --major-engine-version 19
```

For more information, see [describe-option-group-options](../../../cli/latest/reference/rds/describe-option-group-options.md) in the _AWS CLI Command Reference_.

## Memory requirements for specific options

Some options require additional memory to run on your DB instance. For example, Oracle Enterprise Manager
Database Control uses about 300 MB of RAM. If you enable this option for a small DB instance, you might
encounter performance problems due to memory constraints. You can adjust the Oracle parameters so that the
database requires less RAM. Alternatively, you can scale up to a larger DB instance.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Options for Oracle

Amazon S3 integration

All content copied from https://docs.aws.amazon.com/.
