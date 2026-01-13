<p align="center">
    <img src="public/go.png" alt="Go" height="80" />
    <img src="public/graphql.png" alt="GraphQL" height="80" />
    <img src="public/sqlc.png" alt="sqlc" height="80" />
</p>
<p align="center">Collection of Go backend servers paired with various tools and frameworks</p>

This repository contains Collection of servers built in Go with different frameworks and approaches following best practices and recommended abstractions.

## Tools & Frameworks

The following tools/frameworks are used in this project:

| Tool/Framework | Description |
|----------------|-------------|
| **GraphQL** | Implements a GraphQL API layer implemented using gqlgen (schema-first code generation for resolvers and models). |
| **sqlc** | The sqlc compiler generates type-safe Go code from raw SQL queries → implements a clean **repository pattern** for the data layer. |
| **Atlas** | Declarative schema migrations using Atlas. |
| **net/http** | Implements go's net/http standard lib for advanced routing, middleware, and REST endpoint grouping. |

## Development Notes

### GraphQL Server Development Flow

1. Add the relational tables in SQL schema file.
2. Migrate the SQL schema using atlas.
3. Create GraphQL schema: Add the stucts and queries by composing the entities of data according to whats being needed in the client/frontent. Enable autobind and mark separate resolvers in gqlgen.yml models config for the fields which are relationally in a different table in the DB but are composed in each other in the graphql schema.
4. Generate models and resolvers using gqlgen.
5. Write the raw SQL queries according to the generated resolvers.
6. Run sqlc generate to generate the go bindings for queries written in previous step. Use these bindings in the graphql resolvers.

### GraphQL Query Examples

**Create User Mutation:**

```graphql
mutation createUser {
  createUser (input: {id: "id1", name: "John"}){
    name
  }
}
```

**Get Users Query:**

```graphql
query getUsers {
  users {
    id,
    name
  }
}
```

**Create Todo Mutation:**

```graphql
mutation createTodo {
  createTodo (input: {id: "todoid1", text: "First todo from John", userId: "id1"}) {
    text,
    done
  }
}
```

**Get Todos Query:**

```graphql
query getTodos {
  todos {
    id,
    text,
    done,
    user {
      id,
      name
    }
  }
}
```

---

## License

MIT
