# SQLC

- Use `sqlc.arg`, and `sqlc.narg` only, don't use `@`
- created_at, updated_at, deleted_at should be optional argument (use `sqlc.narg`) and default now
- Create returns one, List returns many, Get returns one, Update returns on/many up to the name of the query, Delete execrows
- If passing multiple values (array of an entity), you may use CTE (with) and ref that in
