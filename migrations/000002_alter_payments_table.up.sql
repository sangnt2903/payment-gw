CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

ALTER TABLE payments
    ALTER COLUMN id TYPE UUID USING (uuid_generate_v4()),
    ALTER COLUMN id SET DEFAULT uuid_generate_v4(),
    ADD COLUMN description TEXT,
    ADD COLUMN provider VARCHAR(50),
    ADD COLUMN provider_transaction_id VARCHAR(255),
    ADD COLUMN metadata JSONB DEFAULT '{}'::jsonb;

CREATE INDEX idx_payments_provider ON payments(provider);
CREATE INDEX idx_payments_provider_txn_id ON payments(provider_transaction_id);