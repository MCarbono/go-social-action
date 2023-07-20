DROP TABLE IF EXISTS volunteers;
DROP TABLE IF EXISTS social_actions;
DROP TABLE IF EXISTS social_actions_volunteers;

CREATE TABLE volunteers (
    id TEXT PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    neighborhood TEXT NOT NULL,
    city TEXT NOT NULL,
    created_at timestamp,
    updated_at timestamp
);

CREATE TABLE social_actions (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    organizer TEXT NOT NULL,
    description TEXT NOT NULL,
    street_line TEXT NOT NULL,
    street_number TEXT NOT NULL,
    neighborhood TEXT NOT NULL,
    city TEXT NOT NULL,
    created_at timestamp,
    updated_at timestamp
);

CREATE TABLE social_actions_volunteers (
    id TEXT,
    social_action_id TEXT REFERENCES social_actions (id) ON DELETE CASCADE,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    neighborhood TEXT,
    city TEXT,
    CONSTRAINT unique_id_social_action_id
    UNIQUE (id, social_action_id)
);