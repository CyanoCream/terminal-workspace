CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(255) UNIQUE NOT NULL,
                       password VARCHAR(255) NOT NULL,
                       role VARCHAR(50) NOT NULL DEFAULT 'user',
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE terminals (
                           id SERIAL PRIMARY KEY,
                           name VARCHAR(255) UNIQUE NOT NULL,
                           address TEXT NOT NULL,
                           latitude DECIMAL(10, 8) NOT NULL,
                           longitude DECIMAL(11, 8) NOT NULL,
                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                           updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE gates (
                       id SERIAL PRIMARY KEY,
                       terminal_id INTEGER NOT NULL REFERENCES terminals(id),
                       name VARCHAR(255) NOT NULL,
                       is_active BOOLEAN DEFAULT TRUE,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE terminal_distances (
                                    id SERIAL PRIMARY KEY,
                                    from_terminal_id INTEGER NOT NULL REFERENCES terminals(id),
                                    to_terminal_id INTEGER NOT NULL REFERENCES terminals(id),
                                    distance DECIMAL(10, 2) NOT NULL,
                                    base_price DECIMAL(10, 2) NOT NULL,
                                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                    UNIQUE(from_terminal_id, to_terminal_id)
);

CREATE TABLE cards (
                       id SERIAL PRIMARY KEY,
                       number VARCHAR(255) UNIQUE NOT NULL,
                       balance DECIMAL(10, 2) NOT NULL DEFAULT 0,
                       user_id INTEGER NOT NULL REFERENCES users(id),
                       is_active BOOLEAN DEFAULT TRUE,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE card_transactions (
                                   id SERIAL PRIMARY KEY,
                                   card_id INTEGER NOT NULL REFERENCES cards(id),
                                   amount DECIMAL(10, 2) NOT NULL,
                                   type VARCHAR(20) NOT NULL,
                                   reference_id VARCHAR(255) UNIQUE NOT NULL,
                                   is_synced BOOLEAN DEFAULT FALSE,
                                   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_card_transactions_card_id ON card_transactions(card_id);
CREATE INDEX idx_card_transactions_reference_id ON card_transactions(reference_id);