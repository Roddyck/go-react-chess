-- +goose Up
CREATE TABLE games (
  id uuid PRIMARY KEY,
  board jsonb NOT NULL,
  turn text NOT NULL,
  history jsonb NOT NULL,
  players jsonb NOT NULL
);

-- +goose Down
DROP TABLE games;
