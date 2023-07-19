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
    id TEXT PRIMARY KEY,
    social_action_id TEXT REFERENCES social_actions (id) ON DELETE CASCADE,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    neighborhood TEXT,
    city TEXT
);

-- type SocialActionVolunteer struct {
-- 	ID           string
-- 	FirstName    string
-- 	LastName     string
-- 	Neighborhood string
-- 	City         string
-- 	CreatedAt    time.Time
-- 	UpdatedAt    time.Time
-- }

-- type SocialAction struct {
-- 	ID                    string
-- 	Name                  string
-- 	Organizer             string
-- 	Address               Address
-- 	Description           string
-- 	SocialActionVolunteer []SocialActionVolunteer
-- }

-- type Address struct {
-- 	StreetLine   string
-- 	StreetNumber string
-- 	Neighborhood string
-- 	City         string
-- 	State        string
-- }

