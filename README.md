Convergen
=========

A type-to-type copy function code generator.

Notation Table
--------------

| notation                                  | location           | summary                                                                               |
|-------------------------------------------|--------------------|---------------------------------------------------------------------------------------|
| :match &lt;`name` &#124; `none`>          | interface, method  | Sets the field matcher algorithm (default: `name`).                                   |
| :style &lt;`return` &#124; `arg`>         | interface, method  | Sets the style of the assignee variable input/output (default: `return`).             |
| :recv &lt;_var_>                          | method             | Specifies the source value as a receiver of the generated function.                   |
| :reverse                                  | 	method            | Reverses the copy direction. Might be useful with receiver form.                      |
| :case	                                    | interface, method  | Sets case-sensitive for name match (default).                                         |
| :case:off	                                | interface, method  | Sets case-insensitive for name match.                                                 |
| :getter	                                  | interface, method  | Includes getters for name match.                                                      |
| :getter:off	                              | interface, method  | Excludes getters for name match (default).                                            |
| :stringer                                 | 	interface, method | Calls String() if appropriate in name match.                                          |
| :stringer:off                             | 	interface, method | Calls String() if appropriate in name match (default).                                |
| :typecast	                                | interface, method	 | Allows type casting if appropriate in name match.                                     |
| :typecast:off                             | 	interface, method | Suppresses type casting if appropriate in name match (default).                       |
| :skip &lt;_dst field pattern_>            | method             | Marks the destination field to skip copying. Regex is allowed in /…/ syntax.          |
| :map &lt;_src_> &lt;_dst field_>          | method             | the pair as assign source and destination.                                            |
| :conv &lt;_func_> &lt;_src_> [_to field_] | method             | Converts the source value by the converter and assigns its result to the destination. |
| :literal &lt;_dst_> &lt;_literal_>        | method             | Assigns the literal expression to the destination.                                    |
| :preprocess &lt;_func_>                   | method             | Calls the function at the beginning of the convergen func.                            |
| :postprocess &lt;_func_>                  | method             | Calls the function at the end of the convergen function.                              |

Sample
------

To use Convergen, write a generator code in the following convention:

```go
//go:build convergen

package sample

import (
    "time"

    "github.com/sample/myapp/domain"
    "github.com/sample/myapp/storage"
)

//go:generate go run github.com/reedom/convergen@v0.6.1
type Convergen interface {
    // :typecast
    // :stringer
    // :map Created.UnixMilli() Created
    DomainToStorage(*domain.User) *storage.User
}
```

Convergen generates code similar to the following:

```go
// Code generated by github.com/reedom/convergen
// DO NOT EDIT.

package sample

import (
    "time"

    "github.com/sample/myapp/domain"
    "github.com/sample/myapp/storage"
)

func DomainToStorage(src *domain.User) (dst *storage.User) {
    dst = &storage.User{}
    dst.ID = int64(src.ID)
    dst.Name = src.Name
    dst.Status = src.Status.String()
    dst.Created = src.Created.UnixMilli()

    return
}
```

for these struct types:

```go
package domain

import (
    "time"
)

type User struct {
    ID      int
    Name    string
    Status  Status
    Created time.Time
}

type Status string

func (s Status) String() string {
    return string(s)
}
```

outputs:

```go
package storage

type User struct {
    ID      int64
    Name    string
    Status  string
    Created int64
}
```

Installation and Introduction
-----------------------------

### Use as a Go generator

To use Convergen as a Go generator, install the module in your Go project directory via go get:

```shell
$ go get -u github.com/reedom/convergen@latest
```

Then, write a generator as follows:

```go
//go:generate go run github.com/reedom/convergen@v0.6.1
type Convergen interface {
    …
}
````

### Use as a CLI command

To use Convergen as a CLI command, install the command via go install:

```shell
$ go install github.com/reedom/convergen@latest
```

You can then generate code by calling:

```shell
$ convergen any-codegen-defined-code.go
```

The CLI help shows:

```shell
Usage: convergen [flags] <input path>

By default, the generated code is written to <input path>.gen.go

Flags:
  -dry
        Perform a dry run without writing files.
  -log
        Write log messages to <output path>.log.
  -out string
        Set the output file path.
  -print
        Print the resulting code to STDOUT as well.
