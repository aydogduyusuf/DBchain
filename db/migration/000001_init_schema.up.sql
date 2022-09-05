CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "hashed_password" varchar UNIQUE NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "wallet_public_address" varchar UNIQUE NOT NULL,
  "wallet_private_address" varchar UNIQUE NOT NULL,
  "create_time" timestamptz NOT NULL DEFAULT (now()),
  "update_time" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "delete_time" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "is_active" boolean NOT NULL DEFAULT true
);

CREATE TABLE "tokens" (
  "id" bigserial PRIMARY KEY,
  "u_id" bigserial NOT NULL,
  "token_name" varchar NOT NULL,
  "symbol" varchar NOT NULL,
  "supply" bigint NOT NULL,
  "contract_address" varchar UNIQUE NOT NULL,
  "create_time" timestamptz NOT NULL DEFAULT (now()),
  "update_time" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "delete_time" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "is_active" boolean NOT NULL DEFAULT true
);

CREATE TABLE "transactions" (
  "id" bigserial PRIMARY KEY,
  "transaction_type" varchar NOT NULL,
  "from_address" varchar NOT NULL,
  "to_address" varchar NOT NULL,
  "transfer_data" varchar NOT NULL,
  "hash_value" varchar UNIQUE NOT NULL,
  "create_time" timestamptz NOT NULL DEFAULT (now()),
  "update_time" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "delete_time" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "is_active" boolean NOT NULL DEFAULT true
);

CREATE INDEX ON "tokens" ("u_id");

COMMENT ON COLUMN "transactions"."from_address" IS 'from wallet address';

COMMENT ON COLUMN "transactions"."to_address" IS 'to wallet address, null when contract deploy';

ALTER TABLE "tokens" ADD FOREIGN KEY ("u_id") REFERENCES "users" ("id");