# Schedule expression reference

Use these reference tables to construct schedule expressions for your scheduled
queries. All times are in UTC.

**Cron expression syntax**

Format: `cron(minute hour day-of-month month day-of-week year)`

Use CaseCron ExpressionDescriptionUse When**Daily Schedules**`cron(0 9 * * ? *)`Every day at 9:00 AM UTCDaily reports`cron(0 */6 * * ? *)`Every 6 hours (00:00, 06:00, 12:00, 18:00 UTC)Frequent monitoring`cron(30 2 * * ? *)`Every day at 2:30 AM UTCOff-peak analysis**Business Hours**`cron(0 9-17 ? * MON-FRI *)`Every hour from 9 AM to 5 PM, Monday to Friday UTCBusiness monitoring`cron(0 18 ? * MON-FRI *)`Weekdays at 6:00 PM UTCEnd of business day`cron(0 8,12,17 ? * MON-FRI *)`8 AM, noon, and 5 PM on weekdays UTCKey business times**Weekly Schedules**`cron(0 12 ? * SUN *)`Every Sunday at noon UTCWeekly summaries`cron(0 9 ? * MON *)`Every Monday at 9:00 AM UTCWeek start reports`cron(0 23 ? * FRI *)`Every Friday at 11:00 PM UTCWeek end cleanup**Monthly Schedules**`cron(0 0 1 * ? *)`First day of every month at midnight UTCMonthly reports`cron(0 9 L * ? *)`Last day of every month at 9:00 AM UTCMonth end processing`cron(0 10 1 1,4,7,10 ? *)`First day of each quarter at 10:00 AM UTCQuarterly analysis**High Frequency**`cron(*/15 * * * ? *)`Every 15 minutesReal-time monitoring`cron(0,30 * * * ? *)`Every 30 minutes (at :00 and :30)Frequent checks`cron(0 */2 * * ? *)`Every 2 hoursRegular intervals**Special Cases**`cron(30 8 1 1 ? *)`January 1st at 8:30 AM UTCAnnual reports`cron(0 6 * * SAT,SUN *)`Weekends at 6:00 AM UTCWeekend processing`cron(0 0 ? * MON#1 *)`First Monday of every month at midnight UTCMonthly planning

**Cron expression field reference**

FieldValuesWildcardsExamplesMinute (1st)0-59`* , - /``0` (top of hour), `*/15` (every 15 min), `0,30` (twice hourly)Hour (2nd)0-23`* , - /``9` (9 AM), `*/2` (every 2 hrs), `9-17` (business hours)Day-of-month (3rd)1-31, L, W`* , - / ?``1` (1st day), `L` (last day), `?` (when using day-of-week)Month (4th)1-12 or JAN-DEC`* , - /``1` (January), `JAN`, `1,4,7,10` (quarterly)Day-of-week (5th)1-7 or SUN-SAT`* , - / ? # L``MON-FRI` (weekdays), `SUN`, `MON#1` (1st Monday)Year (6th)1970-2199`* , - /``*` (every year), `2024` (specific year), `2024-2026` (range)

**Wildcard characters and special expressions**

**`*` (asterisk)**

Matches all values in the field. Example: `*` in the hour field means every hour.

**`?` (question mark)**

No specific value. Use in day-of-month or day-of-week when the other is specified. Example: Use `?` in day-of-month when specifying `MON-FRI` in day-of-week.

**`-` (dash)**

Range of values. Example: `MON-FRI` (Monday through Friday), `9-17` (9 AM through 5 PM).

**`,` (comma)**

Multiple specific values. Example: `MON,WED,FRI` (Monday, Wednesday, Friday), `8,12,17` (8 AM, noon, 5 PM).

**`/` (slash)**

Step values or increments. Example: `0/15` in minutes means every 15 minutes starting at minute 0 (0, 15, 30, 45). `*/2` in hours means every 2 hours.

**`L` (last)**

Last day of month or last occurrence of weekday. Example: `L` in day-of-month means last day of month. `FRIL` means last Friday of month.

**`W` (weekday)**

Nearest weekday. Example: `15W` means nearest weekday to the 15th of the month.

**`#` (nth occurrence)**

Nth occurrence of weekday in month. Example: `MON#1` means first Monday of month, `FRI#2` means second Friday of month.

**Common patterns and best practices**

- **For business applications:** Use `MON-FRI` and business hours (e.g., `9-17`) to avoid running queries during weekends or off-hours.

- **For high-frequency monitoring:** Use increments like `*/15` (every 15 minutes) but be mindful of query concurrency limits.

- **For resource efficiency:** Schedule resource-intensive queries during off-peak hours using early morning times like `2-6` UTC.

- **For monthly reports:** Use `L` for last day of month or specific dates like `1` for first day to ensure consistent timing.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Understanding scheduled queries concepts

Best practices
