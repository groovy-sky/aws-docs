# Database load

_Database load (DB load)_ measures the level of session activity in your
database. `DBLoad` is the key metric in Performance Insights, and Performance Insights collects DB load
every second.

###### Topics

- [Active sessions](#USER_PerfInsights.Overview.ActiveSessions.active-sessions)

- [Average active sessions](#USER_PerfInsights.Overview.ActiveSessions.AAS)

- [Average active executions](#USER_PerfInsights.Overview.ActiveSessions.AAE)

- [Dimensions](#USER_PerfInsights.Overview.ActiveSessions.dimensions)

## Active sessions

A _database session_ represents an application's dialogue with a
relational database. An _active session_ is a
connection that has submitted work to the DB engine and is waiting for a response.

A session is active when it's either running on CPU or waiting for a resource to become
available so that it can proceed. For example, an active session might wait for a
page (or block) to be read into memory, and then consume CPU while it reads data from the page.

## Average active sessions

The _average active sessions (AAS)_ is the unit for the
`DBLoad` metric in Performance Insights. It measures how many sessions
are concurrently active on the database.

Every second, Performance Insights samples the number of sessions concurrently running
a query. For each active session, Performance Insights collects the following
data:

- SQL statement

- Session state (running on CPU or waiting)

- Host

- User running the SQL

Performance Insights calculates the AAS by dividing the total number of sessions by
the number of samples for a specific time period. For example, the following table shows
5 consecutive samples of a running query taken at 1-second intervals.

SampleNumber of sessions running queryAASCalculation1222 total sessions / 1 sample2012 total sessions / 2 samples3426 total sessions / 3 samples401.56 total sessions / 4 samples54210 total sessions / 5 samples

In the preceding example, the DB load for the time interval was 2 AAS. This measurement
means that, on average, 2 sessions were active at any given time during the interval
when the 5 samples were taken.

## Average active executions

The _average active executions (AAE)_ per second is related to AAS. To calculate the AAE,
Performance Insights divides the total execution time of a query by the time interval. The following
table shows the AAE calculation for the same query in the preceding table.

Elapsed time (sec)Total execution time (sec)AAECalculation601202120 execution seconds/60 elapsed seconds1201201120 execution seconds/120 elapsed seconds1803802.11380 execution seconds/180 elapsed seconds2403801.58 380 execution seconds/240 elapsed seconds3006002600 execution seconds/300 elapsed seconds

In most cases, the AAS and AAE for a query are approximately the same. However, because the inputs to
the calculations are different data sources, the calculations often vary slightly.

## Dimensions

The `db.load` metric is different from the other time-series metrics because you can break it into
subcomponents called _dimensions_. You can think of dimensions as "slice by" categories for
the different characteristics of the `DBLoad` metric.

When you are diagnosing performance issues, the following dimensions are often the most useful:

###### Topics

- [Wait events](#USER_PerfInsights.Overview.ActiveSessions.waits)

- [Top SQL](#USER_PerfInsights.Overview.ActiveSessions.top-sql)

For a complete list of dimensions for the Aurora engines, see [DB load sliced by dimensions](user-perfinsights-usingdashboard-components.md#USER_PerfInsights.UsingDashboard.Components.AvgActiveSessions.dims).

### Wait events

A _wait event_ causes a SQL statement to wait for a specific event to happen before
it can continue running. Wait events are an important dimension, or category, for DB load because they
indicate where work is impeded.

Every active session is either running on the CPU or waiting. For example, sessions
consume CPU when they search memory for a buffer, perform a calculation, or run
procedural code. When sessions aren't consuming CPU, they might be waiting for a
memory buffer to become free, a data file to be read, or a log to be written to.
The more time that a session waits for resources, the less time it runs on the
CPU.

When you tune a database, you often try to find out the resources that sessions are
waiting for. For example, two or three wait events might account for 90 percent
of DB load. This measure means that, on average, active sessions are spending
most of their time waiting for a small number of resources. If you can find out
the cause of these waits, you can attempt a solution.

Wait events vary by DB engine:

- For a list of the common wait events for Aurora MySQL, see [Aurora MySQL wait events](auroramysql-reference-waitevents.md). To learn how to tune using these wait
events, see [Tuning Aurora MySQL](auroramysql-managing-tuning.md).

- For information about all MySQL wait
events, see [Wait Event Summary Tables](https://dev.mysql.com/doc/refman/8.0/en/performance-schema-wait-summary-tables.html) in the MySQL documentation.

- For a list of common wait events for Aurora PostgreSQL, see [Amazon Aurora PostgreSQL wait events](aurorapostgresql-reference-waitevents.md). To learn how to tune using these
wait events, see [Tuning with wait events for Aurora PostgreSQL](aurorapostgresql-tuning.md).

- For information about all PostgreSQL wait events, see [The Statistics Collector > Wait Event tables](https://www.postgresql.org/docs/current/monitoring-stats.html) in the PostgreSQL documentation.

### Top SQL

Where wait events show bottlenecks, top SQL shows which queries are contributing the most to DB load. For example,
many queries might be currently running on the database, but a single query might consume 99 percent of the
DB load. In this case, the high load might indicate a problem with the query.

By default, the Performance Insights console displays top SQL queries that are contributing to the
database load. The console also shows relevant statistics for each statement. To diagnose performance
problems for a specific statement, you can examine its execution plan.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Overview of Performance Insights

Maximum CPU

All content copied from https://docs.aws.amazon.com/.
