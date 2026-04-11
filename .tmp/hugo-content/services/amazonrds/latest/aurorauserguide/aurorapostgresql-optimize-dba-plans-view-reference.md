# Reference for the apg\_plan\_mgmt.dba\_plans view for Aurora PostgreSQL-Compatible Edition

The columns of plan information in the `apg_plan_mgmt.dba_plans` view
include the following.

dba\_plans columnDescription`cardinality_error`

A measure of the error between the estimated cardinality versus
the actual cardinality. _Cardinality_ is the number of table rows that the
plan is to process. If the cardinality error is large, then it
increases the likelihood that the plan isn't optimal. This
column is populated by the [apg\_plan\_mgmt.evolve\_plan\_baselines](aurorapostgresql-optimize-functions.md#AuroraPostgreSQL.Optimize.Functions.evolve_plan_baselines) function.

`compatibility_level`

This parameter shows when a query plan was last validated. In
Aurora PostgreSQL versions 12.19, 13.15, 14.12, 15.7, 16.3, and later,
it displays the Aurora version number. For earlier versions, it
displays a feature specific version number.

###### Note

Keep this parameter value at its default setting.
Aurora PostgreSQL automatically sets and updates this value.

`created_by`The authenticated user ( `session_user`) who created the
plan.`enabled`

An indicator of whether the plan is enabled or disabled. All plans
are enabled by default. You can disable plans to prevent them from
being used by the optimizer. To modify this value, use the [apg\_plan\_mgmt.set\_plan\_enabled](aurorapostgresql-optimize-functions.md#AuroraPostgreSQL.Optimize.Functions.set_plan_enabled) function.

`environment_variables`

The PostgreSQL Grand Unified Configuration (GUC) parameters and
values that the optimizer has overridden at the time the plan was
captured.

`estimated_startup_cost`The estimated optimizer setup cost before the optimizer delivers rows
of a table.`estimated_total_cost`The estimated optimizer cost to deliver the final table row.`execution_time_benefit_ms`The execution time benefit in milliseconds of enabling the plan. This
column is populated by the [apg\_plan\_mgmt.evolve\_plan\_baselines](aurorapostgresql-optimize-functions.md#AuroraPostgreSQL.Optimize.Functions.evolve_plan_baselines) function. `execution_time_ms`The estimated time in milliseconds that the plan would run. This
column is populated by the [apg\_plan\_mgmt.evolve\_plan\_baselines](aurorapostgresql-optimize-functions.md#AuroraPostgreSQL.Optimize.Functions.evolve_plan_baselines) function. `has_side_effects`A value that indicates that the SQL statement is a data manipulation
language (DML) statement or a SELECT statement that contains a VOLATILE
function. `last_used`This value is updated to the current date whenever the plan is either
executed or when the plan is the query optimizer's minimum-cost
plan. This value is stored in shared memory and periodically flushed to
disk. To get the most up-to-date value, read the date from shared memory
by calling the function `apg_plan_mgmt.plan_last_used(sql_hash,
                                plan_hash)` instead of reading the `last_used`
value. For additional information, see the [apg\_plan\_mgmt.plan\_retention\_period](aurorapostgresql-optimize-parameters.md#AuroraPostgreSQL.Optimize.Parameters.plan_retention_period) parameter. `last_validated`The most recent date and time when it was verified that the plan
could be recreated by either the [apg\_plan\_mgmt.validate\_plans](aurorapostgresql-optimize-functions.md#AuroraPostgreSQL.Optimize.Functions.validate_plans)
function or the [apg\_plan\_mgmt.evolve\_plan\_baselines](aurorapostgresql-optimize-functions.md#AuroraPostgreSQL.Optimize.Functions.evolve_plan_baselines) function.`last_verified`The most recent date and time when a plan was verified to be the
best-performing plan for the specified parameters by the [apg\_plan\_mgmt.evolve\_plan\_baselines](aurorapostgresql-optimize-functions.md#AuroraPostgreSQL.Optimize.Functions.evolve_plan_baselines) function. `origin`

How the plan was captured with the [apg\_plan\_mgmt.capture\_plan\_baselines](aurorapostgresql-optimize-parameters.md#AuroraPostgreSQL.Optimize.Parameters.capture_plan_baselines) parameter. Valid values include the following:

`M` – The plan was captured with manual plan
capture.

`A` – The plan was captured with automatic plan
capture.

`param_list`

The parameter values that were passed to the statement if this is
a prepared statement.

`plan_created`The date and time the plan that was created.`plan_hash`The plan identifier. The combination of `plan_hash` and
`sql_hash` uniquely identifies a specific plan.`plan_outline`A representation of the plan that is used to recreate the actual
execution plan, and that is database-independent. Operators in the tree
correspond to operators that appear in the EXPLAIN output.`planning_time_ms`

The actual time to run the planner, in milliseconds. This column
is populated by the [apg\_plan\_mgmt.evolve\_plan\_baselines](aurorapostgresql-optimize-functions.md#AuroraPostgreSQL.Optimize.Functions.evolve_plan_baselines) function.

`queryId`A statement hash, as calculated by the
`pg_stat_statements` extension. This isn't a stable
or database-independent identifier because it depends on object
identifiers (OIDs). The value will be `0` if
`compute_query_id` is `off` when capturing the
query plan.`sql_hash`A hash value of the SQL statement text, normalized with literals
removed.`sql_text`The full text of the SQL statement.`status`

A plan's status, which determines how the optimizer uses a
plan. Valid values include the following.

- `Approved` – A usable plan that the
optimizer can choose to run. The optimizer runs the
least-cost plan from a managed statement's set of
approved plans (baseline). To reset a plan to approved, use
the [apg\_plan\_mgmt.evolve\_plan\_baselines](aurorapostgresql-optimize-functions.md#AuroraPostgreSQL.Optimize.Functions.evolve_plan_baselines) function.

- `Unapproved` – A captured plan that you
have not verified for use. For more information, see [Evaluating plan performance](aurorapostgresql-optimize-maintenance.md#AuroraPostgreSQL.Optimize.Maintenance.EvaluatingPerformance).

- `Rejected` – A plan that the optimizer
won't use. For more information, see [Rejecting or disabling slower plans](aurorapostgresql-optimize-maintenance.md#AuroraPostgreSQL.Optimize.Maintenance.EvaluatingPerformance.Rejecting).

- `Preferred` – A plan that you have
determined is a preferred plan to use for a managed
statement.

If the optimizer's minimum-cost plan isn't an
approved or preferred plan, you can reduce plan enforcement
overhead. To do so, make a subset of the approved plans
`Preferred`. When the optimizer's
minimum cost isn't an `Approved` plan, a
`Preferred` plan is chosen before an
`Approved` plan.

To reset a plan to `Preferred`, use the [apg\_plan\_mgmt.set\_plan\_status](aurorapostgresql-optimize-functions.md#AuroraPostgreSQL.Optimize.Functions.set_plan_status) function.

`stmt_name`The name of the SQL statement within a PREPARE statement. This value
is an empty string for an unnamed prepared statement. This value is NULL
for a nonprepared statement.`total_time_benefit_ms`

The total time benefit in milliseconds of enabling this plan. This
value considers both planning time and execution time.

If this value is negative, there is a disadvantage to enabling
this plan. This column is populated by the [apg\_plan\_mgmt.evolve\_plan\_baselines](aurorapostgresql-optimize-functions.md#AuroraPostgreSQL.Optimize.Functions.evolve_plan_baselines) function.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Function reference

Advanced features in Query Plan Management

All content copied from https://docs.aws.amazon.com/.
