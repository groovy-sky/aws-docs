Menu

- [Aws](namespace-aws.md)

## HandlerList        in package    - [Aws](package-aws.md)       implements  Countable

Builds a single handler function from zero or more middleware functions and
a handler. The handler function is then used to send command objects and
return a promise that is resolved with an AWS result object.

The "front" of the list is invoked before the "end" of the list. You can add
middleware to the front of the list using one of the "prepend" method, and
the end of the list using one of the "append" method. The last function
invoked in a handler list is the handler (a function that does not accept a
next handler but rather is responsible for returning a promise that is
fulfilled with an Aws\\ResultInterface object).

Handlers are ordered using a "step" that describes the step at which the
SDK is when sending a command. The available steps are:

- init: The command is being initialized, allowing you to do things like add
default options.
- validate: The command is being validated before it is serialized
- build: The command is being serialized into an HTTP request. A middleware
in this step MUST serialize an HTTP request and populate the "@request"
parameter of a command with the request such that it is available to
subsequent middleware.
- sign: The request is being signed and prepared to be sent over the wire.

Middleware can be registered with a name to allow you to easily add a
middleware before or after another middleware by name. This also allows you
to remove a middleware by name (in addition to removing by instance).

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#toc-interfaces)

Countable

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#toc-constants)

[ATTEMPT](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#constant_ATTEMPT)
= 'attempt' [BUILD](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#constant_BUILD)
= 'build' [INIT](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#constant_INIT)
= 'init' [SIGN](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#constant_SIGN)
= 'sign' [VALIDATE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#constant_VALIDATE)
= 'validate'

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method___construct)
: mixed [\_\_toString()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method___toString)
: string Dumps a string representation of the list.[after()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_after)
: mixed Add a middleware after the given middleware by name.[appendAttempt()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_appendAttempt)
: mixed Append a middleware to the attempt step.[appendBuild()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_appendBuild)
: mixed Append a middleware to the build step.[appendInit()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_appendInit)
: mixed Append a middleware to the init step.[appendSign()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_appendSign)
: mixed Append a middleware to the sign step.[appendValidate()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_appendValidate)
: mixed Append a middleware to the validate step.[before()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_before)
: mixed Add a middleware before the given middleware by name.[count()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_count)
: int [hasHandler()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_hasHandler)
: bool Returns true if the builder has a handler.[hasMiddleware()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_hasMiddleware)
: bool Checks if a middleware exists. The middleware
should have been added with a name in order to
use this method.[interpose()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_interpose)
: mixed Interpose a function between each middleware (e.g., allowing for a trace
through the middleware layers).[prependAttempt()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_prependAttempt)
: mixed Prepend a middleware to the attempt step.[prependBuild()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_prependBuild)
: mixed Prepend a middleware to the build step.[prependInit()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_prependInit)
: mixed Prepend a middleware to the init step.[prependSign()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_prependSign)
: mixed Prepend a middleware to the sign step.[prependValidate()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_prependValidate)
: mixed Prepend a middleware to the validate step.[remove()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_remove)
: mixed Remove a middleware by name or by instance from the list.[resolve()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_resolve)
: callable Compose the middleware and handler into a single callable function.[setHandler()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_setHandler)
: mixed Set the HTTP handler that actually returns a response.

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#constants)

#### ATTEMPT  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#constant_ATTEMPT)

`
    public
        mixed
    ATTEMPT
    = 'attempt'
`

#### BUILD  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#constant_BUILD)

`
    public
        mixed
    BUILD
    = 'build'
`

#### INIT  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#constant_INIT)

`
    public
        mixed
    INIT
    = 'init'
`

#### SIGN  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#constant_SIGN)

`
    public
        mixed
    SIGN
    = 'sign'
`

#### VALIDATE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#constant_VALIDATE)

`
    public
        mixed
    VALIDATE
    = 'validate'
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#method___construct)

`
    public
                    __construct([callable $handler = null ]) : mixed`

##### Parameters

$handler
: callable
= null

HTTP handler.

#### \_\_toString()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#method___toString)

Dumps a string representation of the list.

`
    public
                    __toString() : string`

##### Return values

string

#### after()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#method_after)

Add a middleware after the given middleware by name.

`
    public
                    after(string|callable $findName, string $withName, callable $middleware) : mixed`

##### Parameters

$findName
: string\|callable

Add after this

$withName
: string

Optional name to give the middleware

$middleware
: callable

Middleware to add.

#### appendAttempt()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#method_appendAttempt)

Append a middleware to the attempt step.

`
    public
                    appendAttempt(callable $middleware[, string $name = null ]) : mixed`

##### Parameters

$middleware
: callable

Middleware function to add.

$name
: string
= null

Name of the middleware.

#### appendBuild()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#method_appendBuild)

Append a middleware to the build step.

`
    public
                    appendBuild(callable $middleware[, string $name = null ]) : mixed`

##### Parameters

$middleware
: callable

Middleware function to add.

$name
: string
= null

Name of the middleware.

#### appendInit()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#method_appendInit)

Append a middleware to the init step.

`
    public
                    appendInit(callable $middleware[, string $name = null ]) : mixed`

##### Parameters

$middleware
: callable

Middleware function to add.

$name
: string
= null

Name of the middleware.

#### appendSign()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#method_appendSign)

Append a middleware to the sign step.

`
    public
                    appendSign(callable $middleware[, string $name = null ]) : mixed`

##### Parameters

$middleware
: callable

Middleware function to add.

$name
: string
= null

Name of the middleware.

#### appendValidate()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#method_appendValidate)

Append a middleware to the validate step.

`
    public
                    appendValidate(callable $middleware[, string $name = null ]) : mixed`

##### Parameters

$middleware
: callable

Middleware function to add.

$name
: string
= null

Name of the middleware.

#### before()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#method_before)

Add a middleware before the given middleware by name.

`
    public
                    before(string|callable $findName, string $withName, callable $middleware) : mixed`

##### Parameters

$findName
: string\|callable

Add before this

$withName
: string

Optional name to give the middleware

$middleware
: callable

Middleware to add.

#### count()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#method_count)

`
    public
                    count() : int`

##### Return values

int

#### hasHandler()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#method_hasHandler)

Returns true if the builder has a handler.

`
    public
                    hasHandler() : bool`

##### Return values

bool

#### hasMiddleware()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#method_hasMiddleware)

Checks if a middleware exists. The middleware
should have been added with a name in order to
use this method.

`
    public
                    hasMiddleware(string $name) : bool`

##### Parameters

$name
: string

##### Return values

bool

#### interpose()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#method_interpose)

Interpose a function between each middleware (e.g., allowing for a trace
through the middleware layers).

`
    public
                    interpose([callable|null $fn = null ]) : mixed`

The interpose function is a function that accepts a "step" argument as a
string and a "name" argument string. This function must then return a
function that accepts the next handler in the list. This function must
then return a function that accepts a CommandInterface and optional
RequestInterface and returns a promise that is fulfilled with an
Aws\\ResultInterface or rejected with an Aws\\Exception\\AwsException
object.

##### Parameters

$fn
: callable\|null
= null

Pass null to remove any previously set function

#### prependAttempt()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#method_prependAttempt)

Prepend a middleware to the attempt step.

`
    public
                    prependAttempt(callable $middleware[, string $name = null ]) : mixed`

##### Parameters

$middleware
: callable

Middleware function to add.

$name
: string
= null

Name of the middleware.

#### prependBuild()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#method_prependBuild)

Prepend a middleware to the build step.

`
    public
                    prependBuild(callable $middleware[, string $name = null ]) : mixed`

##### Parameters

$middleware
: callable

Middleware function to add.

$name
: string
= null

Name of the middleware.

#### prependInit()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#method_prependInit)

Prepend a middleware to the init step.

`
    public
                    prependInit(callable $middleware[, string $name = null ]) : mixed`

##### Parameters

$middleware
: callable

Middleware function to add.

$name
: string
= null

Name of the middleware.

#### prependSign()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#method_prependSign)

Prepend a middleware to the sign step.

`
    public
                    prependSign(callable $middleware[, string $name = null ]) : mixed`

##### Parameters

$middleware
: callable

Middleware function to add.

$name
: string
= null

Name of the middleware.

#### prependValidate()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#method_prependValidate)

Prepend a middleware to the validate step.

`
    public
                    prependValidate(callable $middleware[, string $name = null ]) : mixed`

##### Parameters

$middleware
: callable

Middleware function to add.

$name
: string
= null

Name of the middleware.

#### remove()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#method_remove)

Remove a middleware by name or by instance from the list.

`
    public
                    remove(string|callable $nameOrInstance) : mixed`

##### Parameters

$nameOrInstance
: string\|callable

Middleware to remove.

#### resolve()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#method_resolve)

Compose the middleware and handler into a single callable function.

`
    public
                    resolve() : callable`

##### Return values

callable

#### setHandler()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html\#method_setHandler)

Set the HTTP handler that actually returns a response.

`
    public
                    setHandler(callable $handler) : mixed`

##### Parameters

$handler
: callable

Function that accepts a request and array of
options and returns a Promise.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#toc-methods)
- Constants
  - [ATTEMPT](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#constant_ATTEMPT)
  - [BUILD](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#constant_BUILD)
  - [INIT](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#constant_INIT)
  - [SIGN](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#constant_SIGN)
  - [VALIDATE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#constant_VALIDATE)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method___construct)
  - [\_\_toString()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method___toString)
  - [after()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_after)
  - [appendAttempt()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_appendAttempt)
  - [appendBuild()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_appendBuild)
  - [appendInit()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_appendInit)
  - [appendSign()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_appendSign)
  - [appendValidate()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_appendValidate)
  - [before()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_before)
  - [count()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_count)
  - [hasHandler()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_hasHandler)
  - [hasMiddleware()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_hasMiddleware)
  - [interpose()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_interpose)
  - [prependAttempt()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_prependAttempt)
  - [prependBuild()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_prependBuild)
  - [prependInit()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_prependInit)
  - [prependSign()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_prependSign)
  - [prependValidate()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_prependValidate)
  - [remove()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_remove)
  - [resolve()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_resolve)
  - [setHandler()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#method_setHandler)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html#top)
