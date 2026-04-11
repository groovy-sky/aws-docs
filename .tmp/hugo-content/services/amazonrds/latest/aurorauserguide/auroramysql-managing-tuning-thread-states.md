# Tuning Aurora MySQL with thread states

The following table summarizes the most common general thread states for Aurora MySQL.

General thread stateDescription

[creating sort index](ams-states-sort-index.md)

This thread state indicates that a thread is
processing a `SELECT` statement that requires the use of an internal temporary
table to sort the data.

[sending data](ams-states-sending-data.md)

This thread state indicates that a thread is reading and filtering rows for a query to
determine the correct result set.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

synch/mutex/innodb/temp\_pool\_manager\_mutex

creating sort index

All content copied from https://docs.aws.amazon.com/.