```

Notations
---------

### `:convergen`

Use the `:convergen` notation to mark an interface as a converter definition.

By default, Convergen only looks for an interface named "`Convergen`" as a converter definition block.
You can use the `:convergen` notation to enable Convergen to recognize other interface names as well.
This is especially useful if you want to define methods with the same name but different receivers.

__Available locations__

interface

__Format__

```text
":convergen"
```

__Examples__

```go
// :convergen
type TransportConvergen interface {
    // :recv t
    ToDomain(*trans.Model) *domain.Model 
}

// :convergen
type PersistentConvergen interface {
    // :recv t
    ToDomain(*persistent.Model) *domain.Model 
}

```


### `:match <algorithm>`

Use the `:match` notation to set the field matcher algorithm.

__Default__

`:match name`

__Available locations__

interface, method

__Format__

```text
":match" <algorithm>

algorithm = "name" | none"
```

__Examples__

With `name` match, the generator matches fields or getter names (and their types) to generate
the conversion code.


```go
package model

type User struct {
    ID   int
    Name string
}
```
```go
package web

type User struct {
    id   int
    name string
}

func (u *User) ID() int {
  return u.id
}
```
```go
// :match name 
type Convergen interface {
    ToStorage(*User) *storage.User
}
```

Convergen generates:

```go
func ToStorage(src *User) (dst *storage.User) {
    dst := &storage.User{}
    dst.ID = src.ID()
    dst.Name = src.name

    return
}
```

With `none` match, Convergen only processes fields or getters that have been explicitly
specified using `:map` and `:conv`.

### `:style <style>`

Use the `:style` notation to set the style of the assignee variable input/output.

__Default__

`:style return`

__Available locations__

interface, method

__Format__

```text
":style" style

