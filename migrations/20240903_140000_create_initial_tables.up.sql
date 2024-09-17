-- File Name: 20240903_140000_create_initial_tables.up.sql
-- Date: 2024-09-03 14:00:00
-- Author: Yunus Emre Alpu

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
    users JSONB DEFAULT NULL,
    shifts JSONB DEFAULT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

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
    users,
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
            "name": "User 1",
            "mail": "",
            "phone": "",
            "description": ""
        },
        {
            "id": 1,
            "name": "User 2",
            "mail": "",
            "phone": "",
            "description": ""
        },
        {
            "id": 2,
            "name": "User 3",
            "mail": "",
            "phone": "",
            "description": ""
        }
    ]',
    '[
        {
            "id": 0,
            "start": "2023-01-01 00:00:00",
            "end": "2023-02-01 00:00:00",
            "user": {
                "id": 0,
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
                "id": 1,
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
                "id": 2,
                "name": "User 3",
                "mail": "",
                "phone": "",
                "description": ""
            }
        }
    ]'
);

-- Production DB (PROD)

-- Alter table to frequency
-- ALTER TABLE shift_schedule ADD COLUMN frequency INTEGER NOT NULL DEFAULT 7;

-- But in PROD DB first I need to change shifts to users and then add shifts column (This is PROD DB)
-- ALTER TABLE shift_schedule RENAME COLUMN shifts TO users;
-- ALTER TABLE shift_schedule ADD COLUMN shifts JSONB DEFAULT NULL;

-- CRON JOB
-- SELECT cron.schedule('soft_delete_expired_shift_schedules', '0 0 * * *', $$UPDATE shift_schedule SET deleted_at = NOW() WHERE end_date < NOW()$$);
-- SELECT * from cron.job;