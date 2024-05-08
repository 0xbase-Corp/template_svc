-- bundles table
CREATE TABLE IF NOT EXISTS bundles (
    bundle_id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- global_wallets table
CREATE TABLE IF NOT EXISTS global_wallets (
    wallet_id SERIAL PRIMARY KEY,
    portfolio_id INTEGER NOT NULL,
    bundle_id INTEGER NULL,
    wallet_address VARCHAR(255) unique NOT NULL,
    blockchain_type VARCHAR(50) NOT NULL,
    api_endpoint TEXT,
    api_version VARCHAR(50),
    last_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