style = "arg" | "return"
```

__Examples__

Examples of `return` style:

Basic:

```go
func ToStorage(src *domain.Pet) (dst *storage.Pet) {
```

With error:

```go
func ToStorage(src *domain.Pet) (dst *storage.Pet, err error) {
```

With receiver:

```go
func (src *domain.Pet) ToStorage() (dst *storage.Pet) {
```

Examples of `arg` style:

Basic:

```go
func ToStorage(dst *storage.Pet, src *domain.Pet) {
```

With error:

```go
func ToStorage(dst *storage.Pet, src *domain.Pet) (err error) {
```

With receiver:

```go
func (src *domain.Pet) ToStorage(dst *storage.Pet) {
```

### `:recv <var>`

Use the `:recv` notation to specify the source value as a receiver of the generated function.

According to the Go language specification, the receiver type must be defined in the same
package as the generated code.

By convention, the &lt;_var_> should be the same identifier as the methods of the type defines.

__Default__

No receiver is used.

__Available locations__

method

__Format__

```text
":recv" var

var = variable-identifier 
```

__Examples__

In the following example, assume that `domain.User` is defined in another file under the same
directory (package). It also assumes that other methods use `u` as their receiver variable name.


```go
package domain

import (
    "github.com/sample/myapp/storage"
)

type Convergen interface {
    // :recv u
    ToStorage(*User) *storage.User  
}
```

The generated code will be:

```go
package domain

import (
    "github.com/sample/myapp/storage"
)

type User struct {
    ID   int
    Name string
}

func (u *User) ToStorage() (dst *storage.User) {
    dst = &storage.User{}
    dst.ID = int64(u.ID)  
    dst.Name = u.Name

    return
}
```

### `:reverse`

Reverse copy direction. Might be useful with receiver form.  
To use `:reverse`, `:style arg` is required. (Otherwise it can't have any data source to copy from.)

__Default__

Copy in normal direction. In receiver form, receiver to a variable in argument.

__Available locations__

method

__Format__

```text
":reverse"
```

__Examples__

```go
package domain

import (
    "github.com/sample/myapp/storage"
)

type Convergen interface {
    // :style arg
    // :recv u
    // :reverse
    FromStorage(*User) *storage.User  
}
```

Will have:

```go
package domain

import (
    "github.com/sample/myapp/storage"
)

type User struct {
    ID   int
    Name string
}

func (u *User) FromStorage(src *storage.User) {
    u.ID = int(src.User)  
    u.Name = src.Name
}
```

### `:case` / `:case:off`

This notation controls case-sensitive or case-insensitive matches in field and method names. 

It is applicable to `:match name`, `:getter`, and `:skip` notations.   
Other notations like `:map` and `:conv` retain case-sensitive matches.

__Default__

":case"

__Available locations__

interface, method

__Format__

```go
":case"
":case:off"
```

__Examples__

```go
// interface level notation makes ":case:off" as default.
// :case:off
type Convergen interface {
    // Turn on case-sensitive match for names.
    // :case
    ToUserModel(*domain.User) storage.User

    // Adopt the default, case-insensitive match in this case.
    ToCategoryModel(*domain.Category) storage.Category
}
```

### `:getter` / `:getter:off`

Include getters for name match.

__Default__

`:getter:off`

__Available locations__

interface, method

__Format__

```text
":getter"
":getter:off"
```

__Examples__

With those models:

```go
package domain

type User struct {
    name string
}

func (u *User) Name() string {
    return u.name
}
```

```go
package storage

type User struct {
    Name string
}
```

The default Convergen behaviour can't find the private `name` and won't notice the getter.  
So, with the following we'll get…

```go
type Convergen interface {
    ToStorageUser(*domain.User) *storage.User
}
````

```go
func ToStorageUser(src *domain.User) (dst *storage.User)
    dst = &storage.User{}
    // no match: dst.Name

    return
}
```

And with `:getter` we'll have…

```go
type Convergen interface {
    // :getter
    ToStorageUser(*domain.User) *storage.User
}
````

```go
func ToStorageUser(src *domain.User) (dst *storage.User)
    dst = &storage.User{}
    dst.Name = src.Name()

    return
}
```

Alternatively, you can get the same result with `:map`.  
This is worth learning since `:getter` affects the entire method - `:map` allows you to get
the result selectively. 

```go
type Convergen interface {
    // :map Name() Name
    ToStorageUser(*domain.User) *storage.User
}
```

### `:stringer` / `:stringer:off`

When matching field names, call the String() method of a custom type if it exists.

By default, Convergen has no way of knowing how to assign a custom type to a string.  
Using the :stringer notation will tell Convergen to look for a String() method on any custom
types and use it when appropriate.

__Default__

`:stringer:off`

__Available locations__

interface, method

__Format__

```text
":stringer"
":stringer:off"
```

__Examples__

Consider the following code:

```go
package domain

type User struct {
    Status Status
}

type Status struct {
    status string
}

func (s Status) String() string {
    return string(s)
}

var (
    NotVerified = Status{"notVerified"}
    Verified    = Status{"verified"}
    Invalidated = Status{"invalidated"}
)
```

```go
package storage

type User struct {
    String string
}
```

Without any additional notations, Convergen has no idea how to assign the `Status` type to a string.
By adding `:stringer` notation to the Convergen interface, we're telling Convergen to look for a
`String()` method on any custom types and use it when appropriate:

```go
type Convergen interface {
    // :stringer
    ToStorageUser(*domain.User) *storage.User
}
```

Convergen will generate the following code:

```go
func ToStorageUser(src *domain.User) (dst *storage.User)
    dst = &storage.User{}
    dst.Status = src.Status.String()

    return
}
```

Alternatively, you can achieve the same result with `:map`. However, `:stringer` affects 
the entire method, while `:map` allows you to specify the fields to map selectively:

```go
type Convergen interface {
    // :map Status.String() Name
    ToStorageUser(*domain.User) *storage.User
}
```

### `:typecast`

Allow type casting if appropriate in name match.

__Default__

`:typecast:off`

__Available locations__

interface, method

__Format__

```text
":typecast"
":typecast:off"
```

__Examples__

With those models:

```go
package domain

type User struct {
    ID     int
    Name   string
    Status Status
}

type Status string

```

```go
package storage

type User struct {
    ID     int64  
    Name   string
    Status string
}
```

Convergen respects types strictly. It will give up copying fields if their types do not match.
Note that Convergen relies on the [types.AssignableTo(V, T Type) bool][] method from the standard
packages. This means that the judgment is done by the type system of Go itself, not by a dumb
string type name match.

[types.AssignableTo(V, T Type) bool]: https://pkg.go.dev/go/types#AssignableTo

Without `:typecast` turned on:

```go
type Convergen interface {
    ToDomainUser(*storage.User) *domain.User
}
````

We'll get:

```go
func ToDomainUser(src *storage.User) (dst *domain.User)
    dst = &domain.User{}
    // no match: dst.ID
    dst.Name = src.Name
    // no match: dst.Status

    return
}
```

With `:typecast` it turned on:

```go
type Convergen interface {
	  // :typecast
    ToDomainUser(*storage.User) *domain.User
}
````

```go
func ToDomainUser(src *storage.User) (dst *domain.User)
    dst = &domain.User{}
    dst.ID = int(src.ID)
    dst.Name = src.Name
    dst.Status = domain.Status(src.Status)

    return
}
```

### `:skip <dst field pattern>`

Mark the destination field to skip copying.

A method can have multiple :skip lines that enable skipping multiple fields.  
Other than field-path match, it accepts [regular expression][] match. To specify,
wrap the expression with `/`.  
`:case` / `:case:off` affects `:skip`.

[regular expression]: https://github.com/google/re2/wiki/Syntax

__Available locations__

method

__Format__

```text
":skip" dst-field-pattern

dst-field-pattern  = field-path | regexp
field-path         = { identifier "." } identifier
regexp             = "/" regular-expression "/" 
```

__Examples__

Suppose we have the following domain and storage structs:

```go
package domain

type User struct {
    ID      int
    Name    string
    Email   string
    Address Address
}

type Address struct {
    Street  string
    City    string
    ZipCode string
}
```

If we want to skip copying the Name field of the storage.User struct,
we can use the :skip notation as follows:

```go
type Convergen interface {
    // :skip Name
    ToStorage(*domain.User) *storage.User
}
```

If we want to skip copying multiple fields, we can use multiple :skip notations:

```go
type Convergen interface {
    // :skip Name
    // :skip Email
    ToStorage(*domain.User) *storage.User
}
```

We can also use regular expressions to match multiple fields:

```go
type Convergen interface {
    // :skip /^Name|Email$/
    ToStorage(*domain.User) *storage.User
}
```

This will result in the same generated code as the previous example.

### `:map <src> <dst field>`

Specify a field mapping rule.

When to use:
- copying a value between fields having different names.
- assigning a method's result value to a destination field.

A method can have multiple `:map` lines that enable mapping multiple fields.

`:case:off` does not affect `:map`;
&lt;src> and &lt;dst field> are compared in a case-sensitive manner.  

__Available locations__

method

__Format__

```text
":map" src dst-field

src                   = field-or-method-chain
dst-field             = field-path
field-path            = { identifier "." } identifier
field-or-getter-chain = { (identifier | getter) "." } (identifier | getter)
getter                = identifier "()"  
```

__Examples__

In the following example, two fields have the same meaning but different names.

```go
package domain

type User struct {
    ID   int
    Name string
}
```
```go
package storage

type User struct {
    UserID int
    Name   string
}
```

We can use `:map` to connect them:

```go
type Convergen interface {
    // Map the "ID" field in domain.User to the "UserID" field in storage.User.
    // :map ID UserID
    ToStorage(*domain.User) *storage.User
}
```

```go
func ToStorage(src *domain.User) (dst *storage.User) {
    dst = storage.User{}
    dst.UserID = src.ID
    dst.Name = src.Name
    
    return
}
```

In the following example, Status is a custom type with a method to retrieve its raw value.

```go
package domain

type User struct {
    ID     int
    Name   string
    Status Status
}

type Status int

func (s Status) Int() int {
    return int(s)
}

var (
    NotVerified = Status(1)
    Verified    = Status(2)
    Invalidated = Status(3)
)
```
```go
package storage

type User struct {
    UserID int
    Name   string
    Status int
}
```

We can use `:map` to apply the method's return value to assign:

```go
type Convergen interface {
    // Map the "ID" field in domain.User to the "UserID" field in storage.User.
    // Map the result of the "Status.Int()" method in domain.User to the "Status" field in storage.User.
    // :map ID UserID
    // :map Status.Int() Status
    ToStorage(*domain.User) *storage.User
}
```

```go
func ToStorage(src *domain.User) (dst *storage.User) {
    dst = storage.User{}
    dst.UserID = src.ID
    dst.Name = src.Name
    dst.Status = src.Status.Int()

    return
}
```

Note that the method's return value should be compatible with the destination field.  
If they are not compatible, you can use `:typecast` or `:stringer` to help Convergen
with the conversion.  
Alternatively, you can use `:conv` notation to define a custom conversion function.

### `:conv <func> <src> [dst field]`

Convert the source value by the converter and assign its result to the destination.

_func_ must accept _src_ value as the sole argument and return either   
  a) a single value that is compatible with the _dst_, or  
  a) a pair of variables as (_dst_, error).   
For the latter case, the method definition should have `error` in return value(s). 

You can omit _dst field_ if the source and destination field paths are exactly the same.

`:case:off` does not take effect on `:conv` as  &lt;src> and &lt;dst field> are compared
in a case-sensitive manner.

__Available locations__

method

__Format__

```text
":conv" func src [dst-field]

func                  = identifier
src                   = field-or-method-chain
dst-field             = field-path
field-path            = { identifier "." } identifier
field-or-getter-chain = { (identifier | getter) "." } (identifier | getter)
getter                = identifier "()"  
```

__Examples__

```go
package domain

type User struct {
    ID    int
    Email string
}
```
```go
package storage

type User struct {
    ID    int
    Email string
}
```

To store an encrypted Email field, we can use a converter function:

```go
import (
    // The referenced library should have been imported anyhow.
    _ "github.com/sample/myapp/crypto"
)

type Convergen interface {
    // :conv crypto.Encrypt Email
    ToStorage(*domain.User) *storage.User
}
```

This results in:

```go
import (
    "github.com/sample/myapp/crypto"
    _ "github.com/sample/myapp/crypto"
)

func ToStorage(src *domain.User) (dst *storage.User) {
    dst = storage.User{}
    dst.ID = src.ID
    dst.Email = crypto.Encrypt(src.Email)

    return
}
```

If you want to use a converter function that returns an error, you should add `error`
to the return values of the converter method as well:

```go
import (
    // The referenced library should have been imported anyhow.
    _ "github.com/sample/myapp/crypto"
)

type Convergen interface {
    // :conv crypto.Decrypt Email
    FromStorage(*storage.User) (*domain.User, error)
}
```

This results in:

```go
import (
    "github.com/sample/myapp/crypto"
    _ "github.com/sample/myapp/crypto"
)

func ToStorage(src *storage.User) (dst *domain.User, err error) {
    dst = domain.User{}
    dst.ID = src.ID
    dst.Email, err = crypto.Decrypt(src.Email)
    if err != nil {
        return
    }

    return
}
```

### `:literal <dst> <literal>`

Assign a literal expression to the destination field.

__Available locations__

method

__Format__

```text
":literal"  dst literal
```

__Examples__

```go
type Convergen interface {
    // :literal Created time.Now()
    FromStorage(*storage.User) *domain.User()
}
```

### `:preprocess <func>` / `:postprocess <func>`

Call the function at the beginning(`preprocess`) or at the end(`postprocess`) of the convergen function.

__Available locations__

method

__Format__

```text
":preprocess"  func
":postprocess" func

func  = identifier
```

__Examples__

```go
type Convergen interface {
    // :preprocess prepareInput
    // :postprocess cleanUpOutput
    FromStorage(*storage.User) *domain.User
}

func prepareInput(src *storage.User) *storage.User {
    // modify the input source before conversion
    return src
}

func cleanUpOutput(dst *domain.User) *domain.User {
    // modify the output destination after conversion
    return dst
}
```
``
When FromStorage is called, the prepareInput function will be called with the input
argument before the conversion takes place. Then the FromStorage method will be executed. 
Finally, the cleanUpOutput function will be called with the output result after the 
conversion has taken place.

```go
type Convergen interface {
    // :preprocess prepareInput
    // :postprocess cleanUpOutput
    FromStorage(*storage.User) (*domain.User, error)
}

func prepareInput(src *storage.User) (*storage.User, error) {
    // modify the input source before conversion
    return src, nil
}

func cleanUpOutput(dst *domain.User) (*domain.User, error) {
    // modify the output destination after conversion
    return dst, nil
}
```


Contributing
------------

For those who want to contribute, there are several ways to do it, including:

- Reporting bugs or issues that you encounter while using Convergen.
- Suggesting new features or improvements to the existing ones.
- Implementing new features or fixing bugs by making a pull request to the project.
- Improving the documentation or examples to make it easier for others to use Convergen.
- Creating a project's logo to help with its branding.
- Showing your support by giving the project a star.

- By contributing to the project, you can help make it better and more useful for everyone.
- So, if you're interested, feel free to get involved!
