# convex

**Flexible Type Conversion and Marshaling for Go**

A lightweight Go library for seamless type conversions, marshaling/unmarshaling, and database integration. Handle JSON, SQL, and custom types with minimal boilerplate.

---

## Features

- **Custom Types**: Extend with custom conversion logic, including `dynamic_map`, `unix_time`, and `pg_interval`.
- **Custom Type Conversions**: `dynamic_map`, `unix_time`, `pg_interval`, and more soon!
- **Database Support**: Implements `Scanner` and `Valuer` for SQL operations.
- **JSON/XML/YAML**: Built-in marshaling and unmarshaling.

---

## Installation

```bash
go get -v github.com/yourusername/convex