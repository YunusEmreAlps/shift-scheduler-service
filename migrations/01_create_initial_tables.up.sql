DROP TABLE IF EXISTS shift_schedule CASCADE;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS CITEXT;
-- CREATE EXTENSION IF NOT EXISTS pg_cron;
-- CREATE EXTENSION IF NOT EXISTS postgis;
-- CREATE EXTENSION IF NOT EXISTS postgis_topology;

CREATE TABLE IF NOT EXISTS shift_schedule (
 	id SERIAL PRIMARY KEY,
    alias VARCHAR(255) NOT NULL,
    description VARCHAR(1024) DEFAULT NULL,
    frequency INTEGER NOT NULL DEFAULT 7,
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE NOT NULL,
    year INTEGER NOT NULL,
    status INTEGER NOT NULL,
    organization JSONB NOT NULL, -- group name, mail, phone, description
    manager JSONB NOT NULL,
    shifts JSONB DEFAULT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

-- Alter table to frequency
-- ALTER TABLE shift_schedule ADD COLUMN frequency INTEGER NOT NULL DEFAULT 7;


-- CRON JOB
-- SELECT cron.schedule('0 0 * * *', $$UPDATE shift_schedule SET deleted_at = NOW() WHERE end_date < NOW()$$, 'delete_expired_shift_schedule', false);

-- Shift Schedule 1
INSERT INTO shift_schedule (
    alias,
    description,
    frequency,
    start_date,
    end_date,
    year,
    status,
    organization,
    manager,
    shifts
) VALUES (
    'Shift 1',
    'Description 1',
    7,
    '2023-01-01 00:00:00',
    '2023-10-01 00:00:00',
    2023,
    0,
    '[{"id": 1, "name": "Group 1", "mail": "", "phone": "", "description": ""}]',
    '[{"id": 1, "name": "Manager 1", "mail": "", "phone": "", "description": ""}]',
    '[
        {
            "id": 0,
            "start": "2023-01-01 00:00:00",
            "end": "2023-02-01 00:00:00",
            "user": {
                "id": 21304362,
                "name": "User 1",
                "mail": "",
                "phone": "",
                "description": ""
            }
        },
        {
            "id": 1,
            "start": "2023-01-01 00:00:00",
            "end": "2023-02-01 00:00:00",
            "user": {
                "id": 21304362,
                "name": "User 2",
                "mail": "",
                "phone": "",
                "description": ""

            }
        },
        {
            "id": 2,
            "start": "2023-01-01 00:00:00",
            "end": "2023-02-01 00:00:00",
            "user": {
                "id": 21304362,
                "name": "User 3",
                "mail": "",
                "phone": "",
                "description": ""
            }
        }
    ]'
);