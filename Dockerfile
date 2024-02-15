FROM golang:1.20

WORKDIR /app

COPY contributors.go /app/.
COPY default_variables.go /app/.
COPY go.mod /app/.
COPY header.go /app/.
COPY main.go /app/.
COPY readme_generation.go /app/.
COPY structure_tree.go /app/.
COPY tags.go /app/.

ENTRYPOINT go run ./main.go ./contributors.go ./default_variables.go ./header.go ./readme_generation.go ./structure_tree.go ./tags.go 