---
name: version-constraints
description: Use when working with version constraint expressions, compatibility checks, and semver range logic. Provides expert guidance on constraint syntax, operators, and usage via Go SDK, CLI, and MCP.
argument-hint: <constraint-expression-or-task>
---

# Version Constraints Skill

## When to Use

- User needs to check if a version satisfies a constraint expression
- User is resolving dependency version compatibility
- User needs to filter versions by semver ranges (caret, tilde, wildcards)
- User is working with version expressions that combine AND/OR logic

## Constraint Syntax

### Operators

| Operator | Meaning | Example | Matches |
|----------|---------|---------|---------|
| `=` | Exact match | `=1.2.3` | Only 1.2.3 |
| `!=` | Not equal | `!=1.2.3` | Anything except 1.2.3 |
| `>` | Greater than | `>1.2.3` | 1.2.4, 2.0.0, ... |
| `>=` | Greater or equal | `>=1.2.3` | 1.2.3, 1.2.4, ... |
| `<` | Less than | `<2.0.0` | 1.9.9, 1.0.0, ... |
| `<=` | Less or equal | `<=2.0.0` | 2.0.0, 1.9.9, ... |
| `^` | Caret (compatible) | `^1.2.3` | >=1.2.3, <2.0.0 |
| `~` | Tilde (approximate) | `~1.2.3` | >=1.2.3, <1.3.0 |
| `x/X/*` | Wildcard | `1.x` | >=1.0.0, <2.0.0 |

### Combining Constraints

| Syntax | Logic | Example |
|--------|-------|---------|
| Comma `,` | AND | `>=1.0.0,<2.0.0` |
| Double-pipe `\|\|` | OR | `>=1.0.0 \|\| >=3.0.0` |

### Caret `^` Semantics

- `^1.2.3` → `>=1.2.3, <2.0.0` (compatible with 1.x)
- `^0.2.3` → `>=0.2.3, <0.3.0` (0.x is treated specially)
- `^0.0.3` → `>=0.0.3, <0.0.4` (0.0.x is treated specially)

### Tilde `~` Semantics

- `~1.2.3` → `>=1.2.3, <1.3.0` (allows patch-level changes)
- `~1.2` → `>=1.2.0, <1.3.0` (same as above)

### Wildcard Semantics

- `1.x` or `1.*` or `1.X` → `>=1.0.0, <2.0.0`
- `1.2.x` or `1.2.*` → `>=1.2.0, <1.3.0`

## API Reference

### Go SDK

**ParseConstraint(expr string) (*Constraint, error)**
Parse a single constraint expression.

**ParseConstraintSet(expr string) (*ConstraintSet, error)**
Parse comma-separated AND constraints.

**ParseConstraintUnion(expr string) (*ConstraintUnion, error)**
Parse ||-separated OR constraint sets.

**func (c *Constraint) Match(v *Version) bool**
Check if a version matches a single constraint.

**func (cs *ConstraintSet) Match(v *Version) bool / Satisfies(v *Version) bool**
Check if a version matches all constraints in the set.

**func (cu *ConstraintUnion) Match(v *Version) bool / Satisfies(v *Version) bool**
Check if a version matches any constraint set in the union.

**func (v *Version) Satisfies(constraint *Constraint) bool**
Check if this version satisfies a constraint.

**func (v *Version) Matches(expr string) (bool, error)**
Parse expression and check if this version satisfies it.

### Filtering

**func FilterByConstraint(versions []*Version, constraint *Constraint) []*Version**
Filter versions by a single Constraint.

**func FilterByConstraintSet(versions []*Version, cs *ConstraintSet) []*Version**
Filter versions by a ConstraintSet.

### CLI

```bash
# Check constraint — single type
versions constraint ">=1.0.0" 1.5.0 --type single

# Check constraint — ConstraintSet (AND logic, default)
versions constraint ">=1.0.0,<2.0.0" 1.5.0

# Check constraint — ConstraintUnion (OR logic)
versions constraint ">=1.0.0 || >=3.0.0" 3.5.0 --type union

# Version-centric check (auto-detects type)
versions satisfies 1.5.0 ">=1.0.0,<2.0.0"

# Filter by constraint — single
versions filter --constraint ">=1.0.0" --constraint-type single 0.5.0 1.0.0 2.0.0

# Filter by constraint — set (default)
versions filter --constraint ">=1.0.0,<2.0.0" 1.0.0 1.5.0 2.0.0

# Filter by constraint — union
versions filter --constraint ">=1.0.0 || >=3.0.0" --constraint-type union 1.0.0 2.0.0 3.0.0
```

### MCP

```
# Check constraint
version_constraint_check(expression=">=1.0.0,<2.0.0", version="1.5.0")

# Check constraint with type
version_constraint_check(expression=">=1.0.0 || >=3.0.0", version="3.5.0", type="union")

# Filter versions by constraint
version_filter(versions=["1.0.0", "1.5.0", "2.0.0"], constraint=">=1.0.0,<2.0.0")
```

## Code Examples

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions-skills"
)

func main() {
    // Single constraint
    c, _ := versions.ParseConstraint(">=1.0.0")
    v := versions.NewVersion("1.5.0")
    fmt.Println(c.Match(v))  // true

    // ConstraintSet (AND logic, comma-separated)
    cs, _ := versions.ParseConstraintSet(">=1.0.0,<2.0.0")
    fmt.Println(cs.Satisfies(versions.NewVersion("1.5.0")))  // true
    fmt.Println(cs.Satisfies(versions.NewVersion("2.0.0")))  // false

    // ConstraintUnion (OR logic, ||-separated)
    cu, _ := versions.ParseConstraintUnion(">=1.0.0,<2.0.0 || >=3.0.0")
    fmt.Println(cu.Satisfies(versions.NewVersion("3.5.0")))  // true
    fmt.Println(cu.Satisfies(versions.NewVersion("2.5.0")))  // false

    // Version.Matches shortcut
    ok, _ := versions.NewVersion("1.5.0").Matches(">=1.0.0,<2.0.0")
    fmt.Println(ok)  // true

    // Caret constraint
    c2, _ := versions.ParseConstraint("^1.2.3")
    fmt.Println(c2.Match(versions.NewVersion("1.9.9")))  // true
    fmt.Println(c2.Match(versions.NewVersion("2.0.0")))  // false

    // Tilde constraint
    c3, _ := versions.ParseConstraint("~1.2.3")
    fmt.Println(c3.Match(versions.NewVersion("1.2.9")))  // true
    fmt.Println(c3.Match(versions.NewVersion("1.3.0")))  // false
}
```

## Important Notes

- Caret `^` follows npm semver semantics: `^0.x.y` is more restrictive than `^1.x.y`
- Tilde `~` allows patch-level changes only: `~1.2.3` matches 1.2.x but not 1.3.0
- Wildcards `x/X/*` are equivalent: `1.x` = `1.X` = `1.*`
- ConstraintSet uses comma for AND; ConstraintUnion uses `||` for OR
- ParseConstraintSet wraps errors with context: `parse constraint "BAD": invalid version in constraint`
- ParseConstraintUnion does not wrap errors — the `||` context is lost in error messages
- Spaces around operators are not supported: use `>=1.0.0` not `>= 1.0.0`
