Menu

- [GuzzleHttp](namespace-guzzlehttp.md)
- [Promise](namespace-guzzlehttp-promise.md)

## Utils        in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](class-guzzlehttp-promise-utils-toc.md)

#### Methods  [header link](class-guzzlehttp-promise-utils-toc-methods.md)

[all()](class-guzzlehttp-promise-utils-method-all.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Given an array of promises, return a promise that is fulfilled when all
the items in the array are fulfilled.[any()](class-guzzlehttp-promise-utils-method-any.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Like some(), with 1 as count. However, if the promise fulfills, the
fulfillment value is not an array of 1 but the value directly.[inspect()](class-guzzlehttp-promise-utils-method-inspect.md)
: array<string\|int, mixed> Synchronously waits on a promise to resolve and returns an inspection
state array.[inspectAll()](class-guzzlehttp-promise-utils-method-inspectall.md)
: array<string\|int, mixed> Waits on all of the provided promises, but does not unwrap rejected
promises as thrown exception.[queue()](class-guzzlehttp-promise-utils-method-queue.md)
: [TaskQueueInterface](class-guzzlehttp-promise-taskqueueinterface.md)Get the global task queue used for promise resolution.[settle()](class-guzzlehttp-promise-utils-method-settle.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise that is fulfilled when all of the provided promises have
been fulfilled or rejected.[some()](class-guzzlehttp-promise-utils-method-some.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Initiate a competitive race between multiple promises or values (values
will become immediately fulfilled promises).[task()](class-guzzlehttp-promise-utils-method-task.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Adds a function to run in the task queue when it is next \`run()\` and
returns a promise that is fulfilled or rejected with the result.[unwrap()](class-guzzlehttp-promise-utils-method-unwrap.md)
: array<string\|int, mixed> Waits on all of the provided promises and returns the fulfilled values.

### Methods  [header link](class-guzzlehttp-promise-utils-methods.md)

#### all()  [header link](class-guzzlehttp-promise-utils-method-all.md)

Given an array of promises, return a promise that is fulfilled when all
the items in the array are fulfilled.

`
    public
            static        all(mixed $promises[, bool $recursive = false ]) : PromiseInterface`

The promise's fulfillment value is an array with fulfillment values at
respective positions to the original array. If any promise in the array
rejects, the returned promise is rejected with the rejection reason.

##### Parameters

$promises
: mixed

Promises or values.

$recursive
: bool
= false

If true, resolves new promises that might have been added to the stack during its own resolution.

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### any()  [header link](class-guzzlehttp-promise-utils-method-any.md)

Like some(), with 1 as count. However, if the promise fulfills, the
fulfillment value is not an array of 1 but the value directly.

`
    public
            static        any(mixed $promises) : PromiseInterface`

##### Parameters

$promises
: mixed

Promises or values.

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### inspect()  [header link](class-guzzlehttp-promise-utils-method-inspect.md)

Synchronously waits on a promise to resolve and returns an inspection
state array.

`
    public
            static        inspect(PromiseInterface $promise) : array<string|int, mixed>`

Returns a state associative array containing a "state" key mapping to a
valid promise state. If the state of the promise is "fulfilled", the
array will contain a "value" key mapping to the fulfilled value of the
promise. If the promise is rejected, the array will contain a "reason"
key mapping to the rejection reason of the promise.

##### Parameters

$promise
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

Promise or value.

##### Return values

array<string\|int, mixed>

#### inspectAll()  [header link](class-guzzlehttp-promise-utils-method-inspectall.md)

Waits on all of the provided promises, but does not unwrap rejected
promises as thrown exception.

`
    public
            static        inspectAll(array<string|int, PromiseInterface> $promises) : array<string|int, mixed>`

Returns an array of inspection state arrays.

##### Parameters

$promises
: array<string\|int, [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md) >

Traversable of promises to wait upon.

##### Tags  [header link](class-guzzlehttp-promise-utils-method-inspectall-tags.md)

seeinspect

for the inspection state array format.

##### Return values

array<string\|int, mixed>

#### queue()  [header link](class-guzzlehttp-promise-utils-method-queue.md)

Get the global task queue used for promise resolution.

`
    public
            static        queue([TaskQueueInterface|null $assign = null ]) : TaskQueueInterface`

This task queue MUST be run in an event loop in order for promises to be
settled asynchronously. It will be automatically run when synchronously
waiting on a promise.

`
while ($eventLoop->isRunning()) {
    GuzzleHttp\Promise\Utils::queue()->run();
}
`

##### Parameters

$assign
: [TaskQueueInterface](class-guzzlehttp-promise-taskqueueinterface.md) \|null
= null

Optionally specify a new queue instance.

##### Return values

[TaskQueueInterface](class-guzzlehttp-promise-taskqueueinterface.md)

#### settle()  [header link](class-guzzlehttp-promise-utils-method-settle.md)

Returns a promise that is fulfilled when all of the provided promises have
been fulfilled or rejected.

`
    public
            static        settle(mixed $promises) : PromiseInterface`

The returned promise is fulfilled with an array of inspection state arrays.

##### Parameters

$promises
: mixed

Promises or values.

##### Tags  [header link](class-guzzlehttp-promise-utils-method-settle-tags.md)

seeinspect

for the inspection state array format.

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### some()  [header link](class-guzzlehttp-promise-utils-method-some.md)

Initiate a competitive race between multiple promises or values (values
will become immediately fulfilled promises).

`
    public
            static        some(int $count, mixed $promises) : PromiseInterface`

When count amount of promises have been fulfilled, the returned promise
is fulfilled with an array that contains the fulfillment values of the
winners in order of resolution.

This promise is rejected with a [AggregateException](class-guzzlehttp-promise-aggregateexception.md) if the number
of fulfilled promises is less than the desired $count.

##### Parameters

$count
: int

Total number of promises.

$promises
: mixed

Promises or values.

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### task()  [header link](class-guzzlehttp-promise-utils-method-task.md)

Adds a function to run in the task queue when it is next \`run()\` and
returns a promise that is fulfilled or rejected with the result.

`
    public
            static        task(callable $task) : PromiseInterface`

##### Parameters

$task
: callable

Task function to run.

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### unwrap()  [header link](class-guzzlehttp-promise-utils-method-unwrap.md)

Waits on all of the provided promises and returns the fulfilled values.

`
    public
            static        unwrap(iterable<string|int, PromiseInterface> $promises) : array<string|int, mixed>`

Returns an array that contains the value of each promise (in the same
order the promises were provided). An exception is thrown if any of the
promises are rejected.

##### Parameters

$promises
: iterable<string\|int, [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md) >

Iterable of PromiseInterface objects to wait on.

##### Tags  [header link](class-guzzlehttp-promise-utils-method-unwrap-tags.md)

throwsThrowable

on error

##### Return values

array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-guzzlehttp-promise-utils-toc-methods.md)
- Methods
  - [all()](class-guzzlehttp-promise-utils-method-all.md)
  - [any()](class-guzzlehttp-promise-utils-method-any.md)
  - [inspect()](class-guzzlehttp-promise-utils-method-inspect.md)
  - [inspectAll()](class-guzzlehttp-promise-utils-method-inspectall.md)
  - [queue()](class-guzzlehttp-promise-utils-method-queue.md)
  - [settle()](class-guzzlehttp-promise-utils-method-settle.md)
  - [some()](class-guzzlehttp-promise-utils-method-some.md)
  - [task()](class-guzzlehttp-promise-utils-method-task.md)
  - [unwrap()](class-guzzlehttp-promise-utils-method-unwrap.md)

[Back To Top](class-guzzlehttp-promise-utils-top.md)
